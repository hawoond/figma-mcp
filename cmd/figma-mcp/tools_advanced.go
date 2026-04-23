package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hawoond/figma-mcp/pkg/figma/types"
	"github.com/hawoond/figma-mcp/pkg/figma/util"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) handleGetLocalVariables(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := s.figma.Variables.GetLocalVariables(ctx, fileKey)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get local variables failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetPublishedVariables(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := s.figma.Variables.GetPublishedVariables(ctx, fileKey)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get published variables failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetVariableSummary(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	summary, err := s.editor.GetVariableSummary(ctx, fileKey)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get variable summary failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(summary, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleExportDesignTokens(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	format := "css"
	if f, ok := req.GetArguments()["format"]; ok {
		if fStr, ok := f.(string); ok {
			format = fStr
		}
	}

	modeFilter := ""
	if m, ok := req.GetArguments()["mode"]; ok {
		if mStr, ok := m.(string); ok {
			modeFilter = mStr
		}
	}

	resp, err := s.figma.Variables.GetLocalVariables(ctx, fileKey)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get variables failed: %v", err)), nil
	}

	tokens := util.ExtractDesignTokens(resp)

	switch format {
	case "css":
		css := util.TokensToCSSVariables(tokens, modeFilter)
		return mcp.NewToolResultText(css), nil
	case "json":
		jsonTokens := util.TokensToJSON(tokens)
		data, _ := json.MarshalIndent(jsonTokens, "", "  ")
		return mcp.NewToolResultText(string(data)), nil
	default:
		data, _ := json.MarshalIndent(tokens, "", "  ")
		return mcp.NewToolResultText(string(data)), nil
	}
}

func (s *Server) handleCreateVariable(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	collectionID, err := req.RequireString("collection_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	name, err := req.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	resolvedType, err := req.RequireString("resolved_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	postReq := &types.PostVariablesRequest{
		Variables: []types.VariableChange{
			{
				Action:               "CREATE",
				Name:                 name,
				VariableCollectionID: collectionID,
				ResolvedType:         resolvedType,
			},
		},
	}

	resp, err := s.figma.Variables.PostVariables(ctx, fileKey, postReq)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("create variable failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleCreateVariableCollection(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	name, err := req.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	postReq := &types.PostVariablesRequest{
		VariableCollections: []types.VariableCollectionChange{
			{
				Action: "CREATE",
				Name:   name,
			},
		},
	}

	resp, err := s.figma.Variables.PostVariables(ctx, fileKey, postReq)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("create variable collection failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetWebhooks(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	teamID, err := req.RequireString("team_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := s.figma.Webhooks.GetTeamWebhooks(ctx, teamID)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get webhooks failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleCreateWebhook(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	teamID, err := req.RequireString("team_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	eventType, err := req.RequireString("event_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	endpoint, err := req.RequireString("endpoint")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	passcode, err := req.RequireString("passcode")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	webhookReq := &types.WebhookRequest{
		TeamID:    teamID,
		EventType: eventType,
		Endpoint:  endpoint,
		Passcode:  passcode,
	}

	if desc, ok := req.GetArguments()["description"]; ok {
		if d, ok := desc.(string); ok {
			webhookReq.Description = d
		}
	}

	resp, err := s.figma.Webhooks.CreateWebhook(ctx, webhookReq)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("create webhook failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleDeleteWebhook(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	webhookID, err := req.RequireString("webhook_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if err := s.figma.Webhooks.DeleteWebhook(ctx, webhookID); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("delete webhook failed: %v", err)), nil
	}

	return mcp.NewToolResultText("Webhook deleted successfully"), nil
}

func (s *Server) handleGetDevResources(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	var nodeIDs []string
	if ids, ok := req.GetArguments()["node_ids"]; ok {
		if idsStr, ok := ids.(string); ok && idsStr != "" {
			nodeIDs = splitAndTrim(idsStr)
		}
	}

	resp, err := s.figma.DevResources.GetDevResources(ctx, fileKey, nodeIDs)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get dev resources failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleCreateDevResource(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	nodeID, err := req.RequireString("node_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	name, err := req.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	resourceURL, err := req.RequireString("url")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	createReq := &types.CreateDevResourcesRequest{
		DevResources: []types.CreateDevResource{
			{
				Name:    name,
				URL:     resourceURL,
				FileKey: fileKey,
				NodeID:  nodeID,
			},
		},
	}

	resp, err := s.figma.DevResources.CreateDevResources(ctx, createReq)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("create dev resource failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleDeleteDevResource(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	devResourceID, err := req.RequireString("dev_resource_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if err := s.figma.DevResources.DeleteDevResource(ctx, fileKey, devResourceID); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("delete dev resource failed: %v", err)), nil
	}

	return mcp.NewToolResultText("Dev resource deleted successfully"), nil
}

func splitAndTrim(s string) []string {
	parts := make([]string, 0)
	for _, p := range splitString(s, ",") {
		trimmed := trimSpace(p)
		if trimmed != "" {
			parts = append(parts, trimmed)
		}
	}
	return parts
}

func splitString(s, sep string) []string {
	var result []string
	start := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == sep {
			result = append(result, s[start:i])
			start = i + 1
		}
	}
	result = append(result, s[start:])
	return result
}

func trimSpace(s string) string {
	start := 0
	end := len(s)
	for start < end && (s[start] == ' ' || s[start] == '\t') {
		start++
	}
	for end > start && (s[end-1] == ' ' || s[end-1] == '\t') {
		end--
	}
	return s[start:end]
}
