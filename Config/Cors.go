package Config

import (
	"github.com/globalxtreme/gobaseconf/config"
	"github.com/rs/cors"
)

var (
	CorsOptions cors.Options
)

func InitCors() {
	CorsOptions.AllowedOrigins = []string{config.HostFull}
	CorsOptions.AllowCredentials = false
	CorsOptions.AllowedMethods = []string{"GET", "POST", "PUT", "DELETE"}
	CorsOptions.AllowedHeaders = []string{"*"}
}
