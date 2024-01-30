package Config

import (
	"fmt"
	"github.com/globalxtreme/gobaseconf/helpers/xtremelog"
	"os"
)

func ErrorHandler() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "panic: %v\n", r)
			xtremelog.Error(r)
		}
	}()
}
