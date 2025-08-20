# Contributing to HaikuGo

Thank you for your interest in contributing to HaikuGo! This document provides guidelines and information for contributors.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Making Changes](#making-changes)
- [Testing](#testing)
- [Submitting Changes](#submitting-changes)
- [Code Style](#code-style)
- [Project Structure](#project-structure)

## Code of Conduct

This project is committed to providing a welcoming and inclusive environment for all contributors. Please be respectful and considerate in all interactions.

**Note**: This project was created with AI assistance and is community-maintained. We welcome contributions from anyone who finds it useful, regardless of their background or experience level.

## Getting Started

1. **Fork the repository** on GitHub
2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/haikugo.git
   cd haikugo
   ```
3. **Add the upstream remote**:
   ```bash
   git remote add upstream https://github.com/thornzero/haikugo.git
   ```

## Development Setup

### Prerequisites

- Go 1.24 or later
- Make (optional, for convenience)

### Setup Commands

```bash
# Install development dependencies
make dev-deps

# Tidy up modules
make tidy

# Run full setup
make setup
```

## Making Changes

### 1. Create a Feature Branch

```bash
git checkout -b feature/your-feature-name
```

### 2. Make Your Changes

- Follow the [Code Style](#code-style) guidelines
- Add tests for new functionality
- Update documentation as needed

### 3. Test Your Changes

```bash
# Run all tests
make test

# Run with coverage
make coverage

# Run all checks
make all
```

### 4. Commit Your Changes

```bash
git add .
git commit -m "feat: add your feature description"
```

Use conventional commit messages:
- `feat:` for new features
- `fix:` for bug fixes
- `docs:` for documentation changes
- `test:` for test additions
- `refactor:` for code refactoring
- `chore:` for maintenance tasks

## Testing

### Running Tests

```bash
# All tests
make test

# Specific package
go test ./internal/analyzer

# With verbose output
go test -v ./...

# With coverage
make coverage
```

### Test Guidelines

- Write tests for all new functionality
- Aim for high test coverage
- Test edge cases and error conditions
- Use descriptive test names
- Follow Go testing best practices

## Submitting Changes

### 1. Push Your Branch

```bash
git push origin feature/your-feature-name
```

### 2. Create a Pull Request

1. Go to your fork on GitHub
2. Click "New Pull Request"
3. Select your feature branch
4. Fill out the PR template
5. Submit the PR

### 3. PR Review Process

- All PRs require review and approval
- Address feedback and requested changes
- Ensure CI checks pass
- Update documentation if needed

## Code Style

### Go Code

- Follow [Effective Go](https://golang.org/doc/effective_go.html)
- Use `gofmt` for formatting
- Run `go vet` to check for common issues
- Follow Go naming conventions

### Documentation

- Add GoDoc comments for all public functions
- Update README.md for user-facing changes
- Keep examples up to date
- Document any new configuration options

### Error Handling

- Use structured errors with context
- Return meaningful error messages
- Handle edge cases gracefully
- Log appropriate information for debugging

## Project Structure

```
haikugo/
├── cmd/haikuctl/          # CLI application
├── internal/              # Internal packages (not importable)
│   ├── analyzer/          # Core analysis logic
│   ├── haiku/            # Haiku data structures
│   └── input/            # Input parsing
├── pkg/haikugo/          # Public API (importable)
├── testdata/             # Test fixtures
├── .github/workflows/    # CI/CD configuration
└── docs/                 # Additional documentation
```

### Package Guidelines

- **`internal/`**: Private packages, not importable by external code
- **`pkg/`**: Public packages, safe for external import
- **`cmd/`**: Command-line applications
- Keep packages focused and cohesive
- Minimize dependencies between packages

## Areas for Contribution

These are suggestions for areas where contributors might find interesting work. Since this is community-maintained, priorities are determined by what people actually want to work on.

### Potential Areas

- [ ] Additional language support (Japanese, other languages)
- [ ] Enhanced syllable counting accuracy
- [ ] Web interface
- [ ] Additional output formats (YAML, CSV)
- [ ] Performance optimizations
- [ ] Extended kigo (season word) lexicon
- [ ] Advanced linguistic analysis
- [ ] Plugin system
- [ ] Additional UI themes
- [ ] Integration with text editors
- [ ] Mobile applications

### Community Priorities

- **No Fixed Priorities**: Work on what interests you or what you need
- **Real Usage**: Features that solve actual problems get more attention
- **Collaboration**: Find others working on similar interests
- **Documentation**: Help others understand and use the project

## Getting Help

- **Issues**: Use GitHub Issues for bug reports and feature requests
- **Discussions**: Use GitHub Discussions for questions and ideas
- **Code Review**: Ask questions in PR comments
- **Documentation**: Check the README and code comments first

## Recognition

Contributors will be recognized in:
- The project README
- Release notes
- The changelog
- GitHub contributors list

## Project Philosophy

This project embodies the spirit of collaborative creation:
- **AI-Assisted**: The initial implementation was created with AI assistance
- **Community-Driven**: Ongoing development depends on community interest and contributions
- **Open Source**: Shared freely for anyone who finds it useful
- **No Pressure**: No deadlines, no obligations, just shared interest in poetry and technology

Thank you for contributing to HaikuGo! 🎋
