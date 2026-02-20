package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/portainer/portainer-mcp/pkg/portainer/models"
)

// GetGitCredentials retrieves all shared git credentials from the Portainer server.
//
// Returns:
//   - A slice of GitCredential objects
//   - An error if the operation fails
func (c *PortainerClient) GetGitCredentials() ([]models.GitCredential, error) {
	var creds []models.GitCredential
	if err := c.doJSONAPIRequest(http.MethodGet, "/cloud/gitcredentials", nil, &creds); err != nil {
		return nil, fmt.Errorf("failed to list git credentials: %w", err)
	}

	return creds, nil
}

// GetGitCredential retrieves a specific shared git credential by ID.
//
// Parameters:
//   - id: The ID of the git credential to retrieve
//
// Returns:
//   - The GitCredential object
//   - An error if the operation fails
func (c *PortainerClient) GetGitCredential(id int) (models.GitCredential, error) {
	var cred models.GitCredential
	if err := c.doJSONAPIRequest(http.MethodGet, fmt.Sprintf("/cloud/gitcredentials/%d", id), nil, &cred); err != nil {
		return models.GitCredential{}, fmt.Errorf("failed to get git credential: %w", err)
	}

	return cred, nil
}

// CreateGitCredential creates a new shared git credential.
//
// Parameters:
//   - req: The git credential creation request
//
// Returns:
//   - The ID of the created git credential
//   - An error if the operation fails
func (c *PortainerClient) CreateGitCredential(req models.GitCredentialCreateRequest) (int, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal create request: %w", err)
	}

	var result struct {
		ID int `json:"Id"`
	}
	if err := c.doJSONAPIRequest(http.MethodPost, "/cloud/gitcredentials", bytes.NewReader(body), &result); err != nil {
		return 0, fmt.Errorf("failed to create git credential: %w", err)
	}

	return result.ID, nil
}

// UpdateGitCredential updates an existing shared git credential.
//
// Parameters:
//   - id: The ID of the git credential to update
//   - req: The git credential update request
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) UpdateGitCredential(id int, req models.GitCredentialUpdateRequest) error {
	body, err := json.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal update request: %w", err)
	}

	if err := c.doJSONAPIRequest(http.MethodPut, fmt.Sprintf("/cloud/gitcredentials/%d", id), bytes.NewReader(body), nil); err != nil {
		return fmt.Errorf("failed to update git credential: %w", err)
	}

	return nil
}

// DeleteGitCredential deletes a shared git credential.
//
// Parameters:
//   - id: The ID of the git credential to delete
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeleteGitCredential(id int) error {
	if err := c.doAPIDelete(fmt.Sprintf("/cloud/gitcredentials/%d", id)); err != nil {
		return fmt.Errorf("failed to delete git credential: %w", err)
	}

	return nil
}
