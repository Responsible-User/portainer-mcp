# Portainer MCP (Extended Fork)

> **This is an extended fork of the [original Portainer MCP server](https://github.com/portainer/portainer-mcp).** The upstream project supported Portainer up to version 2.31.2 with a limited set of tools. This fork significantly extends support through **Portainer 2.38** and adds **50+ new MCP tools** covering git credentials, alerting/observability, fleetwide policies, policy templates, Kubernetes custom resources, registry connectivity testing, and more.

## Overview

Portainer MCP connects your AI assistant directly to your Portainer environments via the [Model Context Protocol (MCP)](https://modelcontextprotocol.io/introduction). Manage Portainer resources such as users, environments, stacks, registries, and policies — or dive deeper by executing any Docker or Kubernetes command directly through the AI.

MCP is an open protocol that standardizes how applications provide context to LLMs. This implementation exposes Portainer's container management capabilities through MCP, allowing AI assistants to interact with your containerized infrastructure in a secure and standardized way.

> [!NOTE]
> This tool is designed to work with Portainer versions **2.27.0 through 2.38.x**. If your Portainer version is outside this range, you can use the `--disable-version-check` flag to attempt connection anyway. See [Portainer Version Support](#portainer-version-support) for details.

It is currently designed to work with a Portainer administrator API token.

## Installation

You can download pre-built binaries for Linux (amd64, arm64) and macOS (arm64) from the [**Latest Release Page**](https://github.com/portainer/portainer-mcp/releases/latest), or build from source:

```bash
git clone https://github.com/portainer/portainer-mcp.git
cd portainer-mcp
make build
```

The binary will be in `dist/portainer-mcp`.

## Usage

Configure your MCP client (e.g. Claude Desktop, Claude Code, Cursor) like so:

```json
{
    "mcpServers": {
        "portainer": {
            "command": "/path/to/portainer-mcp",
            "args": [
                "-server",
                "https://your-portainer:9443",
                "-token",
                "your-api-token",
                "-tools",
                "/tmp/tools.yaml"
            ]
        }
    }
}
```

> [!NOTE]
> By default, the tool looks for `tools.yaml` in the same directory as the binary. If the file does not exist, it will be created with the default tool definitions. You may need to specify a custom path with `-tools` when using AI assistants that have restricted write permissions to the working directory.

### Command Line Flags

| Flag | Required | Description |
|------|----------|-------------|
| `-server` | Yes | The Portainer server URL (e.g. `https://portainer.example.com:9443`) |
| `-token` | Yes | API access token for the Portainer server |
| `-tools` | No | Path to a custom tools.yaml file |
| `-read-only` | No | Run in read-only mode (only list/get tools available) |
| `-disable-version-check` | No | Skip Portainer server version validation at startup |

## Read-Only Mode

For security-conscious users, the application can be run in read-only mode. This ensures only read operations are available, completely preventing any modifications to your Portainer resources.

```json
{
    "mcpServers": {
        "portainer": {
            "command": "/path/to/portainer-mcp",
            "args": [
                "-server", "https://your-portainer:9443",
                "-token", "your-api-token",
                "-read-only"
            ]
        }
    }
}
```

When using read-only mode:
- Only read tools (list, get) will be available to the AI model
- All write tools (create, update, delete) are not loaded
- The Docker and Kubernetes proxy request tools are not loaded

## Disable Version Check

By default, the application validates that your Portainer server version falls within the supported range and will fail to start if there's a mismatch. You can disable this:

```json
{
    "mcpServers": {
        "portainer": {
            "command": "/path/to/portainer-mcp",
            "args": [
                "-server", "https://your-portainer:9443",
                "-token", "your-api-token",
                "-disable-version-check"
            ]
        }
    }
}
```

> [!WARNING]
> Disabling the version check may result in unexpected behavior or API incompatibilities. Some features may not work correctly with unsupported versions.

## Tool Customization

The tool definitions are embedded in the binary by default. You can customize them by specifying a custom tools file:

```json
"args": ["-server", "...", "-token", "...", "-tools", "/path/to/custom/tools.yaml"]
```

The default tools file is available at `internal/tooldef/tools.yaml` in the source code. You can modify tool descriptions and parameter descriptions to alter how AI models interpret and use them, or remove tools you don't need.

> [!WARNING]
> Do not change tool names or parameter definitions (other than descriptions), as this will prevent the tools from functioning correctly.

## Portainer Version Support

This fork supports Portainer versions **2.27.0 through 2.38.x**. The version is validated at startup (can be bypassed with `-disable-version-check`).

### Version History

| Fork Version | Portainer Support | New Tools Added |
|-------------|-------------------|-----------------|
| 1.0.0 | 2.27.0 - 2.38.x | All tools below |

#### Upstream (original portainer-mcp)

| Upstream Version | Portainer Version | Notes |
|-----------------|-------------------|-------|
| 0.1.0 | 2.28.1 | Initial release |
| 0.2.0 | 2.28.1 | Docker proxy |
| 0.3.0 | 2.28.1 | Kubernetes proxy |
| 0.4.0 | 2.29.2 | — |
| 0.4.1 | 2.29.2 | — |
| 0.5.0 | 2.30.0 | — |
| 0.6.0 | 2.31.2 | Last upstream release |

### What This Fork Adds Beyond Upstream

| Portainer Version | Features Added |
|-------------------|----------------|
| 2.33 (LTS) | Docker stacks, settings management, registries, edge jobs, custom templates, webhooks, environment updates, access groups, delete operations |
| 2.34 | Shared git credentials (CRUD), alerting/observability (rules, silences, settings) |
| 2.35 | Version bump only (no new API endpoints) |
| 2.36 | Kubernetes custom resources (CRDs), registry connection testing |
| 2.37 | Fleetwide policies (CRUD), agent version listing |
| 2.38 | Policy templates, policy metadata, policy conflict detection |

## Supported Capabilities

The following table lists all supported operations (80+ MCP tools):

| Resource | Operation | Description |
|----------|-----------|-------------|
| **Environments** | | |
| | listEnvironments | List all available environments |
| | updateEnvironment | Update environment name, public URL, or group |
| | updateEnvironmentTags | Update tags associated with an environment |
| | updateEnvironmentUserAccesses | Update user access policies for an environment |
| | updateEnvironmentTeamAccesses | Update team access policies for an environment |
| | listAgentVersions | List available agent versions |
| **Environment Groups** | | |
| | listEnvironmentGroups | List all environment groups (edge groups) |
| | createEnvironmentGroup | Create a new environment group |
| | updateEnvironmentGroupName | Update environment group name |
| | updateEnvironmentGroupEnvironments | Update environments in a group |
| | updateEnvironmentGroupTags | Update tags on a group |
| | deleteEnvironmentGroup | Delete an environment group |
| **Access Groups** | | |
| | listAccessGroups | List all access groups (endpoint groups) |
| | createAccessGroup | Create a new access group |
| | updateAccessGroupName | Update access group name |
| | updateAccessGroupUserAccesses | Update user accesses for an access group |
| | updateAccessGroupTeamAccesses | Update team accesses for an access group |
| | addEnvironmentToAccessGroup | Add an environment to an access group |
| | removeEnvironmentFromAccessGroup | Remove an environment from an access group |
| | deleteAccessGroup | Delete an access group |
| **Edge Stacks** | | |
| | listStacks | List all edge stacks |
| | getStackFile | Get the compose file for an edge stack |
| | createStack | Create a new edge stack |
| | updateStack | Update an existing edge stack |
| | deleteStack | Delete an edge stack |
| **Docker Stacks** | | |
| | listDockerStacks | List all Docker stacks |
| | getDockerStackFile | Get the compose file for a Docker stack |
| | createDockerStack | Create a new Docker stack |
| | updateDockerStack | Update an existing Docker stack |
| | deleteDockerStack | Delete a Docker stack |
| | startDockerStack | Start a stopped Docker stack |
| | stopDockerStack | Stop a running Docker stack |
| **Tags** | | |
| | listEnvironmentTags | List all environment tags |
| | createEnvironmentTag | Create a new environment tag |
| | deleteTag | Delete a tag |
| **Teams** | | |
| | listTeams | List all teams |
| | createTeam | Create a new team |
| | updateTeamName | Update team name |
| | updateTeamMembers | Update team members |
| | deleteTeam | Delete a team |
| **Users** | | |
| | listUsers | List all users |
| | updateUserRole | Update a user's role |
| **Settings** | | |
| | getSettings | Get Portainer instance settings |
| | updateSettings | Update Portainer instance settings |
| **Registries** | | |
| | listRegistries | List all registries |
| | createRegistry | Create a new registry |
| | deleteRegistry | Delete a registry |
| | testRegistryConnection | Test connectivity to a registry |
| **Edge Jobs** | | |
| | listEdgeJobs | List all edge jobs |
| | getEdgeJob | Get details of a specific edge job |
| | createEdgeJob | Create a new edge job |
| | deleteEdgeJob | Delete an edge job |
| **Custom Templates** | | |
| | listCustomTemplates | List all custom templates |
| | createCustomTemplate | Create a new custom template |
| | deleteCustomTemplate | Delete a custom template |
| **Webhooks** | | |
| | listWebhooks | List all webhooks |
| | createWebhook | Create a new webhook |
| | deleteWebhook | Delete a webhook |
| **Git Credentials** | | |
| | listGitCredentials | List all shared git credentials |
| | getGitCredential | Get details of a specific git credential |
| | createGitCredential | Create a new git credential |
| | updateGitCredential | Update an existing git credential |
| | deleteGitCredential | Delete a git credential |
| **Alerting** | | |
| | listAlerts | List active or silenced alerts |
| | listAlertRules | List all alert rules |
| | getAlertRule | Get details of a specific alert rule |
| | updateAlertRule | Update an alert rule |
| | deleteAlertRule | Delete an alert rule |
| | getAlertingSettings | Get alerting configuration settings |
| | createAlertSilence | Create a new alert silence |
| | deleteAlertSilence | Delete an alert silence |
| **Fleetwide Policies** | | |
| | listPolicies | List all fleetwide policies |
| | getPolicy | Get details of a specific policy |
| | createPolicy | Create a new fleetwide policy |
| | updatePolicy | Update an existing policy |
| | deletePolicy | Delete a policy |
| | listPolicyTemplates | List available policy templates |
| | getPolicyTemplate | Get details of a specific policy template |
| | getPolicyMetadata | Get policy metadata (minimum agent versions) |
| | getPolicyConflicts | Preview conflicts for a proposed policy |
| **Kubernetes Custom Resources** | | |
| | listCustomResourceDefinitions | List all CRDs in a Kubernetes environment |
| | getCustomResourceDefinition | Get details of a specific CRD |
| | deleteCustomResourceDefinition | Delete a CRD |
| | listCustomResources | List custom resource instances |
| | getCustomResource | Get a specific custom resource (JSON/YAML) |
| | deleteCustomResource | Delete a custom resource |
| **Docker Proxy** | | |
| | dockerProxy | Proxy any Docker API request |
| **Kubernetes Proxy** | | |
| | kubernetesProxy | Proxy any Kubernetes API request |
| | getKubernetesResourceStripped | Proxy GET Kubernetes requests with verbose metadata stripped |

## Development

### Building

```bash
make build                          # Build for current platform
make PLATFORM=darwin ARCH=arm64 build  # Cross-compile
```

### Testing

```bash
make test           # Unit tests
make test-integration  # Integration tests (requires Docker)
make test-all       # All tests
```

### Other Commands

```bash
go vet ./...    # Lint
gofmt -s -w .   # Format
make inspector  # MCP Inspector
```

### Token Counting

To estimate how many tokens your tool definitions consume:

```bash
# Generate tools JSON
go run ./cmd/token-count -input internal/tooldef/tools.yaml -output .tmp/tools.json

# Query Anthropic API (requires API key and jq)
./token.sh -k sk-ant-xxxxxxxx -i .tmp/tools.json
```
