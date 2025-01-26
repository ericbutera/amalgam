package feed_add

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/bucket"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/internal/transforms"
	"github.com/ericbutera/amalgam/internal/http/fetch"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	rpc "github.com/ericbutera/amalgam/services/rpc/pkg/client"
	"github.com/samber/lo"
)

const (
	BucketName    = "feed_add" // TODO: ensure files have a delete TTL (low chance these will ever be reused)
	RssPathFormat = "feed_add/%s.xml"
)

var (
	ErrInvalidURL     = errors.New("invalid URL")
	ErrInvalidContent = errors.New("invalid content")
)

type Activities struct {
	transforms transforms.Transforms
	fetch      fetch.Fetch
	bucket     bucket.Bucket
	rpc        pb.FeedServiceClient
	Closers    func()
}

func NewActivities(fetch fetch.Fetch, bucket bucket.Bucket, rpc pb.FeedServiceClient) *Activities {
	return &Activities{
		transforms: transforms.New(),
		fetch:      fetch,
		bucket:     bucket,
		rpc:        rpc,
		Closers:    func() {},
	}
}

// Create a new Activities struct using environment variables. Will panic if any errors occur.
func NewActivitiesFromEnv() *Activities {
	rpcClient, closer := lo.Must2(rpc.NewFromEnv())

	a := NewActivities(
		lo.Must(fetch.New()),
		lo.Must(bucket.NewMinioFromEnv()),
		rpcClient,
	)
	a.Closers = func() {
		closer() //nolint:errcheck
	} // TODO: revisit this pattern
	return a
}

func (a *Activities) CreateVerifyRecord(ctx context.Context, verification FeedVerification) (*FeedVerification, error) {
	// verification
	// - will fail if the URL is invalid
	// - normalizes the URL (imagine ?p=1&x=1 vs ?x=1&p=1)
	resp, err := a.rpc.CreateFeedVerification(ctx, &pb.CreateFeedVerificationRequest{
		Verification: &pb.FeedVerification{
			Url:        verification.URL,
			UserId:     verification.UserID,
			WorkflowId: verification.WorkflowID,
		},
	})
	if err != nil {
		return nil, err
	}
	v := resp.GetVerification()
	return &FeedVerification{
		ID:         v.GetId(),
		WorkflowID: v.GetWorkflowId(),
		URL:        v.GetUrl(),
		UserID:     v.GetUserId(),
	}, nil
}

func (a *Activities) Fetch(ctx context.Context, verification FeedVerification) (string, error) {
	// blob := fmt.Sprintf(RssPathFormat, workflowID)
	// entry := slog.Default().With( "file", blob, "url", verification.URL,)

	// TODO: ensure fetch.Url makes a fetch_history entry
	err := a.fetch.Url(ctx, verification.URL, func(params fetch.CallbackParams) error {
		// retry on 500
		// expo backoff on 429
		// proceed if 200 (possibly other 2xx)
		// stop on everything else
		if params.StatusCode != http.StatusOK {
			return ErrInvalidContent
		}

		// TODO: reuse content from this fetch during the feed_fetch workflow (not possible at the moment due to missing "feed_id")
		// TODO: prevent abuse
		// - limit number of feeds a user_id can add (rate limit by user_id)
		// - limit size of content (large content size, malformed tags)
		// - limit total size of content user_id can add across feeds (user could add 1000 1k feeds to get around limit of per feed size)
		// upload, err := a.bucket.WriteStream(ctx, BucketName, blob, params.Reader, params.ContentType)
		// if err != nil {
		// 	return err
		// }
		// entry.Debug("validating new feed", "key", upload.Key, "bucket", upload.Bucket, "size", upload.Size)

		// Hack to validate RSS. In the future this work won't get thrown away.
		_, err := a.transforms.RssToArticles(params.Reader)
		return err
	}, nil)

	return "", err // TODO: return blob filename
}

func (a *Activities) CreateFeed(ctx context.Context, verification FeedVerification) error {
	resp, err := a.rpc.CreateFeed(ctx, &pb.CreateFeedRequest{
		Feed: &pb.CreateFeedRequest_Feed{
			Url: verification.URL,
		},
		User: &pb.User{Id: verification.UserID},
	})
	if err != nil {
		return err
	}
	slog.Debug("created feed", "feed_id", resp.GetId(), "verification_id", verification.ID)
	return nil
}
