// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": twitter Resource Client
//
// Command:
// $ main

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CallbackTwitterPath computes a request path to the callback action of twitter.
func CallbackTwitterPath() string {

	return fmt.Sprintf("/twitter/callback")
}

// OAuth callback endpoint for Twitter
func (c *Client) CallbackTwitter(ctx context.Context, path string, oauthToken string, oauthVerifier *string) (*http.Response, error) {
	req, err := c.NewCallbackTwitterRequest(ctx, path, oauthToken, oauthVerifier)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCallbackTwitterRequest create the request corresponding to the callback action endpoint of the twitter resource.
func (c *Client) NewCallbackTwitterRequest(ctx context.Context, path string, oauthToken string, oauthVerifier *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("oauthToken", oauthToken)
	if oauthVerifier != nil {
		values.Set("oauthVerifier", *oauthVerifier)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
