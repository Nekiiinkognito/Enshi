package middleware

import "github.com/gin-gonic/gin"

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Next()
	}
}
