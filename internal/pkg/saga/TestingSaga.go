package saga

import (
	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
	error2 "service/internal/pkg/error"
	"service/internal/pkg/grpc/example"
	"service/internal/pkg/saga/client"
)

type TestingSaga struct {
	TestingClient *client.TestingClient

	testingCleanup func()

	testingRollBack []byte
}

/** --- NEW CLIENT --- */

func (saga *TestingSaga) NewTestingClient() {
	saga.TestingClient, saga.testingCleanup = client.NewTestingClient()
}

/** --- ITEM SERVICE CLIENT --- */

func (saga *TestingSaga) TestingStore(request *example.TestingRequest) (string, []byte) {
	check, err := saga.TestingClient.Testing.Store(saga.TestingClient.Ctx, request)
	if err != nil {
		saga.TestingClient = nil
		saga.testingCleanup()

		error2.ErrXtremeTestingSave(err.Error())
	}

	saga.testingRollBack = check.Result

	return check.Message, check.Result
}

func (saga *TestingSaga) TestingRollbackStore() {
	_, err := saga.TestingClient.Testing.RollbackStore(saga.TestingClient.Ctx, &example.RollBackRequest{Data: saga.testingRollBack})
	if err != nil {
		xtremepkg.LogError(err, true)
	}
}

/** --- DEFER FUNCTION --- */

func (saga *TestingSaga) Close() {
	if r := recover(); r != nil {
		if saga.TestingClient != nil && len(saga.testingRollBack) > 0 {
			saga.TestingRollbackStore()
			saga.testingCleanup()
		}

		panic(r)
	}

	saga.Cleanup()
}

func (saga *TestingSaga) Cleanup() {
	if saga.TestingClient != nil {
		saga.testingCleanup()
	}
}
