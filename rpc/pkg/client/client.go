package client

import (
	pb "github.com/ericbutera/amalgam/pkg/rpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	Client pb.FeedServiceClient
	Conn   *grpc.ClientConn
}

func NewClient(target string) (*client, error) {
	creds := insecure.NewCredentials() // TODO: use secure by default!
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, err
	}

	return &client{
		Client: pb.NewFeedServiceClient(conn),
		Conn:   conn,
	}, nil

	/*
		Example usage:

		client, _ := NewClient("localhost:50055")
		defer client.Close()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		feeds, err := client.ListFeeds(ctx, &pb.Empty{})
		if err != nil {
			log.Fatalf("Failed to list feeds: %v", err)
		}
		fmt.Println("Feeds:", feeds)
	*/
}
