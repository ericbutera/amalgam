package bucket

import (
	"context"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/lifecycle"
)

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

func NewMinioClient(config *config.Config) (*MinioBucket, error) {
	bucket := &MinioBucket{
		Region: config.MinioRegion,
	}
	transport, err := minio.DefaultTransport(config.MinioUseSsl)
	if err != nil {
		return nil, err
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

func (bucket *MinioBucket) Create(ctx context.Context, bucketName string, expire bool) error {
	opts := minio.MakeBucketOptions{
		Region: bucket.Region,
	}

	if err := bucket.client.MakeBucket(ctx, bucketName, opts); err != nil {
		if err := bucket.handleBucketExists(ctx, bucketName, err); err != nil {
			return err
		}
	}

	if expire {
		return bucket.setBucketExpiry(ctx, bucketName)
	}

	return nil
}

func (bucket *MinioBucket) handleBucketExists(ctx context.Context, bucketName string, err error) error {
	exists, errBucketExists := bucket.client.BucketExists(ctx, bucketName)
	if errBucketExists == nil && exists {
		return nil
	}
	return err
}

func (bucket *MinioBucket) setBucketExpiry(ctx context.Context, bucketName string) error {
	if err := bucket.Expiry(ctx, bucketName); err != nil {
		return err
	}
	return nil
}

func (bucket *MinioBucket) Exists(ctx context.Context, bucketName string, fileName string) (bool, error) {
	info, err := bucket.client.StatObject(ctx, bucketName, fileName, minio.StatObjectOptions{})
	if err != nil {
		return false, err
	}
	if info.Key == fileName {
		return true, nil
	}
	return false, nil
}

func (bucket *MinioBucket) Expiry(ctx context.Context, bucketName string) error {
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
	return bucket.client.SetBucketLifecycle(ctx, bucketName, config)
}

func (bucket *MinioBucket) WriteStream(ctx context.Context, bucketName string, fileName string, reader io.Reader, contentType string, size int64) (minio.UploadInfo, error) {
	opts := minio.PutObjectOptions{
		ContentType: contentType,
		// TODO: adjust retention & storage class to reduce cloud spend
		// Less frequently accessed objects should go in cheaper storage classes (e.g. Glacier/archive)
		// Mode, RetainUntilDate, Expires, StorageClass
	}
	return bucket.client.PutObject(ctx, bucketName, fileName, reader, size, opts)
}

func (bucket *MinioBucket) Read(ctx context.Context, bucketName string, fileName string) (io.ReadCloser, error) {
	return bucket.client.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
}
