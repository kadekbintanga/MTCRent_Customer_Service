package error

import (
	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"net/http"
)

func ErrXtremeActivityNotFound() {
	xtremeres.Error(http.StatusNotFound, "Service not found", "", nil)
}

func ErrXtremeActivitySave(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to save service", internalMsg, nil)
}

func ErrXtremeActivityActionType() {
	xtremeres.Error(http.StatusInternalServerError, "Activity action doesn't exists!!", "", nil)
}
