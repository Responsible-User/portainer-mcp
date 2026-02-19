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

func (s *PortainerMCPServer) AddEdgeJobFeatures() {
	s.addToolIfExists(ToolListEdgeJobs, s.HandleListEdgeJobs())
	s.addToolIfExists(ToolGetEdgeJob, s.HandleGetEdgeJob())

	if !s.readOnly {
		s.addToolIfExists(ToolCreateEdgeJob, s.HandleCreateEdgeJob())
		s.addToolIfExists(ToolDeleteEdgeJob, s.HandleDeleteEdgeJob())
	}
}

func (s *PortainerMCPServer) HandleListEdgeJobs() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		jobs, err := s.cli.GetEdgeJobs()
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get edge jobs", err), nil
		}

		data, err := json.Marshal(jobs)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal edge jobs", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

func (s *PortainerMCPServer) HandleGetEdgeJob() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		job, err := s.cli.GetEdgeJob(id)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to get edge job", err), nil
		}

		data, err := json.Marshal(job)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal edge job", err), nil
		}

		return mcp.NewToolResultText(string(data)), nil
	}
}

func (s *PortainerMCPServer) HandleCreateEdgeJob() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		name, err := parser.GetString("name", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid name parameter", err), nil
		}

		cronExpression, err := parser.GetString("cronExpression", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid cronExpression parameter", err), nil
		}

		recurring, err := parser.GetBoolean("recurring", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid recurring parameter", err), nil
		}

		scriptContent, err := parser.GetString("scriptContent", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid scriptContent parameter", err), nil
		}

		edgeGroupIds, err := parser.GetArrayOfIntegers("edgeGroupIds", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid edgeGroupIds parameter", err), nil
		}

		req := models.EdgeJobCreateRequest{
			Name:           name,
			CronExpression: cronExpression,
			Recurring:      recurring,
			ScriptContent:  scriptContent,
			EdgeGroups:     edgeGroupIds,
		}

		id, err := s.cli.CreateEdgeJob(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to create edge job", err), nil
		}

		return mcp.NewToolResultText(fmt.Sprintf("Edge job created successfully with ID %d", id)), nil
	}
}

func (s *PortainerMCPServer) HandleDeleteEdgeJob() server.ToolHandlerFunc {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		parser := toolgen.NewParameterParser(request)

		id, err := parser.GetInt("id", true)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("invalid id parameter", err), nil
		}

		err = s.cli.DeleteEdgeJob(id)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to delete edge job", err), nil
		}

		return mcp.NewToolResultText("Edge job deleted successfully"), nil
	}
}
