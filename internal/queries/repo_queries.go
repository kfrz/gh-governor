package queries

type Codeowners struct {
	Errors []struct {
		Suggestion string `graphql:"suggestion"`
	} `graphql:"errors"`
}

type CodeownersBlob struct {
	Blob struct {
		Text string `graphql:"text"`
	} `graphql:"... on Blob"`
}

type CodeownersQuery struct {
	Repository struct {
		RootCodeowners   CodeownersBlob `graphql:"rootCodeowners: object(expression: \"HEAD:CODEOWNERS\")"`
		GithubCodeowners CodeownersBlob `graphql:"githubCodeowners: object(expression: \"HEAD:.github/CODEOWNERS\")"`
		DocsCodeowners   CodeownersBlob `graphql:"docsCodeowners: object(expression: \"HEAD:docs/CODEOWNERS\")"`
		Codeowners       *Codeowners
	} `graphql:"repository(owner: $owner, name: $repoName)"`
}
