package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yjydist/github-activity-cli/util"
)

var RootCmd = &cobra.Command{
	Use:   "github-actvity",
	Short: "github-activity is a CLI tool to get the latest activity of a user on GitHub",
	Long:  "github-activity is a CLI tool to get the latest activity of a user on GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Error: Please provide a GitHub username.")
			return
		}
		username := args[0]
		activities, err := util.GetActivity(username)
		if err != nil {
			fmt.Printf("Error fetching activities for user '%s': %v\n", username, err)
			return
		}
		if len(activities) == 0 {
			fmt.Printf("No activities found for user '%s'.\n", username)
			return
		}
		fmt.Printf("Latest activities for user '%s':\n", username)
		fmt.Println("===================================")
		for _, activity := range activities {
			fmt.Printf("Type: %s\n", activity.Type)
			fmt.Printf("Repository: %s\n", activity.Repo.Name)
			fmt.Println("-----------------------------------")
		}
	},
}
