package service

import (
	"context"
	"fmt"

	"github.com/alompson/lazygit/internal/ai"
	"github.com/alompson/lazygit/internal/git"
)

// CommitService handles commit message generation logic
type CommitService struct {
	gitRepo    git.Repository
	aiProvider ai.Provider
}

// NewCommitService creates a new commit service
func NewCommitService(repo git.Repository, provider ai.Provider) *CommitService {
	return &CommitService{
		gitRepo:    repo,
		aiProvider: provider,
	}
}

// GenerateCommitMessage generates a commit message based on staged changes
func (s *CommitService) GenerateCommitMessage(ctx context.Context) (string, error) {
	// Check if we're in a git repository
	if !s.gitRepo.IsGitRepository() {
		return "", fmt.Errorf("not a git repository")
	}

	// Get staged changes
	diff, err := s.gitRepo.GetStagedDiff()
	if err != nil {
		return "", fmt.Errorf("failed to get staged changes: %w", err)
	}

	// Check if there are any changes
	if !diff.HasChanges {
		return "", fmt.Errorf("no staged changes found. Use 'git add' to stage changes")
	}

	// Generate commit message using AI
	message, err := s.aiProvider.GenerateCommitMessage(ctx, diff.Content)
	if err != nil {
		return "", fmt.Errorf("failed to generate commit message: %w", err)
	}

	return message, nil
}

// GenerateBranchName generates a branch name based on working directory changes
func (s *CommitService) GenerateBranchName(ctx context.Context) (string, error) {
	// Check if we're in a git repository
	if !s.gitRepo.IsGitRepository() {
		return "", fmt.Errorf("not a git repository")
	}

	// Get working directory changes (can be staged or unstaged)
	diff, err := s.gitRepo.GetStagedDiff()
	if err != nil {
		return "", fmt.Errorf("failed to get changes: %w", err)
	}

	// If no staged changes, try unstaged
	if !diff.HasChanges {
		diff, err = s.gitRepo.GetWorkingDiff()
		if err != nil {
			return "", fmt.Errorf("failed to get working changes: %w", err)
		}
	}

	// Check if there are any changes
	if !diff.HasChanges {
		return "", fmt.Errorf("no changes found")
	}

	// Generate branch name using AI
	branchName, err := s.aiProvider.GenerateBranchName(ctx, diff.Content)
	if err != nil {
		return "", fmt.Errorf("failed to generate branch name: %w", err)
	}

	return branchName, nil
}
