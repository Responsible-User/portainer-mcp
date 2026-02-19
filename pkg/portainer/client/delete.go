package client

import (
	"fmt"
)

// DeleteAccessGroup deletes an access group (endpoint group) from the Portainer server.
//
// Parameters:
//   - id: The ID of the access group to delete
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeleteAccessGroup(id int) error {
	if err := c.doAPIDelete(fmt.Sprintf("/endpoint_groups/%d", id)); err != nil {
		return fmt.Errorf("failed to delete access group: %w", err)
	}

	return nil
}

// DeleteEnvironmentGroup deletes an environment group (edge group) from the Portainer server.
//
// Parameters:
//   - id: The ID of the environment group to delete
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeleteEnvironmentGroup(id int) error {
	if err := c.doAPIDelete(fmt.Sprintf("/edge_groups/%d", id)); err != nil {
		return fmt.Errorf("failed to delete environment group: %w", err)
	}

	return nil
}

// DeleteEdgeStack deletes an edge stack from the Portainer server.
//
// Parameters:
//   - id: The ID of the edge stack to delete
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeleteEdgeStack(id int) error {
	if err := c.doAPIDelete(fmt.Sprintf("/edge_stacks/%d", id)); err != nil {
		return fmt.Errorf("failed to delete edge stack: %w", err)
	}

	return nil
}

// DeleteTag deletes a tag from the Portainer server.
//
// Parameters:
//   - id: The ID of the tag to delete
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeleteTag(id int) error {
	if err := c.doAPIDelete(fmt.Sprintf("/tags/%d", id)); err != nil {
		return fmt.Errorf("failed to delete tag: %w", err)
	}

	return nil
}

// DeleteTeam deletes a team from the Portainer server.
//
// Parameters:
//   - id: The ID of the team to delete
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeleteTeam(id int) error {
	if err := c.doAPIDelete(fmt.Sprintf("/teams/%d", id)); err != nil {
		return fmt.Errorf("failed to delete team: %w", err)
	}

	return nil
}
