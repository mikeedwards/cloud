package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"sort"
	"strings"
	"time"

	"github.com/lib/pq"

	"github.com/goadesign/goa"

	"github.com/conservify/sqlxcache"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	pb "github.com/fieldkit/data-protocol"

	"github.com/fieldkit/cloud/server/api/app"
	"github.com/fieldkit/cloud/server/backend"
	"github.com/fieldkit/cloud/server/backend/ingestion"
	"github.com/fieldkit/cloud/server/data"
)

type FilesControllerOptions struct {
	Session  *session.Session
	Database *sqlxcache.DB
	Backend  *backend.Backend
}

type FilesController struct {
	*goa.Controller
	options FilesControllerOptions
}

var (
	DataFileTypeIDs = []string{"4"}
	LogFileTypeIDs  = []string{"2", "3"}
)

func NewFilesController(service *goa.Service, options FilesControllerOptions) *FilesController {
	return &FilesController{
		Controller: service.NewController("FilesController"),
		options:    options,
	}
}

func DeviceFileSummaryType(s *data.DeviceFile) *app.DeviceFile {
	return &app.DeviceFile{
		DeviceID:   s.DeviceID,
		FileID:     s.StreamID,
		Firmware:   s.Firmware,
		ID:         int(s.ID),
		Meta:       s.Meta.String(),
		Size:       int(s.Size),
		FileTypeID: s.FileID,
		Time:       s.Time,
		URL:        s.URL,
		Urls: &app.DeviceFileUrls{
			Csv:  fmt.Sprintf("/files/%v/csv", s.ID),
			Raw:  fmt.Sprintf("/files/%v/raw", s.ID),
			JSON: fmt.Sprintf("/files/%v/json", s.ID),
		},
	}
}

func DeviceFilesType(files []*data.DeviceFile) *app.DeviceFiles {
	summaries := make([]*app.DeviceFile, len(files))
	for i, summary := range files {
		summaries[i] = DeviceFileSummaryType(summary)
	}
	return &app.DeviceFiles{
		Files: summaries,
	}
}

func (c *FilesController) listDeviceFiles(ctx context.Context, fileTypeIDs []string, deviceID string, page *int) (*app.DeviceFiles, error) {
	log := Logger(ctx).Sugar()

	log.Infow("Device", "device_id", deviceID, "file_type_ids", fileTypeIDs)

	pageSize := 100
	offset := 0
	if page != nil {
		offset = *page * pageSize
	}

	files := []*data.DeviceFile{}
	if err := c.options.Database.SelectContext(ctx, &files,
		`SELECT s.* FROM fieldkit.device_stream AS s WHERE (s.file_id = ANY($1)) AND (s.device_id = $2) ORDER BY time DESC LIMIT $3 OFFSET $4`, pq.StringArray(fileTypeIDs), deviceID, pageSize, offset); err != nil {
		return nil, err
	}

	return DeviceFilesType(files), nil
}

func (c *FilesController) ListDeviceDataFiles(ctx *app.ListDeviceDataFilesFilesContext) error {
	files, err := c.listDeviceFiles(ctx, DataFileTypeIDs, ctx.DeviceID, ctx.Page)
	if err != nil {
		return err
	}

	return ctx.OK(files)
}

func (c *FilesController) ListDeviceLogFiles(ctx *app.ListDeviceLogFilesFilesContext) error {
	files, err := c.listDeviceFiles(ctx, LogFileTypeIDs, ctx.DeviceID, ctx.Page)
	if err != nil {
		return err
	}

	return ctx.OK(files)
}

type Device struct {
	DeviceID      string    `db:"device_id"`
	LastFileID    string    `db:"last_stream_id"`
	LastFileTime  time.Time `db:"last_stream_time"`
	NumberOfFiles int       `db:"number_of_files"`
}

func DeviceSummaryType(s *Device) *app.Device {
	return &app.Device{
		DeviceID:      s.DeviceID,
		LastFileID:    s.LastFileID,
		LastFileTime:  s.LastFileTime,
		NumberOfFiles: s.NumberOfFiles,
	}
}

func DevicesType(devices []*Device) *app.Devices {
	summaries := make([]*app.Device, len(devices))
	for i, summary := range devices {
		summaries[i] = DeviceSummaryType(summary)
	}
	return &app.Devices{
		Devices: summaries,
	}
}

func (c *FilesController) ListDevices(ctx *app.ListDevicesFilesContext) error {
	devices := []*Device{}
	if err := c.options.Database.SelectContext(ctx, &devices,
		`SELECT s.device_id,
		    (SELECT stream_id FROM fieldkit.device_stream AS s2 WHERE (s2.device_id = s.device_id) ORDER BY s2.time DESC LIMIT 1) AS last_stream_id,
		    MAX(s.time) AS last_stream_time,
		    COUNT(s.*) AS number_of_files
		 FROM fieldkit.device_stream AS s
                 GROUP BY s.device_id
                 ORDER BY last_stream_time DESC`); err != nil {
		return err
	}

	return ctx.OK(DevicesType(devices))
}

func (c *FilesController) Csv(ctx *app.CsvFilesContext) error {
	iterator, err := c.LookupFile(ctx, ctx.FileID)
	if err != nil {
		return err
	}

	exporter := NewSimpleCsvExporter(ctx.ResponseData)

	return ExportAllFiles(ctx, ctx.ResponseData, iterator, exporter)
}

func (c *FilesController) JSON(ctx *app.JSONFilesContext) error {
	iterator, err := c.LookupFile(ctx, ctx.FileID)
	if err != nil {
		return err
	}

	exporter := NewSimpleJsonExporter(ctx.ResponseData)

	return ExportAllFiles(ctx, ctx.ResponseData, iterator, exporter)
}

func (c *FilesController) Raw(ctx *app.RawFilesContext) error {
	iterator, err := c.LookupFile(ctx, ctx.FileID)
	if err != nil {
		return err
	}

	cs, err := iterator.Next(ctx)
	if err != nil {
		return err
	}

	fileName := fmt.Sprintf("%s.csv", cs.File.StreamID)

	ctx.ResponseData.Header().Set("Content-Type", backend.FkDataBinaryContentType)
	ctx.ResponseData.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", fileName))
	ctx.ResponseData.WriteHeader(http.StatusOK)

	_, err = io.Copy(ctx.ResponseData, cs.Response.Body)
	if err != nil {
		return err
	}

	return nil
}

/*
func (c *FilesController) DeviceData(ctx *app.DeviceDataFilesContext) error {
	iterator, err := c.LookupDeviceFiles(ctx, ctx.DeviceID, DataFileTypeID)
	if err != nil {
		return err
	}

	exporter := NewSimpleCsvExporter(ctx.ResponseData)

	return ExportAllFiles(ctx, iterator, exporter)
}

func (c *FilesController) DeviceLogs(ctx *app.DeviceLogsFilesContext) error {
	iterator, err := c.LookupDeviceFiles(ctx, ctx.DeviceID, LogFileTypeID)
	if err != nil {
		return err
	}

	exporter := NewSimpleCsvExporter(ctx.ResponseData)

	return ExportAllFiles(ctx, iterator, exporter)
}
*/

func ExportAllFiles(ctx context.Context, response *goa.ResponseData, iterator *FileIterator, exporter Exporter) error {
	log := Logger(ctx).Sugar()

	header := false

	for {
		cs, err := iterator.Next(ctx)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if cs == nil {
			break
		}

		if !header {
			fileName := exporter.FileName(cs.File)
			response.Header().Set("Content-Type", exporter.MimeType())
			response.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=\"%s\"", fileName))
			header = true
		}

		defer cs.Response.Body.Close()

		binaryReader := backend.NewFkBinaryReader(exporter.ForFile(cs.File))

		err = binaryReader.Read(ctx, cs.Response.Body)
		if err != nil {
			log.Infow("Error reading stream", "error", err, "file_type_id", cs.File.FileID)
		}
	}

	return nil
}

type Exporter interface {
	ForFile(stream *data.DeviceFile) backend.FormattedMessageReceiver
	FileName(file *data.DeviceFile) string
	MimeType() string
}

type SimpleJsonExporter struct {
	File    *data.DeviceFile
	Writer  io.Writer
	Records int
}

func NewSimpleJsonExporter(writer io.Writer) Exporter {
	return &SimpleJsonExporter{Writer: writer}
}

func (ce *SimpleJsonExporter) MimeType() string {
	return "application/json"
}

func (ce *SimpleJsonExporter) FileName(file *data.DeviceFile) string {
	return fmt.Sprintf("%s.json", file.StreamID)
}

func (ce *SimpleJsonExporter) ForFile(file *data.DeviceFile) backend.FormattedMessageReceiver {
	ce.File = file
	return ce
}

func (ce *SimpleJsonExporter) HandleRecord(ctx context.Context, r *pb.DataRecord) error {
	body, err := json.Marshal(r)
	if err != nil {
		return err
	}

	if ce.Records > 0 {
		fmt.Fprintf(ce.Writer, ",\n")
	} else {
		fmt.Fprintf(ce.Writer, "[\n")
	}

	_, err = ce.Writer.Write(body)
	if err != nil {
		return err
	}

	ce.Records += 1

	return nil
}

func (ce *SimpleJsonExporter) HandleFormattedMessage(ctx context.Context, fm *ingestion.FormattedMessage) (*ingestion.RecordChange, error) {
	return nil, nil
}

func (ce *SimpleJsonExporter) Finish(ctx context.Context) error {
	return nil
}

type SimpleCsvExporter struct {
	File          *data.DeviceFile
	Writer        io.Writer
	HeaderWritten bool
}

func NewSimpleCsvExporter(writer io.Writer) Exporter {
	return &SimpleCsvExporter{Writer: writer}
}

func (ce *SimpleCsvExporter) MimeType() string {
	return "text/csv"
}

func (ce *SimpleCsvExporter) FileName(file *data.DeviceFile) string {
	return fmt.Sprintf("%s.csv", file.StreamID)
}

func (ce *SimpleCsvExporter) ForFile(file *data.DeviceFile) backend.FormattedMessageReceiver {
	ce.File = file
	return ce
}

func (ce *SimpleCsvExporter) HandleRecord(ctx context.Context, r *pb.DataRecord) error {
	if r.Log != nil {
		return ce.ExportLog(ctx, r)
	}

	return nil
}

func (ce *SimpleCsvExporter) ExportLog(ctx context.Context, r *pb.DataRecord) error {
	if !ce.HeaderWritten {
		fmt.Fprintf(ce.Writer, "%v,%v,%v,%v,%v,%v,%v,%v\n", "DeviceID", "FileID", "FileID", "Uptime", "Time", "Level", "Facility", "Message")
		ce.HeaderWritten = true
	}

	fmt.Fprintf(ce.Writer, "%v,%v,%v,%v,%v,%v,%v,%v\n", ce.File.DeviceID, ce.File.FileID, ce.File.ID, r.Log.Uptime, r.Log.Time, r.Log.Level, r.Log.Facility, strings.TrimSpace(r.Log.Message))

	return nil
}

func (ce *SimpleCsvExporter) HandleFormattedMessage(ctx context.Context, fm *ingestion.FormattedMessage) (*ingestion.RecordChange, error) {
	opaqueKeys := reflect.ValueOf(fm.MapValues).MapKeys()
	keys := make([]string, len(opaqueKeys))
	for i := 0; i < len(opaqueKeys); i++ {
		keys[i] = opaqueKeys[i].String()
	}
	sort.Strings(keys)

	if !ce.HeaderWritten {
		fmt.Fprintf(ce.Writer, "%v,%v,%v,%v,%v,%v,%v,%v,", "Device", "File", "File", "Message", "Time", "Longitude", "Latitude", "Fixed")

		for _, key := range keys {
			fmt.Fprintf(ce.Writer, ",%v", key)
		}

		fmt.Fprintf(ce.Writer, "\n")

		ce.HeaderWritten = true
	}

	fmt.Fprintf(ce.Writer, "%v,%v,%v,", ce.File.DeviceID, ce.File.FileID, ce.File.ID)

	fmt.Fprintf(ce.Writer, "%v,%v,", fm.MessageId, fm.Time)

	if fm.Location != nil && len(fm.Location) >= 2 {
		fmt.Fprintf(ce.Writer, "%v,%v", fm.Location[0], fm.Location[1])
	} else {
		fmt.Fprintf(ce.Writer, "%v,%v", 0.0, 0.0)
	}

	fmt.Fprintf(ce.Writer, ",%v", fm.Fixed)

	for _, key := range keys {
		fmt.Fprintf(ce.Writer, ",%v", fm.MapValues[key])
	}

	fmt.Fprintf(ce.Writer, "\n")

	return nil, nil
}

func (ce *SimpleCsvExporter) Finish(ctx context.Context) error {
	return nil
}

type FileIterator struct {
	S3Service *s3.S3
	Offset    int
	Limit     int
	Query     func(s *FileIterator) error
	Files     []*data.DeviceFile
	Index     int
}

type CurrentFile struct {
	File      *data.DeviceFile
	SignedURL string
	Response  *http.Response
}

func (c *FilesController) LookupFile(ctx context.Context, streamID string) (iterator *FileIterator, err error) {
	log := Logger(ctx).Sugar()

	log.Infow("File", "file_id", streamID)

	iterator = &FileIterator{
		Offset:    0,
		Limit:     0,
		Index:     0,
		S3Service: s3.New(c.options.Session),
		Query: func(s *FileIterator) error {
			s.Files = []*data.DeviceFile{}
			if err := c.options.Database.SelectContext(ctx, &s.Files, `SELECT s.* FROM fieldkit.device_stream AS s WHERE s.id = $1`, streamID); err != nil {
				return err
			}
			return nil
		},
	}

	return
}

func (c *FilesController) LookupDeviceFiles(ctx context.Context, deviceID string, fileTypeIDs []string) (iterator *FileIterator, err error) {
	log := Logger(ctx).Sugar()

	log.Infow("File", "device_id", deviceID, "file_type_ids", fileTypeIDs)

	iterator = &FileIterator{
		Offset:    0,
		Limit:     10,
		Index:     0,
		S3Service: s3.New(c.options.Session),
		Query: func(s *FileIterator) error {
			s.Files = []*data.DeviceFile{}
			if err := c.options.Database.SelectContext(ctx, &s.Files,
				`SELECT s.* FROM fieldkit.device_stream AS s WHERE (s.file_id = ANY($1)) AND (s.device_id = $2) ORDER BY time DESC LIMIT $3 OFFSET $4`,
				fileTypeIDs, deviceID, s.Limit, s.Offset); err != nil {
				return err
			}
			return nil
		},
	}

	return
}

func (iterator *FileIterator) Next(ctx context.Context) (cs *CurrentFile, err error) {
	log := Logger(ctx).Sugar()

	if iterator.Files != nil && iterator.Index >= len(iterator.Files) {
		iterator.Files = nil

		// Easy way to handle the single stream per batch.
		if iterator.Limit == 0 {
			log.Infow("No more batches")
			return nil, io.EOF
		}
	}

	if iterator.Files == nil {
		err = iterator.Query(iterator)
		if err != nil {
			return nil, err
		}

		log.Infow("Queried", "batch_size", len(iterator.Files))

		if len(iterator.Files) == 0 {
			log.Infow("No more batches")
			return nil, io.EOF
		}

		iterator.Index = 0
	}

	stream := iterator.Files[iterator.Index]

	log.Infow("File", "file_type_id", stream.FileID, "index", iterator.Index, "size", stream.Size, "url", stream.URL)

	iterator.Index += 1

	signed, err := SignS3URL(iterator.S3Service, stream.URL)
	if err != nil {
		return nil, fmt.Errorf("Error signing stream URL: %v (%v)", stream.URL, err)
	}

	log.Infow("File", "file_type_id", stream.FileID, "signed_url", signed)

	response, err := http.Get(signed)
	if err != nil {
		return nil, fmt.Errorf("Error opening stream URL: %v (%v)", stream.URL, err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("Error retrieving stream: %v (status = %v)", stream.URL, response.StatusCode)
	}

	cs = &CurrentFile{
		File:      stream,
		SignedURL: signed,
		Response:  response,
	}

	return
}
