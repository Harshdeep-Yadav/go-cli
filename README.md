# GitHub CLI Tool (ghcli)

A powerful command-line interface tool for interacting with GitHub, built with Go.

## Features

### Current Features
- Command-line interface for GitHub operations
- Built with Go and Cobra framework
- Colorful terminal output
- Cross-platform support
- Basic issue management

### Upcoming Features

#### Authentication and Personalization
- [ ] GitHub token authentication
- [ ] User profiles and preferences
- [ ] Configuration file for user settings

#### Repository Management
- [ ] `repo create` - Create new repositories
- [ ] `repo clone` - Clone repositories with additional options
- [ ] `repo fork` - Fork repositories
- [ ] `repo star` - Star/unstar repositories
- [ ] `repo watch` - Watch/unwatch repositories

#### Issue Management
- [ ] `issue create` - Create new issues
- [ ] `issue close` - Close issues
- [ ] `issue comment` - Add comments to issues
- [ ] `issue assign` - Assign users to issues
- [ ] `issue label` - Add/remove labels
- [ ] Filter issues by status, labels, assignees

#### Pull Request Features
- [ ] `pr create` - Create pull requests
- [ ] `pr list` - List pull requests
- [ ] `pr review` - Review pull requests
- [ ] `pr merge` - Merge pull requests
- [ ] `pr checkout` - Checkout pull request branches

#### Release Management
- [ ] `release create` - Create new releases
- [ ] `release list` - List releases
- [ ] `release download` - Download release assets

#### User Management
- [ ] `user profile` - View user profiles
- [ ] `user follow` - Follow/unfollow users
- [ ] `user search` - Search for users

#### Advanced Features
- [ ] Interactive mode with TUI (Terminal User Interface)
- [ ] Git workflow integration
- [ ] CI/CD status checking
- [ ] Repository statistics and insights
- [ ] Code search functionality

#### Quality of Life Improvements
- [ ] Better error handling and user feedback
- [ ] Progress bars for long-running operations
- [ ] Colored output for better readability
- [ ] Command aliases for common operations
- [ ] Tab completion support

#### API Rate Limiting and Caching
- [ ] Implement caching for frequently accessed data
- [ ] Handle GitHub API rate limits gracefully
- [ ] Offline mode for basic operations

## Prerequisites

- Go 1.23.2 or higher
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Harshdeep-Yadav/go-cli.git
cd ghcli
```

2. Install dependencies:
```bash
go mod download
```

3. Build the project:
```bash
go build
```

4. (Optional) Add the binary to your PATH:
```bash
# For Linux/macOS
sudo mv ghcli /usr/local/bin/

# For Windows
# Add the directory containing ghcli.exe to your system's PATH
```

## Usage

After installation, you can use the CLI tool with various commands:

```bash
# Basic usage
ghcli [command]

# For help
ghcli --help
```

## Development

### Project Structure

```
ghcli/
├── cmd/         # Command implementations
├── github/      # GitHub API related code
├── main.go      # Entry point
├── go.mod       # Go module file
└── go.sum       # Go module checksum
```

### Adding New Commands

1. Create a new command file in the `cmd` directory
2. Register the command in the root command
3. Implement the command logic

### Implementation Plan

#### Phase 1: Core Features
- Authentication system
- Basic repository operations
- Enhanced issue management
- Pull request support

#### Phase 2: Advanced Features
- Release management
- User management
- Interactive mode
- Caching system

#### Phase 3: Polish
- Comprehensive documentation
- Enhanced error handling
- Progress indicators
- Improved user experience

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- [Cobra](https://github.com/spf13/cobra) - A Commander for modern Go CLI applications
- [fatih/color](https://github.com/fatih/color) - Color package for Go 
