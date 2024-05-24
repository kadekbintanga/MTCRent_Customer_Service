package Runner

import (
	"Service/App/MessageBorker/Testing"
	"Service/App/Service/Constant/MessageBroker"
	"Service/Config"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/console/command"
	"github.com/globalxtreme/gobaseconf/rabbitmq"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "rabbitmq",
		Long: "Running RabbitMQ",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitDevMode()
			Config.InitDB()
			Config.InitRabbitMQ()

			setupRabbitMQConsumer()

			consumeCmd := command.RabbitMQConsumeCommand{}
			consumeCmd.Handle()
		},
	})
}

func setupRabbitMQConsumer() {
	consumer := rabbitmq.Consumer{}
	consumer.Set(map[string]rabbitmq.ConsumerInterface{
		MessageBroker.RABBITMQ_KEY_TESTING_MESSAGE: Testing.TestingConsumer{},
	})
}
