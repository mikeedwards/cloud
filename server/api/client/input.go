// Code generated by goagen v1.1.0, command line:
// $ main
//
// API "fieldkit": input Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
	"strconv"
)

// ListInputPath computes a request path to the list action of input.
func ListInputPath(project string, expedition string) string {
	param0 := project
	param1 := expedition

	return fmt.Sprintf("/projects/@/%s/expeditions/@/%s/inputs", param0, param1)
}

// List a project's inputs
func (c *Client) ListInput(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListInputRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListInputRequest create the request corresponding to the list action endpoint of the input resource.
func (c *Client) NewListInputRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ListIDInputPath computes a request path to the list id action of input.
func ListIDInputPath(expeditionID int) string {
	param0 := strconv.Itoa(expeditionID)

	return fmt.Sprintf("/expeditions/%s/inputs", param0)
}

// List a project's inputs
func (c *Client) ListIDInput(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListIDInputRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListIDInputRequest create the request corresponding to the list id action endpoint of the input resource.
func (c *Client) NewListIDInputRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}
