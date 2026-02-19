package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/portainer/portainer-mcp/pkg/portainer/models"
)

// GetWebhooks retrieves all webhooks from the Portainer server.
//
// Returns:
//   - A slice of Webhook objects
//   - An error if the operation fails
func (c *PortainerClient) GetWebhooks() ([]models.Webhook, error) {
	var webhooks []models.Webhook
	if err := c.doJSONAPIRequest(http.MethodGet, "/webhooks", nil, &webhooks); err != nil {
		return nil, fmt.Errorf("failed to list webhooks: %w", err)
	}

	return webhooks, nil
}

// CreateWebhook creates a new webhook.
//
// Parameters:
//   - req: The webhook creation request
//
// Returns:
//   - The ID of the created webhook
//   - An error if the operation fails
func (c *PortainerClient) CreateWebhook(req models.WebhookCreateRequest) (int, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal create request: %w", err)
	}

	var result struct {
		ID int `json:"Id"`
	}
	if err := c.doJSONAPIRequest(http.MethodPost, "/webhooks", bytes.NewReader(body), &result); err != nil {
		return 0, fmt.Errorf("failed to create webhook: %w", err)
	}

	return result.ID, nil
}

// DeleteWebhook deletes a webhook.
//
// Parameters:
//   - id: The ID of the webhook to delete
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeleteWebhook(id int) error {
	if err := c.doAPIDelete(fmt.Sprintf("/webhooks/%d", id)); err != nil {
		return fmt.Errorf("failed to delete webhook: %w", err)
	}

	return nil
}
