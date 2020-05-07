package tests

import (
	crand "crypto/rand"
	"crypto/sha1"
	"fmt"
	mrand "math/rand"
	"time"

	"github.com/lib/pq"

	"github.com/bxcodec/faker/v3"
	"golang.org/x/crypto/blake2b"

	"github.com/golang/protobuf/proto"

	pb "github.com/fieldkit/data-protocol"

	"github.com/fieldkit/cloud/server/data"
)

type FakeStations struct {
	Owner    *data.User
	Project  *data.Project
	Stations []*data.Station
}

func (e *TestEnv) AddUser(pw string) (*data.User, error) {
	email := faker.Email()
	user := &data.User{
		Name:     faker.Name(),
		Username: email,
		Email:    email,
		Bio:      faker.Sentence(),
	}

	user.SetPassword(pw)

	if err := e.DB.NamedGetContext(e.Ctx, user, `
		INSERT INTO fieldkit.user (name, username, email, password, bio)
		VALUES (:name, :email, :email, :password, :bio)
		RETURNING *
		`, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (e *TestEnv) AddStations(number int) (*FakeStations, error) {
	owner, err := e.AddUser("passwordpassword")
	if err != nil {
		return nil, err
	}

	name := faker.Name()

	project := &data.Project{
		Name: name + " Project",
		Slug: name,
	}

	if err := e.DB.NamedGetContext(e.Ctx, project, `
		INSERT INTO fieldkit.project (name, slug)
		VALUES (:name, :slug)
		RETURNING *
		`, project); err != nil {
		return nil, err
	}

	if _, err := e.DB.ExecContext(e.Ctx, `
		INSERT INTO fieldkit.project_user (project_id, user_id, role) VALUES ($1, $2, $3)
		`, project.ID, owner.ID, data.AdministratorRole.ID); err != nil {
		return nil, err
	}

	stations := []*data.Station{}

	for i := 0; i < number; i += 1 {
		name := fmt.Sprintf("%s #%d", owner.Name, i)

		hasher := sha1.New()
		hasher.Write([]byte(name))
		deviceID := hasher.Sum(nil)

		station := &data.Station{
			OwnerID:  owner.ID,
			Name:     name,
			DeviceID: deviceID,
		}

		if err := e.DB.NamedGetContext(e.Ctx, station, `
			INSERT INTO fieldkit.station (name, device_id, owner_id, status_json)
			VALUES (:name, :device_id, :owner_id, :status_json)
			RETURNING *
		`, station); err != nil {
			return nil, err
		}

		if _, err := e.DB.ExecContext(e.Ctx, `
			INSERT INTO fieldkit.project_station (project_id, station_id) VALUES ($1, $2)
			`, project.ID, station.ID); err != nil {
			return nil, err
		}

		stations = append(stations, station)
	}

	return &FakeStations{
		Owner:    owner,
		Project:  project,
		Stations: stations,
	}, nil
}

func (e *TestEnv) NewStation(owner *data.User) *data.Station {
	name := faker.Name()

	hasher := sha1.New()
	hasher.Write([]byte(name))
	deviceID := hasher.Sum(nil)

	station := &data.Station{
		OwnerID:  owner.ID,
		DeviceID: deviceID,
		Name:     name,
	}

	return station
}

func (e *TestEnv) AddIngestion(user *data.User, url string, deviceID []byte, length int) (*data.Ingestion, error) {
	ingestion := &data.Ingestion{
		URL:          url,
		UserID:       user.ID,
		DeviceID:     deviceID,
		GenerationID: deviceID,
		Type:         "data",
		Size:         int64(length),
		Blocks:       data.Int64Range([]int64{1, 100}),
		Flags:        pq.Int64Array([]int64{}),
	}

	if err := e.DB.NamedGetContext(e.Ctx, ingestion, `
			INSERT INTO fieldkit.ingestion (time, upload_id, user_id, device_id, generation, type, size, url, blocks, flags)
			VALUES (NOW(), :upload_id, :user_id, :device_id, :generation, :type, :size, :url, :blocks, :flags)
			RETURNING *
			`, ingestion); err != nil {
		return nil, err
	}

	return ingestion, nil
}

func (e *TestEnv) AddStationActivity(station *data.Station, user *data.User) error {
	location := data.NewLocation([]float64{0, 0})

	depoyedActivity := &data.StationDeployed{
		StationActivity: data.StationActivity{
			CreatedAt: time.Now(),
			StationID: int64(station.ID),
		},
		DeployedAt: time.Now(),
		Location:   location,
	}

	if _, err := e.DB.NamedExecContext(e.Ctx, `
		INSERT INTO fieldkit.station_deployed (created_at, station_id, deployed_at, location) VALUES (:created_at, :station_id, :deployed_at, ST_SetSRID(ST_GeomFromText(:location), 4326))
		`, depoyedActivity); err != nil {
		return err
	}

	ingestion, err := e.AddIngestion(user, "file:///dev/null", station.DeviceID, 0)
	if err != nil {
		return err
	}

	activity := &data.StationIngestion{
		StationActivity: data.StationActivity{
			CreatedAt: time.Now(),
			StationID: int64(station.ID),
		},
		UploaderID:      int64(user.ID),
		DataIngestionID: ingestion.ID,
		DataRecords:     1,
		Errors:          false,
	}

	if err := e.DB.NamedGetContext(e.Ctx, activity, `
		INSERT INTO fieldkit.station_ingestion (created_at, station_id, uploader_id, data_ingestion_id, data_records, errors)
		VALUES (:created_at, :station_id, :uploader_id, :data_ingestion_id, :data_records, :errors)
		ON CONFLICT (data_ingestion_id) DO NOTHING
		RETURNING *
		`, activity); err != nil {
		return err
	}

	return nil
}

func (e *TestEnv) AddProjectActivity(project *data.Project, station *data.Station, user *data.User) error {
	if err := e.AddStationActivity(station, user); err != nil {
		return err
	}

	projectUpdate := &data.ProjectUpdate{
		ProjectActivity: data.ProjectActivity{
			CreatedAt: time.Now(),
			ProjectID: int64(project.ID),
		},
		AuthorID: int64(user.ID),
		Body:     "Project update",
	}

	if _, err := e.DB.NamedExecContext(e.Ctx, `
		INSERT INTO fieldkit.project_update (created_at, project_id, author_id, body) VALUES (:created_at, :project_id, :author_id, :body)
		`, projectUpdate); err != nil {
		return err
	}

	return nil
}

func (e *TestEnv) NewRandomData(n int) ([]byte, error) {
	data := make([]byte, n)
	actual, err := crand.Read(data)
	if actual != n {
		return nil, fmt.Errorf("unexpected random byte read")
	}
	return data, err
}

func (e *TestEnv) NewMetaLayout(record uint64) *pb.SignedRecord {
	cfg := &pb.DataRecord{
		Modules: []*pb.ModuleInfo{
			&pb.ModuleInfo{
				Name:     "random-module-1",
				Header:   &pb.ModuleHeader{},
				Firmware: &pb.Firmware{},
				Sensors: []*pb.SensorInfo{
					&pb.SensorInfo{
						Name:          "sensor-0",
						UnitOfMeasure: "C",
					},
					&pb.SensorInfo{
						Name:          "sensor-1",
						UnitOfMeasure: "C",
					},
					&pb.SensorInfo{
						Name:          "sensor-2",
						UnitOfMeasure: "C",
					},
					&pb.SensorInfo{
						Name:          "sensor-3",
						UnitOfMeasure: "C",
					},
					&pb.SensorInfo{
						Name:          "sensor-4",
						UnitOfMeasure: "C",
					},
				},
			},
			&pb.ModuleInfo{
				Name:     "random-module-2",
				Header:   &pb.ModuleHeader{},
				Firmware: &pb.Firmware{},
				Sensors: []*pb.SensorInfo{
					&pb.SensorInfo{
						Name:          "sensor-0",
						UnitOfMeasure: "C",
					},
					&pb.SensorInfo{
						Name:          "sensor-1",
						UnitOfMeasure: "C",
					},
					&pb.SensorInfo{
						Name:          "sensor-2",
						UnitOfMeasure: "C",
					},
					&pb.SensorInfo{
						Name:          "sensor-3",
						UnitOfMeasure: "C",
					},
					&pb.SensorInfo{
						Name:          "sensor-4",
						UnitOfMeasure: "C",
					},
					&pb.SensorInfo{
						Name:          "sensor-5",
						UnitOfMeasure: "C",
					},
					&pb.SensorInfo{
						Name:          "sensor-6",
						UnitOfMeasure: "C",
					},
					&pb.SensorInfo{
						Name:          "sensor-7",
						UnitOfMeasure: "C",
					},
					&pb.SensorInfo{
						Name:          "sensor-8",
						UnitOfMeasure: "C",
					},
					&pb.SensorInfo{
						Name:          "sensor-9",
						UnitOfMeasure: "C",
					},
				},
			},
		},
	}

	body := proto.NewBuffer(make([]byte, 0))
	body.EncodeMessage(cfg)

	hash := blake2b.Sum256(body.Bytes())

	return &pb.SignedRecord{
		Kind:   1, /* Modules */
		Time:   0,
		Data:   body.Bytes(),
		Hash:   hash[:],
		Record: record,
	}
}

func (e *TestEnv) NewDataReading(meta, reading uint64) *pb.DataRecord {
	now := time.Now()

	return &pb.DataRecord{
		Readings: &pb.Readings{
			Time:    int64(now.Unix()),
			Reading: uint32(reading),
			Meta:    uint32(meta),
			Flags:   0,
			Location: &pb.DeviceLocation{
				Fix:        1,
				Time:       int64(now.Unix()),
				Longitude:  -118.2709223,
				Latitude:   34.0318047,
				Altitude:   mrand.Float32(),
				Satellites: 6,
			},
			SensorGroups: []*pb.SensorGroup{
				&pb.SensorGroup{
					Module: 0,
					Readings: []*pb.SensorAndValue{
						&pb.SensorAndValue{
							Sensor: 0,
							Value:  mrand.Float32(),
						},
						&pb.SensorAndValue{
							Sensor: 1,
							Value:  mrand.Float32(),
						},
						&pb.SensorAndValue{
							Sensor: 2,
							Value:  mrand.Float32(),
						},
						&pb.SensorAndValue{
							Sensor: 3,
							Value:  mrand.Float32(),
						},
						&pb.SensorAndValue{
							Sensor: 4,
							Value:  mrand.Float32(),
						},
					},
				},
				&pb.SensorGroup{
					Module: 1,
					Readings: []*pb.SensorAndValue{
						&pb.SensorAndValue{
							Sensor: 0,
							Value:  mrand.Float32(),
						},
						&pb.SensorAndValue{
							Sensor: 1,
							Value:  mrand.Float32(),
						},
						&pb.SensorAndValue{
							Sensor: 2,
							Value:  mrand.Float32(),
						},
						&pb.SensorAndValue{
							Sensor: 3,
							Value:  mrand.Float32(),
						},
						&pb.SensorAndValue{
							Sensor: 4,
							Value:  mrand.Float32(),
						},
						&pb.SensorAndValue{
							Sensor: 5,
							Value:  mrand.Float32(),
						},
						&pb.SensorAndValue{
							Sensor: 6,
							Value:  mrand.Float32(),
						},
						&pb.SensorAndValue{
							Sensor: 7,
							Value:  mrand.Float32(),
						},
						&pb.SensorAndValue{
							Sensor: 8,
							Value:  mrand.Float32(),
						},
						&pb.SensorAndValue{
							Sensor: 9,
							Value:  mrand.Float32(),
						},
					},
				},
			},
		},
	}
}

type FilePair struct {
	Meta       []byte
	Data       []byte
	MetaBlocks []uint64
	DataBlocks []uint64
}

func (e *TestEnv) NewFilePair(nmeta, ndata int) (*FilePair, error) {
	dataFile := proto.NewBuffer(make([]byte, 0))
	metaFile := proto.NewBuffer(make([]byte, 0))

	metaNumber := uint64(1)
	dataNumber := uint64(1)

	for m := 0; m < nmeta; m += 1 {
		meta := e.NewMetaLayout(metaNumber)
		if err := metaFile.EncodeMessage(meta); err != nil {
			return nil, err
		}

		metaNumber += 1

		for i := 0; i < ndata; i += 1 {
			dataRecord := e.NewDataReading(meta.Record, dataNumber)
			if err := dataFile.EncodeMessage(dataRecord); err != nil {
				return nil, err
			}

			dataNumber += 1
		}
	}

	return &FilePair{
		Meta:       metaFile.Bytes(),
		Data:       dataFile.Bytes(),
		MetaBlocks: []uint64{1, metaNumber},
		DataBlocks: []uint64{1, dataNumber},
	}, nil
}

func (e *TestEnv) MustDeviceID() []byte {
	id, err := e.NewRandomData(16)
	if err != nil {
		panic(err)
	}
	return id
}
