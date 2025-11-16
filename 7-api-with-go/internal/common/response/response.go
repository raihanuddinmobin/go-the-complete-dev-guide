package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Page       int `json:"page,omitempty"`
	Limit      int `json:"limit,omitempty"`
	TotalItems int `json:"totalItems,omitempty"`
	TotalPages int `json:"totalPages,omitempty"`
}

type APIResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty"`
}

type APIError struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Code    string      `json:"code,omitempty"`
	Details interface{} `json:"details,omitempty"`
}

func Success(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func SuccessWithPagination(ctx *gin.Context, message string, data interface{}, meta *Meta) {
	ctx.JSON(http.StatusOK, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func Error(ctx *gin.Context, statusCode int, message string, code string, details ...interface{}) {
	err := APIError{
		Status:  "error",
		Message: message,
		Code:    code,
	}

	if len(details) > 0 {
		err.Details = details[0]
	}

	ctx.JSON(statusCode, err)
}
