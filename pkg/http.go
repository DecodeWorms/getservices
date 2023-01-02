package pkg

import (
	"getservices/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	StatusInternalServerError = http.StatusInternalServerError
	StatusSuccess             = http.StatusOK
	StatusBadRequest          = http.StatusBadRequest
	CodeDataValidationError   = http.StatusUnprocessableEntity
)

type ResponseBody struct {
	Status     int               `json:"status"`
	Success    bool              `json:"success"`
	Data       interface{}       `json:"data"`
	ErrMessage *errors.UserError `json:"code"`
}

func JsonResponse(c *gin.Context, status bool, code int, errorMessage *errors.UserError, data interface{}) {
	var response ResponseBody
	response.Success = status
	response.Status = code
	response.ErrMessage = errorMessage
	response.Data = data
	c.JSON(code, response)
}
