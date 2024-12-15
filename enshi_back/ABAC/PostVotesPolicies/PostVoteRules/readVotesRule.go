package postvoterules

import (
	"github.com/gin-gonic/gin"
)

func PostVotesReadRule(c *gin.Context) (bool, []error) {
	// rulesToCheck := []rules.RuleFunction{}

	// isAllowed, errors := rules.CheckRules(
	// 	c,
	// 	rulesToCheck,
	// 	rules.ALL_RULES_MUST_BE_COMPLETED,
	// )

	return true, nil
}
