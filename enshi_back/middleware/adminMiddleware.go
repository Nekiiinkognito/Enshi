package middleware

import (
	rest_api_stuff "enshi/REST_API_stuff"
	"enshi/middleware/checkRole"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		isAdmin, err := checkRole.IsAdmin(c)

		if err != nil {
			rest_api_stuff.BadRequestAnswer(c, err)
			c.Abort()
		}

		if !isAdmin {
			rest_api_stuff.UnauthorizedAnswer(c, fmt.Errorf("not allowed"))
			c.Abort()
		}

		c.Next()
	}
}
