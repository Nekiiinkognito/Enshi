package adminpolicies

import (
	globalrules "enshi/ABAC/GlobalRules"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

func AdminPolicies(c *gin.Context) (bool, []error) {
	rulesToCheck := []rules.RuleFunction{
		globalrules.AuthorizedRule,
		globalrules.IsAdminRule,
	}

	return rules.CheckRules(c, rulesToCheck, rules.ALL_RULES_MUST_BE_COMPLETED)
}
