package client

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/portainer/portainer-mcp/pkg/portainer/models"
)

// ListCustomResourceDefinitions retrieves all Custom Resource Definitions for a Kubernetes environment.
//
// Parameters:
//   - environmentID: The ID of the Kubernetes environment
//
// Returns:
//   - A slice of CustomResourceDefinition objects
//   - An error if the operation fails
func (c *PortainerClient) ListCustomResourceDefinitions(environmentID int) ([]models.CustomResourceDefinition, error) {
	var crds []models.CustomResourceDefinition
	if err := c.doJSONAPIRequest(http.MethodGet, fmt.Sprintf("/kubernetes/%d/customresourcedefinitions", environmentID), nil, &crds); err != nil {
		return nil, fmt.Errorf("failed to list custom resource definitions: %w", err)
	}

	return crds, nil
}

// GetCustomResourceDefinition retrieves a specific Custom Resource Definition by name.
//
// Parameters:
//   - environmentID: The ID of the Kubernetes environment
//   - name: The name of the CRD (e.g., "certificates.cert-manager.io")
//
// Returns:
//   - The CustomResourceDefinition object
//   - An error if the operation fails
func (c *PortainerClient) GetCustomResourceDefinition(environmentID int, name string) (models.CustomResourceDefinition, error) {
	var crd models.CustomResourceDefinition
	if err := c.doJSONAPIRequest(http.MethodGet, fmt.Sprintf("/kubernetes/%d/customresourcedefinitions/%s", environmentID, name), nil, &crd); err != nil {
		return models.CustomResourceDefinition{}, fmt.Errorf("failed to get custom resource definition: %w", err)
	}

	return crd, nil
}

// DeleteCustomResourceDefinition deletes a Custom Resource Definition.
//
// Parameters:
//   - environmentID: The ID of the Kubernetes environment
//   - name: The name of the CRD to delete
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeleteCustomResourceDefinition(environmentID int, name string) error {
	if err := c.doAPIDelete(fmt.Sprintf("/kubernetes/%d/customresourcedefinitions/%s", environmentID, name)); err != nil {
		return fmt.Errorf("failed to delete custom resource definition: %w", err)
	}

	return nil
}

// ListCustomResources retrieves all custom resources for a given CRD definition.
//
// Parameters:
//   - environmentID: The ID of the Kubernetes environment
//   - definition: The CRD definition name to list resources for (e.g., "certificates.cert-manager.io")
//
// Returns:
//   - A slice of CustomResource objects
//   - An error if the operation fails
func (c *PortainerClient) ListCustomResources(environmentID int, definition string) ([]models.CustomResource, error) {
	var resources []models.CustomResource
	path := fmt.Sprintf("/kubernetes/%d/customresources?definition=%s", environmentID, definition)
	if err := c.doJSONAPIRequest(http.MethodGet, path, nil, &resources); err != nil {
		return nil, fmt.Errorf("failed to list custom resources: %w", err)
	}

	return resources, nil
}

// GetCustomResource retrieves a specific custom resource.
// If namespace is empty, the cluster-scoped endpoint is used.
// The format parameter controls the response format: empty for JSON summary, "yaml" for full YAML.
//
// Parameters:
//   - environmentID: The ID of the Kubernetes environment
//   - namespace: The namespace of the resource (empty for cluster-scoped resources)
//   - name: The name of the custom resource
//   - definition: The CRD definition name
//   - format: Response format ("yaml" for YAML, empty for JSON summary)
//
// Returns:
//   - The raw JSON response as a byte slice
//   - An error if the operation fails
func (c *PortainerClient) GetCustomResource(environmentID int, namespace, name, definition, format string) (json.RawMessage, error) {
	var path string
	if namespace != "" {
		path = fmt.Sprintf("/kubernetes/%d/customresources/%s/%s?definition=%s", environmentID, namespace, name, definition)
	} else {
		path = fmt.Sprintf("/kubernetes/%d/customresources/%s?definition=%s", environmentID, name, definition)
	}

	if format != "" {
		path += fmt.Sprintf("&format=%s", format)
	}

	var result json.RawMessage
	if err := c.doJSONAPIRequest(http.MethodGet, path, nil, &result); err != nil {
		return nil, fmt.Errorf("failed to get custom resource: %w", err)
	}

	return result, nil
}

// DeleteCustomResource deletes a specific custom resource.
// If namespace is empty, the cluster-scoped endpoint is used.
//
// Parameters:
//   - environmentID: The ID of the Kubernetes environment
//   - namespace: The namespace of the resource (empty for cluster-scoped resources)
//   - name: The name of the custom resource
//   - definition: The CRD definition name
//
// Returns:
//   - An error if the operation fails
func (c *PortainerClient) DeleteCustomResource(environmentID int, namespace, name, definition string) error {
	var path string
	if namespace != "" {
		path = fmt.Sprintf("/kubernetes/%d/customresources/%s/%s?definition=%s", environmentID, namespace, name, definition)
	} else {
		path = fmt.Sprintf("/kubernetes/%d/customresources/%s?definition=%s", environmentID, name, definition)
	}

	if err := c.doAPIDelete(path); err != nil {
		return fmt.Errorf("failed to delete custom resource: %w", err)
	}

	return nil
}
