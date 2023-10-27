package main

import (
	"Service/App/Model/Testing"
	TestingParser "Service/App/Parser/Testing"
	"Service/App/Service/Constant/MessageBroker"
	"Service/Config"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/rabbitmq"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	Config.InitDB()
	Config.InitRabbitMQ()

	testing := Testing.Testing{}
	Config.PgSQL.First(&testing)

	parser := TestingParser.TestingParser{Object: testing}

	amqp := rabbitmq.RabbitMQ{Data: parser.First(), Key: MessageBroker.RABBITMQ_KEY_TESTING_MESSAGE}
	amqp.OnSender(testing.ID.ID, testing.TableName()).
		OnQueue(config.RabbitMQConf.Queue).
		Push()
}
