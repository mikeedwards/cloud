// Code generated by goa v3.0.10, DO NOT EDIT.
//
// test HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	"context"
	"net/http"
	"strconv"

	test "github.com/fieldkit/cloud/server/api/gen/test"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// EncodeGetResponse returns an encoder for responses returned by the test get
// endpoint.
func EncodeGetResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusOK)
		return nil
	}
}

// DecodeGetRequest returns a decoder for requests sent to the test get
// endpoint.
func DecodeGetRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  int64
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewGetPayload(id)

		return payload, nil
	}
}

// EncodeErrorResponse returns an encoder for responses returned by the test
// error endpoint.
func EncodeErrorResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		w.WriteHeader(http.StatusNoContent)
		return nil
	}
}

// EncodeJSONResponse returns an encoder for responses returned by the test
// json endpoint.
func EncodeJSONResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) func(context.Context, http.ResponseWriter, interface{}) error {
	return func(ctx context.Context, w http.ResponseWriter, v interface{}) error {
		res := v.(*test.JSONResult)
		enc := encoder(ctx, w)
		body := res.JSON
		w.WriteHeader(http.StatusOK)
		return enc.Encode(body)
	}
}

// DecodeJSONRequest returns a decoder for requests sent to the test json
// endpoint.
func DecodeJSONRequest(mux goahttp.Muxer, decoder func(*http.Request) goahttp.Decoder) func(*http.Request) (interface{}, error) {
	return func(r *http.Request) (interface{}, error) {
		var (
			id  int64
			err error

			params = mux.Vars(r)
		)
		{
			idRaw := params["id"]
			v, err2 := strconv.ParseInt(idRaw, 10, 64)
			if err2 != nil {
				err = goa.MergeErrors(err, goa.InvalidFieldTypeError("id", idRaw, "integer"))
			}
			id = v
		}
		if err != nil {
			return nil, err
		}
		payload := NewJSONPayload(id)

		return payload, nil
	}
}
