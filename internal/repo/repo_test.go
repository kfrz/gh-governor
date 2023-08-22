package repo

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/kfrz/gh-governor/mocks"
)

func TestPrintCurrentRepoStatus(t *testing.T) {
	// Use the mock that simulates a successful response
	Repo = &mocks.MockRepoSuccess{}

	err := PrintCurrentRepoStatus()
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
}

func TestPrintCurrentRepoStatus_Success(t *testing.T) {
	// Use the mock that simulates a successful response
	Repo = &mocks.MockRepoSuccess{}

	err := PrintCurrentRepoStatus()

	assert.NoError(t, err, "Expected no error for successful repo status retrieval")
}

func TestPrintCurrentRepoStatus_Error(t *testing.T) {
	// Use the mock that returns an error
	Repo = &mocks.MockRepoError{}

	err := PrintCurrentRepoStatus()

	assert.Error(t, err, "Expected an error for failed repo status retrieval")
	assert.Contains(t, err.Error(), "mocked error", "Expected error message to contain 'mocked error'")
}
