package repo

import (
	"fmt"

	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/fatih/color"
	"go.uber.org/zap"

	_ "github.com/kfrz/gh-governor/config"
)

type Repository interface {
	Current() (repository.Repository, error)
}
type DefaultRepository struct{}

var Repo Repository = &DefaultRepository{}

// PrintCurrentRepoStatus prints the current repository status
// Repo.Current() respects the value of the GH_REPO environment variable and reads from git remote configuration as fallback.
func PrintCurrentRepoStatus() error {
	repo, err := Repo.Current()
	if err != nil {
		zap.L().Debug("failed to get current repository", zap.Error(err))
		fmt.Println("💡 Try running 'gh repo' or navigating to a valid git repository.")
		return err
	}
	zap.L().Debug("Repo status", zap.String("host", repo.Host), zap.String("owner", repo.Owner), zap.String("name", repo.Name))
	fmt.Printf("🚀 Current repository: %s/%s\n", color.GreenString(repo.Owner), color.GreenString(repo.Name))
	return nil
}

func (dr *DefaultRepository) Current() (repository.Repository, error) {
	return repository.Current()
}
