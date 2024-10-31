package runner

import (
	xtremeconsole "github.com/globalxtreme/go-core/v2/console"
	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
	"github.com/spf13/cobra"
	"service/internal/app/console"
	"service/internal/pkg/config"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "xtreme:schedule",
		Long: "Running Schedule",
		Run: func(cmd *cobra.Command, args []string) {
			xtremepkg.InitDevMode()

			config.InitTZ()

			DBConn := config.InitDB()
			defer DBConn()

			logRPC := xtremepkg.InitLogRPC()
			defer logRPC()

			xtremeconsole.Schedules(console.RegisterSchedule)
		},
	})
}
