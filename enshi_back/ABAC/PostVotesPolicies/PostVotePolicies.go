package postvotespolicies

import (
	postvoterules "enshi/ABAC/PostVotesPolicies/PostVoteRules"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

const (
	DELETE_VOTE = "delete_vote"
	CREATE_VOTE = "create_vote"
	READ_VOTE   = "read_vote"
)

func PostVotePolicies(c *gin.Context) (bool, []error) {
	target, exists := c.Get("target")
	if !exists {
		return false, nil
	}

	// Permit if one permit
	switch target {
	case DELETE_VOTE:
		return rules.CheckRule(c, postvoterules.PostVoteDeleteRule)

	case CREATE_VOTE:
		return rules.CheckRule(c, postvoterules.PostVoteCreateRule)

	case READ_VOTE:
		return rules.CheckRule(c, postvoterules.PostVoteReadRule)

	}

	return false, nil
}
