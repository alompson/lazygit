# 🤖 Lazy Git

[![Go Version](https://img.shields.io/badge/Go-1.25%2B-00ADD8?logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

**Lazy Git** is a CLI tool that leverages AI to automatically generate meaningful git commit messages and branch names based on your code changes. Never write "fix stuff" or "update" again!
## ✨ Features

- 🎯 **Smart Commit Messages** - Analyzes staged changes and generates conventional commit messages
- 🌿 **Branch Name Generation** - Creates descriptive branch names from your diffs
- 🏗️ **Clean Architecture** - Built with interfaces, dependency injection, and testability in mind


## 📦 Installation

### Using `go install` (Recommended)

```bash
go install github.com/alompson/lazy-git@latest
```

### From Source

```bash
git clone https://github.com/alompson/lazy-git.git
cd lazy-git
go build -o lazy-git
sudo mv lazy-git /usr/local/bin/  # Optional: make it globally available
```

## ⚙️ Configuration

1. **Get an OpenAI API key** from [platform.openai.com/api-keys](https://platform.openai.com/api-keys)

2. **Add to your shell config** (`~/.zshrc` or `~/.bashrc`):

```bash
export AI_API_KEY="sk-proj-your-key-here"
```

3. **Reload your shell:**

```bash
source ~/.zshrc
```

### Optional Configuration

```bash
export AI_MODEL="gpt-4o"  # Specify a different OpenAI model (default: chatgpt-4o-latest)
```

## 🚀 Usage

### Generate Commit Message

```bash
# Stage your changes
git add .

# Generate a commit message
lazygit commit
```

**Example output:**
```
git commit -m "feat: add AI-powered commit message generation

- Implemented OpenAI integration for commit analysis
- Added clean architecture with repository pattern
- Created configuration management system"
```

### Generate Branch Name

```bash
# Make some changes (staged or unstaged)
lazygit branch
```

**Example output:**
```
feature-ai-commit-generation
```

### Advanced: Shell Widget Integration

For the ultimate lazy workflow, add this to your `~/.zshrc` to pre-fill your command line with the generated commit command:

```zsh
# Lazy Git widget - press Ctrl+G to auto-fill commit message
function _lazygit_commit_widget() {
    # Get the full git commit command
    local cmd=$(lazygit commit 2>/dev/null)
    
    if [[ -n "$cmd" ]]; then
        BUFFER="$cmd"
        CURSOR=$#BUFFER
    fi
}

zle -N _lazygit_commit_widget
bindkey '^G' _lazygit_commit_widget  # Ctrl+G to trigger
```

Now just press **Ctrl+G** in your terminal and the commit command appears ready to edit!

## 🏗️ Architecture

```
lazygit/
├── cmd/                    # CLI commands (Cobra)
│   ├── root.go
│   ├── commit.go
│   └── branch.go
├── internal/
│   ├── ai/                 # AI provider interface & implementations
│   │   ├── provider.go
│   │   └── openai.go
│   ├── git/                # Git operations
│   │   └── repository.go
│   ├── config/             # Configuration management
│   │   └── config.go
│   └── service/            # Business logic
│       └── commit.go
└── main.go
```

### Design Principles

- **Interface-based design** for easy testing and extensibility
- **Dependency injection** for loose coupling
- **Repository pattern** to abstract git operations
- **Service layer** for business logic separation
- **Error wrapping** with context for better debugging

## 🛠️ Development

### Prerequisites

- Go 1.25 or later
- Git
- OpenAI API key

### Build

```bash
go build -o lazygit
```

### Run Locally

```bash
go run . commit
go run . branch
```

### Project Structure Explained

- **`cmd/`** - Contains CLI command handlers (thin layer, just parses flags and calls services)
- **`internal/ai/`** - AI provider abstraction (easy to swap OpenAI for other LLMs)
- **`internal/git/`** - Git operations behind an interface (mockable for tests)
- **`internal/service/`** - Business logic (orchestrates git + AI operations)
- **`internal/config/`** - Configuration loading and validation

## 📝 Examples

### Basic Workflow

```bash
# 1. Make some changes
echo "new feature" >> file.go

# 2. Stage them
git add file.go

# 3. Generate commit message
lazygit commit
# Output: git commit -m "feat: add new feature implementation"

# 4. Copy and paste, or use the widget!
```

### Generate Branch Name

```bash
# You're working on a login feature
lazygit branch
# Output: feature-user-authentication-login
```

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details

## 🙏 Acknowledgments

- Built with [Cobra](https://github.com/spf13/cobra) for CLI
- Powered by [OpenAI](https://openai.com/) for AI generation
- Inspired by lazy developers everywhere 🚀

## 🐛 Troubleshooting

### "Failed to load configuration: AI_API_KEY environment variable is required"

Make sure you've exported `AI_API_KEY` in your shell config and reloaded it.

### "not a git repository"

Run the command from inside a git repository.

### "No staged changes found"

Stage your changes first with `git add <files>`

---

Made with ❤️ by developers who hate writing commit messages

