package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/portainer/portainer-mcp/pkg/portainer/models"
)

// GetRegistries retrieves all registries from the Portainer server.
//
// Returns:
//   - A slice of Registry objects
//   - An error if the operation fails
func (c *PortainerClient) GetRegistries() ([]models.Registry, error) {
	var registries []models.Registry
	if err := c.doJSONAPIRequest(http.MethodGet, "/registries", nil, &registries); err != nil {
		return nil, fmt.Errorf("failed to list registries: %w", err)
	}

	return registries, nil
}

// CreateRegistry creates a new container registry.
//
// Parameters:
//   - req: The registry creation request
//
// Returns:
//   - The ID of the created registry
//   - An error if the operation fails
func (c *PortainerClient) CreateRegistry(req models.RegistryCreateRequest) (int, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal create request: %w", err)
	}

	var result struct {
		ID int `json:"Id"`
	}
	if err := c.doJSONAPIRequest(http.MethodPost, "/registries", bytes.NewReader(body), &result); err != nil {
		return 0, fmt.Errorf("failed to create registry: %w", err)
	}

	return result.ID, nil
}

// DeleteRegistry deletes a container registry.
//
// Parameters:
//   - id: The ID of the registry to delete
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeleteRegistry(id int) error {
	if err := c.doAPIDelete(fmt.Sprintf("/registries/%d", id)); err != nil {
		return fmt.Errorf("failed to delete registry: %w", err)
	}

	return nil
}
