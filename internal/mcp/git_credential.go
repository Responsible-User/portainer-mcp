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

// AddGitCredentialFeatures registers all shared git credential related tools.
func (s *PortainerMCPServer) AddGitCredentialFeatures() {
	s.addToolIfExists(ToolListGitCredentials, s.HandleListGitCredentials())
	s.addToolIfExists(ToolGetGitCredential, s.HandleGetGitCredential())

	if !s.readOnly {
		s.addToolIfExists(ToolCreateGitCredential, s.HandleCreateGitCredential())
		s.addToolIfExists(ToolUpdateGitCredential, s.HandleUpdateGitCredential())
		s.addToolIfExists(ToolDeleteGitCredential, s.HandleDeleteGitCredential())
	}
}

// HandleListGitCredentials returns a handler that lists all shared git credentials.
func (s *PortainerMCPServer) HandleListGitCredentials() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		creds, err := s.cli.GetGitCredentials()
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get git credentials", err), nil
		}

		data, err := json.Marshal(creds)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal git credentials", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

// HandleGetGitCredential returns a handler that retrieves a specific shared git credential.
func (s *PortainerMCPServer) HandleGetGitCredential() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		cred, err := s.cli.GetGitCredential(id)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get git credential", err), nil
		}

		data, err := json.Marshal(cred)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal git credential", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

// HandleCreateGitCredential returns a handler that creates a new shared git credential.
func (s *PortainerMCPServer) HandleCreateGitCredential() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		name, err := parser.GetString("name", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid name parameter", err), nil
		}

		username, err := parser.GetString("username", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid username parameter", err), nil
		}

		password, err := parser.GetString("password", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid password parameter", err), nil
		}

		authType, err := parser.GetInt("authorizationType", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid authorizationType parameter", err), nil
		}

		req := models.GitCredentialCreateRequest{
			Name:              name,
			Username:          username,
			Password:          password,
			AuthorizationType: authType,
		}

		id, err := s.cli.CreateGitCredential(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to create git credential", err), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Git credential created successfully with ID %d", id)), nil
	}
}

// HandleUpdateGitCredential returns a handler that updates an existing shared git credential.
func (s *PortainerMCPServer) HandleUpdateGitCredential() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		name, err := parser.GetString("name", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid name parameter", err), nil
		}

		username, err := parser.GetString("username", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid username parameter", err), nil
		}

		password, err := parser.GetString("password", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid password parameter", err), nil
		}

		authType, err := parser.GetInt("authorizationType", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid authorizationType parameter", err), nil
		}

		req := models.GitCredentialUpdateRequest{
			Name:              name,
			Username:          username,
			Password:          password,
			AuthorizationType: authType,
		}

		err = s.cli.UpdateGitCredential(id, req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to update git credential", err), nil
		}

		return mcp.NewToolResultText("Git credential updated successfully"), nil
	}
}

// HandleDeleteGitCredential returns a handler that deletes a shared git credential.
func (s *PortainerMCPServer) HandleDeleteGitCredential() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		err = s.cli.DeleteGitCredential(id)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to delete git credential", err), nil
		}

		return mcp.NewToolResultText("Git credential deleted successfully"), nil
	}
}
