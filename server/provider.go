package server

import (
	"net/http"

	"github.com/DecodeWorms/getservices/errors"
	handler "github.com/DecodeWorms/getservices/handlers"
	"github.com/DecodeWorms/getservices/models"
	"github.com/DecodeWorms/getservices/pkg"

	"github.com/gin-gonic/gin"
)

type ProviderServer struct {
	provider handler.ServiceProviderHandler
}

func NewProviderServer(provider handler.ServiceProviderHandler) ProviderServer {
	return ProviderServer{
		provider: provider,
	}
}

func (provider ProviderServer) SignUpProvider() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.ServiceProviderJson
		if err := ctx.ShouldBindJSON(&data); err != nil {
			errors.CustomError(pkg.CodeDataValidationError, err.Error())
			return
		}
		handlerErr := provider.provider.SignUpProvider(ctx, data)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, nil)

	}
}

func (provider ProviderServer) LoginProvider() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.ServiceProviderLoginJson
		if err := ctx.ShouldBindJSON(&data); err != nil {
			errors.CustomError(pkg.CodeDataValidationError, err.Error())
			return
		}
		clients, handlerErr := provider.provider.LoginProvider(ctx, data)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, clients)

	}
}

func (provider ProviderServer) UpdatePassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var passwordData models.PasswordJson
		email := ctx.Query("email")
		if err := ctx.ShouldBindJSON(&passwordData); err != nil {
			errors.CustomError(pkg.CodeDataValidationError, err.Error())
			return
		}
		handlerErr := provider.provider.UpdatePassword(ctx, email, passwordData)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, nil)
	}
}

func (provider ProviderServer) ServiceProviders() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		providers, handlerErr := provider.provider.ServiceProviders(ctx)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, providers)

	}
}
