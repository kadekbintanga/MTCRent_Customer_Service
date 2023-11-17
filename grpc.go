package main

import (
	"Service/App/GRPC/Server"
	"fmt"
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/globalxtreme/gobaseconf/grpc"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	config.SetHost()

	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)

	server := grpc.GRPCServer{}
	server.NewServer(addr)

	server.Register(
		&Server.ValidationServer{},
	)

	fmt.Println(fmt.Sprintf("gRPC server is running: %s", addr))
	if err := server.Serve(); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
