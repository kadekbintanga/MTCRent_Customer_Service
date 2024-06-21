package error

import (
	"github.com/globalxtreme/gobaseconf/response"
	"net/http"
)

func ErrXtremeTestingGet(internalMsg string) {
	response.Error(http.StatusNotFound, "Testing not found", internalMsg, nil)
}

func ErrXtremeTestingSave(internalMsg string) {
	response.Error(http.StatusInternalServerError, "Unable to save testing", internalMsg, nil)
}

func ErrXtremeTestingDelete(internalMsg string) {
	response.Error(http.StatusInternalServerError, "Unable to delete testing", internalMsg, nil)
}

func ErrXtremeTestingSubSave(internalMsg string) {
	response.Error(http.StatusInternalServerError, "Unable to save testing sub", internalMsg, nil)
}

func ErrXtremeTestingSubDelete(internalMsg string) {
	response.Error(http.StatusInternalServerError, "Unable to delete testing sub", internalMsg, nil)
}
