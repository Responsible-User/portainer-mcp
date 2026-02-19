package client

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/portainer/portainer-mcp/pkg/portainer/models"
)

func (c *PortainerClient) GetSettings() (models.PortainerSettings, error) {
	settings, err := c.cli.GetSettings()
	if err != nil {
		return models.PortainerSettings{}, fmt.Errorf("failed to get settings: %w", err)
	}

	return models.ConvertSettingsToPortainerSettings(settings), nil
}

// UpdateSettings updates the Portainer server settings with the given JSON payload.
// The settingsJSON should be a JSON-encoded string of the settings fields to update.
//
// Parameters:
//   - settingsJSON: A JSON string containing the settings fields to update
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) UpdateSettings(settingsJSON string) error {
	if err := c.doJSONAPIRequest(http.MethodPut, "/settings", bytes.NewReader([]byte(settingsJSON)), nil); err != nil {
		return fmt.Errorf("failed to update settings: %w", err)
	}

	return nil
}
