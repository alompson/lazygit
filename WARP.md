# WARP.md

This file provides guidance to WARP (warp.dev) when working with code in this repository.

## Project purpose

This repository will become a Go-based CLI that acts as an AI assistant for working with Git commits and branches. The CLI should help users:
- Generate high-quality commit messages from staged changes or diffs.
- Propose branch names based on task/feature descriptions or the current work context.
- Optionally summarize recent changes or history to aid in writing commits.

## Tech stack and tooling

- Language: Go `1.25.4` (see `.tool-versions` and `go.mod`).
- CLI framework: [`github.com/spf13/cobra`] for command parsing and UX.
- Dependency management: Go modules (`go.mod`, `go.sum`).

AI provider configuration is expected to be driven via environment variables and/or configuration files (for example, an `.env` file), but exact names may evolve. When integrating or updating AI providers, prefer configuration that can be swapped without touching core command logic.

## Common commands

All commands below should be run from the repository root.

### Build and run

- Build the CLI binary:

  ```sh path=null start=null
  go build -o ai-commit-cli .
  ```

- Run the CLI without a separate build step (useful during development):

  ```sh path=null start=null
  go run ./...
  ```

- After building, invoke the binary directly (name may differ if you change `-o`):

  ```sh path=null start=null
  ./ai-commit-cli --help
  ```

### Tests

When tests are added, use the standard Go tooling:

- Run all tests across the module:

  ```sh path=null start=null
  go test ./...
  ```

- Run tests in a specific package (example for an `internal/commands` package):

  ```sh path=null start=null
  go test ./internal/commands
  ```

- Run a single test by name in a package (replace `TestName` with the actual test function):

  ```sh path=null start=null
  go test ./internal/commands -run ^TestName$
  ```

### Formatting and basic static checks

- Format all Go files in-place:

  ```sh path=null start=null
  gofmt -w .
  ```

- Run basic static analysis with `go vet`:

  ```sh path=null start=null
  go vet ./...
  ```

## Current layout and intended architecture

### Existing structure

- `main.go`: Minimal entrypoint that calls `cmd.Execute()`.
- `cmd/`:
  - `root.go`: Defines the Cobra root command and `Execute` function. This is the top-level entry for all subcommands.
- `internal/commands/`:
  - Currently contains placeholders like `commit.go` and `branch.go`, intended for implementation of domain-specific behavior.

### Intended layering

As the AI-assisted Git CLI evolves, keep a clear separation of concerns:

1. **CLI layer (`cmd/` package)**
   - Defines Cobra commands (e.g., `commit`, `branch`, `summarize`).
   - Handles flags, arguments, and user-facing help/usage.
   - Delegates all non-trivial work to internal packages.

2. **Domain logic (`internal/` packages)**
   - `internal/commands/`: Orchestrates high-level workflows for each CLI command (e.g., "generate commit message for staged changes").
   - Additional internal packages are expected to emerge, for example:
     - `internal/git`: Interactions with the local Git repo (reading diffs, status, branches, commit history).
     - `internal/ai`: Abstractions over AI providers and prompt construction.
   - These packages should be written so they can be unit tested without invoking Cobra or executing shell commands directly.

3. **Integration/adapters layer**
   - Thin wrappers around the OS and external tools (e.g., running `git` commands, reading environment variables, reading/writing config files).
   - Keep these details out of the core domain logic so they can be mocked in tests.

## AI-assisted Git workflows (intended behavior)

The CLI is expected to support flows like the following:

1. **AI-generated commit messages**
   - Read staged changes or a specified diff from Git.
   - Summarize and interpret the changes using an AI provider.
   - Produce a high-quality commit message (subject + body) and print it to stdout or optionally apply it directly via `git commit`.

2. **AI-suggested branch names**
   - Take a description of the work (e.g., from a flag or prompt) and/or inspect the current Git context.
   - Generate candidate branch names based on conventions (e.g., `feature/...`, `fix/...`).
   - Optionally create and switch to the chosen branch via Git.

3. **(Optional) Change summaries**
   - Provide short summaries of staged changes, recent commits, or a branch diff to help the user understand context before committing.

When you extend or modify these flows, prefer designs where:
- Cobra commands remain thin wrappers.
- Core logic lives in `internal/` and is testable without shelling out.
- AI provider–specific details are isolated to a small number of files.

## How future agents should work in this repo

- When adding a new CLI feature, first design the internal API in `internal/` (e.g., functions that take diffs or descriptions and return proposed commit messages or branch names), then wire it into a Cobra command in `cmd/`.
- When integrating new AI providers or changing prompts, confine those changes to the AI-related internal packages rather than editing every command.
- When touching Git-related behavior, centralize shelling-out/`git` calls in a dedicated internal package so that the rest of the code can remain provider-agnostic and easier to test.
