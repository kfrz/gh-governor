package mocks

import (
	"errors"

	"github.com/cli/go-gh/v2/pkg/repository"
)

type MockRepoSuccess struct{}

func (m *MockRepoSuccess) Current() (repository.Repository, error) {
	return repository.Repository{
		Host:  "github.com",
		Owner: "test-owner",
		Name:  "test-repo",
	}, nil
}

type MockRepoError struct{}

func (m *MockRepoError) Current() (repository.Repository, error) {
	return repository.Repository{}, errors.New("mocked error")
}
