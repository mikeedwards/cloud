// Code generated by goa v3.1.2, DO NOT EDIT.
//
// tasks endpoints
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package tasks

import (
	"context"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "tasks" service endpoints.
type Endpoints struct {
	Five          goa.Endpoint
	RefreshDevice goa.Endpoint
}

// NewEndpoints wraps the methods of the "tasks" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Five:          NewFiveEndpoint(s),
		RefreshDevice: NewRefreshDeviceEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "tasks" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Five = m(e.Five)
	e.RefreshDevice = m(e.RefreshDevice)
}

// NewFiveEndpoint returns an endpoint function that calls the method "five" of
// service "tasks".
func NewFiveEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return nil, s.Five(ctx)
	}
}

// NewRefreshDeviceEndpoint returns an endpoint function that calls the method
// "refresh device" of service "tasks".
func NewRefreshDeviceEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*RefreshDevicePayload)
		var err error
		sc := security.JWTScheme{
			Name:           "jwt",
			Scopes:         []string{"api:access", "api:admin", "api:ingestion"},
			RequiredScopes: []string{"api:access"},
		}
		ctx, err = authJWTFn(ctx, p.Auth, &sc)
		if err != nil {
			return nil, err
		}
		return nil, s.RefreshDevice(ctx, p)
	}
}
