package DevTest

import (
	"Service/Config"
	"Service/RPC/gRPC/DevTest"
	"context"
	"github.com/globalxtreme/gobaseconf/grpc"
	"log"
	"time"
)

type ValidationClient struct {
	grpc.GRPCClient
	Validation DevTest.ValidationServiceClient
}

func NewValidationClient(timeout ...time.Duration) (*ValidationClient, context.CancelFunc) {
	client := ValidationClient{}
	cleanup := client.RPCDialClient(Config.DevTestRPC, timeout...)

	client.Validation = DevTest.NewValidationServiceClient(client.Conn)

	return &client, cleanup
}

func (rpc *ValidationClient) ValidationName(name string) (string, []byte) {
	validate, err := rpc.Validation.ValidateName(rpc.Ctx, &DevTest.ValidationNameRequest{Name: name})
	if err != nil {
		log.Panicf("could not validate: %v", err)
	}

	return validate.Message, validate.Result
}

func (rpc *ValidationClient) ValidationMultiField(names []*DevTest.ValidationNameRequest) (string, []byte) {
	validate, err := rpc.Validation.ValidateMultiField(rpc.Ctx, &DevTest.ValidationMultiFieldRequest{Names: names})
	if err != nil {
		log.Panicf("could not validate: %v", err)
	}

	return validate.Message, validate.Result
}
