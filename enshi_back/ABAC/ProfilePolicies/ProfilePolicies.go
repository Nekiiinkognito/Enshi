package profilepolicies

import (
	profilesrules "enshi/ABAC/ProfilePolicies/ProfilesRules"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

const (
	RESET_PROFILE  = "reset_profile"
	UPDATE_PROFILE = "update_profile"
	CREATE_PROFILE = "create_profile"
	GET_PROFILE    = "get_profile"
)

func ProfilePolicies(c *gin.Context) (bool, []error) {
	target, exists := c.Get("target")
	if !exists {
		return false, nil
	}

	// Permit if one permit
	switch target {
	case UPDATE_PROFILE:
		return rules.CheckRule(c, profilesrules.UpdateProfileRule)

	}

	return false, nil
}
