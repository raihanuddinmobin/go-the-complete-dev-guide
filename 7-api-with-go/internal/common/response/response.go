package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type meta struct {
	Page         int `json:"page,omitempty"`
	Limit        int `json:"limit,omitempty"`
	TotalItems   int `json:"totalItems,omitempty"`
	TotalPerPage int `json:"totalPerPage,omitempty"`
}

type apiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *meta       `json:"meta,omitempty"`
}

type apiError struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode,omitempty"`
	ErrorCode  string `json:"errorCode,omitempty"`
	Message    string `json:"message,omitempty"`
}

func Success(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, apiResponse{
		Status:  "Success",
		Message: message,
		Data:    data,
	})
}

func SuccessWithPagination(ctx *gin.Context, message string, data interface{}, meta *meta) {
	ctx.JSON(http.StatusOK, apiResponse{
		Status:  "Success",
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func Error(ctx *gin.Context, statusCode int, message, errorCode string) {
	ctx.JSON(statusCode, apiError{
		Status:     "error",
		StatusCode: statusCode,
		ErrorCode:  errorCode,
		Message:    message,
	})
}
