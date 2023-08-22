// internal/auth/auth_test.go

package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kfrz/gh-governor/internal/client"
	"github.com/kfrz/gh-governor/internal/repo"
	"github.com/kfrz/gh-governor/mocks"
)

func TestCheckAuthStatus_Error_Query(t *testing.T) {
	// Use the mock client that simulates a query execution error
	client.GraphQL = &mocks.MockQueryError{}

	err := CheckAuthStatus(client.GraphQL)
	if err == nil || err.Error() != "failed to fetch current user details: mocked query error" {
		t.Fatalf("Expected 'failed to fetch current user details: mocked query error', but got: %v", err)
	}
}

func TestCheckAuthStatus_Success(t *testing.T) {
	// Use the mock client that simulates a successful query execution
	client.GraphQL = &mocks.MockQuerySuccess{}

	// Mock repo.PrintCurrentRepoStatus to not return an error
	mockRepo := &mocks.MockRepoSuccess{}
	repo.Repo = mockRepo

	if err := CheckAuthStatus(client.GraphQL); err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
}

func TestCheckAuthStatus_NoAuth(t *testing.T) {
	// Set the client to the mock that simulates a query error (no authentication)
	client.GraphQL = &mocks.MockQueryError{}

	err := CheckAuthStatus(client.GraphQL)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to fetch current user details")
}

func TestCheckAuthStatus_Authenticated(t *testing.T) {
	// Set the client to the mock that simulates a successful query
	client.GraphQL = &mocks.MockQuerySuccess{}

	err := CheckAuthStatus(client.GraphQL)
	assert.NoError(t, err)
}
