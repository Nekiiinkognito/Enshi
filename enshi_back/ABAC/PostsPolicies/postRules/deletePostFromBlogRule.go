package postRules

import (
	globalrules "enshi/ABAC/GlobalRules"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

func DeletePostFromBlogRule(c *gin.Context) (bool, []error) {
	rulesToCheck := []rules.RuleFunction{
		globalrules.AuthorizedRule,
		globalrules.IsOwnerOfThePostRule,
		globalrules.IsOwnerOfTheBlogRule,
	}

	isAllowed, errors := rules.CheckRules(
		c,
		rulesToCheck,
		RULES_NUMBER_TO_COMPLETE,
	)

	return isAllowed, errors
}
