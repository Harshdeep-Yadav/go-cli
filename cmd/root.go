package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ghcli",
	Short: "GitHub CLI - fetch GitHub issues",
	Long:  `A command-line tool to fetch GitHub issues for any public repo.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use a subcommand like `issues`")
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
