// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": Query Resource Client
//
// Command:
// $ main

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// ListBySourceQueryPath computes a request path to the list by source action of Query.
func ListBySourceQueryPath(sourceID int) string {
	param0 := strconv.Itoa(sourceID)

	return fmt.Sprintf("/sources/%s/query", param0)
}

// Query data for an source.
func (c *Client) ListBySourceQuery(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListBySourceQueryRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListBySourceQueryRequest create the request corresponding to the list by source action endpoint of the Query resource.
func (c *Client) NewListBySourceQueryRequest(ctx context.Context, path string) (*http.Request, error) {
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
