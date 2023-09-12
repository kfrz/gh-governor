package auth

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/kfrz/gh-governor/internal/client"
	"github.com/kfrz/gh-governor/internal/queries"
	"github.com/kfrz/gh-governor/internal/repo"
)

// checkAuthStatus checks if the user is authenticated via the gh cli
func CheckAuthStatus(client client.GraphQLClient) error {
	query := queries.UserCurrentQuery{}
	if err := client.Query("UserCurrent", &query, nil); err != nil {
		zap.L().Debug("Authentication failed", zap.Error(err))
		fmt.Println("💡 Try running 'gh auth status'")
		return fmt.Errorf("failed to fetch current user details: %v", err)
	}

	err := repo.PrintCurrentRepoStatus()
	if err != nil {
		zap.L().Debug("Failed to get repository status", zap.Error(err))
		return fmt.Errorf("failed while checking repo status: %w", err)
	}

	zap.L().Info("🔐 Authenticated to github", zap.String("user", query.Viewer.Login))
	return nil
}
