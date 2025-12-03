package cmd

import (
	"fmt"
	"os"

	"github.com/alompson/lazy-git/internal/service"
	"github.com/spf13/cobra"
)

var (
	commitService *service.CommitService
)

func SetCommitService(svc *service.CommitService) {
	commitService = svc
}

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Generate a git commit message based on your code changes",
	Long:  `This command analyzes your git diffs and suggests a meaningful commit message using AI.`,
	Run:   commitRunner,
}

func init() {
	RootCmd.AddCommand(commitCmd)
}

func commitRunner(cmd *cobra.Command, args []string) {
	if commitService == nil {
		fmt.Fprintln(os.Stderr, "Error: service not initialized")
		os.Exit(1)
	}

	message, err := commitService.GenerateCommitMessage(cmd.Context())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(message)
}
