package Error

import (
	"github.com/globalxtreme/gobaseconf/response"
	"net/http"
)

func ErrActivityNotFound() {
	response.Error(http.StatusNotFound, "Service not found", "", nil)
}

func ErrActivitySave(internalMsg string) {
	response.Error(http.StatusInternalServerError, "Unable to save service", internalMsg, nil)
}

func ErrActivityActionType() {
	response.Error(http.StatusInternalServerError, "Activity action doesn't exists!!", "", nil)
}
