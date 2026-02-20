package models

// GitCredential represents a shared git credential in Portainer.
type GitCredential struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	Username          string `json:"username"`
	AuthorizationType int    `json:"authorizationType"`
	UserID            int    `json:"userId"`
	CreationDate      int64  `json:"creationDate"`
}

// GitCredentialCreateRequest represents the request body for creating a shared git credential.
type GitCredentialCreateRequest struct {
	Name              string `json:"name"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	AuthorizationType int    `json:"authorizationType"`
}

// GitCredentialUpdateRequest represents the request body for updating a shared git credential.
type GitCredentialUpdateRequest struct {
	Name              string `json:"name"`
	Username          string `json:"username"`
	Password          string `json:"password,omitempty"`
	AuthorizationType int    `json:"authorizationType"`
}
