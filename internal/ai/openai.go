package ai

import (
	"context"
	"fmt"

	openai "github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

// OpenAIProvider implements Provider using OpenAI's API
type OpenAIProvider struct {
	client *openai.Client
	model  openai.ChatModel
}

// NewOpenAIProvider creates a new OpenAI provider
func NewOpenAIProvider(cfg Config) (Provider, error) {
	if cfg.APIKey == "" {
		return nil, fmt.Errorf("API key is required")
	}

	client := openai.NewClient(
		option.WithAPIKey(cfg.APIKey),
	)

	// Default to latest model if not specified
	model := openai.ChatModelChatgpt4oLatest
	if cfg.Model != "" {
		model = openai.ChatModel(cfg.Model)
	}

	return &OpenAIProvider{
		client: &client,
		model:  model,
	}, nil
}

// GenerateCommitMessage generates a commit message from a git diff
func (p *OpenAIProvider) GenerateCommitMessage(ctx context.Context, diff string) (string, error) {
	if diff == "" {
		return "", fmt.Errorf("diff is empty")
	}

	prompt := `Generate a concise and meaningful git commit message based on the following code diffs. 
You should give the answer as "git commit -m "<type>: <message>"" where type is one of: feat, fix, docs, style, refactor, test, chore.

Code diffs:
` + diff

	resp, err := p.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: p.model,
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate commit message: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from AI model")
	}

	return resp.Choices[0].Message.Content, nil
}

// GenerateBranchName generates a branch name from a git diff
func (p *OpenAIProvider) GenerateBranchName(ctx context.Context, diff string) (string, error) {
	if diff == "" {
		return "", fmt.Errorf("diff is empty")
	}

	prompt := `Generate a short, descriptive git branch name based on the following code diffs.
Use kebab-case format (e.g., "feature-user-authentication" or "fix-login-bug").
Maximum 50 characters. Only return the branch name, nothing else.

Code diffs:
` + diff

	resp, err := p.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: p.model,
	})
	if err != nil {
		return "", fmt.Errorf("failed to generate branch name: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from AI model")
	}

	return resp.Choices[0].Message.Content, nil
}
