package mcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/portainer/portainer-mcp/pkg/toolgen"
)

// AddCustomResourceFeatures registers all Kubernetes CRD and custom resource related tools.
func (s *PortainerMCPServer) AddCustomResourceFeatures() {
	s.addToolIfExists(ToolListCustomResourceDefinitions, s.HandleListCustomResourceDefinitions())
	s.addToolIfExists(ToolGetCustomResourceDefinition, s.HandleGetCustomResourceDefinition())
	s.addToolIfExists(ToolListCustomResources, s.HandleListCustomResources())
	s.addToolIfExists(ToolGetCustomResource, s.HandleGetCustomResource())

	if !s.readOnly {
		s.addToolIfExists(ToolDeleteCustomResourceDefinition, s.HandleDeleteCustomResourceDefinition())
		s.addToolIfExists(ToolDeleteCustomResource, s.HandleDeleteCustomResource())
	}
}

// HandleListCustomResourceDefinitions returns a handler that lists all CRDs in a Kubernetes environment.
func (s *PortainerMCPServer) HandleListCustomResourceDefinitions() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		environmentId, err := parser.GetInt("environmentId", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid environmentId parameter", err), nil
		}

		crds, err := s.cli.ListCustomResourceDefinitions(environmentId)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to list custom resource definitions", err), nil
		}

		data, err := json.Marshal(crds)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal custom resource definitions", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

// HandleGetCustomResourceDefinition returns a handler that retrieves a specific CRD.
func (s *PortainerMCPServer) HandleGetCustomResourceDefinition() server.ToolHandlerFunc {
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

		crd, err := s.cli.GetCustomResourceDefinition(environmentId, name)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get custom resource definition", err), nil
		}

		data, err := json.Marshal(crd)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal custom resource definition", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

// HandleDeleteCustomResourceDefinition returns a handler that deletes a CRD.
func (s *PortainerMCPServer) HandleDeleteCustomResourceDefinition() server.ToolHandlerFunc {
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

		err = s.cli.DeleteCustomResourceDefinition(environmentId, name)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to delete custom resource definition", err), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Custom resource definition %s deleted successfully", name)), nil
	}
}

// HandleListCustomResources returns a handler that lists custom resources for a given CRD.
func (s *PortainerMCPServer) HandleListCustomResources() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		environmentId, err := parser.GetInt("environmentId", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid environmentId parameter", err), nil
		}

		definition, err := parser.GetString("definition", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid definition parameter", err), nil
		}

		resources, err := s.cli.ListCustomResources(environmentId, definition)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to list custom resources", err), nil
		}

		data, err := json.Marshal(resources)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal custom resources", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

// HandleGetCustomResource returns a handler that retrieves a specific custom resource.
func (s *PortainerMCPServer) HandleGetCustomResource() server.ToolHandlerFunc {
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

		definition, err := parser.GetString("definition", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid definition parameter", err), nil
		}

		namespace, err := parser.GetString("namespace", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid namespace parameter", err), nil
		}

		format, err := parser.GetString("format", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid format parameter", err), nil
		}

		resource, err := s.cli.GetCustomResource(environmentId, namespace, name, definition, format)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get custom resource", err), nil
		}

		return mcp.NewToolResultText(string(resource)), nil
	}
}

// HandleDeleteCustomResource returns a handler that deletes a custom resource.
func (s *PortainerMCPServer) HandleDeleteCustomResource() server.ToolHandlerFunc {
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

		definition, err := parser.GetString("definition", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid definition parameter", err), nil
		}

		namespace, err := parser.GetString("namespace", false)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid namespace parameter", err), nil
		}

		err = s.cli.DeleteCustomResource(environmentId, namespace, name, definition)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to delete custom resource", err), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Custom resource %s deleted successfully", name)), nil
	}
}
