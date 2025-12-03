package ai

import "context"

// Provider defines the interface for AI-based text generation
type Provider interface {
	GenerateCommitMessage(ctx context.Context, diff string) (string, error)
	GenerateBranchName(ctx context.Context, diff string) (string, error)
}

type Config struct {
	APIKey string
	Model  string
}
