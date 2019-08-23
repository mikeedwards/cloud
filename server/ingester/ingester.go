package ingester

import (
	"context"
	"encoding/base64"
	"fmt"
	"mime"
	"net/http"
	"strconv"
	"strings"
	"time"

	_ "github.com/jmoiron/sqlx/types"
	"github.com/lib/pq"

	jwtgo "github.com/dgrijalva/jwt-go"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"

	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/conservify/sqlxcache"

	"github.com/fieldkit/cloud/server/logging"
)

const (
	FkDataBinaryContentType    = "application/vnd.fk.data+binary"
	FkDataBase64ContentType    = "application/vnd.fk.data+base64"
	MultiPartFormDataMediaType = "multipart/form-data"
	ContentTypeHeaderName      = "Content-Type"
	ContentLengthHeaderName    = "Content-Length"
	XForwardedForHeaderName    = "X-Forwarded-For"
	FkDeviceIdHeaderName       = "Fk-DeviceId"
	FkBlocksIdHeaderName       = "Fk-Blocks"
	FkFlagsIdHeaderName        = "Fk-Flags"
)

var (
	ids = logging.NewIdGenerator()
)

type IngesterOptions struct {
	Database                 *sqlxcache.DB
	AwsSession               *session.Session
	AuthenticationMiddleware goa.Middleware
	Archiver                 StreamArchiver
}

}

func Ingester(ctx context.Context, o *IngesterOptions) http.Handler {
	handler := authentication(o.AuthenticationMiddleware, func(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
		startedAt := time.Now()

		log := logging.Logger(ctx).Sugar()

		token := jwt.ContextJWT(ctx)
		if token == nil {
			return fmt.Errorf("JWT token is missing from context")
		}

		claims, ok := token.Claims.(jwtgo.MapClaims)
		if !ok {
			return fmt.Errorf("JWT claims error")
		}

		headers, err := NewIncomingHeaders(req)
		if err != nil {
			return err
		}

		log.Infow("receiving", "device_id", headers.FkDeviceId, "blocks", headers.FkBlocks)

		if saved, err := o.Archiver.Archive(ctx, headers, req.Body); err != nil {
			return err
		} else {
			if saved != nil {
				if saved.BytesRead != int(headers.ContentLength) {
					log.Warnw("size mismatch", "expected", headers.ContentLength, "actual", saved.BytesRead)
				}

				log.Infow("saved", "stream_id", saved.ID, "time", time.Since(startedAt).String(), "size", saved.BytesRead)
			} else {
				log.Infow("unsaved", "stream_id", saved.ID, "time", time.Since(startedAt).String(), "size", saved.BytesRead)
			}

			// TODO Give information.
			w.WriteHeader(http.StatusOK)
		}

		_ = claims
		_ = headers
		_ = log

		return nil
	})

	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		startedAt := time.Now()

		ctx := logging.WithNewTaskId(req.Context(), ids)

		log := logging.Logger(ctx).Sugar()

		log.Infow("begin")

		err := handler(ctx, w, req)
		if err != nil {
			log.Errorw("completed", "error", err, "time", time.Since(startedAt).String())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Infow("completed", "time", time.Since(startedAt).String())
	})
}

type IncomingHeaders struct {
	ContentType     string
	ContentLength   int32
	MediaType       string
	MediaTypeParams map[string]string
	XForwardedFor   string
	FkDeviceId      []byte
	FkBlocks        []int32
	FkFlags         []int32
}

func NewIncomingHeaders(req *http.Request) (*IncomingHeaders, error) {
	contentType := req.Header.Get(ContentTypeHeaderName)
	mediaType, mediaTypeParams, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, fmt.Errorf("Invalid %s (%s)", ContentTypeHeaderName, contentType)
	}

	contentLengthString := req.Header.Get(ContentLengthHeaderName)
	contentLength, err := strconv.Atoi(contentLengthString)
	if err != nil {
		return nil, fmt.Errorf("Invalid %s (%s)", ContentLengthHeaderName, contentLengthString)
	}

	if contentLength <= 0 {
		return nil, fmt.Errorf("Invalid %s (%v)", ContentLengthHeaderName, contentLength)
	}

	deviceIdRaw := req.Header.Get(FkDeviceIdHeaderName)
	if len(deviceIdRaw) == 0 {
		return nil, fmt.Errorf("Invalid %s (No header)", FkDeviceIdHeaderName)
	}

	deviceId, err := base64.StdEncoding.DecodeString(deviceIdRaw)
	if err != nil {
		return nil, fmt.Errorf("Invalid %s (%v)", FkDeviceIdHeaderName, err)
	}

	blocks, err := parseBlocks(req.Header.Get(FkBlocksIdHeaderName))
	if err != nil {
		return nil, fmt.Errorf("Invalid %s (%v)", FkBlocksIdHeaderName, err)
	}

	headers := &IncomingHeaders{
		ContentType:     contentType,
		ContentLength:   int32(contentLength),
		MediaType:       mediaType,
		MediaTypeParams: mediaTypeParams,
		XForwardedFor:   req.Header.Get(XForwardedForHeaderName),
		FkDeviceId:      deviceId,
		FkBlocks:        blocks,
	}

	return headers, nil
}

func parseBlocks(s string) ([]int32, error) {
	parts := strings.Split(s, ",")

	if len(parts) != 2 {
		return nil, fmt.Errorf("Malformed block range")
	}

	blocks := make([]int32, 2)
	for i, p := range parts {
		b, err := strconv.Atoi(strings.TrimSpace(p))
		if err != nil {
			return nil, err
		}
		blocks[i] = int32(b)
	}

	return blocks, nil
}

func authentication(middleware goa.Middleware, next goa.Handler) goa.Handler {
	return func(ctx context.Context, res http.ResponseWriter, req *http.Request) error {
		ctx = goa.WithRequiredScopes(ctx, []string{"api:access"})
		return middleware(next)(ctx, res, req)
	}
}