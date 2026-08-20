# Network Inspector

![CI](https://github.com/Qyroxen/Network-Inspector/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Network-Inspector/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Network-Inspector?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Network-Inspector)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Network-Inspector)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Network-Inspector?style=social)](https://github.com/Qyroxen/Network-Inspector/stargazers)

## What is it?

Network Inspector is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Network-Inspector.git
cd Network-Inspector
go build -o networkinspector .

# Run
./networkinspector --help
```

## CLI Usage

```bash
# Basic usage
./networkinspector

# With flags
./networkinspector --verbose --output json

# Get help
./networkinspector --help
```

## Examples

```bash
# Example 1
./networkinspector example1

# Example 2
./networkinspector example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o networkinspector .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Network-Inspector/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Network-Inspector?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Network-Inspector/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Network-Inspector?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Network-Inspector/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Network-Inspector" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Network-Inspector/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Network-Inspector" alt="Pull Requests">
  </a>
</p>
