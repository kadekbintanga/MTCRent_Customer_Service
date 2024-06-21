package runner

import (
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/console/command"
	"github.com/spf13/cobra"
	rabbitmq2 "service/internal/app/rabbitmq"
	config2 "service/internal/pkg/config"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "rabbitmq",
		Long: "Running RabbitMQ",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitDevMode()
			config2.InitDB()
			config2.InitRabbitMQ()

			rabbitmq2.Register()

			consumeCmd := command.RabbitMQConsumeCommand{}
			consumeCmd.Handle()
		},
	})
}
