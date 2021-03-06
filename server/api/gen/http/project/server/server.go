// Code generated by goa v3.1.2, DO NOT EDIT.
//
// project HTTP server
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	"context"
	"net/http"
	"regexp"

	project "github.com/fieldkit/cloud/server/api/gen/project"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the project service endpoint HTTP handlers.
type Server struct {
	Mounts       []*MountPoint
	AddUpdate    http.Handler
	DeleteUpdate http.Handler
	ModifyUpdate http.Handler
	Invites      http.Handler
	LookupInvite http.Handler
	AcceptInvite http.Handler
	RejectInvite http.Handler
	CORS         http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the project service endpoints using
// the provided encoder and decoder. The handlers are mounted on the given mux
// using the HTTP verb and path defined in the design. errhandler is called
// whenever a response fails to be encoded. formatter is used to format errors
// returned by the service methods prior to encoding. Both errhandler and
// formatter are optional and can be nil.
func New(
	e *project.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"AddUpdate", "POST", "/projects/{projectId}/updates"},
			{"DeleteUpdate", "DELETE", "/projects/{projectId}/updates/{updateId}"},
			{"ModifyUpdate", "POST", "/projects/{projectId}/updates/{updateId}"},
			{"Invites", "GET", "/projects/invites/pending"},
			{"LookupInvite", "GET", "/projects/invites/{token}"},
			{"AcceptInvite", "POST", "/projects/invites/{id}/accept"},
			{"RejectInvite", "POST", "/projects/invites/{id}/reject"},
			{"CORS", "OPTIONS", "/projects/{projectId}/updates"},
			{"CORS", "OPTIONS", "/projects/{projectId}/updates/{updateId}"},
			{"CORS", "OPTIONS", "/projects/invites/pending"},
			{"CORS", "OPTIONS", "/projects/invites/{token}"},
			{"CORS", "OPTIONS", "/projects/invites/{id}/accept"},
			{"CORS", "OPTIONS", "/projects/invites/{id}/reject"},
		},
		AddUpdate:    NewAddUpdateHandler(e.AddUpdate, mux, decoder, encoder, errhandler, formatter),
		DeleteUpdate: NewDeleteUpdateHandler(e.DeleteUpdate, mux, decoder, encoder, errhandler, formatter),
		ModifyUpdate: NewModifyUpdateHandler(e.ModifyUpdate, mux, decoder, encoder, errhandler, formatter),
		Invites:      NewInvitesHandler(e.Invites, mux, decoder, encoder, errhandler, formatter),
		LookupInvite: NewLookupInviteHandler(e.LookupInvite, mux, decoder, encoder, errhandler, formatter),
		AcceptInvite: NewAcceptInviteHandler(e.AcceptInvite, mux, decoder, encoder, errhandler, formatter),
		RejectInvite: NewRejectInviteHandler(e.RejectInvite, mux, decoder, encoder, errhandler, formatter),
		CORS:         NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "project" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.AddUpdate = m(s.AddUpdate)
	s.DeleteUpdate = m(s.DeleteUpdate)
	s.ModifyUpdate = m(s.ModifyUpdate)
	s.Invites = m(s.Invites)
	s.LookupInvite = m(s.LookupInvite)
	s.AcceptInvite = m(s.AcceptInvite)
	s.RejectInvite = m(s.RejectInvite)
	s.CORS = m(s.CORS)
}

// Mount configures the mux to serve the project endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountAddUpdateHandler(mux, h.AddUpdate)
	MountDeleteUpdateHandler(mux, h.DeleteUpdate)
	MountModifyUpdateHandler(mux, h.ModifyUpdate)
	MountInvitesHandler(mux, h.Invites)
	MountLookupInviteHandler(mux, h.LookupInvite)
	MountAcceptInviteHandler(mux, h.AcceptInvite)
	MountRejectInviteHandler(mux, h.RejectInvite)
	MountCORSHandler(mux, h.CORS)
}

// MountAddUpdateHandler configures the mux to serve the "project" service "add
// update" endpoint.
func MountAddUpdateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleProjectOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/projects/{projectId}/updates", f)
}

// NewAddUpdateHandler creates a HTTP handler which loads the HTTP request and
// calls the "project" service "add update" endpoint.
func NewAddUpdateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeAddUpdateRequest(mux, decoder)
		encodeResponse = EncodeAddUpdateResponse(encoder)
		encodeError    = EncodeAddUpdateError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "add update")
		ctx = context.WithValue(ctx, goa.ServiceKey, "project")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDeleteUpdateHandler configures the mux to serve the "project" service
// "delete update" endpoint.
func MountDeleteUpdateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleProjectOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/projects/{projectId}/updates/{updateId}", f)
}

// NewDeleteUpdateHandler creates a HTTP handler which loads the HTTP request
// and calls the "project" service "delete update" endpoint.
func NewDeleteUpdateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteUpdateRequest(mux, decoder)
		encodeResponse = EncodeDeleteUpdateResponse(encoder)
		encodeError    = EncodeDeleteUpdateError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "delete update")
		ctx = context.WithValue(ctx, goa.ServiceKey, "project")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountModifyUpdateHandler configures the mux to serve the "project" service
// "modify update" endpoint.
func MountModifyUpdateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleProjectOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/projects/{projectId}/updates/{updateId}", f)
}

// NewModifyUpdateHandler creates a HTTP handler which loads the HTTP request
// and calls the "project" service "modify update" endpoint.
func NewModifyUpdateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeModifyUpdateRequest(mux, decoder)
		encodeResponse = EncodeModifyUpdateResponse(encoder)
		encodeError    = EncodeModifyUpdateError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "modify update")
		ctx = context.WithValue(ctx, goa.ServiceKey, "project")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountInvitesHandler configures the mux to serve the "project" service
// "invites" endpoint.
func MountInvitesHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleProjectOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/projects/invites/pending", f)
}

// NewInvitesHandler creates a HTTP handler which loads the HTTP request and
// calls the "project" service "invites" endpoint.
func NewInvitesHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeInvitesRequest(mux, decoder)
		encodeResponse = EncodeInvitesResponse(encoder)
		encodeError    = EncodeInvitesError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "invites")
		ctx = context.WithValue(ctx, goa.ServiceKey, "project")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountLookupInviteHandler configures the mux to serve the "project" service
// "lookup invite" endpoint.
func MountLookupInviteHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleProjectOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/projects/invites/{token}", f)
}

// NewLookupInviteHandler creates a HTTP handler which loads the HTTP request
// and calls the "project" service "lookup invite" endpoint.
func NewLookupInviteHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeLookupInviteRequest(mux, decoder)
		encodeResponse = EncodeLookupInviteResponse(encoder)
		encodeError    = EncodeLookupInviteError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "lookup invite")
		ctx = context.WithValue(ctx, goa.ServiceKey, "project")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountAcceptInviteHandler configures the mux to serve the "project" service
// "accept invite" endpoint.
func MountAcceptInviteHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleProjectOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/projects/invites/{id}/accept", f)
}

// NewAcceptInviteHandler creates a HTTP handler which loads the HTTP request
// and calls the "project" service "accept invite" endpoint.
func NewAcceptInviteHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeAcceptInviteRequest(mux, decoder)
		encodeResponse = EncodeAcceptInviteResponse(encoder)
		encodeError    = EncodeAcceptInviteError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "accept invite")
		ctx = context.WithValue(ctx, goa.ServiceKey, "project")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountRejectInviteHandler configures the mux to serve the "project" service
// "reject invite" endpoint.
func MountRejectInviteHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := handleProjectOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/projects/invites/{id}/reject", f)
}

// NewRejectInviteHandler creates a HTTP handler which loads the HTTP request
// and calls the "project" service "reject invite" endpoint.
func NewRejectInviteHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeRejectInviteRequest(mux, decoder)
		encodeResponse = EncodeRejectInviteResponse(encoder)
		encodeError    = EncodeRejectInviteError(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "reject invite")
		ctx = context.WithValue(ctx, goa.ServiceKey, "project")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service project.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = handleProjectOrigin(h)
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("OPTIONS", "/projects/{projectId}/updates", f)
	mux.Handle("OPTIONS", "/projects/{projectId}/updates/{updateId}", f)
	mux.Handle("OPTIONS", "/projects/invites/pending", f)
	mux.Handle("OPTIONS", "/projects/invites/{token}", f)
	mux.Handle("OPTIONS", "/projects/invites/{id}/accept", f)
	mux.Handle("OPTIONS", "/projects/invites/{id}/reject", f)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// handleProjectOrigin applies the CORS response headers corresponding to the
// origin for the service project.
func handleProjectOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile("(.+[.])?fieldkit.org:\\d+")
	spec1 := regexp.MustCompile("(.+[.])?local.fkdev.org:\\d+")
	spec2 := regexp.MustCompile("(.+[.])?localhost:\\d+")
	origHndlr := h.(http.HandlerFunc)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec1) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec2) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://*.fieldkit.org") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://*.fieldkit.org:8080") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://*.fkdev.org") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://fieldkit.org") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://fieldkit.org:8080") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		if cors.MatchOrigin(origin, "https://fkdev.org") {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "Authorization, Content-Type")
			w.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST, DELETE, PATCH, PUT")
				w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
			}
			origHndlr(w, r)
			return
		}
		origHndlr(w, r)
		return
	})
}
