package bucket

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/lifecycle"
	"github.com/mitchellh/mapstructure"
)

type Config struct {
	MinioEndpoint        string `mapstructure:"minio_endpoint"`
	MinioAccessKey       string `mapstructure:"minio_access_key"`
	MinioSecretAccessKey string `mapstructure:"minio_secret_access_key"`
	MinioRegion          string `mapstructure:"minio_region"`
	MinioUseSsl          bool   `mapstructure:"minio_use_ssl"`
	MinioTrace           bool   `mapstructure:"minio_trace"`
}

type customRoundTripper struct {
	Transport http.RoundTripper
}

func (c *customRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()
	slog.Debug("round-trip request", "url", req.URL, "start", start)

	resp, err := c.Transport.RoundTrip(req)

	slog.Debug("round-trip response", "status", resp.Status, "duration", time.Since(start), "url", req.URL)
	return resp, err
}

// TODO: this struct doesn't make sense (copied from old code)
type MinioBucket struct {
	Region string
	client *minio.Client
}

// use mapstructure to convert any existing struct to a Config
// example: convert a temporal workflow config to a minio config
func NewConfig(data any) (*Config, error) {
	config := &Config{}
	if err := mapstructure.Decode(data, config); err != nil {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}
	return config, nil
}

func NewMinioClient(config *Config) (*MinioBucket, error) {
	bucket := &MinioBucket{
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

func (b *MinioBucket) Create(ctx context.Context, bucketName string) error {
	opts := minio.MakeBucketOptions{
		Region: b.Region,
	}
	if err := b.client.MakeBucket(ctx, bucketName, opts); err != nil {
		if err := b.handleBucketExists(ctx, bucketName, err); err != nil {
			return err
		}
	}
	return nil
}

func (b *MinioBucket) handleBucketExists(ctx context.Context, bucketName string, err error) error {
	exists, errBucketExists := b.client.BucketExists(ctx, bucketName)
	if errBucketExists == nil && exists {
		return nil
	}
	return err
}

func (b *MinioBucket) SetBucketExpiry(ctx context.Context, bucketName string) error {
	return b.Expiry(ctx, bucketName)
}

func (b *MinioBucket) Exists(ctx context.Context, bucketName string, fileName string) (bool, error) {
	info, err := b.client.StatObject(ctx, bucketName, fileName, minio.StatObjectOptions{})
	if err != nil {
		return false, fmt.Errorf("failed to stat object: %w", err)
	}
	if info.Key == fileName {
		return true, nil
	}
	return false, nil
}

func (b *MinioBucket) Expiry(ctx context.Context, bucketName string) error {
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

func (b *MinioBucket) WriteStream(
	ctx context.Context,
	bucketName string,
	fileName string,
	reader io.Reader,
	contentType string,
	size int64,
) (minio.UploadInfo, error) {
	opts := minio.PutObjectOptions{
		ContentType: contentType,
		// TODO: adjust retention & storage class to reduce cloud spend
		// Less frequently accessed objects should go in cheaper storage classes (e.g. Glacier/archive)
		// Mode, RetainUntilDate, Expires, StorageClass
	}
	return b.client.PutObject(ctx, bucketName, fileName, reader, size, opts)
}

func (b *MinioBucket) Read(ctx context.Context, bucketName string, fileName string) (io.ReadCloser, error) {
	obj, err := b.client.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get object: %w", err)
	}
	return obj, nil
}
