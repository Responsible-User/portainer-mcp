package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/portainer/portainer-mcp/pkg/portainer/models"
	"github.com/portainer/portainer-mcp/pkg/toolgen"
)

func (s *PortainerMCPServer) AddCustomTemplateFeatures() {
	s.addToolIfExists(ToolListCustomTemplates, s.HandleListCustomTemplates())

	if !s.readOnly {
		s.addToolIfExists(ToolCreateCustomTemplate, s.HandleCreateCustomTemplate())
		s.addToolIfExists(ToolDeleteCustomTemplate, s.HandleDeleteCustomTemplate())
	}
}

func (s *PortainerMCPServer) HandleListCustomTemplates() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		templates, err := s.cli.GetCustomTemplates()
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get custom templates", err), nil
		}

		data, err := json.Marshal(templates)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal custom templates", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

func (s *PortainerMCPServer) HandleCreateCustomTemplate() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		title, err := parser.GetString("title", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid title parameter", err), nil
		}

		description, err := parser.GetString("description", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid description parameter", err), nil
		}

		fileContent, err := parser.GetString("fileContent", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid fileContent parameter", err), nil
		}

		templateType, err := parser.GetInt("type", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid type parameter", err), nil
		}

		platform, err := parser.GetInt("platform", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid platform parameter", err), nil
		}

		req := models.CustomTemplateCreateRequest{
			Title:       title,
			Description: description,
			FileContent: fileContent,
			Type:        templateType,
			Platform:    platform,
		}

		id, err := s.cli.CreateCustomTemplate(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to create custom template", err), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Custom template created successfully with ID %d", id)), nil
	}
}

func (s *PortainerMCPServer) HandleDeleteCustomTemplate() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		err = s.cli.DeleteCustomTemplate(id)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to delete custom template", err), nil
		}

		return mcp.NewToolResultText("Custom template deleted successfully"), nil
	}
}
