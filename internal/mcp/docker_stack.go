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

func (s *PortainerMCPServer) AddDockerStackFeatures() {
	s.addToolIfExists(ToolListDockerStacks, s.HandleListDockerStacks())
	s.addToolIfExists(ToolGetDockerStackFile, s.HandleGetDockerStackFile())

	if !s.readOnly {
		s.addToolIfExists(ToolCreateDockerStack, s.HandleCreateDockerStack())
		s.addToolIfExists(ToolUpdateDockerStack, s.HandleUpdateDockerStack())
		s.addToolIfExists(ToolDeleteDockerStack, s.HandleDeleteDockerStack())
		s.addToolIfExists(ToolStartDockerStack, s.HandleStartDockerStack())
		s.addToolIfExists(ToolStopDockerStack, s.HandleStopDockerStack())
	}
}

func (s *PortainerMCPServer) HandleListDockerStacks() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		stacks, err := s.cli.GetDockerStacks()
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get docker stacks", err), nil
		}

		data, err := json.Marshal(stacks)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal docker stacks", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

func (s *PortainerMCPServer) HandleGetDockerStackFile() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		file, err := s.cli.GetDockerStackFile(id)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get docker stack file", err), nil
		}

		return mcp.NewToolResultText(file), nil
	}
}

func (s *PortainerMCPServer) HandleCreateDockerStack() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		environmentId, err := parser.GetInt("environmentId", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid environmentId parameter", err), nil
		}

		name, err := parser.GetString("name", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid name parameter", err), nil
		}

		file, err := parser.GetString("file", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid file parameter", err), nil
		}

		id, err := s.cli.CreateDockerStack(environmentId, name, file, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to create docker stack", err), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Docker stack created successfully with ID %d", id)), nil
	}
}

func (s *PortainerMCPServer) HandleUpdateDockerStack() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		environmentId, err := parser.GetInt("environmentId", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid environmentId parameter", err), nil
		}

		file, err := parser.GetString("file", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid file parameter", err), nil
		}

		prune, err := parser.GetBoolean("prune", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid prune parameter", err), nil
		}

		pullImage, err := parser.GetBoolean("pullImage", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid pullImage parameter", err), nil
		}

		err = s.cli.UpdateDockerStack(id, environmentId, file, nil, prune, pullImage)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to update docker stack", err), nil
		}

		return mcp.NewToolResultText("Docker stack updated successfully"), nil
	}
}

func (s *PortainerMCPServer) HandleDeleteDockerStack() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		environmentId, err := parser.GetInt("environmentId", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid environmentId parameter", err), nil
		}

		err = s.cli.DeleteDockerStack(id, environmentId)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to delete docker stack", err), nil
		}

		return mcp.NewToolResultText("Docker stack deleted successfully"), nil
	}
}

func (s *PortainerMCPServer) HandleStartDockerStack() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		environmentId, err := parser.GetInt("environmentId", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid environmentId parameter", err), nil
		}

		err = s.cli.StartDockerStack(id, environmentId)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to start docker stack", err), nil
		}

		return mcp.NewToolResultText("Docker stack started successfully"), nil
	}
}

func (s *PortainerMCPServer) HandleStopDockerStack() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		environmentId, err := parser.GetInt("environmentId", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid environmentId parameter", err), nil
		}

		err = s.cli.StopDockerStack(id, environmentId)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to stop docker stack", err), nil
		}

		return mcp.NewToolResultText("Docker stack stopped successfully"), nil
	}
}

// parseEnvVars parses environment variables from the request parameters.
func parseEnvVars(envMaps []map[string]any) []models.StackEnvVar {
	envVars := make([]models.StackEnvVar, 0, len(envMaps))
	for _, env := range envMaps {
		name, _ := env["name"].(string)
		value, _ := env["value"].(string)
		if name != "" {
			envVars = append(envVars, models.StackEnvVar{Name: name, Value: value})
		}
	}
	return envVars
}
