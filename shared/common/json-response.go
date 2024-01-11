package common

import (
	"net/http"
	"todo-app/shared/shared-model"
	"github.com/gin-gonic/gin"
)

func SendErrorResponse(ctx *gin.Context, code int, message string) {
	ctx.JSON(code, shared-model.Status{
		Code:    code,
		Message: message,
	})
}

func SendPagedResponse(ctx *gin.Context, data interface{}, paging shared-model.Paging, message string) {
	ctx.JSON(http.StatusOK, shared-model.PagedResponse{
		Status: shared-model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data:   data,
		Paging: paging,
	})
}

func SendSingleResponse(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusOK, shared-model.SingleResponse{
		Status: shared-model.Status{
			Code:    http.StatusOK,
			Message: message,
		},
		Data: data,
	})
}

func SendCreatedResponse(ctx *gin.Context, data interface{}, message string) {
	ctx.JSON(http.StatusCreated, shared-model.SingleResponse{
		Status: shared-model.Status{
			Code:    http.StatusCreated,
			Message: message,
		},
		Data: data,
	})
}
