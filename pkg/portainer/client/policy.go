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
