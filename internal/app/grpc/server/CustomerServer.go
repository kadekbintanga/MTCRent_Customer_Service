package server

import (
	"context"
	"encoding/json"

	"google.golang.org/grpc"

	"service/internal/customer/repository"
	"service/internal/pkg/core"
	"service/internal/pkg/form"
	"service/internal/pkg/grpc/customer"
	"service/internal/pkg/parser"
)

type CustomerServer struct {
	customer.UnimplementedCustomerServiceServer
}

func (srv *CustomerServer) Register(serverRPC *grpc.Server) {
	customer.RegisterCustomerServiceServer(serverRPC, srv)
}

func (srv *CustomerServer) FirstByUUID(ctx context.Context, in *customer.FirstCustomerRequest) (*customer.CTResponse, error) {
	res, err := core.GRPCErrorHandler(func() (*customer.CTResponse, error) {
		repo := repository.NewCustomerRepository()
		customer := repo.FirstByForm(form.CustomerFilterForm{UUID: in.Uuid})

		parser := parser.CustomerParser{Object: customer}
		return srv.success(parser.First())

	})

	return res, err
}

/** --- UNEXPORTED FUNCTIONS --- */

func (srv *CustomerServer) success(result ...any) (*customer.CTResponse, error) {
	var data []byte
	if len(result) > 0 {
		data, _ = json.Marshal(result[0])
	}
	return &customer.CTResponse{Message: "Success", Result: data}, nil
}
