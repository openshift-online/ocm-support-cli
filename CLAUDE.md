# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## About This Project

ocm-support-cli is a Go CLI tool that extends the OCM (OpenShift Cluster Manager) CLI with commands specifically designed for support engineers. It's a plugin for the `ocm` command and becomes available as `ocm support`.

## Prerequisites

- OCM CLI must be installed and configured (`ocm login` completed)
- Advanced permissions with roles like `UHCSupport`, `SDBSignalMonitor`

## Common Commands

### Build and Install
- **Build locally**: `make build` (creates `ocm-support` binary)
- **Install to GOPATH**: `make install` (installs to `$GOPATH/bin/ocm-support`)
- **Clean**: `make clean` (removes local binary)
- **Build all commands**: `make cmds`
- **Ensure OCM CLI**: `make ensureOCM` (downloads/installs OCM CLI dependencies)

### Testing and Development
- **Run validation**: `ocm support version` (validate installation)
- **View help**: `ocm support -h` (see all available commands)

## Architecture Overview

### Project Structure
- **`cmd/ocm-support/`**: CLI command definitions using Cobra framework
  - Main entry point: `main.go` → `root.go`
  - Commands organized by operation: `get/`, `create/`, `delete/`, `patch/`, `sync-cloud-resources/`
  - Each command has its own `cmd.go` file in nested directories
- **`pkg/`**: Core business logic packages for each resource type
  - `account/`, `organization/`, `subscription/`, `cluster/` - main resource types
  - `capability/`, `label/`, `role_binding/` - supporting features
  - `registry_credential/`, `quota/`, `machine_pool/` - specialized components
- **`tests/`**: Test files and test dependencies

### Command Architecture
The CLI follows a hierarchical structure:
```
ocm support <operation> <resource> [arguments] [flags]
```

Main operations:
- **get**: Retrieve resources (accounts, organizations, subscriptions, clusters, etc.)
- **create**: Create labels, capabilities, role bindings, registry credentials
- **delete**: Remove labels, capabilities, role bindings, registry credentials, accounts
- **patch**: Modify accounts, organizations, subscriptions (with dry-run support)
- **sync-cloud-resources**: Sync cloud resources and generate quota rules

### Key Dependencies
- **OCM SDK**: Uses `github.com/openshift-online/ocm-sdk-go` for API interactions
- **Cobra**: CLI framework for command structure
- **Logrus**: Structured logging with configurable verbosity levels

### Resource Relationships
- **Accounts** belong to **Organizations**
- **Organizations** have **Subscriptions**
- **Subscriptions** are linked to **Clusters**
- All resources can have **Labels** and **Capabilities**
- **Accounts** can have **Role Bindings** at different scopes (application, organization, subscription)

### Search and Filtering
Most commands support flexible search criteria:
- Resources can be found by ID, external ID, or related resource IDs
- Advanced filtering with SQL-like syntax for complex queries
- Optional secondary search criteria for refined results

### Dry-Run Support
Commands that modify data (patch, delete) include `--dry-run` flags to preview changes before execution.