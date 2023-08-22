package mocks

import (
	"errors"
)

type MockQueryError struct{}

func (m *MockQueryError) Query(query string, result interface{}, variables map[string]interface{}) error {
	return errors.New("mocked query error")
}

type MockQuerySuccess struct{}

func (m *MockQuerySuccess) Query(query string, result interface{}, variables map[string]interface{}) error {
	return nil
}
