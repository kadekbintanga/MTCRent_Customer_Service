package grpc

import (
	"github.com/globalxtreme/gobaseconf/grpc"
	server2 "service/internal/app/grpc/server"
)

func Register(server grpc.GRPCServer) {
	server.Register(
		&server2.TestingServer{},
	)
}
