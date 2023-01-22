package errors

import "net/http"

type UserError struct {
	Code        int    `json:"erro"`
	Description string `json:"description"`
}

func NewUserError(c int, descrip string) *UserError {
	return &UserError{
		Code:        c,
		Description: descrip,
	}
}

func CustomError(c int, description string) *UserError {
	return &UserError{
		Code:        c,
		Description: description,
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

	ErrResourceNotFound          = NewUserError(StatusNotFound, "error resource not found")                      // error resource not found usually from the server
	ErrBadRequest                = NewUserError(StatusBadRequest, "error bad request")                           //error from the client
	ErrForbiddenRequest          = NewUserError(http.StatusForbidden, "error forbit request")                    //server error, unauthorized request
	ErrUnauthorizedRequest       = NewUserError(StatusUnAuthorized, "error unauthoriyed request")                // server error, request lack or expired token
	ErrTooManyRequest            = NewUserError(StatusTooManyRequests, "error too many request")                 // client error , too many request at a given time .
	ErrStatusConflict            = NewUserError(StatusConflict, "error status conflicts")                        // client error, made a request for a target resource and make same request again
	ErrStatusInternalServerError = NewUserError(StatusInternalServerError, "error internal server error")        //server error, probably too much time processing request
	ErrEmailAlreadyExist         = NewUserError(http.StatusForbidden, "error user email  already exist")         //client error, user email already exist
	ErrPhoneNumberAlreadyExist   = NewUserError(http.StatusForbidden, "error phone number already exist")        //client error , user phone number already exist
	ErrHashingPassword           = NewUserError(StatusInternalServerError, "error hashing password")             //server error, unable to hash user password
	ErrValidattingUserData       = NewUserError(http.StatusBadRequest, "error validating user data")             //client error , error validating user error
	ErrCreatingUser              = NewUserError(http.StatusInternalServerError, "error creating user")           //error creating user
	ErrCreatingServices          = NewUserError(http.StatusInternalServerError, "error creating service")        //error creating service
	ErrCreatingAddress           = NewUserError(http.StatusInternalServerError, "error creating address")        //server error in creating user address
	ErrValidatingPassword        = NewUserError(http.StatusInternalServerError, "error validating password")     //server error , error validating password
	ErrLoginUser                 = NewUserError(http.StatusInternalServerError, "error login in user")           //server error ,error login in user
	ErrUpdatingUserResource      = NewUserError(http.StatusInternalServerError, "error updating user resources") //server error, error updating user resources
	ErrResourceIsDeactivated     = NewUserError(http.StatusBadRequest, "error resource is deactivated")          //client error , resource already deactivated
	ErrDeactivatingResource      = NewUserError(http.StatusInternalServerError, "error deactivating an account") //server error, error deactivating an account
	ErrActivatingResource        = NewUserError(http.StatusInternalServerError, "error activating resource")     //server error, error activating resources
	ErrExistingPassword          = NewUserError(http.StatusBadRequest, "error existing password")                //client error , error new password equal to existing password
	ErrUnknownCategory = NewUserError(http.StatusBadRequest, "error parsing service category") //client error, error parsing servoce category
	ErrCompanyNameAlreadyExist = NewUserError(http.StatusBadRequest, "error company name is registered") //client error , error company name is already exist
)
