// Code generated by goagen v1.2.0, DO NOT EDIT.
//
// API "fieldkit": user Resource Client
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

// AddUserPath computes a request path to the add action of user.
func AddUserPath() string {

	return fmt.Sprintf("/users")
}

// Add a user
func (c *Client) AddUser(ctx context.Context, path string, payload *AddUserPayload) (*http.Response, error) {
	req, err := c.NewAddUserRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddUserRequest create the request corresponding to the add action endpoint of the user resource.
func (c *Client) NewAddUserRequest(ctx context.Context, path string, payload *AddUserPayload) (*http.Request, error) {
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
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}

// GetUserPath computes a request path to the get action of user.
func GetUserPath(username string) string {
	param0 := username

	return fmt.Sprintf("/users/@/%s", param0)
}

// Get a user
func (c *Client) GetUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetUserRequest create the request corresponding to the get action endpoint of the user resource.
func (c *Client) NewGetUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// GetCurrentUserPath computes a request path to the get current action of user.
func GetCurrentUserPath() string {

	return fmt.Sprintf("/user")
}

// Get the authenticated user
func (c *Client) GetCurrentUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetCurrentUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetCurrentUserRequest create the request corresponding to the get current action endpoint of the user resource.
func (c *Client) NewGetCurrentUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// GetIDUserPath computes a request path to the get id action of user.
func GetIDUserPath(userID int) string {
	param0 := strconv.Itoa(userID)

	return fmt.Sprintf("/users/%s", param0)
}

// Get a user
func (c *Client) GetIDUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetIDUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetIDUserRequest create the request corresponding to the get id action endpoint of the user resource.
func (c *Client) NewGetIDUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListUserPath computes a request path to the list action of user.
func ListUserPath() string {

	return fmt.Sprintf("/users")
}

// List users
func (c *Client) ListUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListUserRequest create the request corresponding to the list action endpoint of the user resource.
func (c *Client) NewListUserRequest(ctx context.Context, path string) (*http.Request, error) {
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

// LoginUserPath computes a request path to the login action of user.
func LoginUserPath() string {

	return fmt.Sprintf("/login")
}

// Creates a valid JWT given login credentials.
func (c *Client) LoginUser(ctx context.Context, path string, payload *LoginPayload) (*http.Response, error) {
	req, err := c.NewLoginUserRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLoginUserRequest create the request corresponding to the login action endpoint of the user resource.
func (c *Client) NewLoginUserRequest(ctx context.Context, path string, payload *LoginPayload) (*http.Request, error) {
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
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}

// LogoutUserPath computes a request path to the logout action of user.
func LogoutUserPath() string {

	return fmt.Sprintf("/logout")
}

// Creates a valid JWT given login credentials.
func (c *Client) LogoutUser(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewLogoutUserRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLogoutUserRequest create the request corresponding to the logout action endpoint of the user resource.
func (c *Client) NewLogoutUserRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, err
	}
	if c.JWTSigner != nil {
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// RefreshUserPayload is the user refresh action payload.
type RefreshUserPayload struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token" xml:"refresh_token"`
}

// RefreshUserPath computes a request path to the refresh action of user.
func RefreshUserPath() string {

	return fmt.Sprintf("/refresh")
}

// Creates a valid JWT given a refresh token.
func (c *Client) RefreshUser(ctx context.Context, path string, payload *RefreshUserPayload) (*http.Response, error) {
	req, err := c.NewRefreshUserRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewRefreshUserRequest create the request corresponding to the refresh action endpoint of the user resource.
func (c *Client) NewRefreshUserRequest(ctx context.Context, path string, payload *RefreshUserPayload) (*http.Request, error) {
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
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	header.Set("Content-Type", "application/json")
	return req, nil
}

// UpdateUserPath computes a request path to the update action of user.
func UpdateUserPath(userID int) string {
	param0 := strconv.Itoa(userID)

	return fmt.Sprintf("/users/%s", param0)
}

// Update a user
func (c *Client) UpdateUser(ctx context.Context, path string, payload *UpdateUserPayload) (*http.Response, error) {
	req, err := c.NewUpdateUserRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateUserRequest create the request corresponding to the update action endpoint of the user resource.
func (c *Client) NewUpdateUserRequest(ctx context.Context, path string, payload *UpdateUserPayload) (*http.Request, error) {
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
		c.JWTSigner.Sign(req)
	}
	return req, nil
}

// ValidateUserPath computes a request path to the validate action of user.
func ValidateUserPath() string {

	return fmt.Sprintf("/validate")
}

// Validate a user's email address.
func (c *Client) ValidateUser(ctx context.Context, path string, token string) (*http.Response, error) {
	req, err := c.NewValidateUserRequest(ctx, path, token)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewValidateUserRequest create the request corresponding to the validate action endpoint of the user resource.
func (c *Client) NewValidateUserRequest(ctx context.Context, path string, token string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("token", token)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}