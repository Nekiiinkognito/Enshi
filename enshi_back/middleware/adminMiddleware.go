package middleware

import (
	"context"
	rest_api_stuff "enshi/REST_API_stuff"
	db_repo "enshi/db/go_queries"
	"enshi/db_connection"
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

		user, err :=
			db_repo.New(db_connection.Dbx).
				GetUserById(context.Background(), userId)

		if err != nil || user.UserID == 0 {
			rest_api_stuff.BadRequestAnswer(c, err)
			c.Abort()
		}

		if !user.IsAdmin {
			rest_api_stuff.UnauthorizedAnswer(c, fmt.Errorf("not allowed"))
			c.Abort()
		}

		c.Next()
	}
}
