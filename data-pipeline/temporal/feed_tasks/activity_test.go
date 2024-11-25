package feed_tasks_test

import (
	"context"
	"testing"

	"github.com/Khan/genqlient/graphql"
	"github.com/ericbutera/amalgam/data-pipeline/temporal/feed_tasks"
	graph_client "github.com/ericbutera/amalgam/pkg/clients/graphql"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.temporal.io/sdk/mocks"
	"google.golang.org/grpc"
)

type MockFeedServiceClient struct {
	mock.Mock
	pb.FeedServiceClient
}

func (m *MockFeedServiceClient) CreateFeed(ctx context.Context, req *pb.CreateFeedRequest, opts ...grpc.CallOption) (*pb.CreateFeedResponse, error) {
	args := m.Called(ctx, req, opts)
	response, ok := args.Get(0).(*pb.CreateFeedResponse)
	if !ok {
		return nil, args.Error(1)
	}
	return response, args.Error(1)
}

type MockGraphClient struct {
	mock.Mock
	graphql.Client
}

func (m *MockGraphClient) MakeRequest(ctx context.Context, req *graphql.Request, resp *graphql.Response) error {
	args := m.Called(ctx, req, resp)
	return args.Error(0)
}

func TestGenerateFeedsActivity(t *testing.T) {
	t.Parallel()
	host := "faker:8080"
	count := 2

	// graph mock: yikes, this is a lot of work
	// this might be better: https://github.com/Yamashou/gqlgenc
	expectedResponse := graph_client.AddFeedResponse{
		AddFeed: graph_client.AddFeedAddFeedAddResponse{
			Id: "feed-id-123",
		},
	}
	graphMock := new(MockGraphClient)
	graphMock.On("MakeRequest", mock.Anything, mock.AnythingOfType("*graphql.Request"), mock.AnythingOfType("*graphql.Response")).
		Run(func(args mock.Arguments) {
			resp, ok := args.Get(2).(*graphql.Response)
			require.True(t, ok, "expected *graphql.Response")
			if responseData, ok := resp.Data.(*graph_client.AddFeedResponse); ok {
				responseData.AddFeed.Id = expectedResponse.AddFeed.Id
			}
		}).
		Return(nil).
		Times(count)

	activities := feed_tasks.NewActivities(graphMock, nil)
	err := activities.GenerateFeeds(context.Background(), host, count)

	require.NoError(t, err)
	graphMock.AssertExpectations(t)
}

func TestRefreshFeedsActivity(t *testing.T) {
	t.Parallel()

	mockRun := new(mocks.WorkflowRun)
	mockRun.On("GetID").Return("123")

	feedMock := new(mocks.Client)
	feedMock.On("ExecuteWorkflow",
		mock.Anything,
		mock.Anything,
		"FetchFeedsWorkflow",
	).Return(mockRun, nil)

	activities := feed_tasks.NewActivities(nil, feedMock)
	err := activities.RefreshFeeds(context.TODO())

	require.NoError(t, err)

	feedMock.AssertExpectations(t)
}
