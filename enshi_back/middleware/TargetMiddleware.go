package middleware

import (
	"github.com/gin-gonic/gin"
)

const (
	GET    = "GET"
	PUT    = "PUT"
	POST   = "POST"
	DELETE = "DELETE"
)

func TargetMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		switch c.Request.Method {
		case "DELETE":
			c.Set("target", DELETE)
		case "PUT":
			c.Set("target", PUT)
		case "POST":
			c.Set("target", POST)
		case "GET":
			c.Set("target", DELETE)
		}

		c.Next()
	}
}
