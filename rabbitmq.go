package main

import (
	"Service/App/MessageBorker/Testing"
	"Service/App/Service/Constant/MessageBroker"
	"Service/Config"
	"github.com/globalxtreme/gobaseconf/rabbitmq"
	"github.com/globalxtreme/gobaseconf/rabbitmq/command"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	Config.InitDB()
	Config.InitRabbitMQ()

	setupRabbitMQConsumer()

	cmd := command.RabbitMQConsumeCommand{}
	cmd.Handle()
}

func setupRabbitMQConsumer() {
	consumer := rabbitmq.Consumer{}
	consumer.Set(map[string]rabbitmq.ConsumerInterface{
		MessageBroker.RABBITMQ_KEY_TESTING_MESSAGE: Testing.TestingConsumer{},
	})
}
