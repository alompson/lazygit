package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Repository defines operations on a git repository
type Repository interface {
	GetStagedDiff() (*DiffResult, error)
	GetWorkingDiff() (*DiffResult, error)
	IsGitRepository() bool
}

type DiffResult struct {
	Content    string
	FilesCount int
	HasChanges bool
}

type LocalRepository struct {
}

func NewLocalRepository() Repository {
	return &LocalRepository{}
}

// GetStagedDiff returns the diff of staged changes
func (r *LocalRepository) GetStagedDiff() (*DiffResult, error) {
	cmd := exec.Command("git", "diff", "--staged")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to get staged diff: %w (stderr: %s)", err, stderr.String())
	}

	content := out.String()
	return &DiffResult{
		Content:    content,
		HasChanges: len(content) > 0,
		FilesCount: countFiles(content),
	}, nil
}

func (r *LocalRepository) GetWorkingDiff() (*DiffResult, error) {
	cmd := exec.Command("git", "diff")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("failed to get working diff: %w (stderr: %s)", err, stderr.String())
	}

	content := out.String()
	return &DiffResult{
		Content:    content,
		HasChanges: len(content) > 0,
		FilesCount: countFiles(content),
	}, nil
}

// IsGitRepository checks if the current directory is inside a git repository
func (r *LocalRepository) IsGitRepository() bool {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	return cmd.Run() == nil
}

func countFiles(diff string) int {
	if diff == "" {
		return 0
	}
	// Simple count of "diff --git" occurrences
	count := 0
	for i := 0; i < len(diff)-10; i++ {
		if diff[i:i+10] == "diff --git" {
			count++
		}
	}
	return count
}
