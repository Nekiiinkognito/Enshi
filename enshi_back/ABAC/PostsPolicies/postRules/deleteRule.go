package postRules

import (
	globalrules "enshi/ABAC/GlobalRules"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

const RULES_NUMBER_TO_COMPLETE = 2

// Only owner or admin can delete post
func DeleteRule(c *gin.Context) (bool, []error) {
	rulesToCheck := []rules.RuleFunction{
		globalrules.AuthorizedRule,
		globalrules.IsOwnerOfThePostRule,
		globalrules.IsAdminRule,
	}

	isAllowed, errors := rules.CheckRules(
		c,
		rulesToCheck,
		RULES_NUMBER_TO_COMPLETE,
	)

	return isAllowed, errors
}
