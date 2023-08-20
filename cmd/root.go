package cmd

import (
	"fmt"

	gh "github.com/cli/go-gh/v2/pkg/api"
	"github.com/cli/go-gh/v2/pkg/repository"
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var RootCmd = &cobra.Command{
	Use:   "governor",
	Short: "🧐 Governor: Your control center for managing governance.",
	Long: `
|\\//|   ,---.  ,---.,--.  ,--.,---. ,--.--.,--,--,  ,---. ,--.--.
|//\\|  | .-. || .-. |\  ''  /| .-. :|  .--'|      \| .-. ||  .--'
|\\//|  ' '-' '' '-' ' \    / \   --.|  |   |  ||  |' '-' '|  |
|//\\|  ..-  /  '---'   '--'   '----''--'   '--''--' '---' '--'
|\\//|  '---'  governance X auditor X enforcer` + "\n\n" + `
Governor is a CLI tool that allows you to select repositories and
interactively apply and audit governance rules. It provides options to
update the CODEOWNERS file, enforce branch naming conventions, check
repository name by pattern, and more. It is designed to be used in
continuous integration pipelines, but can also be used interactively.

See README.md for more information, including usage examples.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	// RunE is the main entry point for the root command.
	RunE: func(cmd *cobra.Command, args []string) error {
		zap.L().Debug("running gh-governor RunE()")
		return nil
	},
}

func Execute() error {
	// Add color to the help command
	cc.Init(&cc.Config{
		RootCmd:  RootCmd,
		Headings: cc.HiCyan + cc.Bold + cc.Underline,
		Commands: cc.HiYellow + cc.Bold,
		Example:  cc.Italic,
		ExecName: cc.HiYellow + cc.Bold + cc.Underline,
		Flags:    cc.Bold,
	})

	if err := checkAuthStatus(RootCmd, nil); err != nil {
		zap.L().Error("error occurred while checking auth status", zap.Error(err))
		return err
	}

	if err := RootCmd.Execute(); err != nil {
		zap.L().Fatal("error occurred while executing root command: %v\\n", zap.Error(err))
		return err
	}
	return nil
}

// checkAuthStatus checks if the user is authenticated to github.com
// and prints the current user status if so
func checkAuthStatus(cmd *cobra.Command, args []string) error {
	client, err := gh.DefaultGraphQLClient()
	if err != nil {
		return fmt.Errorf("failed to create client: %v", err)
	}

	var query struct {
		Viewer struct {
			Login string
		}
	}
	err = client.Query("UserCurrent", &query, nil)
	if err != nil {
		return fmt.Errorf("failed to fetch current user details: %v", err)
	}

	printCurrentRepoStatus()
	zap.L().Info("Authenticated to github.com", zap.String("user", query.Viewer.Login))
	return nil
}

func printCurrentRepoStatus() {
	// repository.Current() respects the value of the GH_REPO environment variable and reads from git remote configuration as fallback.
	repo, err := repository.Current()
	if err != nil {
		zap.L().Fatal("failed to get current repository", zap.Error(err))
		return
	}
	zap.L().Info("Repo status", zap.String("host", repo.Host), zap.String("owner", repo.Owner), zap.String("name", repo.Name))
}
