package postRules

import (
	"enshi/middleware/checkRole"

	"github.com/gin-gonic/gin"
)

// Only owner or admin can delete post
func DeleteRule(c *gin.Context) (bool, error) {
	isOwner, err := checkRole.IsOwnerOfThePost(c)
	if err != nil {
		return false, err
	}

	isAdmin, err := checkRole.IsAdmin(c)
	if err != nil {
		return false, err
	}

	return isAdmin || isOwner, err
}
