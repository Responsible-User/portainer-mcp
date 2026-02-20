package mcp

import "slices"

// Tool names as defined in the YAML file
const (
	// Access Groups
	ToolCreateAccessGroup                = "createAccessGroup"
	ToolListAccessGroups                 = "listAccessGroups"
	ToolUpdateAccessGroupName            = "updateAccessGroupName"
	ToolUpdateAccessGroupUserAccesses    = "updateAccessGroupUserAccesses"
	ToolUpdateAccessGroupTeamAccesses    = "updateAccessGroupTeamAccesses"
	ToolAddEnvironmentToAccessGroup      = "addEnvironmentToAccessGroup"
	ToolRemoveEnvironmentFromAccessGroup = "removeEnvironmentFromAccessGroup"
	ToolDeleteAccessGroup                = "deleteAccessGroup"

	// Environments
	ToolListEnvironments              = "listEnvironments"
	ToolUpdateEnvironment             = "updateEnvironment"
	ToolUpdateEnvironmentTags         = "updateEnvironmentTags"
	ToolUpdateEnvironmentUserAccesses = "updateEnvironmentUserAccesses"
	ToolUpdateEnvironmentTeamAccesses = "updateEnvironmentTeamAccesses"
	ToolListAgentVersions             = "listAgentVersions"

	// Environment Groups (Edge Groups)
	ToolCreateEnvironmentGroup             = "createEnvironmentGroup"
	ToolListEnvironmentGroups              = "listEnvironmentGroups"
	ToolUpdateEnvironmentGroupName         = "updateEnvironmentGroupName"
	ToolUpdateEnvironmentGroupEnvironments = "updateEnvironmentGroupEnvironments"
	ToolUpdateEnvironmentGroupTags         = "updateEnvironmentGroupTags"
	ToolDeleteEnvironmentGroup             = "deleteEnvironmentGroup"

	// Edge Stacks
	ToolListStacks   = "listStacks"
	ToolGetStackFile = "getStackFile"
	ToolCreateStack  = "createStack"
	ToolUpdateStack  = "updateStack"
	ToolDeleteStack  = "deleteStack"

	// Docker Stacks
	ToolListDockerStacks   = "listDockerStacks"
	ToolGetDockerStackFile = "getDockerStackFile"
	ToolCreateDockerStack  = "createDockerStack"
	ToolUpdateDockerStack  = "updateDockerStack"
	ToolDeleteDockerStack  = "deleteDockerStack"
	ToolStartDockerStack   = "startDockerStack"
	ToolStopDockerStack    = "stopDockerStack"

	// Tags
	ToolCreateEnvironmentTag = "createEnvironmentTag"
	ToolListEnvironmentTags  = "listEnvironmentTags"
	ToolDeleteTag            = "deleteTag"

	// Teams
	ToolCreateTeam        = "createTeam"
	ToolListTeams         = "listTeams"
	ToolUpdateTeamName    = "updateTeamName"
	ToolUpdateTeamMembers = "updateTeamMembers"
	ToolDeleteTeam        = "deleteTeam"

	// Users
	ToolListUsers      = "listUsers"
	ToolUpdateUserRole = "updateUserRole"

	// Settings
	ToolGetSettings    = "getSettings"
	ToolUpdateSettings = "updateSettings"

	// Registries
	ToolListRegistries         = "listRegistries"
	ToolCreateRegistry         = "createRegistry"
	ToolDeleteRegistry         = "deleteRegistry"
	ToolTestRegistryConnection = "testRegistryConnection"

	// Edge Jobs
	ToolListEdgeJobs  = "listEdgeJobs"
	ToolGetEdgeJob    = "getEdgeJob"
	ToolCreateEdgeJob = "createEdgeJob"
	ToolDeleteEdgeJob = "deleteEdgeJob"

	// Custom Templates
	ToolListCustomTemplates  = "listCustomTemplates"
	ToolCreateCustomTemplate = "createCustomTemplate"
	ToolDeleteCustomTemplate = "deleteCustomTemplate"

	// Webhooks
	ToolListWebhooks  = "listWebhooks"
	ToolCreateWebhook = "createWebhook"
	ToolDeleteWebhook = "deleteWebhook"

	// Git Credentials
	ToolListGitCredentials  = "listGitCredentials"
	ToolGetGitCredential    = "getGitCredential"
	ToolCreateGitCredential = "createGitCredential"
	ToolUpdateGitCredential = "updateGitCredential"
	ToolDeleteGitCredential = "deleteGitCredential"

	// Alerting
	ToolListAlerts          = "listAlerts"
	ToolListAlertRules      = "listAlertRules"
	ToolGetAlertRule        = "getAlertRule"
	ToolUpdateAlertRule     = "updateAlertRule"
	ToolDeleteAlertRule     = "deleteAlertRule"
	ToolGetAlertingSettings = "getAlertingSettings"
	ToolCreateAlertSilence  = "createAlertSilence"
	ToolDeleteAlertSilence  = "deleteAlertSilence"

	// Policies
	ToolListPolicies        = "listPolicies"
	ToolGetPolicy           = "getPolicy"
	ToolCreatePolicy        = "createPolicy"
	ToolUpdatePolicy        = "updatePolicy"
	ToolDeletePolicy        = "deletePolicy"
	ToolListPolicyTemplates = "listPolicyTemplates"
	ToolGetPolicyTemplate   = "getPolicyTemplate"
	ToolGetPolicyMetadata   = "getPolicyMetadata"
	ToolGetPolicyConflicts  = "getPolicyConflicts"

	// Kubernetes Custom Resources
	ToolListCustomResourceDefinitions  = "listCustomResourceDefinitions"
	ToolGetCustomResourceDefinition    = "getCustomResourceDefinition"
	ToolDeleteCustomResourceDefinition = "deleteCustomResourceDefinition"
	ToolListCustomResources            = "listCustomResources"
	ToolGetCustomResource              = "getCustomResource"
	ToolDeleteCustomResource           = "deleteCustomResource"

	// Docker Proxy
	ToolDockerProxy = "dockerProxy"

	// Kubernetes Proxy
	ToolKubernetesProxy         = "kubernetesProxy"
	ToolKubernetesProxyStripped = "getKubernetesResourceStripped"
)

// Access levels for users and teams
const (
	// AccessLevelEnvironmentAdmin represents the environment administrator access level
	AccessLevelEnvironmentAdmin = "environment_administrator"
	// AccessLevelHelpdeskUser represents the helpdesk user access level
	AccessLevelHelpdeskUser = "helpdesk_user"
	// AccessLevelStandardUser represents the standard user access level
	AccessLevelStandardUser = "standard_user"
	// AccessLevelReadonlyUser represents the readonly user access level
	AccessLevelReadonlyUser = "readonly_user"
	// AccessLevelOperatorUser represents the operator user access level
	AccessLevelOperatorUser = "operator_user"
)

// User roles
const (
	// UserRoleAdmin represents an admin user role
	UserRoleAdmin = "admin"
	// UserRoleUser represents a regular user role
	UserRoleUser = "user"
	// UserRoleEdgeAdmin represents an edge admin user role
	UserRoleEdgeAdmin = "edge_admin"
)

// All available access levels
var AllAccessLevels = []string{
	AccessLevelEnvironmentAdmin,
	AccessLevelHelpdeskUser,
	AccessLevelStandardUser,
	AccessLevelReadonlyUser,
	AccessLevelOperatorUser,
}

// All available user roles
var AllUserRoles = []string{
	UserRoleAdmin,
	UserRoleUser,
	UserRoleEdgeAdmin,
}

// isValidAccessLevel checks if a given string is a valid access level
func isValidAccessLevel(access string) bool {
	return slices.Contains(AllAccessLevels, access)
}

// isValidUserRole checks if a given string is a valid user role
func isValidUserRole(role string) bool {
	return slices.Contains(AllUserRoles, role)
}
