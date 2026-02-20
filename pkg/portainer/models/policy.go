package models

import "encoding/json"

// Policy represents a fleetwide policy in Portainer.
type Policy struct {
	ID                int             `json:"Id"`
	Name              string          `json:"Name"`
	Type              string          `json:"Type"`
	EnvironmentType   string          `json:"EnvironmentType"`
	EnvironmentGroups []int           `json:"EnvironmentGroups"`
	Data              json.RawMessage `json:"Data,omitempty"`
	CreatedAt         string          `json:"CreatedAt"`
	UpdatedAt         string          `json:"UpdatedAt"`
}

// PolicyCreateRequest represents the request body for creating a policy.
type PolicyCreateRequest struct {
	Name              string          `json:"name"`
	Type              string          `json:"type"`
	EnvironmentType   string          `json:"environmentType"`
	EnvironmentGroups []int           `json:"environmentGroups,omitempty"`
	Data              json.RawMessage `json:"data,omitempty"`
}

// PolicyUpdateRequest represents the request body for updating a policy.
type PolicyUpdateRequest struct {
	Name              string          `json:"name,omitempty"`
	Type              string          `json:"type,omitempty"`
	EnvironmentType   string          `json:"environmentType,omitempty"`
	EnvironmentGroups []int           `json:"environmentGroups,omitempty"`
	Data              json.RawMessage `json:"data,omitempty"`
}

// PolicyTemplate represents a preconfigured policy template.
type PolicyTemplate struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Type        string          `json:"type"`
	Category    string          `json:"category"`
	Data        json.RawMessage `json:"data,omitempty"`
}

// PolicyMetadata represents policy metadata including minimum agent versions.
type PolicyMetadata struct {
	MinimumAgentVersions map[string]string `json:"minimumAgentVersions"`
}

// PolicyConflictInfo represents a conflict between policies.
type PolicyConflictInfo struct {
	EnvironmentCount        int    `json:"environmentCount"`
	EnvironmentGroupID      int    `json:"environmentGroupId"`
	EnvironmentGroupName    string `json:"environmentGroupName"`
	ExistingPolicyID        int    `json:"existingPolicyId"`
	ExistingPolicyName      string `json:"existingPolicyName"`
	SupportedEnvironments   int    `json:"supportedEnvironments"`
	UnsupportedEnvironments int    `json:"unsupportedEnvironments"`
}

// PolicyNewGroupInfo represents information about new groups in a conflicts response.
type PolicyNewGroupInfo struct {
	EnvironmentCount        int    `json:"environmentCount"`
	EnvironmentGroupID      int    `json:"environmentGroupId"`
	EnvironmentGroupName    string `json:"environmentGroupName"`
	SupportedEnvironments   int    `json:"supportedEnvironments"`
	UnsupportedEnvironments int    `json:"unsupportedEnvironments"`
}

// PolicyConflictsResponse represents the response from the policy conflicts endpoint.
type PolicyConflictsResponse struct {
	Conflicts               []PolicyConflictInfo `json:"conflicts"`
	NewGroups               []PolicyNewGroupInfo `json:"newGroups"`
	SupportedEnvironments   int                  `json:"supportedEnvironments"`
	TotalEnvironments       int                  `json:"totalEnvironments"`
	UnsupportedEnvironments int                  `json:"unsupportedEnvironments"`
}

// PolicyConflictsRequest represents the request body for checking policy conflicts.
type PolicyConflictsRequest struct {
	Name              string          `json:"name"`
	Type              string          `json:"type"`
	EnvironmentType   string          `json:"environmentType"`
	EnvironmentGroups []int           `json:"environmentGroups,omitempty"`
	Data              json.RawMessage `json:"data,omitempty"`
}
