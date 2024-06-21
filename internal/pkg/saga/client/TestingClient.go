package client

import (
	"context"
	"github.com/globalxtreme/gobaseconf/grpc"
	"service/internal/pkg/config"
	"service/internal/pkg/grpc/example"
	"time"
)

type TestingClient struct {
	grpc.GRPCClient
	Testing example.TestingServiceClient
}

func NewTestingClient(timeout ...time.Duration) (*TestingClient, context.CancelFunc) {
	client := TestingClient{}
	cleanup := client.RPCDialClient(config.TestingRPC, timeout...)

	client.Testing = example.NewTestingServiceClient(client.Conn)

	return &client, cleanup
}
