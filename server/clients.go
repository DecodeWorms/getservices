package server

import (
	"net/http"

	"github.com/DecodeWorms/getservices/errors"
	handler "github.com/DecodeWorms/getservices/handlers"
	"github.com/DecodeWorms/getservices/models"
	"github.com/DecodeWorms/getservices/pkg"

	"github.com/gin-gonic/gin"
)

type ClientServer struct {
	clienthandler handler.ClientHandler
}

func NewClientServer(clientHandler handler.ClientHandler) ClientServer {
	return ClientServer{
		clienthandler: clientHandler,
	}
}

func (client ClientServer) SignUpClient() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.ClientJson
		if err := ctx.ShouldBindJSON(&data); err != nil {
			errors.CustomError(pkg.CodeDataValidationError, err.Error())
			return
		}
		handlerErr := client.clienthandler.SignUpClient(ctx, data)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, nil)
	}
}

func (client ClientServer) UserLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.ClientLoginJson
		if err := ctx.ShouldBindJSON(&data); err != nil {
			errors.CustomError(pkg.CodeDataValidationError, err.Error())
			return
		}
		clients, handlerErr := client.clienthandler.UserLogin(ctx, data)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, clients)

	}
}

func (client ClientServer) UpdateClient() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.ClientJson
		id := ctx.Query("id")
		if err := ctx.ShouldBindJSON(&data); err != nil {
			errors.CustomError(pkg.CodeDataValidationError, err.Error())
			return
		}
		handlerErr := client.clienthandler.UpdateClient(ctx, id, data)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, nil)
	}

}

func (client ClientServer) Clients() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clients, handlerErr := client.clienthandler.Clients(ctx)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, clients)

	}
}

func (client ClientServer) DeactivateAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")
		handlerErr := client.clienthandler.DeactivateClientAccount(ctx, id)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, nil)

	}
}

func (client ClientServer) ActivateAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")
		handlerErr := client.clienthandler.ActivateAccount(ctx, id)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, nil)

	}
}

func (client ClientServer) UpdateClientPassword() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var passwordData models.PasswordJson
		email := ctx.Query("email")
		if err := ctx.ShouldBindJSON(&passwordData); err != nil {
			errors.CustomError(pkg.CodeDataValidationError, err.Error())
			return
		}
		handlerErr := client.clienthandler.UpdateClientPassword(ctx, email, passwordData)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, nil)
	}
}
