package config

import (
	"github.com/globalxtreme/gobaseconf/config"
	"os"
	"time"
)

var (
	TestingRPC string
)

func InitRPC() {
	config.RPCDialTimeout = 5 * time.Second
	TestingRPC = os.Getenv("GRPC_TESTING_HOST")
}
