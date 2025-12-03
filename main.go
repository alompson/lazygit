/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"log"
	"path/filepath"

	"github.com/alompson/lazygit/cmd"
	"github.com/alompson/lazygit/internal/ai"
	"github.com/alompson/lazygit/internal/config"
	"github.com/alompson/lazygit/internal/git"
	"github.com/alompson/lazygit/internal/service"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	if execPath, err := filepath.Abs("."); err == nil {
		_ = godotenv.Load(filepath.Join(execPath, ".env"))
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	gitRepo := git.NewLocalRepository()

	aiProvider, err := ai.NewOpenAIProvider(ai.Config{
		APIKey: cfg.AIAPIKey,
		Model:  cfg.AIModel,
	})
	if err != nil {
		log.Fatalf("Failed to initialize AI provider: %v", err)
	}

	commitService := service.NewCommitService(gitRepo, aiProvider)

	cmd.SetCommitService(commitService)

	cmd.Execute()
}
