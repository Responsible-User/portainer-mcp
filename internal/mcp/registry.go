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

func (s *PortainerMCPServer) AddRegistryFeatures() {
	s.addToolIfExists(ToolListRegistries, s.HandleListRegistries())
	s.addToolIfExists(ToolTestRegistryConnection, s.HandleTestRegistryConnection())

	if !s.readOnly {
		s.addToolIfExists(ToolCreateRegistry, s.HandleCreateRegistry())
		s.addToolIfExists(ToolDeleteRegistry, s.HandleDeleteRegistry())
	}
}

func (s *PortainerMCPServer) HandleListRegistries() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		registries, err := s.cli.GetRegistries()
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get registries", err), nil
		}

		data, err := json.Marshal(registries)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal registries", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

func (s *PortainerMCPServer) HandleCreateRegistry() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		name, err := parser.GetString("name", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid name parameter", err), nil
		}

		registryType, err := parser.GetInt("type", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid type parameter", err), nil
		}

		url, err := parser.GetString("url", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid url parameter", err), nil
		}

		authentication, err := parser.GetBoolean("authentication", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid authentication parameter", err), nil
		}

		username, err := parser.GetString("username", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid username parameter", err), nil
		}

		password, err := parser.GetString("password", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid password parameter", err), nil
		}

		req := models.RegistryCreateRequest{
			Name:           name,
			Type:           registryType,
			URL:            url,
			Authentication: authentication,
			Username:       username,
			Password:       password,
		}

		id, err := s.cli.CreateRegistry(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to create registry", err), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Registry created successfully with ID %d", id)), nil
	}
}

func (s *PortainerMCPServer) HandleDeleteRegistry() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		err = s.cli.DeleteRegistry(id)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to delete registry", err), nil
		}

		return mcp.NewToolResultText("Registry deleted successfully"), nil
	}
}

// HandleTestRegistryConnection returns a handler that tests a registry connection.
func (s *PortainerMCPServer) HandleTestRegistryConnection() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		url, err := parser.GetString("url", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid url parameter", err), nil
		}

		registryType, err := parser.GetInt("type", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid type parameter", err), nil
		}

		username, err := parser.GetString("username", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid username parameter", err), nil
		}

		password, err := parser.GetString("password", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid password parameter", err), nil
		}

		req := models.RegistryPingRequest{
			URL:      url,
			Type:     registryType,
			Username: username,
			Password: password,
		}

		result, err := s.cli.PingRegistry(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to test registry connection", err), nil
		}

		data, err := json.Marshal(result)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal ping result", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}
