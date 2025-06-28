package rabbitmq

import (
	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
	"service/internal/pkg/core"
	form2 "service/internal/pkg/form"
	"service/internal/testing/service"
	"sync"
)

type TestingProcessedConsumer struct {
	mutex sync.Mutex
}

func (consume *TestingProcessedConsumer) Consume(message xtrememodel.RabbitMQMessage) (interface{}, error) {
	consume.mutex.Lock()
	defer consume.mutex.Unlock()

	return core.RabbitMQErrorHandler(func() (interface{}, error) {
		form := form2.TestingForm{}
		response, err := form.RabbitMQProcessedParse(message)
		if err != nil {
			return nil, err
		}

		xtremepkg.LogInfo(response.Status.Name)

		srv := service.NewTestingService()
		testing := srv.CreateConsumer(form)

		// Taruh di parser
		result := map[string]interface{}{
			"name":     testing.Name,
			"totalSub": len(testing.Subs),
		}

		return result, nil
	})
}
