package middleware

import (
	"github.com/gin-gonic/gin"
)

func ContentTypeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		accept := c.GetHeader("Accept")

		switch accept {
		case "application/xml":
			c.Set("responseType", "xml")
		case "application/json", "*/*":
			c.Set("responseType", "json")
		default:
			c.String(406, "Content type not supported")
			c.Abort()
			return
		}

		c.Next()
	}
}
