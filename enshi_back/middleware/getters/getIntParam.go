package getters

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetInt64Param(c *gin.Context, paramName string) (int64, error) {
	int64ParamValue, err := strconv.ParseInt(c.Param(paramName), 10, 64)

	if err != nil {
		return -1, err
	}

	return int64ParamValue, nil
}
