package runner

import (
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/spf13/cobra"
	"service/internal/app/database/seeder"
	config2 "service/internal/pkg/config"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "seeder",
		Long: "Running Seeder",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitDevMode()
			config2.InitDB()

			seeder.Run()
		},
	})
}
