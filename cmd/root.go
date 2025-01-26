package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github-activity-cli/util"
)

var RootCmd = &cobra.Command{
	Use:   "github-actvity",
	Short: "github-activity is a CLI tool to get the latest activity of a user on GitHub",
	Long:  "github-activity is a CLI tool to get the latest activity of a user on GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		activities := util.GetActivity(username)
		for _, activity := range activities {
			fmt.Println(activity.Type)
			fmt.Println(activity.Repo.Name)
			fmt.Println()

		}
	},
}
