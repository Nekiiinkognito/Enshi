package rules

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RuleFunction func(*gin.Context) (bool, []error)

const (
	ALL_RULES_MUST_BE_COMPLETED = iota
)

func CheckRule(
	c *gin.Context,
	ruleChecker RuleFunction,
) (bool, []error) {
	IsAllowed, err := ruleChecker(c)
	if err != nil {
		return false, err
	}

	return IsAllowed, nil
}

func CheckRules(
	c *gin.Context,
	rules []RuleFunction,
	completedRulesCount int,
) (bool, []error) {
	var allowancesIndexes []int
	var errors []error

	if len(rules) < completedRulesCount {
		return false, []error{fmt.Errorf("there is less rules, that should be completed")}
	}

	for i, rule := range rules {
		if isAllowed, err := CheckRule(c, rule); err != nil {
			errors = append(
				errors,
				err...,
			)
		} else if !isAllowed {
			errors = append(
				errors,
				fmt.Errorf("rule "+
					strconv.Itoa(i)+
					" was rejected"),
			)
		} else {
			allowancesIndexes = append(allowancesIndexes, i)
		}
	}

	switch completedRulesCount {
	case ALL_RULES_MUST_BE_COMPLETED:
		if len(allowancesIndexes) == len(rules) {
			return true, nil
		} else {
			return false, errors
		}
	default:
		if len(allowancesIndexes) >= completedRulesCount {
			return true, nil
		} else {
			return false, errors
		}
	}
}
