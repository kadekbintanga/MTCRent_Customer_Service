package rabbitmq

import (
	xtremerabbitmq "github.com/globalxtreme/go-core/v2/rabbitmq"
	"service/internal/app/rabbitmq/consumer"
	"service/internal/pkg/constant"
)

func Register() {
	rabbitConsumer := xtremerabbitmq.Consumer{}
	rabbitConsumer.Set(map[string]xtremerabbitmq.RabbitMQConsumerInterface{
		constant.RABBITMQ_KEY_TESTING_MESSAGE: consumer.TestingConsumer{},
	})
}
