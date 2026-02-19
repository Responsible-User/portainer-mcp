package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/portainer/portainer-mcp/pkg/portainer/models"
)

// GetCustomTemplates retrieves all custom templates from the Portainer server.
//
// Returns:
//   - A slice of CustomTemplate objects
//   - An error if the operation fails
func (c *PortainerClient) GetCustomTemplates() ([]models.CustomTemplate, error) {
	var templates []models.CustomTemplate
	if err := c.doJSONAPIRequest(http.MethodGet, "/custom_templates", nil, &templates); err != nil {
		return nil, fmt.Errorf("failed to list custom templates: %w", err)
	}

	return templates, nil
}

// CreateCustomTemplate creates a new custom template from a string.
//
// Parameters:
//   - req: The custom template creation request
//
// Returns:
//   - The ID of the created custom template
//   - An error if the operation fails
func (c *PortainerClient) CreateCustomTemplate(req models.CustomTemplateCreateRequest) (int, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal create request: %w", err)
	}

	var result struct {
		ID int `json:"Id"`
	}
	if err := c.doJSONAPIRequest(http.MethodPost, "/custom_templates/create/string", bytes.NewReader(body), &result); err != nil {
		return 0, fmt.Errorf("failed to create custom template: %w", err)
	}

	return result.ID, nil
}

// DeleteCustomTemplate deletes a custom template.
//
// Parameters:
//   - id: The ID of the custom template to delete
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeleteCustomTemplate(id int) error {
	if err := c.doAPIDelete(fmt.Sprintf("/custom_templates/%d", id)); err != nil {
		return fmt.Errorf("failed to delete custom template: %w", err)
	}

	return nil
}
