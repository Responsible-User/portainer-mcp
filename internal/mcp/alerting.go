package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/portainer/portainer-mcp/pkg/toolgen"
)

// AddAlertingFeatures registers all observability alerting related tools.
func (s *PortainerMCPServer) AddAlertingFeatures() {
	s.addToolIfExists(ToolListAlerts, s.HandleListAlerts())
	s.addToolIfExists(ToolListAlertRules, s.HandleListAlertRules())
	s.addToolIfExists(ToolGetAlertRule, s.HandleGetAlertRule())
	s.addToolIfExists(ToolGetAlertingSettings, s.HandleGetAlertingSettings())

	if !s.readOnly {
		s.addToolIfExists(ToolUpdateAlertRule, s.HandleUpdateAlertRule())
		s.addToolIfExists(ToolDeleteAlertRule, s.HandleDeleteAlertRule())
		s.addToolIfExists(ToolCreateAlertSilence, s.HandleCreateAlertSilence())
		s.addToolIfExists(ToolDeleteAlertSilence, s.HandleDeleteAlertSilence())
	}
}

// HandleListAlerts returns a handler that lists active or silenced alerts.
func (s *PortainerMCPServer) HandleListAlerts() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		status, err := parser.GetString("status", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid status parameter", err), nil
		}

		alerts, err := s.cli.GetAlerts(status)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get alerts", err), nil
		}

		return mcp.NewToolResultText(string(alerts)), nil
	}
}

// HandleListAlertRules returns a handler that lists all alert rules.
func (s *PortainerMCPServer) HandleListAlertRules() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		rules, err := s.cli.GetAlertRules()
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get alert rules", err), nil
		}

		data, err := json.Marshal(rules)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal alert rules", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

// HandleGetAlertRule returns a handler that retrieves a specific alert rule.
func (s *PortainerMCPServer) HandleGetAlertRule() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		rule, err := s.cli.GetAlertRule(id)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get alert rule", err), nil
		}

		data, err := json.Marshal(rule)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal alert rule", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

// HandleUpdateAlertRule returns a handler that updates an existing alert rule.
func (s *PortainerMCPServer) HandleUpdateAlertRule() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		ruleJSON, err := parser.GetString("ruleJSON", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid ruleJSON parameter", err), nil
		}

		err = s.cli.UpdateAlertRule(id, ruleJSON)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to update alert rule", err), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Alert rule %d updated successfully", id)), nil
	}
}

// HandleDeleteAlertRule returns a handler that deletes an alert rule.
func (s *PortainerMCPServer) HandleDeleteAlertRule() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		err = s.cli.DeleteAlertRule(id)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to delete alert rule", err), nil
		}

		return mcp.NewToolResultText("Alert rule deleted successfully"), nil
	}
}

// HandleGetAlertingSettings returns a handler that retrieves alerting settings.
func (s *PortainerMCPServer) HandleGetAlertingSettings() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		settings, err := s.cli.GetAlertingSettings()
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get alerting settings", err), nil
		}

		data, err := json.Marshal(settings)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal alerting settings", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

// HandleCreateAlertSilence returns a handler that creates a new alert silence.
func (s *PortainerMCPServer) HandleCreateAlertSilence() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		silenceJSON, err := parser.GetString("silenceJSON", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid silenceJSON parameter", err), nil
		}

		alertManagerURL, err := parser.GetString("alertManagerURL", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid alertManagerURL parameter", err), nil
		}

		err = s.cli.CreateAlertSilence(silenceJSON, alertManagerURL)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to create alert silence", err), nil
		}

		return mcp.NewToolResultText("Alert silence created successfully"), nil
	}
}

// HandleDeleteAlertSilence returns a handler that deletes an alert silence.
func (s *PortainerMCPServer) HandleDeleteAlertSilence() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetString("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		err = s.cli.DeleteAlertSilence(id)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to delete alert silence", err), nil
		}

		return mcp.NewToolResultText("Alert silence deleted successfully"), nil
	}
}
