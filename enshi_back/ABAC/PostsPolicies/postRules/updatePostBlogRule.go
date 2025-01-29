package postRules

import (
	globalrules "enshi/ABAC/GlobalRules"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

// Only user that own target post and blog can do that
func UpdatePostBlogRule(c *gin.Context) (bool, []error) {
	rulesToCheck := []rules.RuleFunction{
		globalrules.AuthorizedRule,
		globalrules.IsOwnerOfThePostRule,
		globalrules.IsOwnerOfTheBlogRule,
	}

	isAllowed, errors := rules.CheckRules(
		c,
		rulesToCheck,
		rules.ALL_RULES_MUST_BE_COMPLETED,
	)

	return isAllowed, errors
}
