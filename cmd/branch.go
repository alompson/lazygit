package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Generate a git branch name based on your code changes",
	Long:  `This command analyzes your git diffs and suggests a meaningful branch name using AI.`,
	Run:   branchRunner,
}

func init() {
	RootCmd.AddCommand(branchCmd)
}

func branchRunner(cmd *cobra.Command, args []string) {
	if commitService == nil {
		fmt.Fprintln(os.Stderr, "Error: service not initialized")
		os.Exit(1)
	}

	branchName, err := commitService.GenerateBranchName(cmd.Context())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(branchName)
}
