// Code generated by goa v3.1.2, DO NOT EDIT.
//
// station endpoints
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package station

import (
	"context"
	"io"

	goa "goa.design/goa/v3/pkg"
	"goa.design/goa/v3/security"
)

// Endpoints wraps the "station" service endpoints.
type Endpoints struct {
	Add         goa.Endpoint
	Get         goa.Endpoint
	Update      goa.Endpoint
	ListMine    goa.Endpoint
	ListProject goa.Endpoint
	Photo       goa.Endpoint
}

// PhotoResponseData holds both the result and the HTTP response body reader of
// the "photo" method.
type PhotoResponseData struct {
	// Result is the method result.
	Result *PhotoResult
	// Body streams the HTTP response body.
	Body io.ReadCloser
}

// NewEndpoints wraps the methods of the "station" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	// Casting service to Auther interface
	a := s.(Auther)
	return &Endpoints{
		Add:         NewAddEndpoint(s, a.JWTAuth),
		Get:         NewGetEndpoint(s, a.JWTAuth),
		Update:      NewUpdateEndpoint(s, a.JWTAuth),
		ListMine:    NewListMineEndpoint(s, a.JWTAuth),
		ListProject: NewListProjectEndpoint(s, a.JWTAuth),
		Photo:       NewPhotoEndpoint(s, a.JWTAuth),
	}
}

// Use applies the given middleware to all the "station" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.Add = m(e.Add)
	e.Get = m(e.Get)
	e.Update = m(e.Update)
	e.ListMine = m(e.ListMine)
	e.ListProject = m(e.ListProject)
	e.Photo = m(e.Photo)
}

// NewAddEndpoint returns an endpoint function that calls the method "add" of
// service "station".
func NewAddEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*AddPayload)
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
		res, err := s.Add(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStationFull(res, "default")
		return vres, nil
	}
}

// NewGetEndpoint returns an endpoint function that calls the method "get" of
// service "station".
func NewGetEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*GetPayload)
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
		res, err := s.Get(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStationFull(res, "default")
		return vres, nil
	}
}

// NewUpdateEndpoint returns an endpoint function that calls the method
// "update" of service "station".
func NewUpdateEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*UpdatePayload)
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
		res, err := s.Update(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStationFull(res, "default")
		return vres, nil
	}
}

// NewListMineEndpoint returns an endpoint function that calls the method "list
// mine" of service "station".
func NewListMineEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListMinePayload)
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
		res, err := s.ListMine(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStationsFull(res, "default")
		return vres, nil
	}
}

// NewListProjectEndpoint returns an endpoint function that calls the method
// "list project" of service "station".
func NewListProjectEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*ListProjectPayload)
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
		res, err := s.ListProject(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedStationsFull(res, "default")
		return vres, nil
	}
}

// NewPhotoEndpoint returns an endpoint function that calls the method "photo"
// of service "station".
func NewPhotoEndpoint(s Service, authJWTFn security.AuthJWTFunc) goa.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		p := req.(*PhotoPayload)
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
		res, body, err := s.Photo(ctx, p)
		if err != nil {
			return nil, err
		}
		return &PhotoResponseData{Result: res, Body: body}, nil
	}
}
