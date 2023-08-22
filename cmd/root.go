// Package cmd contains the root command and all subcommands.
package cmd

import (
	cc "github.com/ivanpirog/coloredcobra"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	_ "github.com/kfrz/gh-governor/config"
	"github.com/kfrz/gh-governor/internal/auth"
	"github.com/kfrz/gh-governor/internal/client"
	"github.com/kfrz/gh-governor/internal/repo"
)

var RootCmd = &cobra.Command{
	Use:   "governor",
	Short: "Governor: Your control center for managing governance.",
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
		// if len(args) == 0 {
		// 	err := cmd.Help()
		// 	if err != nil {
		// 		zap.L().Error("error occurred while running cmd.Help()", zap.Error(err))
		// 		return err
		// 	}
		// }
		zap.L().Debug("running gh-governor RunE()")
		if err := repo.PrintCurrentRepoStatus(); err != nil {
			zap.L().Error("error occurred while fetching repo status", zap.Error(err))
			return err
		}

		return nil
	},
}

func Execute() error {
	if err := auth.CheckAuthStatus(client.GraphQL); err != nil {
		zap.L().Error("error occurred while checking auth status", zap.Error(err))
		return err
	}

	if err := RootCmd.Execute(); err != nil {
		zap.L().Fatal("error occurred while executing root command: %v\\n", zap.Error(err))
		return err
	}
	return nil
}

func init() {
	// Add color to the help command
	cc.Init(&cc.Config{
		RootCmd:  RootCmd,
		Headings: cc.HiCyan + cc.Bold + cc.Underline,
		Commands: cc.HiYellow + cc.Bold,
		Example:  cc.Italic,
		ExecName: cc.HiYellow + cc.Bold + cc.Underline,
		Flags:    cc.Bold,
	})
}
