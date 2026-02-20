package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/portainer/portainer-mcp/pkg/portainer/models"
	"github.com/portainer/portainer-mcp/pkg/portainer/utils"
)

// GetEnvironments retrieves all environments from the Portainer server.
//
// Returns:
//   - A slice of Environment objects
//   - An error if the operation fails
func (c *PortainerClient) GetEnvironments() ([]models.Environment, error) {
	endpoints, err := c.cli.ListEndpoints()
	if err != nil {
		return nil, fmt.Errorf("failed to list endpoints: %w", err)
	}

	environments := make([]models.Environment, len(endpoints))
	for i, endpoint := range endpoints {
		environments[i] = models.ConvertEndpointToEnvironment(endpoint)
	}

	return environments, nil
}

// UpdateEnvironmentTags updates the tags associated with an environment.
//
// Parameters:
//   - id: The ID of the environment to update
//   - tagIds: A slice of tag IDs to associate with the environment
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) UpdateEnvironmentTags(id int, tagIds []int) error {
	tags := utils.IntToInt64Slice(tagIds)
	err := c.cli.UpdateEndpoint(int64(id),
		&tags,
		nil,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to update environment tags: %w", err)
	}
	return nil
}

// UpdateEnvironmentUserAccesses updates the user access policies of an environment.
//
// Parameters:
//   - id: The ID of the environment to update
//   - userAccesses: Map of user IDs to their access level
//
// Valid access levels are:
//   - environment_administrator
//   - helpdesk_user
//   - standard_user
//   - readonly_user
//   - operator_user
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) UpdateEnvironmentUserAccesses(id int, userAccesses map[int]string) error {
	uac := utils.IntToInt64Map(userAccesses)
	err := c.cli.UpdateEndpoint(int64(id),
		nil,
		&uac,
		nil,
	)
	if err != nil {
		return fmt.Errorf("failed to update environment user accesses: %w", err)
	}
	return nil
}

// UpdateEnvironmentTeamAccesses updates the team access policies of an environment.
//
// Parameters:
//   - id: The ID of the environment to update
//   - teamAccesses: Map of team IDs to their access level
//
// Valid access levels are:
//   - environment_administrator
//   - helpdesk_user
//   - standard_user
//   - readonly_user
//   - operator_user
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) UpdateEnvironmentTeamAccesses(id int, teamAccesses map[int]string) error {
	tac := utils.IntToInt64Map(teamAccesses)
	err := c.cli.UpdateEndpoint(int64(id),
		nil,
		nil,
		&tac,
	)
	if err != nil {
		return fmt.Errorf("failed to update environment team accesses: %w", err)
	}
	return nil
}

// UpdateEnvironment updates an environment's name, public URL, and/or group ID.
//
// Parameters:
//   - id: The ID of the environment to update
//   - name: The new name (empty string to skip)
//   - publicURL: The new public URL (empty string to skip)
//   - groupID: The new group ID (0 to skip)
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) UpdateEnvironment(id int, name, publicURL string, groupID int) error {
	payload := map[string]any{}
	if name != "" {
		payload["name"] = name
	}
	if publicURL != "" {
		payload["publicURL"] = publicURL
	}
	if groupID > 0 {
		payload["groupID"] = groupID
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal update request: %w", err)
	}

	if err := c.doJSONAPIRequest(http.MethodPut, fmt.Sprintf("/endpoints/%d", id), bytes.NewReader(body), nil); err != nil {
		return fmt.Errorf("failed to update environment: %w", err)
	}

	return nil
}

// GetAgentVersions retrieves the list of available agent versions.
//
// Returns:
//   - A slice of version strings
//   - An error if the operation fails
func (c *PortainerClient) GetAgentVersions() ([]string, error) {
	var versions []string
	if err := c.doJSONAPIRequest(http.MethodGet, "/endpoints/agent_versions", nil, &versions); err != nil {
		return nil, fmt.Errorf("failed to list agent versions: %w", err)
	}

	return versions, nil
}
