# Agrivanna Common

[![Production Release](https://github.com/Vortex-Livestock/agrivanna-common/actions/workflows/release.yaml/badge.svg)](https://github.com/Vortex-Livestock/agrivanna-common/actions/workflows/release.yaml)
[![Continuous Integration](https://github.com/Vortex-Livestock/agrivanna-common/actions/workflows/continuous_integration.yaml/badge.svg)](https://github.com/Vortex-Livestock/agrivanna-common/actions/workflows/continuous_integration.yaml)

Agrivanna Common is a Go package that contains common functions and utilities that are used across all Agrivanna services.

## Documentation

- [Setup](#setup)
- [Quick Start](#quick-start)
- [Git Workflow](#git-workflow)
- [Branch Naming Convention](#branch-naming-convention)
- [Commit Message Convention](#commit-message-convention)
- [Contributors](#contributors)
- [License](#license)

## Setup

> [!IMPORTANT]
>
> If you are using Windows or Linux, you need to install the following programs using the package manager of your operating system. The following instructions use Homebrew, which is a package manager for Mac OS.

The following programs should be installed:

- [Homebrew](https://brew.sh/)
- [Go](https://golang.org/doc/install)

## Quick Start

1. Clone the repository

```bash
$ git clone git@github.com:Vortex-Livestock/agrivanna-common.git
```

2. Change directory to the project folder

```bash
$ cd agrivanna-common
```

## Git Workflow

- The `devel` branch is the default branch.
- The `main` branch is the production branch.

## Branch Naming Convention

- Feature branches should be named as `feature/AG-<jira-ticket-number>-<feature-name>`.
- Bugfix branches should be named as `bugfix/AG-<jira-ticket-number>-<bugfix-name>`.
- Hotfix branches should be named as `hotfix/AG-<jira-ticket-number>-<hotfix-name>`.

The branche name should be in the following format:

```bash
git checkout -b feature/AG-76-add-protobuff-definitions
```

## Commit Message Convention

The basic structure includes:

- `fix`: for bug fixes
- `feat`: for new features

The commit message should be in the following format:

```bash
git commit -m "feat: AG-76 Add protobuff definitions"
```

## Contributors

- [Axel Sanchez](https://github.com/Axeloooo)

## License

[MIT](https://opensource.org/licenses/MIT)
