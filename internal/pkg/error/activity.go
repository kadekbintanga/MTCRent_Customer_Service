package error

import (
	"github.com/globalxtreme/gobaseconf/response"
	"net/http"
)

func ErrXtremeActivityNotFound() {
	response.Error(http.StatusNotFound, "Service not found", "", nil)
}

func ErrXtremeActivitySave(internalMsg string) {
	response.Error(http.StatusInternalServerError, "Unable to save service", internalMsg, nil)
}

func ErrXtremeActivityActionType() {
	response.Error(http.StatusInternalServerError, "Activity action doesn't exists!!", "", nil)
}
