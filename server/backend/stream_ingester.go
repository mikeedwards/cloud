package backend

import (
	"context"
	"crypto/sha1"
	"hash"
	"io"
	"net/http"
	"time"

	"github.com/Conservify/sqlxcache"

	"github.com/fieldkit/cloud/server/backend/ingestion"
	"github.com/fieldkit/cloud/server/logging"
)

const (
	FkDataBinaryContentType = "application/vnd.fk.data+binary"
	ContentTypeHeaderName   = "Content-Type"
	ContentLengthHeaderName = "Content-Length"
	FkProcessingHeaderName  = "Fk-Processing"
)

type StreamIngester struct {
	backend        *Backend
	db             *sqlxcache.DB
	streamArchiver StreamArchiver
	sourceChanges  ingestion.SourceChangesPublisher
}

func NewStreamIngester(b *Backend, streamArchiver StreamArchiver, sourceChanges ingestion.SourceChangesPublisher) (si *StreamIngester, err error) {
	si = &StreamIngester{
		backend:        b,
		db:             b.db,
		streamArchiver: streamArchiver,
		sourceChanges:  sourceChanges,
	}

	return
}

type ReaderWrapper struct {
	BytesRead int64
	Target    io.Reader
	Hash      hash.Hash
}

func (rw *ReaderWrapper) Read(p []byte) (n int, err error) {
	n, err = rw.Target.Read(p)
	rw.BytesRead += int64(n)
	sliced := p[:n]
	rw.Hash.Write(sliced)
	return n, err
}

func (si *StreamIngester) synchronous(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	log := logging.Logger(ctx).Sugar()

	startedAt := time.Now()
	contentLength := req.Header.Get(ContentLengthHeaderName)
	status := http.StatusOK

	log.Infow("started", ContentLengthHeaderName, contentLength)

	reader := &ReaderWrapper{
		BytesRead: 0,
		Target:    req.Body,
		Hash:      sha1.New(),
	}

	err := si.backend.db.WithNewTransaction(ctx, func(txCtx context.Context) error {
		saver := NewFormattedMessageSaver(si.backend)
		binaryReader := NewFkBinaryReader(saver)

		if err := binaryReader.Read(txCtx, reader); err != nil {
			status = http.StatusInternalServerError
			log.Errorw("error", "error", err, "bytesRead", reader.BytesRead)
			return nil
		}

		saver.EmitChanges(txCtx, si.sourceChanges)

		return nil
	})

	if err != nil {
		status = http.StatusInternalServerError
		log.Infow("completed", "error", err, "bytesRead", reader.BytesRead, "hash", reader.Hash.Sum(nil), "time", time.Since(startedAt).String())
	} else {
		log.Infow("completed", "bytesRead", reader.BytesRead, "hash", reader.Hash.Sum(nil), "time", time.Since(startedAt).String())
	}

	w.WriteHeader(status)
}

func (si *StreamIngester) asynchronous(ctx context.Context, w http.ResponseWriter, req *http.Request) {
	log := logging.Logger(ctx).Sugar()

	startedAt := time.Now()
	contentType := req.Header.Get(ContentTypeHeaderName)
	contentLength := req.Header.Get(ContentLengthHeaderName)

	log.Infof("started (async)", ContentLengthHeaderName, contentLength, ContentTypeHeaderName, contentType)

	if err := si.streamArchiver.Archive(ctx, contentType, req.Body); err != nil {
		log.Infow("completed", "error", err, "time", time.Since(startedAt).String())
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		log.Infow("completed", "time", time.Since(startedAt).String())
		w.WriteHeader(http.StatusOK)
	}
}

var (
	ids = logging.NewIdGenerator()
)

func (si *StreamIngester) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get(ContentTypeHeaderName)
	fkProcessing := req.Header.Get(FkProcessingHeaderName)

	ctx := logging.WithTaskId(req.Context(), ids)
	log := logging.Logger(ctx).Sugar()

	if contentType != FkDataBinaryContentType {
		log.Infof("Stream: unknown content type: %v", contentType)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if fkProcessing == "" {
		si.synchronous(ctx, w, req)
	} else if fkProcessing == "async" {
		si.asynchronous(ctx, w, req)
	}
}
