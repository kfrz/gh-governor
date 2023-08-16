package main

import (
	"log"

	gh "github.com/cli/go-gh/v2/pkg/api"
	"github.com/spf13/cobra"
)


func main() {
	// Initialize Cobra for CLI structure
	rootCmd := &cobra.Command{
		Use:   "check-user",
		Short: "Check if current user is authenticated and print username",
		RunE:   checkCurrentUser,
	}

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func checkCurrentUser(cmd *cobra.Command, args []string) error {
	client, err := gh.DefaultGraphQLClient()
	if err != nil {
		return err
	}

	var query struct {
		Viewer struct {
			Login string
		}
	}
	err = client.Query("UserCurrent", &query, nil)
	if err != nil {
		return err
	}
	log.Printf("Logged in as %s", query.Viewer.Login)
	return nil
}