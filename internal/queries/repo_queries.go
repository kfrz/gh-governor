package queries

// TODO: this query needs fixing, check the explorer
type RepoCodeOwnersQuery struct {
	Repository struct {
		Object struct {
			Text string `graphql:"text"`
		} `graphql:"object(expression: \"HEAD:.github/CODEOWNERS\")"`
	} `graphql:"repository(id: $nodeID)"`
}
