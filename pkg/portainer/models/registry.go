package models

// Registry represents a container registry in Portainer.
type Registry struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Type           int    `json:"type"`
	URL            string `json:"url"`
	Authentication bool   `json:"authentication"`
	Username       string `json:"username,omitempty"`
}

// RegistryCreateRequest represents the request body for creating a registry.
type RegistryCreateRequest struct {
	Name           string `json:"name"`
	Type           int    `json:"type"`
	URL            string `json:"url"`
	Authentication bool   `json:"authentication"`
	Username       string `json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
}
