package server

import (
	"context"
	"encoding/json"

	"github.com/globalxtreme/go-identifier/data"
	"google.golang.org/grpc"

	"service/internal/customer/repository"
	"service/internal/customer/service"
	"service/internal/pkg/core"
	"service/internal/pkg/form"
	"service/internal/pkg/grpc/customer"
	"service/internal/pkg/parser"
)

type CustomerServer struct {
	customer.UnimplementedCustomerServiceServer

	rollbackData map[string]interface{}
}

func (srv *CustomerServer) Register(serverRPC *grpc.Server) {
	customer.RegisterCustomerServiceServer(serverRPC, srv)
}

func (srv *CustomerServer) FirstByUUID(ctx context.Context, in *customer.FirstCustomerRequest) (*customer.CTResponse, error) {
	res, err := core.GRPCErrorHandler(func() (*customer.CTResponse, error) {
		repo := repository.NewCustomerRepository()
		customer := repo.FirstByForm(form.CustomerFilterForm{UUID: in.Uuid})

		parser := parser.CustomerParser{Object: customer}
		return srv.success(parser.FirstGRPC())

	})

	return res, err
}

func (srv *CustomerServer) UpdateStatus(ctx context.Context, in *customer.CustomerUpdateStatusRequest) (*customer.CTResponse, error) {
	res, err := core.GRPCErrorHandler(func() (*customer.CTResponse, error) {
		service := service.NewCustomerService()
		service.SetEmployeeIdentifier(data.EmployeeIdentifierData{ID: in.CreatedBy, FullName: in.CreatedByName})
		_, rollbackCustomer := service.UpdateStatus(in.Uuid, form.CustomerStatusForm{
			StatusId:        int(in.StatusId),
			BlacklistReason: in.BlacklistReason,
		})

		srv.rollbackData = map[string]interface{}{
			"uuid":            rollbackCustomer.UUID,
			"statusId":        rollbackCustomer.StatusId,
			"blacklistReason": rollbackCustomer.BlacklistReason,
		}

		return srv.success(srv.rollbackData)
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
