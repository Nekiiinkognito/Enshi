package middleware

import (
	profilepolicies "enshi/ABAC/ProfilePolicies"
	rest_api_stuff "enshi/REST_API_stuff"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProfileMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.Request.Method {
		case "PUT":
			c.Set("target", profilepolicies.UPDATE_PROFILE)
		}

		isAllowed, errors := profilepolicies.ProfilePolicies(c)

		var errorsMap = map[int]string{}
		for i, error := range errors {
			errorsMap[i] = error.Error()
		}

		if errors != nil {
			c.IndentedJSON(http.StatusUnauthorized, errorsMap)
			c.Abort()
			return
		}

		if !isAllowed {
			rest_api_stuff.UnauthorizedAnswer(c, fmt.Errorf("you have no permission"))
			c.Abort()
			return
		}

		c.Next()
	}
}
