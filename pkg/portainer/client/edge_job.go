package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/portainer/portainer-mcp/pkg/portainer/models"
)

// GetEdgeJobs retrieves all edge jobs from the Portainer server.
//
// Returns:
//   - A slice of EdgeJob objects
//   - An error if the operation fails
func (c *PortainerClient) GetEdgeJobs() ([]models.EdgeJob, error) {
	var jobs []models.EdgeJob
	if err := c.doJSONAPIRequest(http.MethodGet, "/edge_jobs", nil, &jobs); err != nil {
		return nil, fmt.Errorf("failed to list edge jobs: %w", err)
	}

	return jobs, nil
}

// GetEdgeJob retrieves a specific edge job by ID.
//
// Parameters:
//   - id: The ID of the edge job
//
// Returns:
//   - The EdgeJob object
//   - An error if the operation fails
func (c *PortainerClient) GetEdgeJob(id int) (models.EdgeJob, error) {
	var job models.EdgeJob
	if err := c.doJSONAPIRequest(http.MethodGet, fmt.Sprintf("/edge_jobs/%d", id), nil, &job); err != nil {
		return models.EdgeJob{}, fmt.Errorf("failed to get edge job: %w", err)
	}

	return job, nil
}

// CreateEdgeJob creates a new edge job.
//
// Parameters:
//   - req: The edge job creation request
//
// Returns:
//   - The ID of the created edge job
//   - An error if the operation fails
func (c *PortainerClient) CreateEdgeJob(req models.EdgeJobCreateRequest) (int, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return 0, fmt.Errorf("failed to marshal create request: %w", err)
	}

	var result struct {
		ID int `json:"Id"`
	}
	if err := c.doJSONAPIRequest(http.MethodPost, "/edge_jobs/create/string", bytes.NewReader(body), &result); err != nil {
		return 0, fmt.Errorf("failed to create edge job: %w", err)
	}

	return result.ID, nil
}

// DeleteEdgeJob deletes an edge job.
//
// Parameters:
//   - id: The ID of the edge job to delete
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeleteEdgeJob(id int) error {
	if err := c.doAPIDelete(fmt.Sprintf("/edge_jobs/%d", id)); err != nil {
		return fmt.Errorf("failed to delete edge job: %w", err)
	}

	return nil
}
