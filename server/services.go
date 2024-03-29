package server

import (
	"net/http"

	"github.com/DecodeWorms/getservices/errors"
	"github.com/DecodeWorms/getservices/handlers"
	"github.com/DecodeWorms/getservices/models"
	"github.com/DecodeWorms/getservices/pkg"

	"github.com/gin-gonic/gin"
)

type ServiceServer struct {
	service handlers.ServiceHandler
}

func NewServiceServer(serv handlers.ServiceHandler) ServiceServer {
	return ServiceServer{
		service: serv,
	}
}

func (service ServiceServer) CreateService() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var jsonData models.ServiceJson
		id := ctx.Query("id")

		if err := ctx.ShouldBindJSON(&jsonData); err != nil {
			errors.CustomError(pkg.CodeDataValidationError, err.Error())
			return
		}
		handlerErr := service.service.CreateService(ctx, id, jsonData)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, nil)
	}

}

func (service ServiceServer) UpdateService() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var data models.ServiceJson
		id := ctx.Query("id")

		if err := ctx.ShouldBindJSON(&data); err != nil {
			errors.CustomError(pkg.CodeDataValidationError, err.Error())
			return
		}

		if handlerErr := service.service.UpdateService(ctx, id, data); handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, nil, nil)
	}
}

func (service ServiceServer) GetServicesCategories() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data := service.service.GetServicesCategories(ctx)
		pkg.JsonResponse(ctx, true, http.StatusOK, nil, data)

	}
}

func (service ServiceServer) GetServices() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		servic := ctx.Query("service")
		data, handlerErr := service.service.GetServices(ctx, servic)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, data)

	}
}

func (service ServiceServer) GetService() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")
		data, handlerErr := service.service.GetService(ctx, id)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, data)
	}
}

func (service ServiceServer) UpdateAddress() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")
		var data models.ServiceAddressJson
		if err := ctx.ShouldBindJSON(&data); err != nil {
			errors.CustomError(pkg.CodeDataValidationError, err.Error())
			return
		}
		handlerErr := service.service.UpdateAddress(ctx, id, data)
		if handlerErr != nil {
			pkg.JsonResponse(ctx, false, handlerErr.Code, handlerErr, nil)
			return
		}
		pkg.JsonResponse(ctx, true, http.StatusOK, handlerErr, nil)

	}
}
