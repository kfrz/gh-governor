package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"github.com/kfrz/gh-governor/internal/auth"
	"github.com/kfrz/gh-governor/internal/repo"
	"github.com/kfrz/gh-governor/mocks"
)

func TestRootCmd_Execute(t *testing.T) {
	// Use the mock client that simulates a successful query execution
	auth.Client = &mocks.MockQuerySuccess{}
	repo.Repo = &mocks.MockRepoSuccess{}

	err := Execute()
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}
}

func TestExecute_Error(t *testing.T) {
	// Mocking for error scenario
	auth.Client = &mocks.MockQueryError{}
	repo.Repo = &mocks.MockRepoError{}

	err := Execute()
	if err == nil {
		t.Fatalf("Expected an error, but got none")
	}
	// You can further check the error message if needed
}

func TestExecute(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "test execute",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use the mock client that simulates a successful query execution
			auth.Client = &mocks.MockQuerySuccess{}
			repo.Repo = &mocks.MockRepoSuccess{} // Add this line

			buf := new(bytes.Buffer)
			RootCmd.SetOutput(buf)
			err := Execute()
			if (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			zap.L().Info(buf.String())
			assert.Contains(t, buf.String(), "Governor is a CLI tool")
		})
	}
}
