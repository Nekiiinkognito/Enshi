package bookmarkspolicies

import (
	bookmarksrules "enshi/ABAC/BookmarkPolicies/bookmarkRules"
	"enshi/ABAC/rules"

	"github.com/gin-gonic/gin"
)

const (
	DELETE_BOOKMARK = "delete_bookmark"
	CREATE_BOOKMARK = "create_bookmark"
	READ_BOOKMARK   = "read_bookmark"
)

func BlogPolicies(c *gin.Context) (bool, []error) {
	target, exists := c.Get("target")
	if !exists {
		return false, nil
	}

	// Permit if one permit
	switch target {
	case DELETE_BOOKMARK:
		return rules.CheckRule(c, bookmarksrules.BookmarkDeleteRule)

	case CREATE_BOOKMARK:
		return rules.CheckRule(c, bookmarksrules.BookmarkCreateRule)

	case READ_BOOKMARK:
		return rules.CheckRule(c, bookmarksrules.BookmarkReadRule)

	}

	return false, nil
}
