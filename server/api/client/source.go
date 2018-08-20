// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "fieldkit": source Resource Client
//
// Command:
// $ main

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ListSourcePath computes a request path to the list action of source.
func ListSourcePath(project string, expedition string) string {
	param0 := project
	param1 := expedition

	return fmt.Sprintf("/projects/@/%s/expeditions/@/%s/sources", param0, param1)
}

// List a project's sources
func (c *Client) ListSource(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListSourceRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListSourceRequest create the request corresponding to the list action endpoint of the source resource.
func (c *Client) NewListSourceRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListExpeditionIDSourcePath computes a request path to the list expedition id action of source.
func ListExpeditionIDSourcePath(expeditionID int) string {
	param0 := strconv.Itoa(expeditionID)

	return fmt.Sprintf("/expeditions/%s/sources", param0)
}

// List an expedition's sources
func (c *Client) ListExpeditionIDSource(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListExpeditionIDSourceRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListExpeditionIDSourceRequest create the request corresponding to the list expedition id action endpoint of the source resource.
func (c *Client) NewListExpeditionIDSourceRequest(ctx context.Context, path string) (*http.Request, error) {
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
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// ListIDSourcePath computes a request path to the list id action of source.
func ListIDSourcePath(sourceID int) string {
	param0 := strconv.Itoa(sourceID)

	return fmt.Sprintf("/sources/%s", param0)
}

// List an source
func (c *Client) ListIDSource(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListIDSourceRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListIDSourceRequest create the request corresponding to the list id action endpoint of the source resource.
func (c *Client) NewListIDSourceRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// SummaryByIDSourcePath computes a request path to the summary by id action of source.
func SummaryByIDSourcePath(sourceID int) string {
	param0 := strconv.Itoa(sourceID)

	return fmt.Sprintf("/sources/%s/summary", param0)
}

// List an source
func (c *Client) SummaryByIDSource(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewSummaryByIDSourceRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSummaryByIDSourceRequest create the request corresponding to the summary by id action endpoint of the source resource.
func (c *Client) NewSummaryByIDSourceRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// TemporalClusterGeometryByIDSourcePath computes a request path to the temporal cluster geometry by id action of source.
func TemporalClusterGeometryByIDSourcePath(sourceID int, clusterID int) string {
	param0 := strconv.Itoa(sourceID)
	param1 := strconv.Itoa(clusterID)

	return fmt.Sprintf("/sources/%s/temporal/%s/geometry", param0, param1)
}

// Retrieve temporal cluster geometry
func (c *Client) TemporalClusterGeometryByIDSource(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewTemporalClusterGeometryByIDSourceRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewTemporalClusterGeometryByIDSourceRequest create the request corresponding to the temporal cluster geometry by id action endpoint of the source resource.
func (c *Client) NewTemporalClusterGeometryByIDSourceRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// UpdateSourcePath computes a request path to the update action of source.
func UpdateSourcePath(sourceID int) string {
	param0 := strconv.Itoa(sourceID)

	return fmt.Sprintf("/sources/%s", param0)
}

// Update an source
func (c *Client) UpdateSource(ctx context.Context, path string, payload *UpdateSourcePayload) (*http.Response, error) {
	req, err := c.NewUpdateSourceRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateSourceRequest create the request corresponding to the update action endpoint of the source resource.
func (c *Client) NewUpdateSourceRequest(ctx context.Context, path string, payload *UpdateSourcePayload) (*http.Request, error) {
	var body bytes.Buffer
	err := c.Encoder.Encode(payload, &body, "*/*")
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("PATCH", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}
