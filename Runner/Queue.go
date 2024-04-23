package Runner

import (
	"Service/App/Job/Resizing"
	"Service/App/Service/Constant/Queue"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/queue"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "queue",
		Long: "Running Queue",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitDevMode()

			queueNames := cmd.Flags().String("q", Queue.QUEUE_HIGH, "Queue name")

			configurations := [...]queue.JobConf{
				{
					Context:     Resizing.Resizing{},
					JobFunc:     (*Resizing.Resizing).Image,
					Concurrency: 10,
					QueueName:   Queue.QUEUE_HIGH,
					JobName:     Queue.JOB_REZISE_IMAGE,
					Priority:    1,
				},
			}

			worker := queue.Queue{Names: *queueNames}
			worker.Work(configurations[:])
		},
	})
}
