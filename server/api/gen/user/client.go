// Code generated by goa v3.1.2, DO NOT EDIT.
//
// user client
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package user

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "user" service client.
type Client struct {
	RolesEndpoint goa.Endpoint
}

// NewClient initializes a "user" service client given the endpoints.
func NewClient(roles goa.Endpoint) *Client {
	return &Client{
		RolesEndpoint: roles,
	}
}

// Roles calls the "roles" endpoint of the "user" service.
func (c *Client) Roles(ctx context.Context, p *RolesPayload) (res *AvailableRoles, err error) {
	var ires interface{}
	ires, err = c.RolesEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*AvailableRoles), nil
}