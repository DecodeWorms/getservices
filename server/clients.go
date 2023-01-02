package server

import (
	"getservices/errors"
	handler "getservices/handlers"
	"getservices/models"
	"getservices/pkg"
	"net/http"

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
		if err := ctx.ShouldBindJSON(&data); err != nil {
			errors.CustomError(pkg.CodeDataValidationError, err.Error())
			return
		}
		handlerErr := client.clienthandler.UpdateClient(ctx, data)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, nil)
	}

}
