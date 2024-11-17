package getters

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Returns -1, error if there is no such param or value invalid
func GetInt64Param(c *gin.Context, paramName string) (int64, error) {
	int64ParamValue, err := strconv.ParseInt(strings.Trim(c.Param(paramName), "/"), 10, 64)

	if err != nil {
		return -1, err
	}

	return int64ParamValue, nil
}
