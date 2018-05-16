package backend

import (
	"context"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"mime"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/Conservify/sqlxcache"

	"github.com/fieldkit/cloud/server/backend/ingestion"
	"github.com/fieldkit/cloud/server/backend/ingestion/formatting"
	"github.com/fieldkit/cloud/server/logging"
)

const (
	FkDataBinaryContentType = "application/vnd.fk.data+binary"
	FkDataBase64ContentType = "application/vnd.fk.data+base64"
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

func (si *StreamIngester) synchronous(ctx context.Context, mediaType string, w http.ResponseWriter, req *http.Request) {
	log := Logger(ctx).Sugar()

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

		if mediaType == FkDataBinaryContentType {
			binaryReader := NewFkBinaryReader(saver)
			if err := binaryReader.Read(txCtx, reader); err != nil {
				return err
			}
		} else if mediaType == FkDataBase64ContentType {
			decoder := base64.NewDecoder(base64.StdEncoding, reader)
			binaryReader := NewFkBinaryReader(saver)
			if err := binaryReader.Read(txCtx, decoder); err != nil {
				return err
			}
		} else {
			decoder := json.NewDecoder(req.Body)
			message := &formatting.HttpJsonMessage{}
			err := decoder.Decode(message)
			if err != nil {
				return fmt.Errorf("JSON Error: '%v'", err)
			}

			messageId, err := uuid.NewRandom()
			if err != nil {
				return err
			}

			fm, err := message.ToFormattedMessage(ingestion.MessageId(messageId.String()))
			if err != nil {
				return err
			}

			recordChange, err := saver.HandleFormattedMessage(ctx, fm)
			if err != nil {
				return err
			}

			_ = recordChange
		}

		saver.EmitChanges(txCtx, si.sourceChanges)

		return nil
	})

	if err != nil {
		status = http.StatusInternalServerError
		log.Errorw("completed", "error", err, "bytesRead", reader.BytesRead, "hash", reader.Hash.Sum(nil), "time", time.Since(startedAt).String())
	} else {
		log.Infow("completed", "bytesRead", reader.BytesRead, "hash", reader.Hash.Sum(nil), "time", time.Since(startedAt).String())
	}

	w.WriteHeader(status)
}

func (si *StreamIngester) asynchronous(ctx context.Context, mediaType string, w http.ResponseWriter, req *http.Request) {
	log := Logger(ctx).Sugar()

	startedAt := time.Now()
	contentType := req.Header.Get(ContentTypeHeaderName)
	contentLength := req.Header.Get(ContentLengthHeaderName)

	log.Infof("started (async)", ContentLengthHeaderName, contentLength, ContentTypeHeaderName, contentType)

	if err := si.streamArchiver.Archive(ctx, contentType, req.Body); err != nil {
		log.Errorw("completed", "error", err, "time", time.Since(startedAt).String())
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
	ctx := logging.WithNewTaskId(req.Context(), ids)
	log := Logger(ctx).Sugar()

	contentType := req.Header.Get(ContentTypeHeaderName)
	fkProcessing := req.Header.Get(FkProcessingHeaderName)

	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		log.Errorw("completed", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if mediaType != FkDataBinaryContentType && mediaType != FkDataBase64ContentType && mediaType != formatting.HttpProviderJsonContentType {
		log.Infow(fmt.Sprintf("Unknown content type '%v'", contentType), "Content-Type", contentType)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if fkProcessing == "" {
		si.synchronous(ctx, mediaType, w, req)
	} else if fkProcessing == "async" {
		si.asynchronous(ctx, mediaType, w, req)
	}
}
