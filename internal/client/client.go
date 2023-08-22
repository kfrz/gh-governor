// Package client provides interfaces and default implementations for making GraphQL and REST API calls to GitHub.
package client

import (
	"net/http"

	gh "github.com/cli/go-gh/v2/pkg/api"
)

type GraphQLClient interface {
	Query(query string, result interface{}, variables map[string]interface{}) error
}

type RESTClient interface {
	Request(method, path string, body interface{}) (*http.Response, error)
	Get(path string, result interface{}) error
}

type DefaultGraphQLClient struct{}

type DefaultRESTClient struct {
	client *gh.RESTClient
}

func NewDefaultRESTClient() (*DefaultRESTClient, error) {
	client, err := gh.DefaultRESTClient()
	if err != nil {
		return nil, err
	}
	return &DefaultRESTClient{client: client}, nil
}

func (d *DefaultGraphQLClient) Query(query string, result interface{}, variables map[string]interface{}) error {
	client, err := gh.DefaultGraphQLClient()
	if err != nil {
		return err
	}
	return client.Query(query, result, variables)
}

func (d *DefaultRESTClient) Request(method, path string, body interface{}) (*http.Response, error) {
	return d.client.Request(method, path, nil)
}

func (d *DefaultRESTClient) Get(path string, result interface{}) error {
	return d.client.Get(path, result)
}

var GraphQL GraphQLClient = &DefaultGraphQLClient{}

var REST RESTClient
