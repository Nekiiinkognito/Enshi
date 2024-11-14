package middleware

import (
	rest_api_stuff "enshi/REST_API_stuff"
	"enshi/middleware/checkRole"
	"enshi/middleware/getters"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, err := getters.GetUserIdFromContext(c)

		if err != nil || userId == 0 {
			rest_api_stuff.BadRequestAnswer(c, err)
			c.Abort()
		}

		isAdmin, err := checkRole.IsAdmin(userId)

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
