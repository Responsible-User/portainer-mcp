package mcp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/portainer/portainer-mcp/pkg/portainer/client"
	"github.com/portainer/portainer-mcp/pkg/portainer/models"
	"github.com/portainer/portainer-mcp/pkg/toolgen"
	"golang.org/x/mod/semver"
)

const (
	// MinimumToolsVersion is the minimum supported version of the tools.yaml file
	MinimumToolsVersion = "1.0"
	// MinSupportedPortainerVersion is the minimum version of Portainer supported by this tool
	MinSupportedPortainerVersion = "2.27.0"
	// MaxSupportedPortainerVersion is the maximum version of Portainer supported by this tool
	MaxSupportedPortainerVersion = "2.36"
)

// PortainerClient defines the interface for the wrapper client used by the MCP server
type PortainerClient interface {
	// Tag methods
	GetEnvironmentTags() ([]models.EnvironmentTag, error)
	CreateEnvironmentTag(name string) (int, error)
	DeleteTag(id int) error

	// Environment methods
	GetEnvironments() ([]models.Environment, error)
	UpdateEnvironmentTags(id int, tagIds []int) error
	UpdateEnvironmentUserAccesses(id int, userAccesses map[int]string) error
	UpdateEnvironmentTeamAccesses(id int, teamAccesses map[int]string) error
	UpdateEnvironment(id int, name, publicURL string, groupID int) error

	// Environment Group methods
	GetEnvironmentGroups() ([]models.Group, error)
	CreateEnvironmentGroup(name string, environmentIds []int) (int, error)
	UpdateEnvironmentGroupName(id int, name string) error
	UpdateEnvironmentGroupEnvironments(id int, environmentIds []int) error
	UpdateEnvironmentGroupTags(id int, tagIds []int) error
	DeleteEnvironmentGroup(id int) error

	// Access Group methods
	GetAccessGroups() ([]models.AccessGroup, error)
	CreateAccessGroup(name string, environmentIds []int) (int, error)
	UpdateAccessGroupName(id int, name string) error
	UpdateAccessGroupUserAccesses(id int, userAccesses map[int]string) error
	UpdateAccessGroupTeamAccesses(id int, teamAccesses map[int]string) error
	AddEnvironmentToAccessGroup(id int, environmentId int) error
	RemoveEnvironmentFromAccessGroup(id int, environmentId int) error
	DeleteAccessGroup(id int) error

	// Edge Stack methods
	GetStacks() ([]models.Stack, error)
	GetStackFile(id int) (string, error)
	CreateStack(name string, file string, environmentGroupIds []int) (int, error)
	UpdateStack(id int, file string, environmentGroupIds []int) error
	DeleteEdgeStack(id int) error

	// Docker Stack methods
	GetDockerStacks() ([]models.DockerStack, error)
	GetDockerStackFile(id int) (string, error)
	CreateDockerStack(endpointID int, name, composeFileContent string, env []models.StackEnvVar) (int, error)
	UpdateDockerStack(id, endpointID int, composeFileContent string, env []models.StackEnvVar, prune, pullImage bool) error
	DeleteDockerStack(id, endpointID int) error
	StartDockerStack(id, endpointID int) error
	StopDockerStack(id, endpointID int) error

	// Team methods
	CreateTeam(name string) (int, error)
	GetTeams() ([]models.Team, error)
	UpdateTeamName(id int, name string) error
	UpdateTeamMembers(id int, userIds []int) error
	DeleteTeam(id int) error

	// User methods
	GetUsers() ([]models.User, error)
	UpdateUserRole(id int, role string) error

	// Settings methods
	GetSettings() (models.PortainerSettings, error)
	UpdateSettings(settingsJSON string) error

	// Version methods
	GetVersion() (string, error)

	// Registry methods
	GetRegistries() ([]models.Registry, error)
	CreateRegistry(req models.RegistryCreateRequest) (int, error)
	DeleteRegistry(id int) error
	PingRegistry(req models.RegistryPingRequest) (models.RegistryPingResponse, error)

	// Edge Job methods
	GetEdgeJobs() ([]models.EdgeJob, error)
	GetEdgeJob(id int) (models.EdgeJob, error)
	CreateEdgeJob(req models.EdgeJobCreateRequest) (int, error)
	DeleteEdgeJob(id int) error

	// Custom Template methods
	GetCustomTemplates() ([]models.CustomTemplate, error)
	CreateCustomTemplate(req models.CustomTemplateCreateRequest) (int, error)
	DeleteCustomTemplate(id int) error

	// Webhook methods
	GetWebhooks() ([]models.Webhook, error)
	CreateWebhook(req models.WebhookCreateRequest) (int, error)
	DeleteWebhook(id int) error

	// Git Credential methods
	GetGitCredentials() ([]models.GitCredential, error)
	GetGitCredential(id int) (models.GitCredential, error)
	CreateGitCredential(req models.GitCredentialCreateRequest) (int, error)
	UpdateGitCredential(id int, req models.GitCredentialUpdateRequest) error
	DeleteGitCredential(id int) error

	// Alerting methods
	GetAlerts(status string) (json.RawMessage, error)
	GetAlertRules() ([]models.AlertingRule, error)
	GetAlertRule(id int) (models.AlertingRule, error)
	UpdateAlertRule(id int, ruleJSON string) error
	DeleteAlertRule(id int) error
	GetAlertingSettings() ([]models.AlertingSettings, error)
	CreateAlertSilence(silenceJSON string, alertManagerURL string) error
	DeleteAlertSilence(id string) error

	// Kubernetes Custom Resource methods
	ListCustomResourceDefinitions(environmentID int) ([]models.CustomResourceDefinition, error)
	GetCustomResourceDefinition(environmentID int, name string) (models.CustomResourceDefinition, error)
	DeleteCustomResourceDefinition(environmentID int, name string) error
	ListCustomResources(environmentID int, definition string) ([]models.CustomResource, error)
	GetCustomResource(environmentID int, namespace, name, definition, format string) (json.RawMessage, error)
	DeleteCustomResource(environmentID int, namespace, name, definition string) error

	// Docker Proxy methods
	ProxyDockerRequest(opts models.DockerProxyRequestOptions) (*http.Response, error)

	// Kubernetes Proxy methods
	ProxyKubernetesRequest(opts models.KubernetesProxyRequestOptions) (*http.Response, error)
}

// PortainerMCPServer is the main server that handles MCP protocol communication
// with AI assistants and translates them into Portainer API calls.
type PortainerMCPServer struct {
	srv      *server.MCPServer
	cli      PortainerClient
	tools    map[string]mcp.Tool
	readOnly bool
}

// ServerOption is a function that configures the server
type ServerOption func(*serverOptions)

// serverOptions contains all configurable options for the server
type serverOptions struct {
	client              PortainerClient
	readOnly            bool
	disableVersionCheck bool
}

// WithClient sets a custom client for the server.
// This is primarily used for testing to inject mock clients.
func WithClient(client PortainerClient) ServerOption {
	return func(opts *serverOptions) {
		opts.client = client
	}
}

// WithReadOnly sets the server to read-only mode.
// This will prevent the server from registering write tools.
func WithReadOnly(readOnly bool) ServerOption {
	return func(opts *serverOptions) {
		opts.readOnly = readOnly
	}
}

// WithDisableVersionCheck disables the Portainer server version check.
// This allows connecting to unsupported Portainer versions.
func WithDisableVersionCheck(disable bool) ServerOption {
	return func(opts *serverOptions) {
		opts.disableVersionCheck = disable
	}
}

// NewPortainerMCPServer creates a new Portainer MCP server.
//
// This server provides an implementation of the MCP protocol for Portainer,
// allowing AI assistants to interact with Portainer through a structured API.
//
// Parameters:
//   - serverURL: The base URL of the Portainer server (e.g., "https://portainer.example.com")
//   - token: The API token for authenticating with the Portainer server
//   - toolsPath: Path to the tools.yaml file that defines the available MCP tools
//   - options: Optional functional options for customizing server behavior (e.g., WithClient)
//
// Returns:
//   - A configured PortainerMCPServer instance ready to be started
//   - An error if initialization fails
//
// Possible errors:
//   - Failed to load tools from the specified path
//   - Failed to communicate with the Portainer server
//   - Incompatible Portainer server version
func NewPortainerMCPServer(serverURL, token, toolsPath string, options ...ServerOption) (*PortainerMCPServer, error) {
	opts := &serverOptions{}

	for _, option := range options {
		option(opts)
	}

	tools, err := toolgen.LoadToolsFromYAML(toolsPath, MinimumToolsVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to load tools: %w", err)
	}

	var portainerClient PortainerClient
	if opts.client != nil {
		portainerClient = opts.client
	} else {
		portainerClient = client.NewPortainerClient(serverURL, token, client.WithSkipTLSVerify(true))
	}

	if !opts.disableVersionCheck {
		version, err := portainerClient.GetVersion()
		if err != nil {
			return nil, fmt.Errorf("failed to get Portainer server version: %w", err)
		}

		if err := checkPortainerVersion(version); err != nil {
			return nil, err
		}
	}

	return &PortainerMCPServer{
		srv: server.NewMCPServer(
			"Portainer MCP Server",
			"0.5.1",
			server.WithToolCapabilities(true),
			server.WithLogging(),
		),
		cli:      portainerClient,
		tools:    tools,
		readOnly: opts.readOnly,
	}, nil
}

// checkPortainerVersion validates that the given Portainer server version
// falls within the supported range [MinSupportedPortainerVersion, MaxSupportedPortainerVersion].
func checkPortainerVersion(version string) error {
	// semver requires a "v" prefix
	v := "v" + version
	min := "v" + MinSupportedPortainerVersion
	max := "v" + MaxSupportedPortainerVersion

	if !semver.IsValid(v) {
		return fmt.Errorf("invalid Portainer server version format: %s", version)
	}

	if semver.Compare(v, min) < 0 {
		return fmt.Errorf("unsupported Portainer server version: %s, minimum supported version is %s", version, MinSupportedPortainerVersion)
	}

	// Compare only major.minor for the max check so any patch version of 2.33.x is accepted
	if semver.Compare(semver.MajorMinor(v), semver.MajorMinor(max)) > 0 {
		return fmt.Errorf("unsupported Portainer server version: %s, maximum supported version is %s", version, MaxSupportedPortainerVersion)
	}

	return nil
}

// Start begins listening for MCP protocol messages on standard input/output.
// This is a blocking call that will run until the connection is closed.
func (s *PortainerMCPServer) Start() error {
	return server.ServeStdio(s.srv)
}

// addToolIfExists adds a tool to the server if it exists in the tools map
func (s *PortainerMCPServer) addToolIfExists(toolName string, handler server.ToolHandlerFunc) {
	if tool, exists := s.tools[toolName]; exists {
		s.srv.AddTool(tool, handler)
	} else {
		log.Printf("Tool %s not found, will not be registered for MCP usage", toolName)
	}
}
