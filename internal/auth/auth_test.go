// internal/auth/auth_test.go

package auth

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"

	"github.com/kfrz/gh-governor/internal/repo"
	"github.com/kfrz/gh-governor/mocks"
)

func TestCheckAuthStatus_Error_Query(t *testing.T) {
	// Use the mock client that simulates a query execution error
	Client = &mocks.MockQueryError{}

	err := CheckAuthStatus(nil, nil)
	if err == nil || err.Error() != "failed to fetch current user details: mocked query error" {
		t.Fatalf("Expected 'failed to fetch current user details: mocked query error', but got: %v", err)
	}
}

// Mock for the GraphQL client that simulates a successful query execution
type mockQuerySuccess struct{}

func (m *mockQuerySuccess) Query(query string, result interface{}, variables map[string]interface{}) error {
	return nil
}

func TestCheckAuthStatus_Success(t *testing.T) {
	// Use the mock client that simulates a successful query execution
	Client = &mockQuerySuccess{}

	// Mock repo.PrintCurrentRepoStatus to not return an error
	mockRepo := &mocks.MockRepoSuccess{}
	repo.Repo = mockRepo

	err := CheckAuthStatus(&cobra.Command{}, nil)
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
}

func TestCheckAuthStatus_NoAuth(t *testing.T) {
	// Set the client to the mock that simulates a query error (no authentication)
	Client = &mocks.MockQueryError{}

	cmd := &cobra.Command{}
	err := CheckAuthStatus(cmd, nil)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to fetch current user details")
}

func TestCheckAuthStatus_Authenticated(t *testing.T) {
	// Set the client to the mock that simulates a successful query
	Client = &mocks.MockQuerySuccess{}

	cmd := &cobra.Command{}
	err := CheckAuthStatus(cmd, nil)

	assert.NoError(t, err)
}
