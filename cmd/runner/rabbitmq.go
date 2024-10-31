package runner

import (
	"github.com/globalxtreme/go-core/v2/console/command"
	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
	"github.com/spf13/cobra"
	"service/internal/app/rabbitmq"
	"service/internal/pkg/config"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "xtreme:rabbitmq",
		Long: "Running RabbitMQ",
		Run: func(cmd *cobra.Command, args []string) {
			xtremepkg.InitDevMode()

			config.InitTZ()

			DBClose := config.InitDB()
			defer DBClose()

			rabbitMQClose := config.InitRabbitMQ()
			defer rabbitMQClose()

			logCleanup := xtremepkg.InitLogRPC()
			defer logCleanup()

			rabbitmq.Register()

			consumeCmd := command.RabbitMQConsumeCommand{}
			consumeCmd.Handle()
		},
	})
}
