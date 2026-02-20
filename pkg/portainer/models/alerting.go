package models

// AlertingRule represents an alerting rule in Portainer.
type AlertingRule struct {
	ID                        int               `json:"id"`
	Name                      string            `json:"name"`
	Description               string            `json:"description,omitempty"`
	Severity                  string            `json:"severity"`
	ConditionOperator         string            `json:"conditionOperator"`
	Threshold                 float64           `json:"threshold"`
	Duration                  int               `json:"duration"`
	Enabled                   bool              `json:"enabled"`
	IsEditable                bool              `json:"isEditable"`
	IsInternal                bool              `json:"isInternal"`
	MetricType                string            `json:"metricType"`
	AlertManagerID            int               `json:"alertManagerID"`
	CreatedAt                 string            `json:"createdAt,omitempty"`
	UpdatedAt                 string            `json:"updatedAt,omitempty"`
	Labels                    map[string]string `json:"labels,omitempty"`
	Summary                   string            `json:"summary,omitempty"`
	CreatedBy                 string            `json:"createdBy,omitempty"`
	SupportedAgentVersion     string            `json:"supportedAgentVersion,omitempty"`
	SupportedEnvironmentTypes string            `json:"supportedEnvironmentTypes,omitempty"`
}

// AlertingSettings represents alerting configuration (alert manager instance).
type AlertingSettings struct {
	ID                   int                           `json:"id"`
	Name                 string                        `json:"name"`
	Enabled              bool                          `json:"enabled"`
	IsInternal           bool                          `json:"isInternal"`
	Status               string                        `json:"status"`
	URL                  string                        `json:"url,omitempty"`
	PortainerURL         string                        `json:"portainerURL,omitempty"`
	NotificationChannels []AlertingNotificationChannel `json:"notificationChannels,omitempty"`
	CreatedAt            string                        `json:"createdAt,omitempty"`
	CreatedBy            string                        `json:"createdBy,omitempty"`
	Uptime               string                        `json:"uptime,omitempty"`
}

// AlertingNotificationChannel represents a notification channel in alerting settings.
type AlertingNotificationChannel struct {
	ID      int                    `json:"id"`
	Name    string                 `json:"name"`
	Type    string                 `json:"type"`
	Config  map[string]interface{} `json:"config,omitempty"`
	Enabled bool                   `json:"enabled"`
}
