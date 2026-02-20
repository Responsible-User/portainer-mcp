package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/portainer/portainer-mcp/pkg/portainer/models"
)

// policyListResponse wraps the API response for listing policies.
type policyListResponse struct {
	Policies []models.Policy `json:"policies"`
}

// GetPolicies retrieves all fleetwide policies.
//
// Returns:
//   - A slice of Policy objects
//   - An error if the operation fails
func (c *PortainerClient) GetPolicies() ([]models.Policy, error) {
	var result policyListResponse
	if err := c.doJSONAPIRequest(http.MethodGet, "/policies", nil, &result); err != nil {
		return nil, fmt.Errorf("failed to list policies: %w", err)
	}

	return result.Policies, nil
}

// GetPolicy retrieves a specific policy by ID.
//
// Parameters:
//   - id: The ID of the policy to retrieve
//
// Returns:
//   - The Policy object
//   - An error if the operation fails
func (c *PortainerClient) GetPolicy(id int) (models.Policy, error) {
	var policy models.Policy
	if err := c.doJSONAPIRequest(http.MethodGet, fmt.Sprintf("/policies/%d", id), nil, &policy); err != nil {
		return models.Policy{}, fmt.Errorf("failed to get policy: %w", err)
	}

	return policy, nil
}

// CreatePolicy creates a new fleetwide policy.
//
// Parameters:
//   - req: The policy creation request
//
// Returns:
//   - The ID of the created policy
//   - An error if the operation fails
func (c *PortainerClient) CreatePolicy(req models.PolicyCreateRequest) (int, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal create request: %w", err)
	}

	var result models.Policy
	if err := c.doJSONAPIRequest(http.MethodPost, "/policies", bytes.NewReader(body), &result); err != nil {
		return 0, fmt.Errorf("failed to create policy: %w", err)
	}

	return result.ID, nil
}

// UpdatePolicy updates an existing fleetwide policy.
//
// Parameters:
//   - id: The ID of the policy to update
//   - req: The policy update request
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) UpdatePolicy(id int, req models.PolicyUpdateRequest) error {
	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal update request: %w", err)
	}

	if err := c.doJSONAPIRequest(http.MethodPut, fmt.Sprintf("/policies/%d", id), bytes.NewReader(body), nil); err != nil {
		return fmt.Errorf("failed to update policy: %w", err)
	}

	return nil
}

// DeletePolicy deletes a fleetwide policy.
//
// Parameters:
//   - id: The ID of the policy to delete
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeletePolicy(id int) error {
	if err := c.doAPIDelete(fmt.Sprintf("/policies/%d", id)); err != nil {
		return fmt.Errorf("failed to delete policy: %w", err)
	}

	return nil
}

// templateListResponse wraps the API response for listing policy templates.
type templateListResponse struct {
	Templates []models.PolicyTemplate `json:"templates"`
}

// GetPolicyTemplates retrieves all policy templates, optionally filtered by category and type.
//
// Parameters:
//   - category: Optional filter by policy category (e.g., "security", "rbac")
//   - policyType: Optional filter by policy type (e.g., "security-k8s")
//
// Returns:
//   - A slice of PolicyTemplate objects
//   - An error if the operation fails
func (c *PortainerClient) GetPolicyTemplates(category, policyType string) ([]models.PolicyTemplate, error) {
	path := "/policies/templates"
	params := []string{}

	if category != "" {
		params = append(params, "category="+category)
	}
	if policyType != "" {
		params = append(params, "type="+policyType)
	}
	if len(params) > 0 {
		path += "?"
		for i, p := range params {
			if i > 0 {
				path += "&"
			}
			path += p
		}
	}

	var result templateListResponse
	if err := c.doJSONAPIRequest(http.MethodGet, path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to list policy templates: %w", err)
	}

	return result.Templates, nil
}

// GetPolicyTemplate retrieves a specific policy template by ID.
//
// Parameters:
//   - id: The string ID of the policy template to retrieve
//
// Returns:
//   - The PolicyTemplate object
//   - An error if the operation fails
func (c *PortainerClient) GetPolicyTemplate(id string) (models.PolicyTemplate, error) {
	var template models.PolicyTemplate
	if err := c.doJSONAPIRequest(http.MethodGet, fmt.Sprintf("/policies/templates/%s", id), nil, &template); err != nil {
		return models.PolicyTemplate{}, fmt.Errorf("failed to get policy template: %w", err)
	}

	return template, nil
}

// policyMetadataResponse wraps the API response for policy metadata.
type policyMetadataResponse struct {
	MinimumAgentVersions map[string]string `json:"minimumAgentVersions"`
}

// GetPolicyMetadata retrieves policy metadata including minimum agent versions.
//
// Returns:
//   - The PolicyMetadata object
//   - An error if the operation fails
func (c *PortainerClient) GetPolicyMetadata() (models.PolicyMetadata, error) {
	var result policyMetadataResponse
	if err := c.doJSONAPIRequest(http.MethodGet, "/policies/metadata", nil, &result); err != nil {
		return models.PolicyMetadata{}, fmt.Errorf("failed to get policy metadata: %w", err)
	}

	return models.PolicyMetadata{
		MinimumAgentVersions: result.MinimumAgentVersions,
	}, nil
}

// GetPolicyConflicts checks for policy conflicts based on a proposed policy configuration.
//
// Parameters:
//   - req: The policy conflicts request containing the proposed policy configuration
//
// Returns:
//   - The PolicyConflictsResponse with conflict details
//   - An error if the operation fails
func (c *PortainerClient) GetPolicyConflicts(req models.PolicyConflictsRequest) (models.PolicyConflictsResponse, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return models.PolicyConflictsResponse{}, fmt.Errorf("failed to marshal conflicts request: %w", err)
	}

	var result models.PolicyConflictsResponse
	if err := c.doJSONAPIRequest(http.MethodPost, "/policies/conflicts", bytes.NewReader(body), &result); err != nil {
		return models.PolicyConflictsResponse{}, fmt.Errorf("failed to get policy conflicts: %w", err)
	}

	return result, nil
}
