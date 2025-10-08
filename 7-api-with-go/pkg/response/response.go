package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mobin.dev/internals/dtos"
)

func ResponseList[T any, M any](ctx *gin.Context, status int, success bool, msg string, data T, meta M) {
	ctx.JSON(status, dtos.ApiResponseList[T, M]{
		Success: success,
		Data:    data,
		Message: msg,
		Meta:    &meta,
	})
}

func RespondSingle[T any](ctx *gin.Context, status int, success bool, msg string, data *T) {
	ctx.JSON(status, dtos.ApiResponseSingle[T]{
		Success: success,
		Message: msg,
		Data:    data,
	})
}

func OK[T any](ctx *gin.Context, msg string, data T) {
	RespondSingle(ctx, http.StatusOK, true, msg, &data)
}

func Error(ctx *gin.Context, status int, msg string) {
	RespondSingle[any](ctx, status, false, msg, nil)
}
