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

// AddPolicyFeatures registers all fleetwide policy related tools.
func (s *PortainerMCPServer) AddPolicyFeatures() {
	s.addToolIfExists(ToolListPolicies, s.HandleListPolicies())
	s.addToolIfExists(ToolGetPolicy, s.HandleGetPolicy())

	if !s.readOnly {
		s.addToolIfExists(ToolCreatePolicy, s.HandleCreatePolicy())
		s.addToolIfExists(ToolUpdatePolicy, s.HandleUpdatePolicy())
		s.addToolIfExists(ToolDeletePolicy, s.HandleDeletePolicy())
	}
}

// HandleListPolicies returns a handler that lists all fleetwide policies.
func (s *PortainerMCPServer) HandleListPolicies() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		policies, err := s.cli.GetPolicies()
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get policies", err), nil
		}

		data, err := json.Marshal(policies)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal policies", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

// HandleGetPolicy returns a handler that retrieves a specific policy.
func (s *PortainerMCPServer) HandleGetPolicy() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		policy, err := s.cli.GetPolicy(id)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get policy", err), nil
		}

		data, err := json.Marshal(policy)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal policy", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

// HandleCreatePolicy returns a handler that creates a new fleetwide policy.
func (s *PortainerMCPServer) HandleCreatePolicy() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		name, err := parser.GetString("name", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid name parameter", err), nil
		}

		policyType, err := parser.GetString("type", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid type parameter", err), nil
		}

		environmentType, err := parser.GetString("environmentType", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid environmentType parameter", err), nil
		}

		environmentGroups, err := parser.GetArrayOfIntegers("environmentGroups", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid environmentGroups parameter", err), nil
		}

		dataJSON, err := parser.GetString("dataJSON", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid dataJSON parameter", err), nil
		}

		req := models.PolicyCreateRequest{
			Name:              name,
			Type:              policyType,
			EnvironmentType:   environmentType,
			EnvironmentGroups: environmentGroups,
		}

		if dataJSON != "" {
			req.Data = json.RawMessage(dataJSON)
		}

		id, err := s.cli.CreatePolicy(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to create policy", err), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Policy created successfully with ID %d", id)), nil
	}
}

// HandleUpdatePolicy returns a handler that updates an existing fleetwide policy.
func (s *PortainerMCPServer) HandleUpdatePolicy() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		name, err := parser.GetString("name", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid name parameter", err), nil
		}

		policyType, err := parser.GetString("type", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid type parameter", err), nil
		}

		environmentType, err := parser.GetString("environmentType", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid environmentType parameter", err), nil
		}

		environmentGroups, err := parser.GetArrayOfIntegers("environmentGroups", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid environmentGroups parameter", err), nil
		}

		dataJSON, err := parser.GetString("dataJSON", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid dataJSON parameter", err), nil
		}

		req := models.PolicyUpdateRequest{
			Name:              name,
			Type:              policyType,
			EnvironmentType:   environmentType,
			EnvironmentGroups: environmentGroups,
		}

		if dataJSON != "" {
			req.Data = json.RawMessage(dataJSON)
		}

		err = s.cli.UpdatePolicy(id, req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to update policy", err), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Policy %d updated successfully", id)), nil
	}
}

// HandleDeletePolicy returns a handler that deletes a fleetwide policy.
func (s *PortainerMCPServer) HandleDeletePolicy() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		err = s.cli.DeletePolicy(id)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to delete policy", err), nil
		}

		return mcp.NewToolResultText("Policy deleted successfully"), nil
	}
}
