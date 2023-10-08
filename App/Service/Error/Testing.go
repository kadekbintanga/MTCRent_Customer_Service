package Error

import (
	"github.com/globalxtreme/gobaseconf/response"
	"net/http"
)

func ErrXtremeTestingSave(internalMsg string) {
	response.Error(http.StatusInternalServerError, "Unable to save testing", internalMsg, nil)
}

func ErrXtremeTestingSubSave(internalMsg string) {
	response.Error(http.StatusInternalServerError, "Unable to save testing sub", internalMsg, nil)
}
