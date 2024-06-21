package runner

import (
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/database/migration"
	"github.com/spf13/cobra"
	migration2 "service/internal/app/database/migration"
	config2 "service/internal/pkg/config"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "migration",
		Long: "Running Migration",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitDevMode()
			config2.InitDB()

			migration.Migrate(migration2.Tables(), migration2.Columns())
		},
	})
}
