package response

import (
	"encoding/xml"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Meta struct {
	Page       int `json:"page" xml:"page"`
	Limit      int `json:"limit" xml:"limit"`
	TotalItems int `json:"totalItems" xml:"totalItems"`
	TotalPages int `json:"totalPages" xml:"totalPages"`
}

type APIResponse struct {
	XMLName xml.Name    `xml:"response" json:"-"`
	Status  string      `json:"status" xml:"status"`
	Message string      `json:"message,omitempty" xml:"message,omitempty"`
	Data    interface{} `json:"data,omitempty" xml:"data,omitempty"`
	Meta    *Meta       `json:"meta,omitempty" xml:"meta,omitempty"`
}

type APIError struct {
	Status  string      `json:"status" xml:"status"`
	Message string      `json:"message" xml:"message"`
	Code    string      `json:"code,omitempty" xml:"code,omitempty"`
	TraceID string      `json:"traceId,omitempty" xml:"traceId,omitempty"`
	Details interface{} `json:"details,omitempty" xml:"details,omitempty"`
}

func isXML(c *gin.Context) bool {
	v, ok := c.Get("responseType")
	return ok && v == "xml"
}

func respond(c *gin.Context, status int, payload any) {
	if isXML(c) {
		c.XML(status, payload)
		return
	}
	c.JSON(status, payload)
}

func Success(c *gin.Context, message string, data interface{}, statusCodes ...int) {

	statusCode := http.StatusOK
	if len(statusCodes) > 0 {
		statusCode = statusCodes[0]
	}

	respond(c, statusCode, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

func SuccessWithPagination(c *gin.Context, message string, data interface{}, meta *Meta) {
	respond(c, http.StatusOK, APIResponse{
		Status:  "success",
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func Error(c *gin.Context, statusCode int, message, traceID, code string, details ...interface{}) {
	err := APIError{
		Status:  "error",
		Message: message,
		Code:    code,
		TraceID: traceID,
	}
	if len(details) > 0 {
		err.Details = details[0]
	}

	respond(c, statusCode, err)
}
