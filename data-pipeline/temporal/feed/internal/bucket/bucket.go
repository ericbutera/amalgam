package bucket

import (
	"context"
	"io"
	"log/slog"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed/internal/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/minio/minio-go/v7/pkg/lifecycle"
)

// TODO: this struct doesn't make sense (copied from old code)
type MinioBucket struct {
	Region string
	client *minio.Client
}

func NewMinioClient(config *config.Config) (*MinioBucket, error) {
	// TODO: refactor to use options pattern
	bucket := &MinioBucket{
		Region: "us-east-1",
	}

	client, err := minio.New(config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretAccessKey, ""),
		Secure: config.MinioUseSsl,
	})
	if err != nil {
		return nil, err
	}

	bucket.client = client
	return bucket, nil
}

func (bucket *MinioBucket) Create(ctx context.Context, bucketName string, expire bool) error {
	opts := minio.MakeBucketOptions{
		Region: bucket.Region,
	}
	// TODO: refactor (complexity)
	if err := bucket.client.MakeBucket(ctx, bucketName, opts); err != nil {
		exists, errBucketExists := bucket.client.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			slog.Debug("Bucket already exists", "bucket", bucketName)
		} else {
			return errBucketExists
		}
	} else if expire {
		if err := bucket.Expiry(ctx, bucketName); err != nil {
			return err
		}
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

func (bucket *MinioBucket) WriteStream(ctx context.Context, bucketName string, fileName string, reader io.Reader, contentType string, size int64) (*minio.UploadInfo, error) {
	opts := minio.PutObjectOptions{
		ContentType: contentType,
		//RetainUntilDate
		//Expires
	}
	info, err := bucket.client.PutObject(ctx, bucketName, fileName, reader, size, opts)
	if err != nil {
		return nil, err
	}

	return &info, err
}

func (bucket *MinioBucket) Read(ctx context.Context, bucketName string, fileName string) (io.ReadCloser, error) {
	reader, err := bucket.client.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return reader, nil
}
