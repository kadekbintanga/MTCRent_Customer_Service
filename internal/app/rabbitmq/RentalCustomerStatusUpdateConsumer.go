package rabbitmq

import (
	"errors"
	"sync"

	xtrememodel "github.com/globalxtreme/go-core/v2/model"
	xtremerabbitmq "github.com/globalxtreme/go-core/v2/rabbitmq"

	"service/internal/customer/service"
	"service/internal/pkg/core"
	form2 "service/internal/pkg/form"
)

type RentalCustomerStatusUpdateConsumer struct {
	xtremerabbitmq.AsyncWorkflowConsumerBase
	mutex sync.Mutex
}

func (consume *RentalCustomerStatusUpdateConsumer) Consume(message xtrememodel.RabbitMQMessage) (interface{}, error, []byte) {
	consume.mutex.Lock()
	defer consume.mutex.Unlock()

	return core.RabbitMQErrorHandler(func() (interface{}, error) {
		dataRaw, ok := message.Payload["data"]
		if !ok {
			return nil, errors.New("Payload does not contain data")
		}

		data, ok := dataRaw.(map[string]interface{})

		form := form2.CustomerStatusForm{
			StatusId:        data["statusId"].(int),
			BlacklistReason: data["blacklistReason"].(string),
		}
		form.Validate()

		srv := service.NewCustomerService()
		srv.UpdateStatus(data["uuid"].(string), form)

		return nil, nil
	})
}
