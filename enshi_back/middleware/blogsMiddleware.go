package middleware

import (
	blogspolicies "enshi/ABAC/blogsPolicies"
	rest_api_stuff "enshi/REST_API_stuff"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BlogsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		switch c.Request.Method {
		case "DELETE":
			c.Set("target", blogspolicies.DELETE_BLOG)
		case "PUT":
			c.Set("target", blogspolicies.UPDATE_BLOG)
		case "POST":
			c.Set("target", blogspolicies.CREATE_BLOG)
		case "GET":
			c.Set("target", blogspolicies.GET_BLOG)
		}

		isAllowed, errors := blogspolicies.BlogPolicies(c)

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
