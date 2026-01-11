# Contributing to ffrelayctl

Hello! Thank you for your interest in contributing to ffrelayctl!

This document provides guidelines for contributing to the project.

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:
- [Go](https://go.dev/doc/install) (version 1.25 or later)
- [Git](https://git-scm.com/install/)
- [Docker](https://docs.docker.com/get-started/get-docker/) (required for commit hooks)

Additionally, you'll need an account and API key from [Firefox Relay](https://relay.firefox.com).

## Development Setup

After cloning the repository, set up your development environment:

```bash
$ make setup
```

This command will:
- Check for required dependencies (Docker)
- Configure git hooks for commit message linting
- Set up the development environment

### Building

Build the project locally:

```bash
$ go build -o ffrelayctl .
```

### Running the CLI

You can run the built binary directly:

```bash
$ ./ffrelayctl --help
```

Or install it to your `$GOPATH/bin`:

```bash
$ go install .
```

### Testing

Ensure you have a Firefox Relay API key set as an environment variable:

```bash
$ export FFRELAYCTL_KEY=replace-me
```

Test commands manually:

```bash
$ ./ffrelayctl profiles list
$ ./ffrelayctl masks list
```

## Making Changes

### Creating a Branch

Create a new branch for your changes:

```bash
$ git checkout -b feature/feature-name
# or
$ git checkout -b fix/fix-name
```

### Code Style

- Follow standard Go conventions
- Run `$ go fmt` to format your code
- Ensure your code passes `$ go vet`
- Keep things [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself)

## Commit Guidelines

[Conventional Commits](https://www.conventionalcommits.org/) are enforced by commitlint.

### Commit Scopes

Common scopes in this project:
- `cli`: CLI interface and commands
- `api`: API client code
- `masks`: Email masks
- `phones`: Phone masks
- `profiles`: Profile management
- `contacts`: Contact management
- `dev`: Development tools, setup, experience
- `release`: Release process
- `readme`: README documentation

### Commit Examples

```bash
feat(masks): add support for filtering masks by status
fix(api): handle rate limiting errors correctly
docs(readme): update installation instructions
chore(deps): update dependencies to latest versions
```

## Pull Requests

1. Update your branch with the latest changes from upstream:
   ```bash
   $ git fetch upstream
   $ git rebase upstream/main
   ```

2. Ensure all commits follow the commit guidelines described above

3. Open a Pull Request on GitHub with:
   - A clear title following conventional commit format
   - A brief summary of the changes and why they're needed
   - Any related issues or resources for reference

### PR Requirements

- All commits must follow conventional commit format
- Code must be properly formatted (`go fmt`)
- No new warnings from `go vet`
- The PR should focus on a single feature or fix
- Keep PRs reasonably sized for easier review

## Reporting Issues

Feel free to report issues (bugs, feature requests, etc) using GitHub.

## License

By contributing to ffrelayctl, you agree that your contributions will be licensed under the [MIT License](LICENSE).
