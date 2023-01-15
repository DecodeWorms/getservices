package handler

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
)

type ServiceProviderHandler struct {
	providers storage.ProviderServices
}

func NewServiceProviderHandler(provider storage.ProviderServices) ServiceProviderHandler {
	return ServiceProviderHandler{
		providers: provider,
	}
}

func (providers ServiceProviderHandler) SignUpProvider(ctx *gin.Context, data models.ServiceProviderJson) *errors.UserError {
	_, err := providers.providers.ProviderByEmail(data.Email)
	if err == nil {
		custom := errors.ErrEmailAlreadyExist
		return custom
	}
	_, err = providers.providers.ProviderByPhoneNumber(data.PhoneNumber)
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
	providerId := pkg.GenerateUserId(name)

	resources := models.ServiceProvider{
		ServiceProviderId:  providerId,
		FirstName:          data.FirstName,
		LastName:           data.LastName,
		PhoneNumber:        data.PhoneNumber,
		Email:              trimmedEmail,
		Password:           hashedPassword,
		IsFullyOnboarded:   true,
		IsAccountConfirmed: true,
	}

	err = providers.providers.Create(resources)
	if err != nil {
		custom := errors.ErrCreatingUser
		return custom
	}

	address := models.ServiceProviderAddress{
		ServiceProviderId: resources.ServiceProviderId,
		Name:              data.Address.Name,
		City:              data.Address.City,
		ZipCode:           data.Address.ZipCode,
	}

	if err = providers.providers.CreateAddress(address); err != nil {
		custom := errors.ErrCreatingUser
		return custom
	}

	return nil
}
