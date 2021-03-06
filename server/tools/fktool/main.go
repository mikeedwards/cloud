package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	_ "github.com/google/uuid"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	fk "github.com/fieldkit/cloud/server/api/client"
	fktesting "github.com/fieldkit/cloud/server/tools"
)

type options struct {
	Scheme   string
	Host     string
	Email    string
	Password string

	Module  string
	Profile string

	FirmwareFile string

	DryRun bool
}

type Metadata struct {
	ETag    string
	Module  string
	Profile string
	Map     map[string]*string
}

func getModuleFromJobName(name string) (string, error) {
	return strings.Replace(name, "/", "-", 1), nil
}

func getProfileFromFile(module, path string) (string, error) {
	file := filepath.Base(path)
	re := regexp.MustCompile(fmt.Sprintf("%s-(.+).bin", module))
	m := re.FindAllStringSubmatch(file, -1)
	if len(m) == 0 {
		return "", fmt.Errorf("malformed file name %s (%s), no profile for %s", path, file, module)
	}
	return m[0][1], nil
}

func getMetaFromEnvironment(moduleOverride, profileOverride string, file string) (metadata *Metadata, err error) {
	jobName := os.Getenv("JOB_NAME")
	if jobName == "" {
		return nil, fmt.Errorf("ENV[JOB_NAME] missing.")
	}

	module := moduleOverride
	if moduleOverride == "" {
		module, err = getModuleFromJobName(jobName)
		if err != nil {
			return nil, fmt.Errorf("Error getting module from job name: %v", err)
		}
		log.Printf("found module name: '%s'", module)
	} else {
		log.Printf("using module override: '%s'", module)
	}

	buildTime := os.Getenv("BUILD_TIMESTAMP")
	layout := "20060102_150405"
	_, err = time.Parse(layout, buildTime)
	if err != nil {
		return nil, fmt.Errorf("ENV[BUILD_TIMESTAMP] parse failed.")
	}

	profile := profileOverride
	if profile == "" {
		profile, err = getProfileFromFile(module, file)
		if err != nil {
			return nil, err
		}
	}
	etag := fmt.Sprintf("%s_%s_%s", module, profile, buildTime)

	metadata = &Metadata{
		ETag:    etag,
		Module:  module,
		Profile: profile,
		Map:     make(map[string]*string),
	}

	metadata.Map["Build-Id"] = aws.String(os.Getenv("BUILD_ID"))
	metadata.Map["Build-Number"] = aws.String(os.Getenv("BUILD_NUMBER"))
	metadata.Map["Build-Tag"] = aws.String(os.Getenv("BUILD_TAG"))
	metadata.Map["Build-Time"] = aws.String(buildTime)
	metadata.Map["Build-Commit"] = aws.String(os.Getenv("GIT_COMMIT"))
	metadata.Map["Build-Branch"] = aws.String(os.Getenv("GIT_BRANCH"))
	metadata.Map["Build-Job-Base"] = aws.String(os.Getenv("JOB_BASE_NAME"))
	metadata.Map["Build-Job-Name"] = aws.String(jobName)
	metadata.Map["Build-Module"] = aws.String(module)
	metadata.Map["Build-Profile"] = aws.String(profile)

	return
}

func createAwsSession() (s *session.Session, err error) {
	configs := []aws.Config{
		aws.Config{
			Region:                        aws.String("us-east-1"),
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
		aws.Config{
			Region:                        aws.String("us-east-1"),
			Credentials:                   credentials.NewEnvCredentials(),
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
	}

	for _, config := range configs {
		sessionOptions := session.Options{
			Config: config,
		}
		session, err := session.NewSessionWithOptions(sessionOptions)
		if err == nil {
			return session, nil
		}
	}

	return nil, fmt.Errorf("error creating AWS session: %v", err)
}

func getFileHash(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	hasher := sha256.New()

	if _, err := io.Copy(hasher, f); err != nil {
		return "", err
	}

	h := hex.EncodeToString(hasher.Sum(nil))

	return h, nil
}

func hasFile(session *session.Session, id string) (string, error) {
	bucket := "conservify-firmware"

	hoi := &s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(id),
	}

	svc := s3.New(session)

	_, err := svc.HeadObject(hoi)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case "NotFound": // TODO Make this s3.ErrCodeNoSuchKey eventually.
				return "", nil
			}
		}

		return "", err
	}

	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, id), nil
}

func uploadFirmware(ctx context.Context, c *fk.Client, moduleOverride, profileOverride, filename string, dryRun bool) error {
	id, err := getFileHash(filename)
	if err != nil {
		return fmt.Errorf("error getting file hash: %v", filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening file: %v", filename)
	}

	defer file.Close()

	metadata, err := getMetaFromEnvironment(moduleOverride, profileOverride, filename)
	if err != nil {
		return err
	}

	session, err := createAwsSession()
	if err != nil {
		return err
	}

	uploader := s3manager.NewUploader(session)

	url := ""

	existingUrl, err := hasFile(session, id)
	if err != nil {
		return err
	}

	if existingUrl != "" {
		log.Printf("object already uploaded")
		url = existingUrl
	} else {
		log.Printf("uploading %s...", filename)

		if !dryRun {
			r, err := uploader.Upload(&s3manager.UploadInput{
				ACL:         nil,
				ContentType: aws.String("application/octet-stream"),
				Bucket:      aws.String("conservify-firmware"),
				Key:         aws.String(id),
				Body:        file,
				Metadata:    metadata.Map,
				Tagging:     nil,
			})
			if err != nil {
				return fmt.Errorf("error uploading firmware: %v", err)
			}

			log.Printf("uploaded %s", r.Location)

			url = r.Location
		}
	}

	log.Printf("creating database entry...")

	jsonData, err := json.Marshal(metadata.Map)
	if err != nil {
		return fmt.Errorf("error serializing metadata: %v", err)
	}

	addFirmwarePayload := fk.AddFirmwarePayload{
		Etag:    metadata.ETag,
		Module:  metadata.Module,
		Profile: metadata.Profile,
		URL:     url,
		Meta:    string(jsonData),
	}

	if !dryRun {
		r, err := c.AddFirmware(ctx, fk.AddFirmwarePath(), &addFirmwarePayload)
		if err != nil {
			return fmt.Errorf("error adding firmware: %v", err)
		}

		log.Printf("added: %v", r)
	}

	log.Printf("done!")

	return nil
}

func main() {
	ctx := context.TODO()

	o := options{}

	flag.StringVar(&o.Scheme, "scheme", "http", "fk instance scheme")
	flag.StringVar(&o.Host, "host", "127.0.0.1:8080", "fk instance hostname")
	flag.StringVar(&o.Email, "email", "info@conservify.org", "email")
	flag.StringVar(&o.Password, "password", "asdfasdfasdf", "password")
	flag.StringVar(&o.Module, "module", "", "override module")
	flag.StringVar(&o.Profile, "profile", "", "override profile")
	flag.StringVar(&o.FirmwareFile, "firmware-file", "", "firmware file")
	flag.BoolVar(&o.DryRun, "dry", false, "dry run")

	flag.Parse()

	c, err := fktesting.CreateAndAuthenticate(ctx, o.Host, o.Scheme, o.Email, o.Password)
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("authenticated as %s (%s)", o.Email, o.Host)

	if o.FirmwareFile != "" {
		err := uploadFirmware(ctx, c, o.Module, o.Profile, o.FirmwareFile, o.DryRun)
		if err != nil {
			log.Fatalf("error adding firmware: %v", err)
		}
	}
}
