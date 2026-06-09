package rabbitmq

import (
	"fmt"
	data2 "github.com/globalxtreme/go-identifier/data"
	"sync"

	xtremerabbitmq "github.com/globalxtreme/go-core/v2/rabbitmq"

	"service/internal/customer/service"
	"service/internal/pkg/core"
	form2 "service/internal/pkg/form"
)

type RentalCustomerStatusUpdateConsumer struct {
	xtremerabbitmq.AsyncWorkflowConsumerBase
	mutex sync.Mutex
}

func (consume *RentalCustomerStatusUpdateConsumer) Consume(payload interface{}) (interface{}, error, []byte) {
	consume.mutex.Lock()
	defer consume.mutex.Unlock()

	return core.RabbitMQErrorHandler(func() (interface{}, error) {
		form := form2.CustomerStatusForm{}
		form.AsyncWorkflowParse(payload)
		// form.Validate()

		srv := service.NewCustomerService()
		srv.SetEmployeeIdentifier(data2.EmployeeIdentifierData{ID: form.CreatedByUUID, FullName: form.CreatedByName})
		srv.UpdateStatus(form.uu, form)

		return nil, nil
	})
}

func (consume *RentalCustomerStatusUpdateConsumer) Response(payload interface{}, data ...interface{}) interface{} {
	return nil

}

// func (consume *RentalCustomerStatusUpdateConsumer) Consume(message xtrememodel.RabbitMQMessage) (interface{}, error, []byte) {
// 	consume.mutex.Lock()
// 	defer consume.mutex.Unlock()

// 	return core.RabbitMQErrorHandler(func() (interface{}, error) {
// 		dataRaw, ok := message.Payload["data"]
// 		if !ok {
// 			return nil, errors.New("Payload does not contain data")
// 		}

// 		data, ok := dataRaw.(map[string]interface{})

// 		form := form2.CustomerStatusForm{
// 			StatusId:        data["statusId"].(int),
// 			BlacklistReason: data["blacklistReason"].(string),
// 		}
// 		form.Validate()

// 		srv := service.NewCustomerService()
// 		srv.UpdateStatus(data["uuid"].(string), form)

// 		return nil, nil
// 	})
// }
