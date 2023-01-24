package handlers

import (
	"fmt"
	"getservices/errors"
	"getservices/models"
	"getservices/pkg"
	"getservices/storage"
	"getservices/validations"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

type ServiceHandler struct {
	Service  storage.ServiceServices
	Provider storage.ProviderServices
}

func NewServiceHandler(serv storage.ServiceServices, prov storage.ProviderServices) ServiceHandler {
	return ServiceHandler{
		Service:  serv,
		Provider: prov,
	}
}

func (service ServiceHandler) CreateService(ctx *gin.Context, email string, data models.ServiceJson) *errors.UserError {
	//check if the user by email exist
	provider, err := service.Provider.ProviderByEmail(email)
	if err != nil {
		custom := errors.ErrResourceNotFound
		return custom
	}

	//validate if the phone number is unique
	_, err = service.Service.ServiceByPhoneNumber(data.PhoneNumber)
	if err == nil {
		custom := errors.ErrPhoneNumberAlreadyExist
		return custom
	}

	//parse category
	if b := pkg.ParseCategory(data.Service); !b{
		custom := errors.ErrUnknownCategory
		return custom
	}

	//check if the company name is unique
	_ , err = service.Service.ServiceByCompanyName(data.CompanyName)
	if err == nil{
		custom := errors.ErrCompanyNameAlreadyExist
		return custom
	}

	//validate the JSON data
	v := validations.Validate{Validate: validations.NewVaLidate()}
	valErr := ValidatedData(v, data)
	if len(valErr) > 0 {
		return errors.NewUserError(errors.StatusBadRequest, valErr[0].Error())
	}

	//create service by service provider
	resource := models.Services{
		ServiceProviderId: provider.ServiceProviderId,
		PhoneNumber:       data.PhoneNumber,
		YearOfExperience:  data.YearOfExperience,
		Service:           data.Service,
		CompanyName:       data.CompanyName,
	}
	if err = service.Service.Create(resource); err != nil {
		custom := errors.ErrCreatingServices
		return custom
	}

	//creating service address
	serviceAddress := models.ServiceAddress{
		ServiceProviderId: provider.ServiceProviderId,
		Name:              data.Address.Name,
		City:              data.Address.City,
		ZipCode:           data.Address.ZipCode,
	}

	if err = service.Service.CreateAddress(serviceAddress); err != nil {
		custom := errors.ErrCreatingAddress
		return custom
	}
	return nil
}

func(service ServiceHandler)GetServicesCategories(ctx *gin.Context)[]string{
	return pkg.ServiceCategory
}

func(service ServiceHandler)GetServices(ctx *gin.Context, serviceName string)([]*models.ServiceProviderDetail, *errors.UserError){
	result := make([]*models.ServiceProviderDetail , 0)

	services , err := service.Service.ServiceByService(serviceName)
	if err != nil{
		custom := errors.ErrResourceNotFound
		return nil ,custom
	}

	for _ , serv := range services{
		res := &models.ServiceProviderDetail{}
		var g errgroup.Group
		g.Go(func() error {
			provid , err := service.Provider.Provider(serv.ServiceProviderId)
			if err != nil{
				return err
			}
			res.PhoneNumber = provid.PhoneNumber
			res.FullName = fmt.Sprintf("%s %s",provid.FirstName,provid.LastName)
			res.Email = provid.Email
			return nil
		})

		g.Go(func() error {
			add , err := service.Service.AddressByProviderId(serv.ServiceProviderId)
			if err != nil{
				return err
			}
			res.CompanyName = serv.CompanyName
			res.CompanyPhoneNumber = serv.PhoneNumber
			res.Service = serv.Service
			res.AddressName = add.Name
			res.AddressCity = add.City
			return nil

		})
		if err := g.Wait(); err != nil{
			return nil , errors.ErrResourceNotFound
		}
		result = append(result, res)
	}
	return result , nil
}

func(service ServiceHandler)GetService(ctx *gin.Context ,email string)(*models.ServiceProviderDetail, *errors.UserError){
	res := &models.ServiceProviderDetail{}
	prov , err := service.Provider.ProviderByEmail(email)
	if err != nil{
		custom := errors.ErrResourceNotFound
		return nil, custom
	}
	
	var g errgroup.Group
	g.Go(func() error {
		serv , err := service.Service.Service(prov.ServiceProviderId)
		if err != nil{
			return err
		}
		res.CompanyName = serv.CompanyName
		res.Service = serv.Service
		res.CompanyPhoneNumber = serv.PhoneNumber
		return nil
	})

	g.Go(func() error {
		add , err := service.Service.AddressByProviderId(prov.ServiceProviderId)
		if err != nil{
			return err
		}
	     res.AddressName = add.Name
	     res.AddressCity = add.City
		 return nil
	})

	if err := g.Wait(); err !=nil{
		return nil , errors.ErrResourceNotFound
	}
	res.FullName = fmt.Sprintf("%s %s",prov.FirstName, prov.LastName)
	res.PhoneNumber = prov.PhoneNumber
	res.Email = prov.Email

	return res, nil
}
