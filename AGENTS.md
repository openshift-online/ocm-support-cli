# AGENTS.md

This file provides guidance to AI coding assistants when working with this repository.

## Project Overview

OCM Support CLI — extends `ocm-cli` with commands for support engineers managing OpenShift clusters. Provides tools for investigating cluster issues, managing subscriptions, and performing support operations.

## Build & Test Commands

```bash
make build           # Build the ocm-support binary
make install         # Install to GOPATH/bin
make cmds            # Build all command binaries
make clean           # Remove build artifacts
```

## Architecture

- **cmd/**: CLI entry points organized by command group
  - **cmd/ocm-support/**: Main binary entry point
- **pkg/**: Supporting libraries and utilities
- **tests/**: Test suites

## Key Conventions

- Module path: `github.com/openshift-online/ocm-support-cli`
- Plugin for `ocm-cli` — follows OCM CLI plugin conventions
- Uses Cobra for command structure
- Requires OCM CLI to be installed for integration
