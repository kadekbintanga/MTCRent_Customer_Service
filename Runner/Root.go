package Runner

import (
	"Service/Config"
	"Service/Router"
	"fmt"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"net/http"
	"os"
)

var rootCmd = &cobra.Command{
	Use:  "root",
	Long: "Running service api",
	Run: func(cmd *cobra.Command, args []string) {
		config.InitDevMode()
		config.SetHost()

		Config.InitDB()
		Config.InitCors()
		Config.InitRabbitMQ()
		Config.InitMail()
		Config.InitRPC()
		Config.InitValidation()

		newCors := cors.New(Config.CorsOptions)

		newRoute := mux.NewRouter()
		router.RegisterRouter(newRoute, Router.Register)

		fmt.Println(fmt.Sprintf("Server started on %s", config.HostFull))

		err := http.ListenAndServe(config.Host+":"+config.Port, newCors.Handler(newRoute))
		if err != nil {
			panic(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&config.DevMode, "dev", false, "Set for development mode")
}
