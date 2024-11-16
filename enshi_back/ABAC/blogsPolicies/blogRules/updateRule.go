package blogrules

import (
	globalrules "enshi/ABAC/GlobalRules"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

func BlogUpdateRule(c *gin.Context) (bool, []error) {
	rulesToCheck := []rules.RuleFunction{
		globalrules.AuthorizedRule,
		globalrules.IsOwnerOfTheBlogRule,
	}

	isAllowed, errors := rules.CheckRules(
		c,
		rulesToCheck,
		rules.ALL_RULES_MUST_BE_COMPLETED,
	)

	return isAllowed, errors
}
