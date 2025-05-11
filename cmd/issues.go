package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"

	"github.com/harshdeep/ghcli/github"
	"github.com/spf13/cobra"
)

var repo string
var state string
var pages int

var issuesCmd = &cobra.Command{
	Use:   "issues",
	Short: "Fetch GitHub issues for a repository",
	Run: func(cmd *cobra.Command, args []string) {
		if repo == "" {
			fmt.Println("You must specify a repository using --repo")
			os.Exit(1)
		}

		issues, err := github.FetchIssues(repo, state, pages)
		if err != nil {
			fmt.Println("Error fetching issues:", err)
			os.Exit(1)
		}

		for _, issue := range issues {
			color.New(color.FgRed).Printf("#%d ", issue.Number)
			color.New(color.FgGreen).Printf("%s\n", issue.Title)
			color.New(color.FgCyan).Println(issue.URL, "\n")
		}
		
		
	},

}
func init() {
	RootCmd.AddCommand(issuesCmd)

	issuesCmd.Flags().StringVarP(&repo, "repo", "r", "", "Repository in the format owner/repo")
	issuesCmd.Flags().StringVarP(&state, "state", "s", "open", "State of the issues (open, closed)")
	issuesCmd.Flags().IntVar(&pages, "pages", 1, "Number of pages to fetch (30 issues per page)")

}