package error

import (
	xtremeres "github.com/globalxtreme/go-core/v2/response"
	"net/http"
)

func ErrXtremeTestingGet(internalMsg string) {
	xtremeres.Error(http.StatusNotFound, "Testing not found", internalMsg, nil)
}

func ErrXtremeTestingSave(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to save testing", internalMsg, nil)
}

func ErrXtremeTestingDelete(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to delete testing", internalMsg, nil)
}

func ErrXtremeTestingSubSave(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to save testing sub", internalMsg, nil)
}

func ErrXtremeTestingSubDelete(internalMsg string) {
	xtremeres.Error(http.StatusInternalServerError, "Unable to delete testing sub", internalMsg, nil)
}
