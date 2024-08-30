package config

import (
	"fmt"
	"os"
	"time"
)

func InitTZ() {
	tz := os.Getenv("TZ")
	if tz == "" {
		tz = "Asia/Makassar"
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Printf("Error loading location: %v\n", err)
		return
	}

	time.Local = loc
}
