package api

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"

	"github.com/goadesign/goa"

	"github.com/fieldkit/cloud/server/api/app"
	"github.com/fieldkit/cloud/server/backend/repositories"
	"github.com/fieldkit/cloud/server/data"
)

type JSONDataController struct {
	*goa.Controller
	options *ControllerOptions
}

func NewJSONDataController(ctx context.Context, service *goa.Service, options *ControllerOptions) *JSONDataController {
	return &JSONDataController{
		options:    options,
		Controller: service.NewController("JSONDataController"),
	}
}

func JSONDataRowsType(dm []*repositories.DataRow, includeMetas bool) []*app.JSONDataRow {
	wm := make([]*app.JSONDataRow, len(dm))
	for i, r := range dm {
		var metas []int
		if includeMetas {
			metas = make([]int, len(r.MetaIDs))
			for i, v := range r.MetaIDs {
				metas[i] = int(v)
			}
		}
		wm[i] = &app.JSONDataRow{
			ID:       int64(r.ID),
			Time:     int64(r.Time * 1000),
			Location: r.Location,
			D:        r.D,
			Metas:    metas,
		}
	}
	return wm
}

func JSONDataMetaModuleType(dm []*repositories.DataMetaModule) []*app.JSONDataMetaModule {
	modules := make([]*app.JSONDataMetaModule, len(dm))
	for i, m := range dm {
		sensors := make([]*app.JSONDataMetaSensor, len(m.Sensors))

		for j, s := range m.Sensors {
			sensors[j] = &app.JSONDataMetaSensor{
				Name:  s.Name,
				Key:   s.Key,
				Units: s.Units,
			}
		}

		modules[i] = &app.JSONDataMetaModule{
			Manufacturer: m.Manufacturer,
			Kind:         m.Kind,
			Version:      m.Version,
			ID:           m.ID,
			Name:         m.Name,
			Sensors:      sensors,
		}
	}
	return modules
}

func JSONDataMetaType(dm *repositories.VersionMeta) *app.JSONDataMeta {
	return &app.JSONDataMeta{
		ID: int(dm.ID),
		Station: &app.JSONDataMetaStation{
			ID:      dm.Station.ID,
			Name:    dm.Station.Name,
			Modules: JSONDataMetaModuleType(dm.Station.Modules),
			Firmware: &app.JSONDataMetaStationFirmware{
				Version:   dm.Station.Firmware.Version,
				Build:     dm.Station.Firmware.Build,
				Number:    dm.Station.Firmware.Number,
				Timestamp: int(dm.Station.Firmware.Timestamp),
				Hash:      dm.Station.Firmware.Hash,
			},
		},
	}
}

func JSONDataResponseType(dm []*repositories.Version) []*app.JSONDataVersion {
	wm := make([]*app.JSONDataVersion, len(dm))
	for i, r := range dm {
		wm[i] = &app.JSONDataVersion{
			Meta: JSONDataMetaType(r.Meta),
			Data: JSONDataRowsType(r.Data, false),
		}
	}
	return wm
}

func JSONDataSummaryResponseType(dm *repositories.ModulesAndData) *app.JSONDataSummaryResponse {
	return &app.JSONDataSummaryResponse{
		Modules: JSONDataMetaModuleType(dm.Modules),
		Data:    JSONDataRowsType(dm.Data, true),
	}
}

func (c *JSONDataController) Summary(ctx *app.SummaryJSONDataContext) error {
	log := Logger(ctx).Sugar()

	log.Infow("summarize", "device_id", ctx.DeviceID)

	r, err := repositories.NewDataRepository(c.options.Database)
	if err != nil {
		return err
	}

	opts := &repositories.SummaryQueryOpts{
		DeviceID:   ctx.DeviceID,
		Page:       0,
		PageSize:   1000,
		Resolution: 1000,
		Internal:   false,
		Start:      0,
		End:        0,
	}

	if ctx.Page != nil {
		opts.Page = *ctx.Page
	}
	if ctx.PageSize != nil {
		opts.PageSize = *ctx.PageSize
	}
	if ctx.Start != nil {
		opts.Start = int64(*ctx.Start)
	}
	if ctx.End != nil {
		opts.End = int64(*ctx.End)
	}
	if ctx.Resolution != nil {
		opts.Resolution = *ctx.Resolution
	}
	if ctx.Internal != nil {
		opts.Internal = *ctx.Internal
	}

	modulesAndData, err := r.QueryDeviceModulesAndData(ctx, opts)
	if err != nil {
		return err
	}

	return ctx.OK(JSONDataSummaryResponseType(modulesAndData))
}

func (c *JSONDataController) Get(ctx *app.GetJSONDataContext) error {
	log := Logger(ctx).Sugar()

	log.Infow("json", "device_id", ctx.DeviceID)

	p, err := NewPermissions(ctx)
	if err != nil {
		return err
	}

	deviceIdBytes, err := data.DecodeBinaryString(ctx.DeviceID)
	if err != nil {
		return err
	}

	err = p.CanViewStationByDeviceID(deviceIdBytes)
	if err != nil {
		return err
	}

	pageNumber := 0
	if ctx.Page != nil {
		pageNumber = *ctx.Page
	}

	pageSize := 200
	if ctx.PageSize != nil {
		pageSize = *ctx.PageSize
	}

	internal := false
	if ctx.Internal != nil {
		internal = *ctx.Internal
	}

	vr, err := repositories.NewVersionRepository(c.options.Database)
	if err != nil {
		return err
	}

	versions, err := vr.QueryDevice(ctx, ctx.DeviceID, deviceIdBytes, internal, pageNumber, pageSize)
	if err != nil {
		return err
	}

	return ctx.OK(&app.JSONDataResponse{
		Versions: JSONDataResponseType(versions),
	})
}

type JSONLine struct {
	ID       int64                  `form:"id" json:"id" yaml:"id" xml:"id"`
	Time     int64                  `form:"time" json:"time" yaml:"time" xml:"time"`
	Meta     int64                  `form:"meta" json:"meta" yaml:"meta" xml:"meta"`
	Location []float64              `form:"location" json:"location" yaml:"location" xml:"location"`
	D        map[string]interface{} `form:"d" json:"d" yaml:"d" xml:"d"`
}

func (c *JSONDataController) GetLines(ctx *app.GetLinesJSONDataContext) error {
	log := Logger(ctx).Sugar()

	_ = log

	p, err := NewPermissions(ctx)
	if err != nil {
		return err
	}

	deviceIdBytes, err := data.DecodeBinaryString(ctx.DeviceID)
	if err != nil {
		return err
	}

	err = p.CanViewStationByDeviceID(deviceIdBytes)
	if err != nil {
		return err
	}

	pageNumber := 0
	if ctx.Page != nil {
		pageNumber = *ctx.Page
	}

	pageSize := 200
	if ctx.PageSize != nil {
		pageSize = *ctx.PageSize
	}

	internal := false
	if ctx.Internal != nil {
		internal = *ctx.Internal
	}

	vr, err := repositories.NewVersionRepository(c.options.Database)
	if err != nil {
		return err
	}

	versions, err := vr.QueryDevice(ctx, ctx.DeviceID, deviceIdBytes, internal, pageNumber, pageSize)
	if err != nil {
		return err
	}

	ctx.ResponseData.Header().Set("Content-Type", "text/plain")
	ctx.ResponseData.WriteHeader(200)

	writer := bufio.NewWriter(ctx.ResponseData)

	for _, version := range versions {
		for _, row := range version.Data {
			line := JSONLine{
				Time:     row.Time * 1000,
				Location: row.Location,
				ID:       row.ID,
				Meta:     version.Meta.ID,
				D:        row.D,
			}
			bytes, err := json.Marshal(line)
			if err != nil {
				return err
			}
			writer.WriteString(fmt.Sprintf("%s\n", string(bytes)))
		}
	}

	writer.Flush()

	return nil
}