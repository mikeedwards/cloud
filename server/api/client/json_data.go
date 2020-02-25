// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": jsonData Resource Client
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

// GetJSONDataPath computes a request path to the get action of jsonData.
func GetJSONDataPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/data/devices/%s/data/json", param0)
}

// Retrieve data
func (c *Client) GetJSONData(ctx context.Context, path string, end *int, internal *bool, page *int, pageSize *int, start *int) (*http.Response, error) {
	req, err := c.NewGetJSONDataRequest(ctx, path, end, internal, page, pageSize, start)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetJSONDataRequest create the request corresponding to the get action endpoint of the jsonData resource.
func (c *Client) NewGetJSONDataRequest(ctx context.Context, path string, end *int, internal *bool, page *int, pageSize *int, start *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if end != nil {
		tmp255 := strconv.Itoa(*end)
		values.Set("end", tmp255)
	}
	if internal != nil {
		tmp256 := strconv.FormatBool(*internal)
		values.Set("internal", tmp256)
	}
	if page != nil {
		tmp257 := strconv.Itoa(*page)
		values.Set("page", tmp257)
	}
	if pageSize != nil {
		tmp258 := strconv.Itoa(*pageSize)
		values.Set("pageSize", tmp258)
	}
	if start != nil {
		tmp259 := strconv.Itoa(*start)
		values.Set("start", tmp259)
	}
	u.RawQuery = values.Encode()
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

// GetLinesJSONDataPath computes a request path to the get lines action of jsonData.
func GetLinesJSONDataPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/data/devices/%s/data/lines", param0)
}

// Retrieve data
func (c *Client) GetLinesJSONData(ctx context.Context, path string, end *int, internal *bool, page *int, pageSize *int, start *int) (*http.Response, error) {
	req, err := c.NewGetLinesJSONDataRequest(ctx, path, end, internal, page, pageSize, start)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetLinesJSONDataRequest create the request corresponding to the get lines action endpoint of the jsonData resource.
func (c *Client) NewGetLinesJSONDataRequest(ctx context.Context, path string, end *int, internal *bool, page *int, pageSize *int, start *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if end != nil {
		tmp260 := strconv.Itoa(*end)
		values.Set("end", tmp260)
	}
	if internal != nil {
		tmp261 := strconv.FormatBool(*internal)
		values.Set("internal", tmp261)
	}
	if page != nil {
		tmp262 := strconv.Itoa(*page)
		values.Set("page", tmp262)
	}
	if pageSize != nil {
		tmp263 := strconv.Itoa(*pageSize)
		values.Set("pageSize", tmp263)
	}
	if start != nil {
		tmp264 := strconv.Itoa(*start)
		values.Set("start", tmp264)
	}
	u.RawQuery = values.Encode()
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

// SummaryJSONDataPath computes a request path to the summary action of jsonData.
func SummaryJSONDataPath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/data/devices/%s/summary/json", param0)
}

// Retrieve summarized data
func (c *Client) SummaryJSONData(ctx context.Context, path string, end *int, internal *bool, interval *int, page *int, pageSize *int, resolution *int, start *int) (*http.Response, error) {
	req, err := c.NewSummaryJSONDataRequest(ctx, path, end, internal, interval, page, pageSize, resolution, start)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewSummaryJSONDataRequest create the request corresponding to the summary action endpoint of the jsonData resource.
func (c *Client) NewSummaryJSONDataRequest(ctx context.Context, path string, end *int, internal *bool, interval *int, page *int, pageSize *int, resolution *int, start *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if end != nil {
		tmp265 := strconv.Itoa(*end)
		values.Set("end", tmp265)
	}
	if internal != nil {
		tmp266 := strconv.FormatBool(*internal)
		values.Set("internal", tmp266)
	}
	if interval != nil {
		tmp267 := strconv.Itoa(*interval)
		values.Set("interval", tmp267)
	}
	if page != nil {
		tmp268 := strconv.Itoa(*page)
		values.Set("page", tmp268)
	}
	if pageSize != nil {
		tmp269 := strconv.Itoa(*pageSize)
		values.Set("pageSize", tmp269)
	}
	if resolution != nil {
		tmp270 := strconv.Itoa(*resolution)
		values.Set("resolution", tmp270)
	}
	if start != nil {
		tmp271 := strconv.Itoa(*start)
		values.Set("start", tmp271)
	}
	u.RawQuery = values.Encode()
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
