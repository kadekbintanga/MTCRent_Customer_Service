package Server

import (
	"Service/RPC/gRPC/DevTest"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

type ValidationServer struct {
	DevTest.UnimplementedValidationServiceServer
}

func (s *ValidationServer) Register(server *grpc.Server) {
	DevTest.RegisterValidationServiceServer(server, s)
}

func (s *ValidationServer) ValidateName(ctx context.Context, in *DevTest.ValidationNameRequest) (*DevTest.ValidationResponse, error) {
	fmt.Println("Name: " + in.GetName())

	return validationSuccess("Your name is valid")
}

func (s *ValidationServer) ValidateMultiField(ctx context.Context, in *DevTest.ValidationMultiFieldRequest) (*DevTest.ValidationResponse, error) {
	for _, name := range in.GetNames() {
		fmt.Println("Multi field: " + name.GetName())
	}

	return validationSuccess("Your multi fields is valid")
}

func validationSuccess(message ...string) (*DevTest.ValidationResponse, error) {
	msg := "Success"
	if len(message) > 0 {
		msg = message[0]
	}

	return &DevTest.ValidationResponse{Message: msg}, nil
}
