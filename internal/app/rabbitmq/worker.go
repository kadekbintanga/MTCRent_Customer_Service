package rabbitmq

import (
	"github.com/globalxtreme/gobaseconf/rabbitmq"
	"service/internal/app/rabbitmq/consumer"
	"service/internal/pkg/constant"
)

func Register() {
	rabbitConsumer := rabbitmq.Consumer{}
	rabbitConsumer.Set(map[string]rabbitmq.ConsumerInterface{
		constant.RABBITMQ_KEY_TESTING_MESSAGE: consumer.TestingConsumer{},
	})
}
