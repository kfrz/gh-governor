package repo

import (
	"errors"
	"testing"

	"github.com/cli/go-gh/v2/pkg/repository"
)

// Mock for the Repository that simulates a successful response
type MockRepoSuccess struct{}

func (m *MockRepoSuccess) Current() (repository.Repository, error) {
	return repository.Repository{
		Host:  "github.com",
		Owner: "test-owner",
		Name:  "test-repo",
	}, nil
}

func TestPrintCurrentRepoStatus(t *testing.T) {
	// Use the mock that simulates a successful response
	Repo = &MockRepoSuccess{}

	err := PrintCurrentRepoStatus()
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
}

type MockRepoError struct{}

func (m *MockRepoError) Current() (repository.Repository, error) {
	return repository.Repository{}, errors.New("mocked error")
}

func TestPrintCurrentRepoStatus_Error(t *testing.T) {
	// Use the mock that returns an error
	Repo = &MockRepoError{}

	err := PrintCurrentRepoStatus()
	if err.Error() != "mocked error" {
		t.Fatalf("Expected 'mocked error', but got: %v", err)
	}
}
