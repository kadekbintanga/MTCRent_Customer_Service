package runner

import (
	"fmt"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/router"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"service/internal/app/api"
	config2 "service/internal/pkg/config"
)

var rootCmd = &cobra.Command{
	Use:  "root",
	Long: "Running service api",
	Run: func(cmd *cobra.Command, args []string) {
		config.InitDevMode()
		config.SetHost()

		config2.InitDB()
		config2.InitCors()
		config2.InitRabbitMQ()
		config2.InitMail()
		config2.InitRPC()
		config2.InitValidation()

		newCors := cors.New(config2.CorsOptions)

		newRoute := mux.NewRouter()
		router.RegisterRouter(newRoute, api.Register)

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
