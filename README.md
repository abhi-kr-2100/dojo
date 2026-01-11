# Dojo

Dojo is a workspace manager for the [Jujutsu (jj)](https://github.com/jj-vcs/jj) version control system. It aims to eliminate every pain point associated with working on multiple workspaces.

## Features

- **Workspace Management**: Easily create, list, and remove workspaces.
- **Hooks**: Run setup/teardown commands during workspace lifecycle events.

## Installation

(Coming soon)

## Usage

### CLI Commands

*   `dojo ls`: List all managed workspaces.
*   `dojo new <name>`: Create a new workspace.
*   `dojo cd <name> [--new]`: Switch to a different workspace. Use `--new` to create it if it doesn't exist.
*   `dojo rm [<name>]`: Remove a workspace. Defaults to the current workspace if no name is provided.

### Examples

**Create a new feature workspace:**
```bash
dojo new feature-login
```

**List all workspaces:**
```bash
dojo ls
```

**Switch to the new workspace:**
```bash
dojo cd feature-login
```

**Create and switch to a workspace in one command:**
```bash
dojo cd feature-payment --new
```

**Remove a workspace by name:**
```bash
dojo rm feature-login
```

**Delete the current workspace:**
```bash
dojo rm
```

### Configuration & Hooks

Dojo is configured via a `dojo.toml` file placed locally inside the project.

#### Hook Definitions

Hooks are defined under the `[hooks]` section in the configuration file. They are executed as shell commands.

**Supported Hooks:**

*   `pre-create`: Executed before a new workspace is created.
*   `post-create`: Executed after a workspace is successfully created.
*   `pre-switch`: Executed before switching context to another workspace.
*   `post-switch`: Executed after switching context.
*   `pre-remove`: Executed before a workspace is removed.
*   `post-remove`: Executed after a workspace is removed.

**Example `dojo.toml`:**

```toml
[hooks]
# Run before creating a new workspace
pre-create = [
  "echo 'Preparing to create workspace...'"
]

# Run after a workspace is successfully created
# Useful for initializing the repo or setting config
post-create = [
  "jj git init --colocate",
  "echo 'Workspace created and initialized!'"
]

# Run after switching to a workspace
post-switch = [
  "jj status"
]

# Run after removing a workspace
post-remove = [
  "echo 'Workspace removed.'"
]
```
