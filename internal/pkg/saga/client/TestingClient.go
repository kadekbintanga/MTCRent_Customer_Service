package client

import (
	"context"
	xtremegrpc "github.com/globalxtreme/go-core/v2/grpc"
	"service/internal/pkg/config"
	"service/internal/pkg/grpc/example"
	"time"
)

type TestingClient struct {
	xtremegrpc.GRPCClient
	Testing example.TestingServiceClient
}

func NewTestingClient(timeout ...time.Duration) (*TestingClient, context.CancelFunc) {
	client := TestingClient{}
	cleanup := client.RPCDialClient(config.TestingRPC, timeout...)

	client.Testing = example.NewTestingServiceClient(client.Conn)

	return &client, cleanup
}
