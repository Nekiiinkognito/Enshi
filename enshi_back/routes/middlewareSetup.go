package routes

import (
	globalrules "enshi/ABAC/GlobalRules"
	"enshi/ABAC/rules"
	"enshi/middleware"
)

const (
	POST_MIDDLEWARE       = "POST_MIDDLEWARE"
	BLOG_MIDDLEWARE       = "BLOG_MIDDLEWARE"
	PROFILE_MIDDLEWARE    = "PROFILE_MIDDLEWARE"
	BOOKMARK_MIDDLEWARE   = "BOOKMARK_MIDDLEWARE"
	POST_BLOG_MIDDLEWARE  = "POST_BLOG_MIDDLEWARE"
	POST_VOTE_MIDDLEWARE  = "POST_VOTE_MIDDLEWARE"
	POST_VOTES_MIDDLEWARE = "POST_VOTES_MIDDLEWARE"
)

var MiddlewareProvider = middleware.MiddlewareProvider{
	Policies: make(map[string]middleware.Policy),
}

var policiesToRegister = map[string]middleware.RulesToCheck{
	POST_MIDDLEWARE: {
		middleware.GET: {
			Rules:           make([]rules.RuleFunction, 0),
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},
		middleware.POST: {
			Rules: []rules.RuleFunction{
				globalrules.AuthorizedRule,
			},
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},
		middleware.PUT: {
			Rules: []rules.RuleFunction{
				globalrules.AuthorizedRule,
				globalrules.IsOwnerOfThePostRule,
			},
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},
		middleware.DELETE: {
			Rules: []rules.RuleFunction{
				globalrules.AuthorizedRule,
				globalrules.IsOwnerOfThePostRule,
				globalrules.IsAdminRule,
			},
			MustBeCompleted: 2,
		},
	},

	BOOKMARK_MIDDLEWARE: {
		middleware.GET: {
			Rules: []rules.RuleFunction{
				globalrules.AuthorizedRule,
			},
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},
		middleware.DELETE: {
			Rules: []rules.RuleFunction{
				globalrules.AuthorizedRule,
			},
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},
		middleware.POST: {
			Rules: []rules.RuleFunction{
				globalrules.AuthorizedRule,
			},
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},
	},

	BLOG_MIDDLEWARE: {
		middleware.GET: {
			Rules:           make([]rules.RuleFunction, 0),
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},
		middleware.POST: {
			Rules: []rules.RuleFunction{
				globalrules.AuthorizedRule,
			},
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},
		middleware.PUT: {
			Rules: []rules.RuleFunction{
				globalrules.AuthorizedRule,
				globalrules.IsOwnerOfTheBlogRule,
			},
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},
		middleware.DELETE: {
			Rules: []rules.RuleFunction{
				globalrules.AuthorizedRule,
				globalrules.IsOwnerOfTheBlogRule,
				globalrules.IsAdminRule,
			},
			MustBeCompleted: 2,
		},
	},

	POST_VOTE_MIDDLEWARE: {
		middleware.GET: {
			Rules: []rules.RuleFunction{
				globalrules.AuthorizedRule,
			},
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},
		middleware.POST: {
			Rules: []rules.RuleFunction{
				globalrules.AuthorizedRule,
			},
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},
		middleware.DELETE: {
			Rules:           make([]rules.RuleFunction, 0),
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},
	},

	POST_VOTES_MIDDLEWARE: {
		middleware.GET: {
			Rules:           make([]rules.RuleFunction, 0),
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},
	},

	PROFILE_MIDDLEWARE: {
		middleware.PUT: {
			Rules: []rules.RuleFunction{
				globalrules.AuthorizedRule,
			},
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},
	},

	POST_BLOG_MIDDLEWARE: {
		middleware.PUT: {
			Rules: []rules.RuleFunction{
				globalrules.AuthorizedRule,
				globalrules.IsOwnerOfThePostRule,
				globalrules.IsOwnerOfTheBlogRule,
			},
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},

		middleware.DELETE: {
			Rules: []rules.RuleFunction{
				globalrules.AuthorizedRule,
				globalrules.IsOwnerOfThePostRule,
				globalrules.IsOwnerOfTheBlogRule,
			},
			MustBeCompleted: rules.ALL_RULES_MUST_BE_COMPLETED,
		},
	},
}

func InitMiddlewareProvider() {
	MiddlewareProvider.InitMiddlewareProvider(policiesToRegister)
}
