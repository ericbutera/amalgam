package feed_tasks

import (
	"context"
	"testing"

	"github.com/Khan/genqlient/graphql"
	graph_client "github.com/ericbutera/amalgam/pkg/clients/graphql"
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type MockFeedServiceClient struct {
	mock.Mock
	pb.FeedServiceClient
}

func (m *MockFeedServiceClient) CreateFeed(ctx context.Context, req *pb.CreateFeedRequest, opts ...grpc.CallOption) (*pb.CreateFeedResponse, error) {
	args := m.Called(ctx, req, opts)
	return args.Get(0).(*pb.CreateFeedResponse), args.Error(1)
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
			resp := args.Get(2).(*graphql.Response)
			if responseData, ok := resp.Data.(*graph_client.AddFeedResponse); ok {
				responseData.AddFeed.Id = expectedResponse.AddFeed.Id
			}
		}).
		Return(nil).
		Times(count)

	activities := NewActivities(graphMock)
	err := activities.GenerateFeeds(context.Background(), host, count)

	assert.NoError(t, err)
	graphMock.AssertExpectations(t)
}
