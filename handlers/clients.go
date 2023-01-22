package handlers

import (
	"fmt"
	"getservices/errors"
	"getservices/hashpassword"
	"getservices/models"
	"getservices/pkg"
	"getservices/storage"
	"getservices/validations"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ClientHandler struct {
	ClientService storage.ClientServices
	//ClientService storage.ClientServices
	//ProviderServices client.ProviderServices
	//Services         client.ServiceServices
}

func NewClientHandler(clientService storage.ClientServices) ClientHandler {
	return ClientHandler{
		ClientService: clientService,
		//ProviderServices: ProviderServices,
		//Services:         Services,
	}
}

func (client ClientHandler) SignUpClient(ctx *gin.Context, data models.ClientJson) *errors.UserError {
	//check if cliet email already exist
	_, err := client.ClientService.ClientByEmail(data.Email)
	if err == nil {
		custom := errors.ErrEmailAlreadyExist
		return custom
	}

	//check if client phone number already exist
	_, err = client.ClientService.ClientByPhoneNumber(data.PhoneNumber)
	if err == nil {
		custom := errors.ErrPhoneNumberAlreadyExist
		return custom
	}

	//trim the email
	trimmedEmail := fmt.Sprint(strings.TrimSpace(data.Email))

	//validate data from the json
	v := validations.Validate{Validate: validations.NewVaLidate()}
	valErr := ValidatedData(v, data)
	if len(valErr) > 0 {
		return errors.NewUserError(errors.StatusBadRequest, valErr[0].Error())
	}

	//hash password from client
	hashedPassword, err := hashpassword.HashPasswordWithGivenCost([]byte(data.Password), hashpassword.MaxCost)
	if err != nil {
		custom := errors.ErrHashingPassword
		return custom
	}

	//compare password and confirm password
	result := hashpassword.ComparePasswordWithConfirmPassword(data.Password, data.ConfirmPassword)
	if !result {
		return errors.NewUserError(errors.StatusInternalServerError, "password and confirm password not match")
	}

	//get user first two characters
	name := pkg.GetUserFirstTwoChar(data.FirstName)
	//generate a unique clientId
	userId := pkg.GenerateUserId(name)

	resources := models.Client{
		ClientId:           userId,
		FirstName:          data.FirstName,
		LastName:           data.LastName,
		PhoneNumber:        data.PhoneNumber,
		Email:              trimmedEmail,
		Password:           hashedPassword,
		IsFullyOnboarded:   true,
		IsAccountConfirmed: true,
	}

	err = client.ClientService.Create(resources)
	if err != nil {
		custom := errors.ErrCreatingUser
		return custom
	}

	address := models.Address{
		ClientId: resources.ClientId,
		Name:     data.Address.Name,
		City:     data.Address.City,
		ZipCode:  data.Address.ZipCode,
	}
	if err = client.ClientService.CreateAddress(address); err != nil {
		custom := errors.ErrCreatingUser
		return custom
	}
	return nil
}

func (client ClientHandler) UserLogin(ctx *gin.Context, data models.ClientLoginJson) (*models.ClientLoginResponse, *errors.UserError) {
	//validate login data from the json
	v := validations.Validate{Validate: validations.NewVaLidate()}
	valErr := ValidatedData(v, data)
	if len(valErr) > 0 {
		return nil, errors.NewUserError(errors.StatusBadRequest, valErr[0].Error())
	}

	//validate the email existense
	cli, err := client.ClientService.ClientByEmail(data.Email)
	if err != nil {
		custom := errors.ErrResourceNotFound
		return nil, custom
	}

	//validate password
	b := hashpassword.ComparePasswordWithHashed(cli.Password, data.Password)
	if !b {
		custom := errors.ErrValidatingPassword
		return nil, custom
	}

	clientData, err := client.ClientService.Login(cli.Email)
	if err != nil {
		custom := errors.ErrLoginUser
		return nil, custom
	}

	clientAddress, err := client.ClientService.AddressByClientId(clientData.ClientId)
	if err != nil {
		custom := errors.ErrResourceNotFound
		return nil, custom
	}
	result := &models.ClientLoginResponse{}
	result.ClientId = clientData.ClientId
	result.FirstName = clientData.FirstName
	result.LastName = clientData.LastName
	result.PhoneNumber = clientData.PhoneNumber
	result.Email = clientData.Email
	result.IsFullyOnboarded = clientData.IsFullyOnboarded
	result.IsAccountConfirmed = clientData.IsAccountConfirmed
	result.Address.Name = clientAddress.Name
	result.Address.City = clientAddress.City
	result.Address.ZipCode = clientAddress.ZipCode

	return result, nil
}

func (client ClientHandler) UpdateClient(ctx *gin.Context, data models.ClientJson) *errors.UserError {
	//TODO : use user id store in the context as an arg instead of email
	cli, err := client.ClientService.ClientByEmail(data.Email)
	if err != nil {
		custom := errors.ErrResourceNotFound
		return custom
	}

	//validate the json data
	v := validations.Validate{Validate: validations.NewVaLidate()}
	valErr := ValidatedData(v, data)
	if len(valErr) > 0 {
		return errors.NewUserError(errors.StatusBadRequest, valErr[0].Error())
	}

	//update database resources
	resources := models.Client{
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Email:       strings.ToLower(data.Email),
		PhoneNumber: data.PhoneNumber,
	}
	err = client.ClientService.Update(cli.ClientId, resources)
	if err != nil {
		custom := errors.ErrUpdatingUserResource
		return custom
	}

	addResources := models.Address{
		Name:    data.Address.Name,
		City:    data.Address.City,
		ZipCode: data.Address.ZipCode,
	}

	if err := client.ClientService.UpdateAddress(cli.ClientId, addResources); err != nil {
		custom := errors.ErrUpdatingUserResource
		return custom
	}
	return nil

}

func (client ClientHandler) DeactivateClientAccount(ctx *gin.Context, email string) *errors.UserError {
	cli, err := client.ClientService.ClientByEmail(email)
	if err != nil {
		custom := errors.ErrResourceNotFound
		return custom
	}
	if err = client.ClientService.DeactivateAccount(cli.ClientId); err != nil {
		custom := errors.ErrDeactivatingResource
		return custom
	}
	return nil
}

func (client ClientHandler) ActivateAccount(ctx *gin.Context, email string) *errors.UserError {
	cli, err := client.ClientService.GetDeletedAgentByEmail(email)
	if err != nil {
		custom := errors.ErrResourceNotFound
		return custom
	}
	if err = client.ClientService.ActivateAccount(cli.ClientId); err != nil {
		custom := errors.ErrActivatingResource
		return custom
	}
	return nil

}

func (client ClientHandler) UpdateClientPassword(ctx *gin.Context, email string, passwordData models.PasswordJson) *errors.UserError {
	//check if the user data exist
	cli, err := client.ClientService.ClientByEmail(email)
	if err != nil {
		custom := errors.ErrResourceNotFound
		return custom
	}
	//check if the new password is equivalent to the existing password
	b := hashpassword.ComparePasswordWithHashed(cli.Password, passwordData.NewPassword)
	if b {
		custom := errors.ErrExistingPassword
		return custom
	}

	//compare new passowrd with the confirm password
	result := hashpassword.ComparePasswordWithConfirmPassword(passwordData.NewPassword, passwordData.ConfirmNewPassword)
	if !result {
		return errors.NewUserError(errors.StatusInternalServerError, "password and confirm password not match")
	}

	//hash the password
	hashPassword, err := hashpassword.HashPasswordWithGivenCost([]byte(passwordData.NewPassword), hashpassword.DefaultCost)
	if err != nil {
		custom := errors.ErrHashingPassword
		return custom
	}

	data := &models.Client{
		Email:    email,
		Password: hashPassword,
	}

	if err = client.ClientService.UpdatePassword(data); err != nil {
		custom := errors.ErrUpdatingUserResource
		return custom
	}

	return nil
}

func ValidatedData(v validations.Validate, data interface{}) []error {
	errDetails := make([]error, 0)

	err := v.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			e := errors.New(fmt.Sprintf("user_data object: a valid %v of type %v is required, but recieved '%v' ", strings.ToLower(err.Field()), err.Kind(), err.Value()))
			errDetails = append(errDetails, e)
		}
		return errDetails
	}

	return errDetails
}
