// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "fieldkit": project Resource Client
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

// AddProjectPath computes a request path to the add action of project.
func AddProjectPath() string {

	return fmt.Sprintf("/projects")
}

// Add a project
func (c *Client) AddProject(ctx context.Context, path string, payload *AddProjectPayload) (*http.Response, error) {
	req, err := c.NewAddProjectRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddProjectRequest create the request corresponding to the add action endpoint of the project resource.
func (c *Client) NewAddProjectRequest(ctx context.Context, path string, payload *AddProjectPayload) (*http.Request, error) {
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
	if c.JWTSigner != nil {
		if err := c.JWTSigner.Sign(req); err != nil {
			return nil, err
		}
	}
	return req, nil
}

// GetProjectPath computes a request path to the get action of project.
func GetProjectPath(project string) string {
	param0 := project

	return fmt.Sprintf("/projects/@/%s", param0)
}

// Get a project
func (c *Client) GetProject(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetProjectRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetProjectRequest create the request corresponding to the get action endpoint of the project resource.
func (c *Client) NewGetProjectRequest(ctx context.Context, path string) (*http.Request, error) {
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

// GetIDProjectPath computes a request path to the get id action of project.
func GetIDProjectPath(projectID int) string {
	param0 := strconv.Itoa(projectID)

	return fmt.Sprintf("/projects/%s", param0)
}

// Get a project
func (c *Client) GetIDProject(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetIDProjectRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetIDProjectRequest create the request corresponding to the get id action endpoint of the project resource.
func (c *Client) NewGetIDProjectRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListProjectPath computes a request path to the list action of project.
func ListProjectPath() string {

	return fmt.Sprintf("/projects")
}

// List projects
func (c *Client) ListProject(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListProjectRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListProjectRequest create the request corresponding to the list action endpoint of the project resource.
func (c *Client) NewListProjectRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListCurrentProjectPath computes a request path to the list current action of project.
func ListCurrentProjectPath() string {

	return fmt.Sprintf("/user/projects")
}

// List the authenticated user's projects
func (c *Client) ListCurrentProject(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListCurrentProjectRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListCurrentProjectRequest create the request corresponding to the list current action endpoint of the project resource.
func (c *Client) NewListCurrentProjectRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateProjectPath computes a request path to the update action of project.
func UpdateProjectPath(projectID int) string {
	param0 := strconv.Itoa(projectID)

	return fmt.Sprintf("/projects/%s", param0)
}

// Update a project
func (c *Client) UpdateProject(ctx context.Context, path string, payload *AddProjectPayload) (*http.Response, error) {
	req, err := c.NewUpdateProjectRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateProjectRequest create the request corresponding to the update action endpoint of the project resource.
func (c *Client) NewUpdateProjectRequest(ctx context.Context, path string, payload *AddProjectPayload) (*http.Request, error) {
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
