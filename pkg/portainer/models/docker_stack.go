package models

// DockerStack represents a Docker standalone stack in Portainer.
type DockerStack struct {
	ID            int               `json:"id"`
	Name          string            `json:"name"`
	Type          int               `json:"type"`
	Status        int               `json:"status"`
	EndpointID    int               `json:"endpoint_id"`
	EntryPoint    string            `json:"entry_point"`
	Env           []StackEnvVar     `json:"env,omitempty"`
	CreatedBy     string            `json:"created_by"`
	CreationDate  int64             `json:"creation_date"`
	UpdateDate    int64             `json:"update_date,omitempty"`
	UpdatedBy     string            `json:"updated_by,omitempty"`
	IsComposeFormat bool            `json:"is_compose_format"`
}

// StackEnvVar represents an environment variable in a stack.
type StackEnvVar struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// DockerStackCreateRequest represents the request body for creating a Docker standalone stack.
type DockerStackCreateRequest struct {
	Name             string        `json:"name"`
	StackFileContent string        `json:"stackFileContent"`
	Env              []StackEnvVar `json:"env,omitempty"`
}

// DockerStackUpdateRequest represents the request body for updating a Docker standalone stack.
type DockerStackUpdateRequest struct {
	StackFileContent string        `json:"stackFileContent"`
	Env              []StackEnvVar `json:"env,omitempty"`
	Prune            bool          `json:"prune"`
	PullImage        bool          `json:"pullImage"`
}
