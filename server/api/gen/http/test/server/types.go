// Code generated by goa v3.1.2, DO NOT EDIT.
//
// test HTTP server types
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package server

import (
	test "github.com/fieldkit/cloud/server/api/gen/test"
)

// GetBadRequestResponseBody is the type of the "test" service "get" endpoint
// HTTP response body for the "bad-request" error.
type GetBadRequestResponseBody string

// GetForbiddenResponseBody is the type of the "test" service "get" endpoint
// HTTP response body for the "forbidden" error.
type GetForbiddenResponseBody string

// GetNotFoundResponseBody is the type of the "test" service "get" endpoint
// HTTP response body for the "not-found" error.
type GetNotFoundResponseBody string

// GetUnauthorizedResponseBody is the type of the "test" service "get" endpoint
// HTTP response body for the "unauthorized" error.
type GetUnauthorizedResponseBody string

// ErrorBadRequestResponseBody is the type of the "test" service "error"
// endpoint HTTP response body for the "bad-request" error.
type ErrorBadRequestResponseBody string

// ErrorForbiddenResponseBody is the type of the "test" service "error"
// endpoint HTTP response body for the "forbidden" error.
type ErrorForbiddenResponseBody string

// ErrorNotFoundResponseBody is the type of the "test" service "error" endpoint
// HTTP response body for the "not-found" error.
type ErrorNotFoundResponseBody string

// ErrorUnauthorizedResponseBody is the type of the "test" service "error"
// endpoint HTTP response body for the "unauthorized" error.
type ErrorUnauthorizedResponseBody string

// EmailBadRequestResponseBody is the type of the "test" service "email"
// endpoint HTTP response body for the "bad-request" error.
type EmailBadRequestResponseBody string

// EmailForbiddenResponseBody is the type of the "test" service "email"
// endpoint HTTP response body for the "forbidden" error.
type EmailForbiddenResponseBody string

// EmailNotFoundResponseBody is the type of the "test" service "email" endpoint
// HTTP response body for the "not-found" error.
type EmailNotFoundResponseBody string

// EmailUnauthorizedResponseBody is the type of the "test" service "email"
// endpoint HTTP response body for the "unauthorized" error.
type EmailUnauthorizedResponseBody string

// NewGetBadRequestResponseBody builds the HTTP response body from the result
// of the "get" endpoint of the "test" service.
func NewGetBadRequestResponseBody(res test.BadRequest) GetBadRequestResponseBody {
	body := GetBadRequestResponseBody(res)
	return body
}

// NewGetForbiddenResponseBody builds the HTTP response body from the result of
// the "get" endpoint of the "test" service.
func NewGetForbiddenResponseBody(res test.Forbidden) GetForbiddenResponseBody {
	body := GetForbiddenResponseBody(res)
	return body
}

// NewGetNotFoundResponseBody builds the HTTP response body from the result of
// the "get" endpoint of the "test" service.
func NewGetNotFoundResponseBody(res test.NotFound) GetNotFoundResponseBody {
	body := GetNotFoundResponseBody(res)
	return body
}

// NewGetUnauthorizedResponseBody builds the HTTP response body from the result
// of the "get" endpoint of the "test" service.
func NewGetUnauthorizedResponseBody(res test.Unauthorized) GetUnauthorizedResponseBody {
	body := GetUnauthorizedResponseBody(res)
	return body
}

// NewErrorBadRequestResponseBody builds the HTTP response body from the result
// of the "error" endpoint of the "test" service.
func NewErrorBadRequestResponseBody(res test.BadRequest) ErrorBadRequestResponseBody {
	body := ErrorBadRequestResponseBody(res)
	return body
}

// NewErrorForbiddenResponseBody builds the HTTP response body from the result
// of the "error" endpoint of the "test" service.
func NewErrorForbiddenResponseBody(res test.Forbidden) ErrorForbiddenResponseBody {
	body := ErrorForbiddenResponseBody(res)
	return body
}

// NewErrorNotFoundResponseBody builds the HTTP response body from the result
// of the "error" endpoint of the "test" service.
func NewErrorNotFoundResponseBody(res test.NotFound) ErrorNotFoundResponseBody {
	body := ErrorNotFoundResponseBody(res)
	return body
}

// NewErrorUnauthorizedResponseBody builds the HTTP response body from the
// result of the "error" endpoint of the "test" service.
func NewErrorUnauthorizedResponseBody(res test.Unauthorized) ErrorUnauthorizedResponseBody {
	body := ErrorUnauthorizedResponseBody(res)
	return body
}

// NewEmailBadRequestResponseBody builds the HTTP response body from the result
// of the "email" endpoint of the "test" service.
func NewEmailBadRequestResponseBody(res test.BadRequest) EmailBadRequestResponseBody {
	body := EmailBadRequestResponseBody(res)
	return body
}

// NewEmailForbiddenResponseBody builds the HTTP response body from the result
// of the "email" endpoint of the "test" service.
func NewEmailForbiddenResponseBody(res test.Forbidden) EmailForbiddenResponseBody {
	body := EmailForbiddenResponseBody(res)
	return body
}

// NewEmailNotFoundResponseBody builds the HTTP response body from the result
// of the "email" endpoint of the "test" service.
func NewEmailNotFoundResponseBody(res test.NotFound) EmailNotFoundResponseBody {
	body := EmailNotFoundResponseBody(res)
	return body
}

// NewEmailUnauthorizedResponseBody builds the HTTP response body from the
// result of the "email" endpoint of the "test" service.
func NewEmailUnauthorizedResponseBody(res test.Unauthorized) EmailUnauthorizedResponseBody {
	body := EmailUnauthorizedResponseBody(res)
	return body
}

// NewGetPayload builds a test service get endpoint payload.
func NewGetPayload(id int64) *test.GetPayload {
	v := &test.GetPayload{}
	v.ID = &id

	return v
}

// NewEmailPayload builds a test service email endpoint payload.
func NewEmailPayload(address string, auth string) *test.EmailPayload {
	v := &test.EmailPayload{}
	v.Address = address
	v.Auth = auth

	return v
}
