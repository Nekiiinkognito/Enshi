package getters

import (
	"enshi/global"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserIdFromContext(c *gin.Context) (int64, error) {
	userId, exists := c.Get(global.ContextUserId)

	if !exists {
		return -1, fmt.Errorf("error getting user id")
	}

	if parsedUserId, err := strconv.ParseInt(userId.(string), 10, 64); err != nil {
		return -1, fmt.Errorf("error parsing user id")
	} else {
		return parsedUserId, nil
	}
}
