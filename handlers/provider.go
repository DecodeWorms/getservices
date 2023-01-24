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

		//parse password
	if b := hashpassword.ParsePassword(data.Password); !b{
		custom := errors.ErrPasswordStrength
		return custom
	}
		//compare password and confirm password
	result := hashpassword.ComparePasswordWithConfirmPassword(data.Password, data.ConfirmPassword)
	if !result {
		return errors.NewUserError(errors.StatusInternalServerError, "password and confirm password not match")
	}

	//hash password from client
	hashedPassword, err := hashpassword.HashPasswordWithGivenCost([]byte(data.Password), hashpassword.MaxCost)
	if err != nil {
		custom := errors.ErrHashingPassword
		return custom
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

func (providers ServiceProviderHandler) LoginProvider(ctx *gin.Context, data models.ServiceProviderLoginJson) (*models.ServiceProviderLoginResponse, *errors.UserError) {
	//validate login data from the json
	v := validations.Validate{Validate: validations.NewVaLidate()}
	valErr := ValidatedData(v, data)
	if len(valErr) > 0 {
		return nil, errors.NewUserError(errors.StatusBadRequest, valErr[0].Error())
	}

	//validate the email existense
	prov, err := providers.providers.ProviderByEmail(data.Email)
	if err != nil {
		custom := errors.ErrResourceNotFound
		return nil, custom
	}

	//validate password
	b := hashpassword.ComparePasswordWithHashed(prov.Password, data.Password)
	if !b {
		custom := errors.ErrValidatingPassword
		return nil, custom
	}

	providerData, err := providers.providers.Login(prov.Email)
	if err != nil {
		custom := errors.ErrLoginUser
		return nil, custom
	}

	providerAddress, err := providers.providers.AddressByProviderId(providerData.ServiceProviderId)
	if err != nil {
		custom := errors.ErrResourceNotFound
		return nil, custom
	}
	result := &models.ServiceProviderLoginResponse{}
	result.ServiceProviderId = providerData.ServiceProviderId
	result.FirstName = providerData.FirstName
	result.LastName = providerData.LastName
	result.PhoneNumber = providerData.PhoneNumber
	result.Email = providerData.Email
	result.IsFullyOnboarded = providerData.IsFullyOnboarded
	result.IsAccountConfirmed = providerData.IsAccountConfirmed
	result.Address.Name = providerAddress.Name
	result.Address.City = providerAddress.City
	result.Address.ZipCode = providerAddress.ZipCode

	return result, nil
}

func (provider ServiceProviderHandler) UpdatePassword(ctx *gin.Context, email string, passwordData models.PasswordJson) *errors.UserError {
	//check if the user data exist
	cli, err := provider.providers.ProviderByEmail(email)
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

	data := &models.ServiceProvider{
		Email:    email,
		Password: hashPassword,
	}

	if err = provider.providers.UpdatePassword(data); err != nil {
		custom := errors.ErrUpdatingUserResource
		return custom
	}
	return nil
}
