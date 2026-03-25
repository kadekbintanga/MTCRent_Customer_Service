package saga

import (
	"context"
	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
	error2 "service/internal/pkg/error"
	"service/internal/pkg/grpc/example"
	"service/internal/pkg/saga/grpc"
	"service/internal/pkg/saga/privateapi"
	"time"
)

// TODO: Hanya contoh. nanti langsung hapus saja
type TestingSaga struct {
	testingAPI privateapi.TestingAPI

	testingRPCRollBack []byte
	testingAPIRollBack interface{}
}

/** --- SETTER --- */

func (sg *TestingSaga) NewTestingSaga() {
	sg.testingAPI = privateapi.NewTestingAPI()
}

/** --- ITEM SERVICE CLIENT --- */

func (sg *TestingSaga) TestingStore(request *example.TestingRequest) (string, []byte) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := grpc.TestingRPCClient.Store(ctx, request)
	if err != nil {
		error2.ErrXtremeTestingSave(err.Error())
	}

	result := resp.GetResult()

	sg.testingRPCRollBack = result

	return resp.GetMessage(), result
}

func (sg *TestingSaga) TestingStoreAPI(request *example.TestingRequest) interface{} {
	resp := sg.testingAPI.Store(request)
	result := resp.Result
	if result != nil {
		sg.testingAPIRollBack = result
	}

	return result
}

func (sg *TestingSaga) TestingRollbackStore() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := grpc.TestingRPCClient.RollbackStore(ctx, &example.RollBackRequest{Data: sg.testingRPCRollBack})
	if err != nil {
		xtremepkg.LogError(err, true)
	}
}

func (sg *TestingSaga) TestingAPIRollbackStore() {
	sg.testingAPI.RollBack(sg.testingAPIRollBack)
}

/** --- DEFER FUNCTION --- */

func (sg *TestingSaga) Close() {
	if r := recover(); r != nil {
		if len(sg.testingRPCRollBack) > 0 {
			sg.TestingRollbackStore()
		}

		if sg.testingAPIRollBack != nil {
			sg.TestingAPIRollbackStore()
		}

		panic(r)
	}
}
