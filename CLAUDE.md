# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build, Test & Run Commands
- Build: `make build`
- Unit tests: `make test`
- Single test: `go test -v ./path/to/package -run TestName`
- Integration tests: `make test-integration` (requires Docker)
- All tests: `make test-all`
- Lint: `go vet ./...`
- Format: `gofmt -s -w .`
- MCP Inspector: `make inspector`
- Cross-platform build: `make PLATFORM=<platform> ARCH=<arch> build`

## Architecture Overview

This is an MCP (Model Context Protocol) server that exposes Portainer container management operations as tools for AI assistants. It communicates over stdio using the `mcp-go` library.

### Request Flow

```text
AI Assistant → MCP Protocol (stdio) → PortainerMCPServer → Wrapper Client → Raw API Client → Portainer API
```

### Key Layers

**Entry point** (`cmd/portainer-mcp/mcp.go`): Parses CLI flags, creates tools.yaml if needed, initializes the server, registers feature groups via `server.AddXXXFeatures()`, then starts stdio listener.

**MCP handlers** (`internal/mcp/`): One file per domain (environment.go, tag.go, stack.go, docker.go, kubernetes.go, etc.). Each file exposes an `AddXXXFeatures()` method that registers read tools unconditionally and write tools only when `!s.readOnly`. Handlers return `ToolHandlerFunc` closures that parse parameters, call the wrapper client, and return JSON or error results.

**Wrapper client** (`pkg/portainer/client/`): Abstraction over the raw API client. Calls the raw client, transforms raw models into local models. Used by all MCP handlers.

**Local models** (`pkg/portainer/models/`): Simplified structs with `ConvertXXX()` functions that transform raw API models. Only contain fields relevant to MCP.

**Tool definitions** (`internal/tooldef/tools.yaml`): YAML file embedded in the binary at build time. Defines tool names, descriptions, parameters, and annotations. External file can override the embedded version.

**Tool name constants** (`internal/mcp/schema.go`): All tool names as constants (e.g., `ToolListEnvironments = "listEnvironments"`). Registration uses `s.addToolIfExists(ToolName, handler)` which silently skips tools not in the YAML.

## Code Style

- Error handling: `fmt.Errorf("failed to X: %w", err)` for wrapping; handlers return `mcp.NewToolResultErrorFromErr("description", err), nil` (error as tool result, not Go error)
- Import groups: stdlib, external, internal. Alias raw models as `apimodels`, local models use default `models`
- Prefix raw model variables with `raw` (e.g., `rawSettings`, `rawEndpoint`)
- Table-driven tests with descriptive case names
- Functional options pattern for server/client configuration
- Document exported functions with Parameters/Returns sections

### Import Conventions
```go
import (
    "github.com/portainer/portainer-mcp/pkg/portainer/models"        // Default: models
    apimodels "github.com/portainer/client-api-go/v2/pkg/models"     // Alias: apimodels
)
```

### Handler Parameter Parsing

```go
parser := toolgen.NewParameterParser(request)
id, err := parser.GetInt("id", true)       // true = required
name, err := parser.GetString("name", false) // false = optional
```

## Client and Model Guidelines

### Two Client Layers

1. **Raw Client** (`github.com/portainer/client-api-go/v2`): Direct Portainer API communication. Used in integration tests as ground truth. Works with raw models from `github.com/portainer/client-api-go/v2/pkg/models`.
2. **Wrapper Client** (`pkg/portainer/client`): Simplified interface returning local models. Used by MCP handlers. The `PortainerClient` interface in `internal/mcp/server.go` defines the contract.

### Model Structure

- **Raw Models** (`github.com/portainer/client-api-go/v2/pkg/models`): Direct API mapping. Prefix variables with `raw`.
- **Local Models** (`pkg/portainer/models`): Simplified structs with `ConvertXXX()` functions from raw models. Only relevant fields.

## Adding a New Feature

1. Add tool definition to `internal/tooldef/tools.yaml`
2. Add tool name constant to `internal/mcp/schema.go`
3. Create or extend handler file in `internal/mcp/` with `AddXXXFeatures()` and handler methods
4. Add wrapper client methods to `pkg/portainer/client/` if new API calls are needed
5. Create local models in `pkg/portainer/models/` with `ConvertXXX()` functions
6. Add mock methods to `internal/mcp/mocks_test.go` and write unit tests
7. Write integration tests in `tests/integration/`
8. Register the feature group in `cmd/portainer-mcp/mcp.go` via `server.AddXXXFeatures()`

## Integration Tests

- Uses `testcontainers-go` to spin up real Portainer instances in Docker
- `tests/integration/helpers/test_env.go` sets up test environment with both raw client and MCP server
- Tests compare MCP handler output against direct raw API calls as ground truth
- Each test gets an isolated environment with automatic cleanup
- **Unit Tests**: Mock the `PortainerClient` interface, verify conversions and expected local model output
- **Integration Tests**: Call MCP handler and compare with ground-truth from raw client

## Design Documentation

- Design decisions are in `docs/design/` with naming convention `YYYYMM-N-short-description.md`
- Summary table in `docs/design_summary.md` — add new entries there
- Review existing decisions before making significant architectural changes
