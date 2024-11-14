package postRules

import (
	"enshi/middleware/checkRole"

	"github.com/gin-gonic/gin"
)

// Only owner of the post can change it
func PostUpdateRule(c *gin.Context) (bool, error) {
	isOwner, err := checkRole.IsOwnerOfThePost(c)
	if err != nil {
		return false, err
	}

	return isOwner, nil
}
