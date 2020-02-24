// Code generated by goa v3.0.10, DO NOT EDIT.
//
// test endpoints
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package test

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "test" service endpoints.
type Endpoints struct {
	Get   goa.Endpoint
	Error goa.Endpoint
	JSON  goa.Endpoint
}

// NewEndpoints wraps the methods of the "test" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		Get:   NewGetEndpoint(s),
		Error: NewErrorEndpoint(s),
		JSON:  NewJSONEndpoint(s),
	}
}

// Use applies the given middleware to all the "test" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Get = m(e.Get)
	e.Error = m(e.Error)
	e.JSON = m(e.JSON)
}

// NewGetEndpoint returns an endpoint function that calls the method "get" of
// service "test".
func NewGetEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetPayload)
		return nil, s.Get(ctx, p)
	}
}

// NewErrorEndpoint returns an endpoint function that calls the method "error"
// of service "test".
func NewErrorEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, s.Error(ctx)
	}
}

// NewJSONEndpoint returns an endpoint function that calls the method "json" of
// service "test".
func NewJSONEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*JSONPayload)
		return s.JSON(ctx, p)
	}
}
