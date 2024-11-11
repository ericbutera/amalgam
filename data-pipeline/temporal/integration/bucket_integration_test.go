package bucket_test

import (
	"context"
	"os"
	"testing"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/bucket"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"
)

const TestBucketName = "test-bucket-name"

func TestMinioCreate(t *testing.T) {
	// https://mickey.dev/posts/go-build-tags-testing/
	t.Parallel()
	ctx := context.Background()

	config := &bucket.Config{
		MinioEndpoint:        os.Getenv("MINIO_ENDPOINT"),
		MinioAccessKey:       os.Getenv("MINIO_ACCESS_KEY"),
		MinioSecretAccessKey: os.Getenv("MINIO_SECRET_ACCESS_KEY"),
		MinioUseSsl:          lo.Ternary(os.Getenv("MINIO_USE_SSL") == "false", false, true),
	}

	if config.MinioEndpoint == "" {
		t.Skip("skipping test; no Minio endpoint")
	}

	b, err := bucket.NewMinio(config)
	require.NoError(t, err)

	err = b.Create(ctx, TestBucketName)
	require.NoError(t, err)
	// TODO: test bucket exists

	client, err := minio.New(config.MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.MinioAccessKey, config.MinioSecretAccessKey, ""),
		Secure: config.MinioUseSsl,
	})
	require.NoError(t, err)

	ok, err := client.BucketExists(ctx, TestBucketName)
	require.NoError(t, err)
	require.True(t, ok)
}
