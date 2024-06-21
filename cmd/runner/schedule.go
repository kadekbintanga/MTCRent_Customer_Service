package runner

import (
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/console"
	"github.com/spf13/cobra"
	console2 "service/internal/app/console"
	config2 "service/internal/pkg/config"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "schedule",
		Long: "Running Schedule",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitDevMode()
			config2.InitDB()

			console.Schedules(console2.RegisterSchedule)
		},
	})
}
