package mcp

import (
	"context"
	"encoding/json"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/portainer/portainer-mcp/pkg/toolgen"
)

func (s *PortainerMCPServer) AddSettingsFeatures() {
	s.addToolIfExists(ToolGetSettings, s.HandleGetSettings())

	if !s.readOnly {
		s.addToolIfExists(ToolUpdateSettings, s.HandleUpdateSettings())
	}
}

func (s *PortainerMCPServer) HandleGetSettings() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		settings, err := s.cli.GetSettings()
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get settings", err), nil
		}

		data, err := json.Marshal(settings)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal settings", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

func (s *PortainerMCPServer) HandleUpdateSettings() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		settingsJSON, err := parser.GetString("settingsJSON", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid settingsJSON parameter", err), nil
		}

		err = s.cli.UpdateSettings(settingsJSON)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to update settings", err), nil
		}

		return mcp.NewToolResultText("Settings updated successfully"), nil
	}
}
