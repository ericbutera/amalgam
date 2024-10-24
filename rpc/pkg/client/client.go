package client

import (
	pb "github.com/ericbutera/amalgam/pkg/feeds/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	Client pb.FeedServiceClient
	Conn   *grpc.ClientConn
}

func NewClient(target string) (*client, error) {
	// TODO: https://github.com/grpc-ecosystem/go-grpc-middleware/blob/main/interceptors/logging/examples/slog/example_test.go#L44-L55
	creds := insecure.NewCredentials() // TODO: use secure by default!
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return &client{
		Client: pb.NewFeedServiceClient(conn),
		Conn:   conn,
	}, nil
}
