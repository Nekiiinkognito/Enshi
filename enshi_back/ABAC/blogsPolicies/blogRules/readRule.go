package blogrules

import (
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

func BlogReadRule(c *gin.Context) (bool, []error) {
	rulesToCheck := []rules.RuleFunction{}

	isAllowed, errors := rules.CheckRules(
		c,
		rulesToCheck,
		rules.ALL_RULES_MUST_BE_COMPLETED,
	)

	return isAllowed, errors
}
