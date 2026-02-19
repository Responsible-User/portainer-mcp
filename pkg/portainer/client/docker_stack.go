package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/portainer/portainer-mcp/pkg/portainer/models"
)

// GetDockerStacks retrieves all Docker standalone stacks from the Portainer server.
//
// Returns:
//   - A slice of DockerStack objects
//   - An error if the operation fails
func (c *PortainerClient) GetDockerStacks() ([]models.DockerStack, error) {
	var stacks []models.DockerStack
	if err := c.doJSONAPIRequest(http.MethodGet, "/stacks", nil, &stacks); err != nil {
		return nil, fmt.Errorf("failed to list docker stacks: %w", err)
	}

	return stacks, nil
}

// GetDockerStackFile retrieves the compose file content for a Docker stack.
//
// Parameters:
//   - id: The ID of the stack
//
// Returns:
//   - The compose file content
//   - An error if the operation fails
func (c *PortainerClient) GetDockerStackFile(id int) (string, error) {
	var result struct {
		StackFileContent string `json:"StackFileContent"`
	}
	if err := c.doJSONAPIRequest(http.MethodGet, fmt.Sprintf("/stacks/%d/file", id), nil, &result); err != nil {
		return "", fmt.Errorf("failed to get docker stack file: %w", err)
	}

	return result.StackFileContent, nil
}

// CreateDockerStack creates a new Docker standalone stack.
//
// Parameters:
//   - endpointID: The ID of the environment to deploy the stack to
//   - name: The name of the stack
//   - composeFileContent: The Docker Compose file content
//   - env: Optional environment variables
//
// Returns:
//   - The ID of the created stack
//   - An error if the operation fails
func (c *PortainerClient) CreateDockerStack(endpointID int, name, composeFileContent string, env []models.StackEnvVar) (int, error) {
	reqBody := models.DockerStackCreateRequest{
		Name:             name,
		StackFileContent: composeFileContent,
		Env:              env,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal create request: %w", err)
	}

	var result struct {
		ID int `json:"Id"`
	}
	path := fmt.Sprintf("/stacks/create/standalone/string?endpointId=%d", endpointID)
	if err := c.doJSONAPIRequest(http.MethodPost, path, bytes.NewReader(body), &result); err != nil {
		return 0, fmt.Errorf("failed to create docker stack: %w", err)
	}

	return result.ID, nil
}

// UpdateDockerStack updates an existing Docker standalone stack.
//
// Parameters:
//   - id: The ID of the stack to update
//   - endpointID: The ID of the environment the stack belongs to
//   - composeFileContent: The updated Docker Compose file content
//   - env: Optional environment variables
//   - prune: Whether to prune services that are no longer in the compose file
//   - pullImage: Whether to pull the latest image before deploying
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) UpdateDockerStack(id, endpointID int, composeFileContent string, env []models.StackEnvVar, prune, pullImage bool) error {
	reqBody := models.DockerStackUpdateRequest{
		StackFileContent: composeFileContent,
		Env:              env,
		Prune:            prune,
		PullImage:        pullImage,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("failed to marshal update request: %w", err)
	}

	path := fmt.Sprintf("/stacks/%d?endpointId=%d", id, endpointID)
	if err := c.doJSONAPIRequest(http.MethodPut, path, bytes.NewReader(body), nil); err != nil {
		return fmt.Errorf("failed to update docker stack: %w", err)
	}

	return nil
}

// DeleteDockerStack deletes a Docker standalone stack.
//
// Parameters:
//   - id: The ID of the stack to delete
//   - endpointID: The ID of the environment the stack belongs to
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeleteDockerStack(id, endpointID int) error {
	path := fmt.Sprintf("/stacks/%d?endpointId=%d", id, endpointID)
	if err := c.doAPIDelete(path); err != nil {
		return fmt.Errorf("failed to delete docker stack: %w", err)
	}

	return nil
}

// StartDockerStack starts a stopped Docker standalone stack.
//
// Parameters:
//   - id: The ID of the stack to start
//   - endpointID: The ID of the environment the stack belongs to
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) StartDockerStack(id, endpointID int) error {
	path := fmt.Sprintf("/stacks/%d/start?endpointId=%d", id, endpointID)
	if err := c.doJSONAPIRequest(http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to start docker stack: %w", err)
	}

	return nil
}

// StopDockerStack stops a running Docker standalone stack.
//
// Parameters:
//   - id: The ID of the stack to stop
//   - endpointID: The ID of the environment the stack belongs to
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) StopDockerStack(id, endpointID int) error {
	path := fmt.Sprintf("/stacks/%d/stop?endpointId=%d", id, endpointID)
	if err := c.doJSONAPIRequest(http.MethodPost, path, nil, nil); err != nil {
		return fmt.Errorf("failed to stop docker stack: %w", err)
	}

	return nil
}
