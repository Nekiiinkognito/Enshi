package middleware

import (
	"enshi/ABAC/rules"
	"fmt"

	"github.com/gin-gonic/gin"
)

type WorkRule struct {
	Rules           []rules.RuleFunction
	MustBeCompleted int
}

type Policy func(c *gin.Context) (bool, []error)
type RuleSets map[string]rules.RuleFunction
type RulesToCheck map[string]WorkRule

type MiddlewareProvider struct {
	Policies map[string]Policy
}

func CreateRuleFunction(rulesToCheck []rules.RuleFunction, mustBeCompleted int) rules.RuleFunction {
	return func(c *gin.Context) (bool, []error) {

		isAllowed, errors := rules.CheckRules(
			c,
			rulesToCheck,
			mustBeCompleted,
		)

		return isAllowed, errors
	}
}

func CreatePolicy(ruleSets RuleSets) Policy {

	return func(c *gin.Context) (bool, []error) {
		targetAction, exists := c.Get("target")
		if !exists {
			return false, nil
		}

		for action, rule := range ruleSets {
			if action == targetAction {
				return rules.CheckRule(c, rule)
			}
		}

		return false, nil
	}

}

// Accepts
//
// ruleSetName -> `string` name of the policy(like old one "postPolicy" etc.)
//
// rulesToCheck -> map where keys like ["GET", "POST", etc.] and values are struct of type {rules: [list of rules to check], mustBeCompleted: how many rules must be completed from the list before}
func (m *MiddlewareProvider) RegisterPolicy(
	ruleSetName string,
	rulesToCheck RulesToCheck,
) error {

	for k := range m.Policies {
		if k == ruleSetName {
			return fmt.Errorf("name: " + ruleSetName + " already exists")
		}
	}

	newRuleSets := make(RuleSets)
	for setName, workRule := range rulesToCheck {
		newRuleFunction := CreateRuleFunction(workRule.Rules, workRule.MustBeCompleted)
		newRuleSets[setName] = newRuleFunction
	}

	newPolicy := CreatePolicy(newRuleSets)

	m.Policies[ruleSetName] = newPolicy

	return nil
}

func (m *MiddlewareProvider) GetMiddleware(
	policyName string,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		validName := false
		for key := range m.Policies {
			if key == policyName {
				validName = true
			}
		}
		if !validName {
			c.Abort()
			fmt.Println("invalid policy name: " + policyName)
			return
		}

		isAllowed, errors := m.Policies[policyName](c)

		if rules.ShouldAbortRequest(c, isAllowed, errors) {
			c.Abort()
			return
		}

		c.Next()
	}

}
