package Runner

import (
	"Service/App/GRPC/Server"
	"Service/Config"
	"fmt"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/grpc"
	"github.com/spf13/cobra"
	"log"
	"os"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "grpc",
		Long: "Running gRPC",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitDevMode()
			Config.InitDB()

			addr := fmt.Sprintf("%s", os.Getenv("GRPC_HOST"))

			server := grpc.GRPCServer{}
			server.NewServer(addr)

			server.Register(
				&Server.ValidationServer{},
			)

			fmt.Println(fmt.Sprintf("gRPC server is running: %s", addr))
			if err := server.Serve(); err != nil {
				log.Fatalf("Failed to server: %v", err)
			}
		},
	})
}
