package bucket

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/ericbutera/amalgam/pkg/config/env"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/lifecycle"
	"github.com/mitchellh/mapstructure"
)

type Config struct {
	MinioEndpoint        string `env:"MINIO_ENDPOINT"`
	MinioAccessKey       string `env:"MINIO_ACCESS_KEY"`
	MinioSecretAccessKey string `env:"MINIO_SECRET_ACCESS_KEY"`
	MinioRegion          string `env:"MINIO_REGION" envDefault:"us-east-1"`
	MinioUseSsl          bool   `env:"MINIO_USE_SSL"`
	MinioTrace           bool   `env:"MINIO_TRACE"`
}

const DefaultWriteSize int64 = -1

// helper for object storage
type Bucket interface {
	Create(ctx context.Context, bucketName string) error
	SetBucketExpiry(ctx context.Context, bucketName string) error
	Exists(ctx context.Context, bucketName string, fileName string) (bool, error)
	WriteStream(ctx context.Context, bucketName string, fileName string, reader io.Reader, contentType string) (*UploadInfo, error)
	Read(ctx context.Context, bucketName string, fileName string) (io.ReadCloser, error)
}

type customRoundTripper struct {
	Transport http.RoundTripper
}

func (c *customRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	slog.Debug("round-trip request", "url", req.URL, "start", start)

	resp, err := c.Transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	slog.Debug("round-trip response", "status", resp.Status, "duration", time.Since(start), "url", req.URL)

	return resp, err
}

// TODO: this struct doesn't make sense (copied from old code)
type Minio struct {
	Region string
	client *minio.Client
}

// use mapstructure to convert any existing struct to a Config
// example: convert a temporal workflow config to a minio config
func NewConfig(data any) (*Config, error) {
	config := &Config{}

	err := mapstructure.Decode(data, config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}

	return config, nil
}

func NewMinio(config *Config) (Bucket, error) {
	bucket := &Minio{
		Region: config.MinioRegion,
	}

	transport, err := minio.DefaultTransport(config.MinioUseSsl)
	if err != nil {
		return nil, fmt.Errorf("failed to create transport: %w", err)
	}

	client, err := minio.New(config.MinioEndpoint, &minio.Options{
		Creds:     credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretAccessKey, ""),
		Secure:    config.MinioUseSsl,
		Transport: &customRoundTripper{Transport: transport},
	})
	if err != nil {
		return nil, err
	}

	if config.MinioTrace {
		client.TraceOn(nil)
	}

	bucket.client = client

	return bucket, nil
}

func NewMinioFromEnv() (Bucket, error) {
	config, err := env.New[Config]()
	if err != nil {
		return nil, err
	}

	return NewMinio(config)
}

func (b *Minio) Create(ctx context.Context, bucketName string) error {
	opts := minio.MakeBucketOptions{
		Region: b.Region,
	}

	err := b.client.MakeBucket(ctx, bucketName, opts)
	if err != nil {
		err := b.handleBucketExists(ctx, bucketName, err)
		if err != nil {
			return err
		}
	}

	return nil
}

func (b *Minio) SetBucketExpiry(ctx context.Context, bucketName string) error {
	return b.Expiry(ctx, bucketName)
}

func (b *Minio) Exists(ctx context.Context, bucketName string, fileName string) (bool, error) {
	info, err := b.client.StatObject(ctx, bucketName, fileName, minio.StatObjectOptions{})
	if err != nil {
		return false, fmt.Errorf("failed to stat object: %w", err)
	}

	if info.Key == fileName {
		return true, nil
	}

	return false, nil
}

func (b *Minio) Expiry(ctx context.Context, bucketName string) error {
	config := lifecycle.NewConfiguration()
	config.Rules = []lifecycle.Rule{
		{
			ID:     "expire-bucket",
			Status: "Enabled",
			Expiration: lifecycle.Expiration{
				Days: 1,
			},
		},
	}

	return b.client.SetBucketLifecycle(ctx, bucketName, config)
}

type UploadInfo = minio.UploadInfo

func (b *Minio) WriteStream(
	ctx context.Context,
	bucketName string,
	fileName string,
	reader io.Reader,
	contentType string,
) (*UploadInfo, error) {
	opts := minio.PutObjectOptions{
		ContentType: contentType,
		// TODO: adjust retention & storage class to reduce cloud spend
		// Less frequently accessed objects should go in cheaper storage classes (e.g. Glacier/archive)
		// Mode, RetainUntilDate, Expires, StorageClass
	}

	upload, err := b.client.PutObject(ctx, bucketName, fileName, reader, DefaultWriteSize, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to put object: %w", err)
	}

	return &upload, nil
}

func (b *Minio) Read(ctx context.Context, bucketName string, fileName string) (io.ReadCloser, error) {
	obj, err := b.client.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %w", err)
	}

	return obj, nil
}

func (b *Minio) handleBucketExists(ctx context.Context, bucketName string, err error) error {
	exists, errBucketExists := b.client.BucketExists(ctx, bucketName)
	if errBucketExists == nil && exists {
		return nil
	}

	return err
}
