package models

// EdgeJob represents an edge job in Portainer.
type EdgeJob struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	CronExpression string `json:"cron_expression"`
	Recurring      bool   `json:"recurring"`
	Created        int64  `json:"created"`
	ScriptPath     string `json:"script_path,omitempty"`
	EdgeGroups     []int  `json:"edge_groups"`
}

// EdgeJobCreateRequest represents the request body for creating an edge job.
type EdgeJobCreateRequest struct {
	Name           string `json:"name"`
	CronExpression string `json:"cronExpression"`
	Recurring      bool   `json:"recurring"`
	ScriptContent  string `json:"fileContent"`
	EdgeGroups     []int  `json:"edgeGroups"`
}
