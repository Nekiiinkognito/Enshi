package blogrules

import (
	globalrules "enshi/ABAC/globalRules"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

func BlogCreateRule(c *gin.Context) (bool, []error) {
	rulesToCheck := []rules.RuleFunction{
		globalrules.AuthorizedRule,
	}

	isAllowed, errors := rules.CheckRules(
		c,
		rulesToCheck,
		rules.ALL_RULES_MUST_BE_COMPLETED,
	)

	return isAllowed, errors
}
