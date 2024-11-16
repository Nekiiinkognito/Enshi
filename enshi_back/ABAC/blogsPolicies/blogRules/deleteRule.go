package blogrules

import (
	globalrules "enshi/ABAC/globalRules"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

func BlogDeleteRule(c *gin.Context) (bool, []error) {
	rulesToCheck := []rules.RuleFunction{
		globalrules.AuthorizedRule,
		globalrules.IsOwnerOfTheBlogRule,
		globalrules.IsAdminRule,
	}

	isAllowed, errors := rules.CheckRules(
		c,
		rulesToCheck,
		2,
	)

	return isAllowed, errors
}
