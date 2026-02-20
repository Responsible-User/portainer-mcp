package mcp

import (
	"encoding/json"
	"net/http"

	"github.com/portainer/portainer-mcp/pkg/portainer/models"
	"github.com/stretchr/testify/mock"
)

// Mock Implementation Patterns:
//
// This file contains mock implementations of the PortainerClient interface.
// The following patterns are used throughout the mocks:
//
// 1. Methods returning (T, error):
//    - Uses m.Called() to record the method call and get mock behavior
//    - Includes nil check on first return value to avoid type assertion panics
//    - Example:
//      func (m *Mock) Method() (T, error) {
//          args := m.Called()
//          if args.Get(0) == nil {
//              return nil, args.Error(1)
//          }
//          return args.Get(0).(T), args.Error(1)
//      }
//
// 2. Methods returning only error:
//    - Uses m.Called() with any parameters
//    - Returns only the error value
//    - Example:
//      func (m *Mock) Method(param string) error {
//          args := m.Called(param)
//          return args.Error(0)
//      }
//
// Usage in Tests:
//   mock := new(MockPortainerClient)
//   mock.On("MethodName").Return(expectedValue, nil)
//   result, err := mock.MethodName()
//   mock.AssertExpectations(t)

// MockPortainerClient is a mock implementation of the PortainerClient interface
type MockPortainerClient struct {
	mock.Mock
}

// Tag methods

func (m *MockPortainerClient) GetEnvironmentTags() ([]models.EnvironmentTag, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.EnvironmentTag), args.Error(1)
}

func (m *MockPortainerClient) CreateEnvironmentTag(name string) (int, error) {
	args := m.Called(name)
	return args.Int(0), args.Error(1)
}

func (m *MockPortainerClient) DeleteTag(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Environment methods

func (m *MockPortainerClient) GetEnvironments() ([]models.Environment, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Environment), args.Error(1)
}

func (m *MockPortainerClient) UpdateEnvironmentTags(id int, tagIds []int) error {
	args := m.Called(id, tagIds)
	return args.Error(0)
}

func (m *MockPortainerClient) UpdateEnvironmentUserAccesses(id int, userAccesses map[int]string) error {
	args := m.Called(id, userAccesses)
	return args.Error(0)
}

func (m *MockPortainerClient) UpdateEnvironmentTeamAccesses(id int, teamAccesses map[int]string) error {
	args := m.Called(id, teamAccesses)
	return args.Error(0)
}

func (m *MockPortainerClient) UpdateEnvironment(id int, name, publicURL string, groupID int) error {
	args := m.Called(id, name, publicURL, groupID)
	return args.Error(0)
}

func (m *MockPortainerClient) GetAgentVersions() ([]string, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]string), args.Error(1)
}

// Environment Group methods

func (m *MockPortainerClient) GetEnvironmentGroups() ([]models.Group, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Group), args.Error(1)
}

func (m *MockPortainerClient) CreateEnvironmentGroup(name string, environmentIds []int) (int, error) {
	args := m.Called(name, environmentIds)
	return args.Int(0), args.Error(1)
}

func (m *MockPortainerClient) UpdateEnvironmentGroupName(id int, name string) error {
	args := m.Called(id, name)
	return args.Error(0)
}

func (m *MockPortainerClient) UpdateEnvironmentGroupEnvironments(id int, environmentIds []int) error {
	args := m.Called(id, environmentIds)
	return args.Error(0)
}

func (m *MockPortainerClient) UpdateEnvironmentGroupTags(id int, tagIds []int) error {
	args := m.Called(id, tagIds)
	return args.Error(0)
}

func (m *MockPortainerClient) DeleteEnvironmentGroup(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Access Group methods

func (m *MockPortainerClient) GetAccessGroups() ([]models.AccessGroup, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.AccessGroup), args.Error(1)
}

func (m *MockPortainerClient) CreateAccessGroup(name string, environmentIds []int) (int, error) {
	args := m.Called(name, environmentIds)
	return args.Int(0), args.Error(1)
}

func (m *MockPortainerClient) UpdateAccessGroupName(id int, name string) error {
	args := m.Called(id, name)
	return args.Error(0)
}

func (m *MockPortainerClient) UpdateAccessGroupUserAccesses(id int, userAccesses map[int]string) error {
	args := m.Called(id, userAccesses)
	return args.Error(0)
}

func (m *MockPortainerClient) UpdateAccessGroupTeamAccesses(id int, teamAccesses map[int]string) error {
	args := m.Called(id, teamAccesses)
	return args.Error(0)
}

func (m *MockPortainerClient) AddEnvironmentToAccessGroup(id int, environmentId int) error {
	args := m.Called(id, environmentId)
	return args.Error(0)
}

func (m *MockPortainerClient) RemoveEnvironmentFromAccessGroup(id int, environmentId int) error {
	args := m.Called(id, environmentId)
	return args.Error(0)
}

func (m *MockPortainerClient) DeleteAccessGroup(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Stack methods

func (m *MockPortainerClient) GetStacks() ([]models.Stack, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Stack), args.Error(1)
}

func (m *MockPortainerClient) GetStackFile(id int) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func (m *MockPortainerClient) CreateStack(name string, file string, environmentGroupIds []int) (int, error) {
	args := m.Called(name, file, environmentGroupIds)
	return args.Int(0), args.Error(1)
}

func (m *MockPortainerClient) UpdateStack(id int, file string, environmentGroupIds []int) error {
	args := m.Called(id, file, environmentGroupIds)
	return args.Error(0)
}

func (m *MockPortainerClient) DeleteEdgeStack(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Docker Stack methods

func (m *MockPortainerClient) GetDockerStacks() ([]models.DockerStack, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.DockerStack), args.Error(1)
}

func (m *MockPortainerClient) GetDockerStackFile(id int) (string, error) {
	args := m.Called(id)
	return args.String(0), args.Error(1)
}

func (m *MockPortainerClient) CreateDockerStack(endpointID int, name, composeFileContent string, env []models.StackEnvVar) (int, error) {
	args := m.Called(endpointID, name, composeFileContent, env)
	return args.Int(0), args.Error(1)
}

func (m *MockPortainerClient) UpdateDockerStack(id, endpointID int, composeFileContent string, env []models.StackEnvVar, prune, pullImage bool) error {
	args := m.Called(id, endpointID, composeFileContent, env, prune, pullImage)
	return args.Error(0)
}

func (m *MockPortainerClient) DeleteDockerStack(id, endpointID int) error {
	args := m.Called(id, endpointID)
	return args.Error(0)
}

func (m *MockPortainerClient) StartDockerStack(id, endpointID int) error {
	args := m.Called(id, endpointID)
	return args.Error(0)
}

func (m *MockPortainerClient) StopDockerStack(id, endpointID int) error {
	args := m.Called(id, endpointID)
	return args.Error(0)
}

// Team methods

func (m *MockPortainerClient) CreateTeam(name string) (int, error) {
	args := m.Called(name)
	return args.Int(0), args.Error(1)
}

func (m *MockPortainerClient) GetTeams() ([]models.Team, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Team), args.Error(1)
}

func (m *MockPortainerClient) UpdateTeamName(id int, name string) error {
	args := m.Called(id, name)
	return args.Error(0)
}

func (m *MockPortainerClient) UpdateTeamMembers(id int, userIds []int) error {
	args := m.Called(id, userIds)
	return args.Error(0)
}

func (m *MockPortainerClient) DeleteTeam(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// User methods

func (m *MockPortainerClient) GetUsers() ([]models.User, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockPortainerClient) UpdateUserRole(id int, role string) error {
	args := m.Called(id, role)
	return args.Error(0)
}

// Settings methods

func (m *MockPortainerClient) GetSettings() (models.PortainerSettings, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return models.PortainerSettings{}, args.Error(1)
	}
	return args.Get(0).(models.PortainerSettings), args.Error(1)
}

func (m *MockPortainerClient) UpdateSettings(settingsJSON string) error {
	args := m.Called(settingsJSON)
	return args.Error(0)
}

func (m *MockPortainerClient) GetVersion() (string, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return "", args.Error(1)
	}
	return args.Get(0).(string), args.Error(1)
}

// Registry methods

func (m *MockPortainerClient) GetRegistries() ([]models.Registry, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Registry), args.Error(1)
}

func (m *MockPortainerClient) CreateRegistry(req models.RegistryCreateRequest) (int, error) {
	args := m.Called(req)
	return args.Int(0), args.Error(1)
}

func (m *MockPortainerClient) DeleteRegistry(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockPortainerClient) PingRegistry(req models.RegistryPingRequest) (models.RegistryPingResponse, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return models.RegistryPingResponse{}, args.Error(1)
	}
	return args.Get(0).(models.RegistryPingResponse), args.Error(1)
}

// Edge Job methods

func (m *MockPortainerClient) GetEdgeJobs() ([]models.EdgeJob, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.EdgeJob), args.Error(1)
}

func (m *MockPortainerClient) GetEdgeJob(id int) (models.EdgeJob, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return models.EdgeJob{}, args.Error(1)
	}
	return args.Get(0).(models.EdgeJob), args.Error(1)
}

func (m *MockPortainerClient) CreateEdgeJob(req models.EdgeJobCreateRequest) (int, error) {
	args := m.Called(req)
	return args.Int(0), args.Error(1)
}

func (m *MockPortainerClient) DeleteEdgeJob(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Custom Template methods

func (m *MockPortainerClient) GetCustomTemplates() ([]models.CustomTemplate, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.CustomTemplate), args.Error(1)
}

func (m *MockPortainerClient) CreateCustomTemplate(req models.CustomTemplateCreateRequest) (int, error) {
	args := m.Called(req)
	return args.Int(0), args.Error(1)
}

func (m *MockPortainerClient) DeleteCustomTemplate(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Webhook methods

func (m *MockPortainerClient) GetWebhooks() ([]models.Webhook, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Webhook), args.Error(1)
}

func (m *MockPortainerClient) CreateWebhook(req models.WebhookCreateRequest) (int, error) {
	args := m.Called(req)
	return args.Int(0), args.Error(1)
}

func (m *MockPortainerClient) DeleteWebhook(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Git Credential methods

func (m *MockPortainerClient) GetGitCredentials() ([]models.GitCredential, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.GitCredential), args.Error(1)
}

func (m *MockPortainerClient) GetGitCredential(id int) (models.GitCredential, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return models.GitCredential{}, args.Error(1)
	}
	return args.Get(0).(models.GitCredential), args.Error(1)
}

func (m *MockPortainerClient) CreateGitCredential(req models.GitCredentialCreateRequest) (int, error) {
	args := m.Called(req)
	return args.Int(0), args.Error(1)
}

func (m *MockPortainerClient) UpdateGitCredential(id int, req models.GitCredentialUpdateRequest) error {
	args := m.Called(id, req)
	return args.Error(0)
}

func (m *MockPortainerClient) DeleteGitCredential(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

// Alerting methods

func (m *MockPortainerClient) GetAlerts(status string) (json.RawMessage, error) {
	args := m.Called(status)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(json.RawMessage), args.Error(1)
}

func (m *MockPortainerClient) GetAlertRules() ([]models.AlertingRule, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.AlertingRule), args.Error(1)
}

func (m *MockPortainerClient) GetAlertRule(id int) (models.AlertingRule, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return models.AlertingRule{}, args.Error(1)
	}
	return args.Get(0).(models.AlertingRule), args.Error(1)
}

func (m *MockPortainerClient) UpdateAlertRule(id int, ruleJSON string) error {
	args := m.Called(id, ruleJSON)
	return args.Error(0)
}

func (m *MockPortainerClient) DeleteAlertRule(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockPortainerClient) GetAlertingSettings() ([]models.AlertingSettings, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.AlertingSettings), args.Error(1)
}

func (m *MockPortainerClient) CreateAlertSilence(silenceJSON string, alertManagerURL string) error {
	args := m.Called(silenceJSON, alertManagerURL)
	return args.Error(0)
}

func (m *MockPortainerClient) DeleteAlertSilence(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

// Policy methods

func (m *MockPortainerClient) GetPolicies() ([]models.Policy, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Policy), args.Error(1)
}

func (m *MockPortainerClient) GetPolicy(id int) (models.Policy, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return models.Policy{}, args.Error(1)
	}
	return args.Get(0).(models.Policy), args.Error(1)
}

func (m *MockPortainerClient) CreatePolicy(req models.PolicyCreateRequest) (int, error) {
	args := m.Called(req)
	return args.Int(0), args.Error(1)
}

func (m *MockPortainerClient) UpdatePolicy(id int, req models.PolicyUpdateRequest) error {
	args := m.Called(id, req)
	return args.Error(0)
}

func (m *MockPortainerClient) DeletePolicy(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockPortainerClient) GetPolicyTemplates(category, policyType string) ([]models.PolicyTemplate, error) {
	args := m.Called(category, policyType)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.PolicyTemplate), args.Error(1)
}

func (m *MockPortainerClient) GetPolicyTemplate(id string) (models.PolicyTemplate, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return models.PolicyTemplate{}, args.Error(1)
	}
	return args.Get(0).(models.PolicyTemplate), args.Error(1)
}

func (m *MockPortainerClient) GetPolicyMetadata() (models.PolicyMetadata, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return models.PolicyMetadata{}, args.Error(1)
	}
	return args.Get(0).(models.PolicyMetadata), args.Error(1)
}

func (m *MockPortainerClient) GetPolicyConflicts(req models.PolicyConflictsRequest) (models.PolicyConflictsResponse, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return models.PolicyConflictsResponse{}, args.Error(1)
	}
	return args.Get(0).(models.PolicyConflictsResponse), args.Error(1)
}

// Kubernetes Custom Resource methods

func (m *MockPortainerClient) ListCustomResourceDefinitions(environmentID int) ([]models.CustomResourceDefinition, error) {
	args := m.Called(environmentID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.CustomResourceDefinition), args.Error(1)
}

func (m *MockPortainerClient) GetCustomResourceDefinition(environmentID int, name string) (models.CustomResourceDefinition, error) {
	args := m.Called(environmentID, name)
	if args.Get(0) == nil {
		return models.CustomResourceDefinition{}, args.Error(1)
	}
	return args.Get(0).(models.CustomResourceDefinition), args.Error(1)
}

func (m *MockPortainerClient) DeleteCustomResourceDefinition(environmentID int, name string) error {
	args := m.Called(environmentID, name)
	return args.Error(0)
}

func (m *MockPortainerClient) ListCustomResources(environmentID int, definition string) ([]models.CustomResource, error) {
	args := m.Called(environmentID, definition)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.CustomResource), args.Error(1)
}

func (m *MockPortainerClient) GetCustomResource(environmentID int, namespace, name, definition, format string) (json.RawMessage, error) {
	args := m.Called(environmentID, namespace, name, definition, format)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(json.RawMessage), args.Error(1)
}

func (m *MockPortainerClient) DeleteCustomResource(environmentID int, namespace, name, definition string) error {
	args := m.Called(environmentID, namespace, name, definition)
	return args.Error(0)
}

// Docker Proxy methods
func (m *MockPortainerClient) ProxyDockerRequest(opts models.DockerProxyRequestOptions) (*http.Response, error) {
	args := m.Called(opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*http.Response), args.Error(1)
}

// Kubernetes Proxy methods
func (m *MockPortainerClient) ProxyKubernetesRequest(opts models.KubernetesProxyRequestOptions) (*http.Response, error) {
	args := m.Called(opts)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*http.Response), args.Error(1)
}
