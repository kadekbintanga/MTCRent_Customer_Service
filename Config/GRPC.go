package Config

import (
	"github.com/globalxtreme/gobaseconf/config"
	"os"
	"time"
)

var (
	DevTestRPC string
)

func InitRPC() {
	config.RPCDialTimeout = 5 * time.Second
	DevTestRPC = os.Getenv("GRPC_DEV_TEST_HOST")
}
