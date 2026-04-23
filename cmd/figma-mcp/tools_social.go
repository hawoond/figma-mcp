package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hawoond/figma-mcp/pkg/figma/types"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) handleGetComments(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	asMarkdown := false
	if md, ok := req.GetArguments()["as_markdown"]; ok {
		if b, ok := md.(bool); ok {
			asMarkdown = b
		}
	}

	resp, err := s.figma.Comments.GetComments(ctx, fileKey, asMarkdown)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get comments failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handlePostComment(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	message, err := req.RequireString("message")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	postReq := &types.PostCommentRequest{
		Message: message,
	}

	if commentID, ok := req.GetArguments()["comment_id"]; ok {
		if id, ok := commentID.(string); ok {
			postReq.CommentID = id
		}
	}

	comment, err := s.figma.Comments.PostComment(ctx, fileKey, postReq)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("post comment failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(comment, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleDeleteComment(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	commentID, err := req.RequireString("comment_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	if err := s.figma.Comments.DeleteComment(ctx, fileKey, commentID); err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("delete comment failed: %v", err)), nil
	}

	return mcp.NewToolResultText("Comment deleted successfully"), nil
}

func (s *Server) handleGetMe(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	user, err := s.figma.Users.GetMe(ctx)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get me failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(user, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetTeamProjects(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	teamID, err := req.RequireString("team_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := s.figma.Projects.GetTeamProjects(ctx, teamID)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get team projects failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetProjectFiles(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	projectID, err := req.RequireString("project_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := s.figma.Projects.GetProjectFiles(ctx, projectID, nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get project files failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetTeamComponents(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	teamID, err := req.RequireString("team_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := s.figma.Projects.GetTeamComponents(ctx, teamID, nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get team components failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetFileComponents(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := s.figma.Projects.GetFileComponents(ctx, fileKey)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get file components failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetComponent(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	key, err := req.RequireString("key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := s.figma.Projects.GetComponent(ctx, key)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get component failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetFileComponentSets(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := s.figma.Projects.GetFileComponentSets(ctx, fileKey)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get file component sets failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetTeamStyles(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	teamID, err := req.RequireString("team_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := s.figma.Projects.GetTeamStyles(ctx, teamID, nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get team styles failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetFileStyles(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := s.figma.Projects.GetFileStyles(ctx, fileKey)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get file styles failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}
