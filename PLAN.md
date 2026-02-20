# Plan: Update Portainer MCP Server to Support Portainer 2.33 LTS

## Key Constraint

The `client-api-go` dependency has no v2.33 release (latest is v2.31.2). Since the 2.31→2.33 API changes are mostly bug fixes and the endpoint signatures are backward compatible, we will:

- **Keep `client-api-go` v2.31.2** for all existing operations
- **Add a generic Portainer API HTTP client** for new endpoints not covered by the SDK
- **Relax version checking** to accept 2.33.x alongside 2.31.x using semver minimum version comparison instead of exact match

## Phase 1: Version Compatibility

### 1a. Relax version check in `internal/mcp/server.go`

- Change `SupportedPortainerVersion` from exact `"2.31.2"` to a minimum like `"2.27.0"` (or the oldest supported)
- Change the version comparison from `!=` (exact match) to semver `<` (minimum version check) using `golang.org/x/mod/semver`
- Add a `MaxSupportedPortainerVersion = "2.33"` constant for documentation clarity
- This allows the MCP server to work with Portainer 2.31.x through 2.33.x

### 1b. Add generic Portainer API client

- Add a `PortainerAPIRequest` method to the wrapper client in `pkg/portainer/client/` that makes raw HTTP requests to `https://{host}/api/{path}` with the API key header
- This reuses the same pattern as the existing `ProxyClient` in `client-api-go` but targets Portainer API endpoints (not Docker/K8s proxied endpoints)
- Add a `PortainerAPIRequest` method to the `PortainerAPIClient` interface
- New features will use this for endpoints not covered by the swagger-generated SDK

## Phase 2: Docker Standalone Stacks

Currently only Edge Stacks are supported. Docker standalone stacks are a major use case.

### New tools

- `listDockerStacks` — `GET /api/stacks` (filtered to non-edge stacks)
- `getDockerStackFile` — `GET /api/stacks/{id}/file`
- `createDockerStack` — `POST /api/stacks/create/standalone/string` (from compose string)
- `updateDockerStack` — `PUT /api/stacks/{id}` with compose file content
- `deleteDockerStack` — `DELETE /api/stacks/{id}`
- `startDockerStack` — `POST /api/stacks/{id}/start`
- `stopDockerStack` — `POST /api/stacks/{id}/stop`

### Files to modify/create

- `internal/tooldef/tools.yaml` — add 7 tool definitions
- `internal/mcp/schema.go` — add 7 tool name constants
- `internal/mcp/docker_stack.go` (new) — handler file with `AddDockerStackFeatures()` + 7 handlers
- `pkg/portainer/client/docker_stack.go` (new) — wrapper client methods using generic API client
- `pkg/portainer/models/docker_stack.go` (new) — `DockerStack` model + conversion
- `internal/mcp/mocks_test.go` — add mock methods
- `internal/mcp/docker_stack_test.go` (new) — unit tests
- `cmd/portainer-mcp/mcp.go` — register `server.AddDockerStackFeatures()`

## Phase 3: Container Registries

### New tools

- `listRegistries` — `GET /api/registries`
- `createRegistry` — `POST /api/registries`
- `deleteRegistry` — `DELETE /api/registries/{id}`

### Files to modify/createnow go thr

- `internal/tooldef/tools.yaml` — add 3 tool definitions
- `internal/mcp/schema.go` — add 3 tool name constants
- `internal/mcp/registry.go` (new) — handler file with `AddRegistryFeatures()`
- `pkg/portainer/client/registry.go` (new) — wrapper client methods
- `pkg/portainer/models/registry.go` (new) — `Registry` model
- `internal/mcp/mocks_test.go` — add mock methods
- `internal/mcp/registry_test.go` (new) — unit tests
- `cmd/portainer-mcp/mcp.go` — register `server.AddRegistryFeatures()`

## Phase 4: Edge Jobs

### New tools:

- `listEdgeJobs` — `GET /api/edge_jobs`
- `createEdgeJob` — `POST /api/edge_jobs`
- `deleteEdgeJob` — `DELETE /api/edge_jobs/{id}`
- `getEdgeJobResults` — `GET /api/edge_jobs/{id}`

### Files:

- Same pattern: tools.yaml + schema.go + `internal/mcp/edge_job.go` + `pkg/portainer/client/edge_job.go` + `pkg/portainer/models/edge_job.go` + tests + registration

## Phase 5: Delete Operations for Existing Resources

Add delete capabilities to currently supported resource types:

### New tools:

- `deleteAccessGroup` — `DELETE /api/endpoint_groups/{id}`
- `deleteEnvironmentGroup` — `DELETE /api/edge_groups/{id}`
- `deleteStack` (edge stack) — `DELETE /api/edge_stacks/{id}`
- `deleteTag` — `DELETE /api/tags/{id}`
- `deleteTeam` — `DELETE /api/teams/{id}`

### Files:

- `internal/tooldef/tools.yaml` — add 5 tool definitions
- `internal/mcp/schema.go` — add 5 tool name constants
- Add handlers to existing files (`access_group.go`, `group.go`, `stack.go`, `tag.go`, `team.go`)
- Add wrapper client delete methods to existing client files
- Add to `PortainerClient` interface in `server.go`
- Add mock methods + unit tests

## Phase 6: Settings Update

### New tools:\]

- `updateSettings` — `PUT /api/settings`

### Files:

- `internal/tooldef/tools.yaml` — add tool definition
- `internal/mcp/schema.go` — add constant
- Extend `internal/mcp/settings.go` handler
- Add wrapper client method to `pkg/portainer/client/settings.go`
- Add to `PortainerClient` interface

## Phase 7: Custom Templates

### New tools:

- `listCustomTemplates` — `GET /api/custom_templates`
- `createCustomTemplate` — `POST /api/custom_templates/create/standalone/string`
- `deleteCustomTemplate` — `DELETE /api/custom_templates/{id}`

### Files:

- Same pattern as registries

## Phase 8: Webhooks

### New tools:

- `listWebhooks` — `GET /api/webhooks`
- `createWebhook` — `POST /api/webhooks`
- `deleteWebhook` — `DELETE /api/webhooks/{id}`

### Files:

- Same pattern

## Phase 9: Update Environment (expanded)

### New tools:

- `updateEnvironment` — `PUT /api/endpoints/{id}` (update name, URL, public IP, group assignment)

### Files:

- Extend existing environment handler and client files

## Phase 10: Tests & Build Verification

- Run `make test` to verify all unit tests pass
- Run `go vet ./...` to check for lint issues
- Run `make build` to verify the binary compiles
- Update `internal/tooldef/tools.yaml` version from `v1.2` to `v1.3` (new tools added)

## Summary of New Tools (total ~28 new tools)

| Category | New Tools | Type |
|----------|-----------|------|
| Docker Stacks | 7 | Mixed read/write |
| Registries | 3 | Mixed |
| Edge Jobs | 4 | Mixed |
| Deletes | 5 | Write |
| Settings Update | 1 | Write |
| Custom Templates | 3 | Mixed |
| Webhooks | 3 | Mixed |
| Environment Update | 1 | Write |
| **Total** | **~27** | |

## Architecture Decision: Generic API Client

The key architectural addition is a generic Portainer API client method that enables calling any Portainer API endpoint without needing it in the swagger-generated SDK. This:

1. Unblocks us from waiting for `client-api-go` v2.33
2. Follows the same HTTP pattern already used by `ProxyDockerRequest`/`ProxyKubernetesRequest`
3. Will be used for: Docker stacks, registries, edge jobs, custom templates, webhooks, settings update, and deletes
4. Existing operations continue using the swagger-generated client (no regression risk)
