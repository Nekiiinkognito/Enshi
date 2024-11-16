package postRules

import (
	globalrules "enshi/ABAC/globalRules"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

// Only owner of the post can change it
func PostUpdateRule(c *gin.Context) (bool, []error) {
	rulesToCheck := []rules.RuleFunction{
		globalrules.AuthorizedRule,
		globalrules.IsOwnerOfThePostRule,
	}

	isAllowed, errors := rules.CheckRules(
		c,
		rulesToCheck,
		rules.ALL_RULES_MUST_BE_COMPLETED,
	)

	return isAllowed, errors
}
