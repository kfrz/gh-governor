package auth

import (
	"fmt"

	gh "github.com/cli/go-gh/v2/pkg/api"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/kfrz/gh-governor/internal/repo"
)

// GraphQLClient is an interface that wraps the Query method
type GraphQLClient interface {
	Query(query string, result interface{}, variables map[string]interface{}) error
}

// DefaultClient is the default implementation of the GraphQLClient
type DefaultClient struct{}

func (d *DefaultClient) Query(query string, result interface{}, variables map[string]interface{}) error {
	client, err := gh.DefaultGraphQLClient()
	if err != nil {
		return err
	}
	return client.Query(query, result, variables)
}

var Client GraphQLClient = &DefaultClient{}

// checkAuthStatus checks if the user is authenticated to github.com
// and prints the current user status if so
func CheckAuthStatus(cmd *cobra.Command, args []string) error {
	var query struct {
		Viewer struct {
			Login string
		}
	}

	// Use the Client variable instead of directly calling gh.DefaultGraphQLClient()
	if err := Client.Query("UserCurrent", &query, nil); err != nil {
		zap.L().Error("Authentication failed", zap.Error(err))
		return fmt.Errorf("failed to fetch current user details: %v", err)
	}

	err := repo.PrintCurrentRepoStatus()
	if err != nil {
		zap.L().Error("Failed to get repository status", zap.Error(err))
		return err
	}

	zap.L().Debug("Authenticated to github.com", zap.String("user", query.Viewer.Login))
	return nil
}
