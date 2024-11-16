package postRules

import (
	globalrules "enshi/ABAC/globalRules"
	"enshi/middleware/checkRole"

	"github.com/gin-gonic/gin"
)

// Only owner of the post can change it
func PostUpdateRule(c *gin.Context) (bool, error) {
	// Sender should be authorized
	isAuthorized, err := globalrules.AuthorizedRule(c)
	if err != nil {
		return false, err
	} else if !isAuthorized {
		return false, nil
	}

	isOwner, err := checkRole.IsOwnerOfThePost(c)
	if err != nil {
		return false, err
	}

	return isOwner, nil
}
