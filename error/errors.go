package error

import "net/http"

type UserError struct {
	Code        int    `json:"erro"`
	Description string `json:"description"`
}

func NewUserError(c int, descrip string) UserError {
	return UserError{
		Code:        c,
		Description: descrip,
	}
}

var (
	StatusForbidden           = http.StatusForbidden
	StatusBadRequest          = http.StatusBadRequest
	StatusUnAuthorized        = http.StatusUnauthorized
	StatusNotFound            = http.StatusNotFound
	StatusTooManyRequests     = http.StatusTooManyRequests
	StatusConflict            = http.StatusConflict
	StatusInternalServerError = http.StatusInternalServerError

	ErrResourceNotFound          = NewUserError(StatusNotFound, "error resource not found")               // error resource not found usually from the server
	ErrBadRequest                = NewUserError(StatusBadRequest, "error bad request")                    //error from the client
	ErrForbiddenRequest          = NewUserError(http.StatusForbidden, "error forbit request")             //server error, unauthorized request
	ErrUnauthorizedRequest       = NewUserError(StatusUnAuthorized, "error unauthoriyed request")         // server error, request lack or expired token
	ErrTooManyRequest            = NewUserError(StatusTooManyRequests, "error too many request")          // client error , too many request at a given time .
	ErrStatusConflict            = NewUserError(StatusConflict, "error status conflicts")                 // client error, made a request for a target resource and make same request again
	ErrStatusInternalServerError = NewUserError(StatusInternalServerError, "error internal server error") //server error, probably too much time processing request
)
