package models

// CustomTemplate represents a custom template in Portainer.
type CustomTemplate struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Type        int    `json:"type"`
	Platform    int    `json:"platform"`
	CreatedBy   string `json:"created_by"`
}

// CustomTemplateCreateRequest represents the request body for creating a custom template.
type CustomTemplateCreateRequest struct {
	Title            string `json:"title"`
	Description      string `json:"description"`
	FileContent      string `json:"fileContent"`
	Type             int    `json:"type"`
	Platform         int    `json:"platform"`
}
