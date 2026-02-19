package models

// Webhook represents a webhook in Portainer.
type Webhook struct {
	ID         int    `json:"id"`
	Token      string `json:"token"`
	ResourceID string `json:"resource_id"`
	EndpointID int    `json:"endpoint_id"`
	Type       int    `json:"type"`
}

// WebhookCreateRequest represents the request body for creating a webhook.
type WebhookCreateRequest struct {
	ResourceID string `json:"resourceID"`
	EndpointID int    `json:"endpointID"`
	Type       int    `json:"webhookType"`
}
