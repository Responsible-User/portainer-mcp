package models

import "encoding/json"

// Policy represents a fleetwide policy in Portainer.
type Policy struct {
	ID                int             `json:"Id"`
	Name              string          `json:"Name"`
	Type              string          `json:"Type"`
	EnvironmentType   string          `json:"EnvironmentType"`
	EnvironmentGroups []int           `json:"EnvironmentGroups"`
	Data              json.RawMessage `json:"Data,omitempty"`
	CreatedAt         string          `json:"CreatedAt"`
	UpdatedAt         string          `json:"UpdatedAt"`
}

// PolicyCreateRequest represents the request body for creating a policy.
type PolicyCreateRequest struct {
	Name              string          `json:"name"`
	Type              string          `json:"type"`
	EnvironmentType   string          `json:"environmentType"`
	EnvironmentGroups []int           `json:"environmentGroups,omitempty"`
	Data              json.RawMessage `json:"data,omitempty"`
}

// PolicyUpdateRequest represents the request body for updating a policy.
type PolicyUpdateRequest struct {
	Name              string          `json:"name,omitempty"`
	Type              string          `json:"type,omitempty"`
	EnvironmentType   string          `json:"environmentType,omitempty"`
	EnvironmentGroups []int           `json:"environmentGroups,omitempty"`
	Data              json.RawMessage `json:"data,omitempty"`
}
