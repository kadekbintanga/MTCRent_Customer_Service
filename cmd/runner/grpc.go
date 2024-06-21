package runner

import (
	"fmt"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/grpc"
	"github.com/spf13/cobra"
	"log"
	"os"
	grpc2 "service/internal/app/grpc"
	config2 "service/internal/pkg/config"
)

func init() {
	rootCmd.AddCommand(&cobra.Command{
		Use:  "grpc",
		Long: "Running gRPC",
		Run: func(cmd *cobra.Command, args []string) {
			config.InitDevMode()
			config2.InitDB()

			addr := fmt.Sprintf("%s", os.Getenv("GRPC_HOST"))

			server := grpc.GRPCServer{}
			server.NewServer(addr)

			grpc2.Register(server)

			fmt.Println(fmt.Sprintf("gRPC server is running: %s", addr))
			if err := server.Serve(); err != nil {
				log.Fatalf("Failed to server: %v", err)
			}
		},
	})
}
