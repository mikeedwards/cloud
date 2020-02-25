// Code generated by goa v3.0.10, DO NOT EDIT.
//
// tasks service
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package tasks

import (
	"context"

	"goa.design/goa/v3/security"
)

// Service is the tasks service interface.
type Service interface {
	// Five implements five.
	Five(context.Context) (err error)
	// RefreshDevice implements refresh device.
	RefreshDevice(context.Context, *RefreshDevicePayload) (err error)
}

// Auther defines the authorization functions to be implemented by the service.
type Auther interface {
	// JWTAuth implements the authorization logic for the JWT security scheme.
	JWTAuth(ctx context.Context, token string, schema *security.JWTScheme) (context.Context, error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "tasks"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"five", "refresh device"}

// RefreshDevicePayload is the payload type of the tasks service refresh device
// method.
type RefreshDevicePayload struct {
	Auth     string
	DeviceID string
}

// credentials are invalid
type Unauthorized string

// token scopes are invalid
type InvalidScopes string

// Error returns an error description.
func (e Unauthorized) Error() string {
	return "credentials are invalid"
}

// ErrorName returns "unauthorized".
func (e Unauthorized) ErrorName() string {
	return "unauthorized"
}

// Error returns an error description.
func (e InvalidScopes) Error() string {
	return "token scopes are invalid"
}

// ErrorName returns "invalid-scopes".
func (e InvalidScopes) ErrorName() string {
	return "invalid-scopes"
}