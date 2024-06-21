package runner

import (
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/queue"
	"github.com/spf13/cobra"
	queue2 "service/internal/app/queue"
	"service/internal/pkg/constant"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "queue",
		Long: "Running Queue",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitDevMode()

			queueNames := cmd.Flags().String("q", constant.QUEUE_HIGH, "Queue name")
			configurations := queue2.Register()

			worker := queue.Queue{Names: *queueNames}
			worker.Work(configurations)
		},
	})
}
