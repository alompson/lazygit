package config

import (
	"fmt"
	"os"
)

type Config struct {
	AIAPIKey string
	AIModel  string
	GitWorkDir string
}

func Load() (*Config, error) {
	cfg := &Config{
		AIAPIKey:   os.Getenv("AI_API_KEY"),
		AIModel:    os.Getenv("AI_MODEL"),
		GitWorkDir: os.Getenv("GIT_WORK_DIR"),
	}

	if cfg.AIAPIKey == "" {
		return nil, fmt.Errorf("AI_API_KEY environment variable is required")
	}

	return cfg, nil
}
