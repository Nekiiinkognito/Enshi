package rules

import (
	rest_api_stuff "enshi/REST_API_stuff"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShouldAbortRequest(c *gin.Context, isAllowed bool, errors []error) bool {
	var errorsMap = map[int]string{}
	for i, error := range errors {
		errorsMap[i] = error.Error()
	}

	if errors != nil {
		c.IndentedJSON(http.StatusUnauthorized, errorsMap)
		return true
	}

	if !isAllowed {
		rest_api_stuff.UnauthorizedAnswer(c, fmt.Errorf("you have no permission"))
		return true
	}

	return false
}
