# Contributing to real-time-forum

Thank you for your interest in contributing to this project! We welcome contributions of all kinds — whether it's bug fixes, new features, documentation improvements, or helping others in issues. This document will guide you on how to effectively contribute.

## How to Contribute

### Reporting Issues
- Please check the existing issues before opening a new one to avoid duplicates.
- Provide clear and detailed information to help us understand and reproduce the problem.

### Submitting Code Changes
- Fork the repository and create your feature branch from `main`.
- Make sure your code passes all tests and follows the project's code style.
- Write tests for any new features or bug fixes where applicable.
- Update documentation if your changes affect usage or configuration.
- Submit a pull request (PR) with a clear description of your changes.
- We review PRs promptly and may suggest improvements or ask questions.

## Project Structure Overview
- Backend Go server code is in the `/cmd/server` directory.
- Frontend SPA is located in the `/frontend` directory, containing HTML, CSS, and JS assets.
- Use the provided Makefile to build, test, and format the code. Run `make help` for commands.

## Development Setup
- Install Go (version 1.22 or higher recommended).
- Install Node.js/npm for building frontend assets.
- Use `make` targets like `make run` to start the server and `make test` to run backend tests.
- Frontend development instructions live in the `/frontend/README.md`.

## Style Guidelines
- Follow Go idiomatic style for backend code (use `go fmt` and `golangci-lint`).
- Follow existing styling conventions in the frontend for HTML, CSS, and JS.
- Keep commits small and focused with descriptive messages.

## Getting Help
- Feel free to open issues for questions or help with development.

## Code of Conduct
By participating, you agree to abide by our [Code of Conduct](CODE_OF_CONDUCT.md). We aim to provide a welcoming, inclusive environment for all contributors.

---

Thank you for helping make this project better! Your contributions are highly valued.
