package config

import (
	"errors"
	"fmt"
	xtremepkg "github.com/globalxtreme/go-core/v2/pkg"
	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"net/http"
	"os"
	"service/internal/pkg/grpc/example"
)

func ErrorHandler(fn func() error) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "panic: %v\n", r)
			xtremepkg.LogError(r)
		}
	}()

	return fn()
}

func GRPCErrorHandler(fn func() (*example.EXResponse, error)) (res *example.EXResponse, err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "panic: %v\n", r)
			xtremepkg.LogError(r)

			if panicData, ok := r.(*xtremeres.ResponseError); ok {
				status := panicData.Status
				err = errors.New(fmt.Sprintf("Code: %d. Message: %s. InternalMsg: %s", status.Code, status.Message, status.InternalMsg))
			} else if panicData, ok := r.(error); ok {
				err = errors.New(fmt.Sprintf("Code: %d. Message: %v", http.StatusInternalServerError, panicData.Error()))
			} else {
				err = errors.New(fmt.Sprintf("Code: %d. Message: An error Occurred.", http.StatusInternalServerError))
			}
		}
	}()

	res, err = fn()
	return
}
