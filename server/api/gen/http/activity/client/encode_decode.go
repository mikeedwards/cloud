// Code generated by goa v3.1.2, DO NOT EDIT.
//
// activity HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	activity "github.com/fieldkit/cloud/server/api/gen/activity"
	activityviews "github.com/fieldkit/cloud/server/api/gen/activity/views"
	goahttp "goa.design/goa/v3/http"
)

// BuildStationRequest instantiates a HTTP request object with method and path
// set to call the "activity" service "station" endpoint
func (c *Client) BuildStationRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id int64
	)
	{
		p, ok := v.(*activity.StationPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("activity", "station", "*activity.StationPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: StationActivityPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("activity", "station", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeStationRequest returns an encoder for requests sent to the activity
// station server.
func EncodeStationRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*activity.StationPayload)
		if !ok {
			return goahttp.ErrInvalidType("activity", "station", "*activity.StationPayload", v)
		}
		if p.Auth != nil {
			head := *p.Auth
			req.Header.Set("Authorization", head)
		}
		values := req.URL.Query()
		if p.Page != nil {
			values.Add("page", fmt.Sprintf("%v", *p.Page))
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeStationResponse returns a decoder for responses returned by the
// activity station endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeStationResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body StationResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("activity", "station", err)
			}
			p := NewStationActivityPageViewOK(&body)
			view := "default"
			vres := &activityviews.StationActivityPage{Projected: p, View: view}
			if err = activityviews.ValidateStationActivityPage(vres); err != nil {
				return nil, goahttp.ErrValidationError("activity", "station", err)
			}
			res := activity.NewStationActivityPage(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("activity", "station", resp.StatusCode, string(body))
		}
	}
}

// BuildProjectRequest instantiates a HTTP request object with method and path
// set to call the "activity" service "project" endpoint
func (c *Client) BuildProjectRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id int64
	)
	{
		p, ok := v.(*activity.ProjectPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("activity", "project", "*activity.ProjectPayload", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: ProjectActivityPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("activity", "project", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeProjectRequest returns an encoder for requests sent to the activity
// project server.
func EncodeProjectRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*activity.ProjectPayload)
		if !ok {
			return goahttp.ErrInvalidType("activity", "project", "*activity.ProjectPayload", v)
		}
		if p.Auth != nil {
			head := *p.Auth
			req.Header.Set("Authorization", head)
		}
		values := req.URL.Query()
		if p.Page != nil {
			values.Add("page", fmt.Sprintf("%v", *p.Page))
		}
		req.URL.RawQuery = values.Encode()
		return nil
	}
}

// DecodeProjectResponse returns a decoder for responses returned by the
// activity project endpoint. restoreBody controls whether the response body
// should be restored after having been read.
func DecodeProjectResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = ioutil.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body ProjectResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("activity", "project", err)
			}
			p := NewProjectActivityPageViewOK(&body)
			view := "default"
			vres := &activityviews.ProjectActivityPage{Projected: p, View: view}
			if err = activityviews.ValidateProjectActivityPage(vres); err != nil {
				return nil, goahttp.ErrValidationError("activity", "project", err)
			}
			res := activity.NewProjectActivityPage(vres)
			return res, nil
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("activity", "project", resp.StatusCode, string(body))
		}
	}
}

// unmarshalActivityEntryResponseBodyToActivityviewsActivityEntryView builds a
// value of type *activityviews.ActivityEntryView from a value of type
// *ActivityEntryResponseBody.
func unmarshalActivityEntryResponseBodyToActivityviewsActivityEntryView(v *ActivityEntryResponseBody) *activityviews.ActivityEntryView {
	res := &activityviews.ActivityEntryView{
		ID:        v.ID,
		Key:       v.Key,
		CreatedAt: v.CreatedAt,
		Type:      v.Type,
		Meta:      v.Meta,
	}
	res.Project = unmarshalProjectSummaryResponseBodyToActivityviewsProjectSummaryView(v.Project)
	res.Station = unmarshalStationSummaryResponseBodyToActivityviewsStationSummaryView(v.Station)

	return res
}

// unmarshalProjectSummaryResponseBodyToActivityviewsProjectSummaryView builds
// a value of type *activityviews.ProjectSummaryView from a value of type
// *ProjectSummaryResponseBody.
func unmarshalProjectSummaryResponseBodyToActivityviewsProjectSummaryView(v *ProjectSummaryResponseBody) *activityviews.ProjectSummaryView {
	res := &activityviews.ProjectSummaryView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}

// unmarshalStationSummaryResponseBodyToActivityviewsStationSummaryView builds
// a value of type *activityviews.StationSummaryView from a value of type
// *StationSummaryResponseBody.
func unmarshalStationSummaryResponseBodyToActivityviewsStationSummaryView(v *StationSummaryResponseBody) *activityviews.StationSummaryView {
	res := &activityviews.StationSummaryView{
		ID:   v.ID,
		Name: v.Name,
	}

	return res
}