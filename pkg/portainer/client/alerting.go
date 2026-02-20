package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/portainer/portainer-mcp/pkg/portainer/models"
)

// GetAlerts retrieves alerts from the Portainer observability alerting system.
//
// Parameters:
//   - status: Filter by alert status ("active" or "silenced"). Empty string returns all alerts.
//
// Returns:
//   - The raw JSON response as a byte slice (alert format varies by source)
//   - An error if the operation fails
func (c *PortainerClient) GetAlerts(status string) (json.RawMessage, error) {
	path := "/observability/alerting/alerts"
	if status != "" {
		path = fmt.Sprintf("%s?status=%s", path, status)
	}

	var alerts json.RawMessage
	if err := c.doJSONAPIRequest(http.MethodGet, path, nil, &alerts); err != nil {
		return nil, fmt.Errorf("failed to list alerts: %w", err)
	}

	return alerts, nil
}

// GetAlertRules retrieves all alert rules.
//
// Returns:
//   - A slice of AlertingRule objects
//   - An error if the operation fails
func (c *PortainerClient) GetAlertRules() ([]models.AlertingRule, error) {
	var rules []models.AlertingRule
	if err := c.doJSONAPIRequest(http.MethodGet, "/observability/alerting/rules", nil, &rules); err != nil {
		return nil, fmt.Errorf("failed to list alert rules: %w", err)
	}

	return rules, nil
}

// GetAlertRule retrieves a specific alert rule by ID.
//
// Parameters:
//   - id: The ID of the alert rule to retrieve
//
// Returns:
//   - The AlertingRule object
//   - An error if the operation fails
func (c *PortainerClient) GetAlertRule(id int) (models.AlertingRule, error) {
	var rule models.AlertingRule
	if err := c.doJSONAPIRequest(http.MethodGet, fmt.Sprintf("/observability/alerting/rules/%d", id), nil, &rule); err != nil {
		return models.AlertingRule{}, fmt.Errorf("failed to get alert rule: %w", err)
	}

	return rule, nil
}

// UpdateAlertRule updates an alert rule.
// The ruleJSON parameter should contain the alert rule fields as a JSON string.
// The client wraps it in the expected API payload format: {"alertingRule": <ruleJSON>}.
//
// Parameters:
//   - id: The ID of the alert rule to update
//   - ruleJSON: JSON string containing the alert rule fields
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) UpdateAlertRule(id int, ruleJSON string) error {
	payload := fmt.Sprintf(`{"alertingRule":%s}`, ruleJSON)

	if err := c.doJSONAPIRequest(http.MethodPut, fmt.Sprintf("/observability/alerting/rules/%d", id), bytes.NewReader([]byte(payload)), nil); err != nil {
		return fmt.Errorf("failed to update alert rule: %w", err)
	}

	return nil
}

// DeleteAlertRule deletes an alert rule.
//
// Parameters:
//   - id: The ID of the alert rule to delete
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeleteAlertRule(id int) error {
	if err := c.doAPIDelete(fmt.Sprintf("/observability/alerting/rules/%d", id)); err != nil {
		return fmt.Errorf("failed to delete alert rule: %w", err)
	}

	return nil
}

// GetAlertingSettings retrieves the alerting settings (alert manager instances).
//
// Returns:
//   - A slice of AlertingSettings objects
//   - An error if the operation fails
func (c *PortainerClient) GetAlertingSettings() ([]models.AlertingSettings, error) {
	var settings []models.AlertingSettings
	if err := c.doJSONAPIRequest(http.MethodGet, "/observability/alerting/settings", nil, &settings); err != nil {
		return nil, fmt.Errorf("failed to get alerting settings: %w", err)
	}

	return settings, nil
}

// CreateAlertSilence creates a new alert silence.
// The silenceJSON parameter should contain the silence fields as a JSON string.
// The client wraps it in the expected API payload format: {"alertManagerURL": "...", "silence": <silenceJSON>}.
//
// Parameters:
//   - silenceJSON: JSON string containing the silence definition
//   - alertManagerURL: The URL of the alert manager instance (optional, can be empty)
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) CreateAlertSilence(silenceJSON string, alertManagerURL string) error {
	alertManagerURLJSON, _ := json.Marshal(alertManagerURL)
	payload := fmt.Sprintf(`{"alertManagerURL":%s,"silence":%s}`, string(alertManagerURLJSON), silenceJSON)

	if err := c.doJSONAPIRequest(http.MethodPost, "/observability/alerting/silence", bytes.NewReader([]byte(payload)), nil); err != nil {
		return fmt.Errorf("failed to create alert silence: %w", err)
	}

	return nil
}

// DeleteAlertSilence deletes an alert silence.
//
// Parameters:
//   - id: The string ID of the alert silence to delete
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeleteAlertSilence(id string) error {
	if err := c.doAPIDelete(fmt.Sprintf("/observability/alerting/silence/%s", id)); err != nil {
		return fmt.Errorf("failed to delete alert silence: %w", err)
	}

	return nil
}
