package cmd

import (
	"testing"

	"github.com/kfrz/gh-governor/internal/client"
	"github.com/kfrz/gh-governor/internal/repo"
	"github.com/kfrz/gh-governor/mocks"
)

func TestRootCmd_Execute(t *testing.T) {
	client.GraphQL = &mocks.MockQuerySuccess{}
	repo.Repo = &mocks.MockRepoSuccess{}

	err := Execute()
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
}

func TestExecute_Error(t *testing.T) {
	client.GraphQL = &mocks.MockQueryError{}
	repo.Repo = &mocks.MockRepoError{}

	err := Execute()
	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}
}
