// Package cmd contains the root command and all subcommands.
package cmd

import (
	"fmt"

	cc "github.com/ivanpirog/coloredcobra"
	"github.com/shurcooL/githubv4"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	_ "github.com/kfrz/gh-governor/config"
	"github.com/kfrz/gh-governor/internal/auth"
	"github.com/kfrz/gh-governor/internal/client"
	"github.com/kfrz/gh-governor/internal/queries"
	"github.com/kfrz/gh-governor/internal/repo"
)

var RootCmd = &cobra.Command{
	Use:   "governor",
	Short: "Governor: Your control center for managing governance.",
	Long: `
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
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				zap.L().Error("error occurred while running cmd.Help()", zap.Error(err))
				return err
			}
		}
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
		return err
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	currentRepo, err := repo.Repo.Current()
	if err != nil {
		return err
	}

	owner := currentRepo.Owner
	repoName := currentRepo.Name

	codeownersContents, err := CheckCodeownersLocations(owner, repoName)
	if err != nil {
		return err
	}

	for path, content := range codeownersContents {
		zap.L().Info(content)
		zap.L().Info(path)
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

// CheckCodeownersLocations checks the locations of the CODEOWNERS files and returns their content
// TODO: this should probably be moved to a different package, and abstracted such that
// it can be used by other commands looking for a set of files in a repo.. but that is
// breaking the pattern of having a CodeownersQuery struct, which would make things more complicated
// to refactor, we need to make sure our queries/schemas don't break... so for now, we'll leave it here.
func CheckCodeownersLocations(owner string, repoName string) (map[string]string, error) {
	if repoName == "" {
		repo, err := repo.Repo.Current()
		if err != nil {
			return nil, err
		}
		repoName = repo.Name
	}

	variables := map[string]interface{}{
		"owner":    githubv4.String(owner),
		"repoName": githubv4.String(repoName),
	}

	// Query uses pointer to a go struct
	query := &queries.CodeownersQuery{}

	err := client.GraphQL.Query("CodeownersQuery", query, variables)
	if err != nil {
		zap.L().Fatal("failed while querying")
	}

	githubCodeownersText := ""
	if query.Repository.GithubCodeowners.Blob.Text != "" {
		githubCodeownersText = query.Repository.GithubCodeowners.Blob.Text
	}

	docsCodeownersText := ""
	if query.Repository.DocsCodeowners.Blob.Text != "" {
		docsCodeownersText = query.Repository.DocsCodeowners.Blob.Text
	}

	rootCodeownersText := ""
	if query.Repository.RootCodeowners.Blob.Text != "" {
		rootCodeownersText = query.Repository.RootCodeowners.Blob.Text
	}

	results := map[string]string{
		"HEAD:.github/CODEOWNERS": githubCodeownersText,
		"HEAD:docs/CODEOWNERS":    docsCodeownersText,
		"HEAD:CODEOWNERS":         rootCodeownersText,
	}

	for key, value := range results {
		if value == "" {
			results[key] = "🚫 file not present"
		} else {
			results[key] = results[key] + "\n✅ File is present"
		}
	}

	if query.Repository.Codeowners != nil {
		for _, err := range query.Repository.Codeowners.Errors {
			if err.Suggestion != "" {
				zap.L().Info("suggestion 😅: ", zap.String("error", err.Suggestion))
			}
		}
	} else {
		zap.L().Info("✅ no codeowners errors found")
	}

	return results, nil
}
