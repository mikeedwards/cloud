// Code generated by goagen v1.4.0, DO NOT EDIT.
//
// API "fieldkit": Firmware Resource Client
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

// AddFirmwarePath computes a request path to the add action of Firmware.
func AddFirmwarePath() string {

	return fmt.Sprintf("/firmware")
}

// Add firmware
func (c *Client) AddFirmware(ctx context.Context, path string, payload *AddFirmwarePayload) (*http.Response, error) {
	req, err := c.NewAddFirmwareRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewAddFirmwareRequest create the request corresponding to the add action endpoint of the Firmware resource.
func (c *Client) NewAddFirmwareRequest(ctx context.Context, path string, payload *AddFirmwarePayload) (*http.Request, error) {
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

// CheckFirmwarePath computes a request path to the check action of Firmware.
func CheckFirmwarePath(deviceID string, module string) string {
	param0 := deviceID
	param1 := module

	return fmt.Sprintf("/devices/%s/%s/firmware", param0, param1)
}

// Return firmware for a device
func (c *Client) CheckFirmware(ctx context.Context, path string, fkCompiled *string, ifNoneMatch *string) (*http.Response, error) {
	req, err := c.NewCheckFirmwareRequest(ctx, path, fkCompiled, ifNoneMatch)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCheckFirmwareRequest create the request corresponding to the check action endpoint of the Firmware resource.
func (c *Client) NewCheckFirmwareRequest(ctx context.Context, path string, fkCompiled *string, ifNoneMatch *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if fkCompiled != nil {

		header.Set("Fk-Compiled", *fkCompiled)
	}
	if ifNoneMatch != nil {

		header.Set("If-None-Match", *ifNoneMatch)
	}
	return req, nil
}

// DeleteFirmwarePath computes a request path to the delete action of Firmware.
func DeleteFirmwarePath(firmwareID int) string {
	param0 := strconv.Itoa(firmwareID)

	return fmt.Sprintf("/firmware/%s", param0)
}

// Delete firmware
func (c *Client) DeleteFirmware(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDeleteFirmwareRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDeleteFirmwareRequest create the request corresponding to the delete action endpoint of the Firmware resource.
func (c *Client) NewDeleteFirmwareRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("DELETE", u.String(), nil)
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

// DownloadFirmwarePath computes a request path to the download action of Firmware.
func DownloadFirmwarePath(firmwareID int) string {
	param0 := strconv.Itoa(firmwareID)

	return fmt.Sprintf("/firmware/%s/download", param0)
}

// DownloadFirmware makes a request to the download action endpoint of the Firmware resource
func (c *Client) DownloadFirmware(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewDownloadFirmwareRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewDownloadFirmwareRequest create the request corresponding to the download action endpoint of the Firmware resource.
func (c *Client) NewDownloadFirmwareRequest(ctx context.Context, path string) (*http.Request, error) {
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

// ListFirmwarePath computes a request path to the list action of Firmware.
func ListFirmwarePath() string {

	return fmt.Sprintf("/firmware")
}

// List firmware
func (c *Client) ListFirmware(ctx context.Context, path string, module *string, page *int, pageSize *int, profile *string) (*http.Response, error) {
	req, err := c.NewListFirmwareRequest(ctx, path, module, page, pageSize, profile)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListFirmwareRequest create the request corresponding to the list action endpoint of the Firmware resource.
func (c *Client) NewListFirmwareRequest(ctx context.Context, path string, module *string, page *int, pageSize *int, profile *string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if module != nil {
		values.Set("module", *module)
	}
	if page != nil {
		tmp238 := strconv.Itoa(*page)
		values.Set("page", tmp238)
	}
	if pageSize != nil {
		tmp239 := strconv.Itoa(*pageSize)
		values.Set("pageSize", tmp239)
	}
	if profile != nil {
		values.Set("profile", *profile)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ListDeviceFirmwarePath computes a request path to the list device action of Firmware.
func ListDeviceFirmwarePath(deviceID string) string {
	param0 := deviceID

	return fmt.Sprintf("/devices/%s/firmware", param0)
}

// List device firmware
func (c *Client) ListDeviceFirmware(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewListDeviceFirmwareRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewListDeviceFirmwareRequest create the request corresponding to the list device action endpoint of the Firmware resource.
func (c *Client) NewListDeviceFirmwareRequest(ctx context.Context, path string) (*http.Request, error) {
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

// UpdateFirmwarePath computes a request path to the update action of Firmware.
func UpdateFirmwarePath() string {

	return fmt.Sprintf("/devices/firmware")
}

// Update an Device firmware
func (c *Client) UpdateFirmware(ctx context.Context, path string, payload *UpdateDeviceFirmwarePayload) (*http.Response, error) {
	req, err := c.NewUpdateFirmwareRequest(ctx, path, payload)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewUpdateFirmwareRequest create the request corresponding to the update action endpoint of the Firmware resource.
func (c *Client) NewUpdateFirmwareRequest(ctx context.Context, path string, payload *UpdateDeviceFirmwarePayload) (*http.Request, error) {
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
