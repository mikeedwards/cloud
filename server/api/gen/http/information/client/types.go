// Code generated by goa v3.1.2, DO NOT EDIT.
//
// information HTTP client types
//
// Command:
// $ goa gen github.com/fieldkit/cloud/server/api/design

package client

import (
	information "github.com/fieldkit/cloud/server/api/gen/information"
	informationviews "github.com/fieldkit/cloud/server/api/gen/information/views"
	goa "goa.design/goa/v3/pkg"
)

// DeviceLayoutResponseBody is the type of the "information" service "device
// layout" endpoint HTTP response body.
type DeviceLayoutResponseBody struct {
	Configurations []*StationConfigurationResponseBody     `form:"configurations,omitempty" json:"configurations,omitempty" xml:"configurations,omitempty"`
	Sensors        map[string][]*StationSensorResponseBody `form:"sensors,omitempty" json:"sensors,omitempty" xml:"sensors,omitempty"`
}

// DeviceLayoutBadRequestResponseBody is the type of the "information" service
// "device layout" endpoint HTTP response body for the "bad-request" error.
type DeviceLayoutBadRequestResponseBody string

// DeviceLayoutForbiddenResponseBody is the type of the "information" service
// "device layout" endpoint HTTP response body for the "forbidden" error.
type DeviceLayoutForbiddenResponseBody string

// DeviceLayoutNotFoundResponseBody is the type of the "information" service
// "device layout" endpoint HTTP response body for the "not-found" error.
type DeviceLayoutNotFoundResponseBody string

// DeviceLayoutUnauthorizedResponseBody is the type of the "information"
// service "device layout" endpoint HTTP response body for the "unauthorized"
// error.
type DeviceLayoutUnauthorizedResponseBody string

// FirmwareStatisticsBadRequestResponseBody is the type of the "information"
// service "firmware statistics" endpoint HTTP response body for the
// "bad-request" error.
type FirmwareStatisticsBadRequestResponseBody string

// FirmwareStatisticsForbiddenResponseBody is the type of the "information"
// service "firmware statistics" endpoint HTTP response body for the
// "forbidden" error.
type FirmwareStatisticsForbiddenResponseBody string

// FirmwareStatisticsNotFoundResponseBody is the type of the "information"
// service "firmware statistics" endpoint HTTP response body for the
// "not-found" error.
type FirmwareStatisticsNotFoundResponseBody string

// FirmwareStatisticsUnauthorizedResponseBody is the type of the "information"
// service "firmware statistics" endpoint HTTP response body for the
// "unauthorized" error.
type FirmwareStatisticsUnauthorizedResponseBody string

// StationConfigurationResponseBody is used to define fields on response body
// types.
type StationConfigurationResponseBody struct {
	ID           *int64                       `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Time         *int64                       `form:"time,omitempty" json:"time,omitempty" xml:"time,omitempty"`
	ProvisionID  *int64                       `form:"provision_id,omitempty" json:"provision_id,omitempty" xml:"provision_id,omitempty"`
	MetaRecordID *int64                       `form:"meta_record_id,omitempty" json:"meta_record_id,omitempty" xml:"meta_record_id,omitempty"`
	SourceID     *int32                       `form:"source_id,omitempty" json:"source_id,omitempty" xml:"source_id,omitempty"`
	Modules      []*StationModuleResponseBody `form:"modules,omitempty" json:"modules,omitempty" xml:"modules,omitempty"`
}

// StationModuleResponseBody is used to define fields on response body types.
type StationModuleResponseBody struct {
	ID           *int64                       `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	HardwareID   *string                      `form:"hardware_id,omitempty" json:"hardware_id,omitempty" xml:"hardware_id,omitempty"`
	MetaRecordID *int64                       `form:"meta_record_id,omitempty" json:"meta_record_id,omitempty" xml:"meta_record_id,omitempty"`
	Name         *string                      `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Position     *int32                       `form:"position,omitempty" json:"position,omitempty" xml:"position,omitempty"`
	Flags        *int32                       `form:"flags,omitempty" json:"flags,omitempty" xml:"flags,omitempty"`
	Internal     *bool                        `form:"internal,omitempty" json:"internal,omitempty" xml:"internal,omitempty"`
	Sensors      []*StationSensorResponseBody `form:"sensors,omitempty" json:"sensors,omitempty" xml:"sensors,omitempty"`
}

// StationSensorResponseBody is used to define fields on response body types.
type StationSensorResponseBody struct {
	Name          *string                    `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	UnitOfMeasure *string                    `form:"unit_of_measure,omitempty" json:"unit_of_measure,omitempty" xml:"unit_of_measure,omitempty"`
	Reading       *SensorReadingResponseBody `form:"reading,omitempty" json:"reading,omitempty" xml:"reading,omitempty"`
	Key           *string                    `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
	Ranges        []*SensorRangeResponseBody `form:"ranges,omitempty" json:"ranges,omitempty" xml:"ranges,omitempty"`
}

// SensorReadingResponseBody is used to define fields on response body types.
type SensorReadingResponseBody struct {
	Last *float32 `form:"last,omitempty" json:"last,omitempty" xml:"last,omitempty"`
	Time *int64   `form:"time,omitempty" json:"time,omitempty" xml:"time,omitempty"`
}

// SensorRangeResponseBody is used to define fields on response body types.
type SensorRangeResponseBody struct {
	Minimum *float32 `form:"minimum,omitempty" json:"minimum,omitempty" xml:"minimum,omitempty"`
	Maximum *float32 `form:"maximum,omitempty" json:"maximum,omitempty" xml:"maximum,omitempty"`
}

// NewDeviceLayoutResponseViewOK builds a "information" service "device layout"
// endpoint result from a HTTP "OK" response.
func NewDeviceLayoutResponseViewOK(body *DeviceLayoutResponseBody) *informationviews.DeviceLayoutResponseView {
	v := &informationviews.DeviceLayoutResponseView{}
	v.Configurations = make([]*informationviews.StationConfigurationView, len(body.Configurations))
	for i, val := range body.Configurations {
		v.Configurations[i] = unmarshalStationConfigurationResponseBodyToInformationviewsStationConfigurationView(val)
	}
	v.Sensors = make(map[string][]*informationviews.StationSensorView, len(body.Sensors))
	for key, val := range body.Sensors {
		tk := key
		tv := make([]*informationviews.StationSensorView, len(val))
		for i, val := range val {
			tv[i] = unmarshalStationSensorResponseBodyToInformationviewsStationSensorView(val)
		}
		v.Sensors[tk] = tv
	}

	return v
}

// NewDeviceLayoutBadRequest builds a information service device layout
// endpoint bad-request error.
func NewDeviceLayoutBadRequest(body DeviceLayoutBadRequestResponseBody) information.BadRequest {
	v := information.BadRequest(body)
	return v
}

// NewDeviceLayoutForbidden builds a information service device layout endpoint
// forbidden error.
func NewDeviceLayoutForbidden(body DeviceLayoutForbiddenResponseBody) information.Forbidden {
	v := information.Forbidden(body)
	return v
}

// NewDeviceLayoutNotFound builds a information service device layout endpoint
// not-found error.
func NewDeviceLayoutNotFound(body DeviceLayoutNotFoundResponseBody) information.NotFound {
	v := information.NotFound(body)
	return v
}

// NewDeviceLayoutUnauthorized builds a information service device layout
// endpoint unauthorized error.
func NewDeviceLayoutUnauthorized(body DeviceLayoutUnauthorizedResponseBody) information.Unauthorized {
	v := information.Unauthorized(body)
	return v
}

// NewFirmwareStatisticsResultOK builds a "information" service "firmware
// statistics" endpoint result from a HTTP "OK" response.
func NewFirmwareStatisticsResultOK(body interface{}) *information.FirmwareStatisticsResult {
	v := body
	res := &information.FirmwareStatisticsResult{
		Object: v,
	}

	return res
}

// NewFirmwareStatisticsBadRequest builds a information service firmware
// statistics endpoint bad-request error.
func NewFirmwareStatisticsBadRequest(body FirmwareStatisticsBadRequestResponseBody) information.BadRequest {
	v := information.BadRequest(body)
	return v
}

// NewFirmwareStatisticsForbidden builds a information service firmware
// statistics endpoint forbidden error.
func NewFirmwareStatisticsForbidden(body FirmwareStatisticsForbiddenResponseBody) information.Forbidden {
	v := information.Forbidden(body)
	return v
}

// NewFirmwareStatisticsNotFound builds a information service firmware
// statistics endpoint not-found error.
func NewFirmwareStatisticsNotFound(body FirmwareStatisticsNotFoundResponseBody) information.NotFound {
	v := information.NotFound(body)
	return v
}

// NewFirmwareStatisticsUnauthorized builds a information service firmware
// statistics endpoint unauthorized error.
func NewFirmwareStatisticsUnauthorized(body FirmwareStatisticsUnauthorizedResponseBody) information.Unauthorized {
	v := information.Unauthorized(body)
	return v
}

// ValidateStationConfigurationResponseBody runs the validations defined on
// StationConfigurationResponseBody
func ValidateStationConfigurationResponseBody(body *StationConfigurationResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.ProvisionID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("provision_id", "body"))
	}
	if body.Time == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("time", "body"))
	}
	if body.Modules == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("modules", "body"))
	}
	for _, e := range body.Modules {
		if e != nil {
			if err2 := ValidateStationModuleResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStationModuleResponseBody runs the validations defined on
// StationModuleResponseBody
func ValidateStationModuleResponseBody(body *StationModuleResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Position == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("position", "body"))
	}
	if body.Flags == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("flags", "body"))
	}
	if body.Internal == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("internal", "body"))
	}
	if body.Sensors == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("sensors", "body"))
	}
	for _, e := range body.Sensors {
		if e != nil {
			if err2 := ValidateStationSensorResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateStationSensorResponseBody runs the validations defined on
// StationSensorResponseBody
func ValidateStationSensorResponseBody(body *StationSensorResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.UnitOfMeasure == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("unit_of_measure", "body"))
	}
	if body.Key == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("key", "body"))
	}
	if body.Ranges == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("ranges", "body"))
	}
	if body.Reading != nil {
		if err2 := ValidateSensorReadingResponseBody(body.Reading); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	for _, e := range body.Ranges {
		if e != nil {
			if err2 := ValidateSensorRangeResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateSensorReadingResponseBody runs the validations defined on
// SensorReadingResponseBody
func ValidateSensorReadingResponseBody(body *SensorReadingResponseBody) (err error) {
	if body.Last == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("last", "body"))
	}
	if body.Time == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("time", "body"))
	}
	return
}

// ValidateSensorRangeResponseBody runs the validations defined on
// SensorRangeResponseBody
func ValidateSensorRangeResponseBody(body *SensorRangeResponseBody) (err error) {
	if body.Minimum == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("minimum", "body"))
	}
	if body.Maximum == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("maximum", "body"))
	}
	return
}
