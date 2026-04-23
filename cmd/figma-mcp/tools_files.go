package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/hawoond/figma-mcp/pkg/figma/api"
	"github.com/hawoond/figma-mcp/pkg/figma/util"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) handleGetFile(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	opts := &api.GetFileOptions{}

	if depth, ok := req.GetArguments()["depth"]; ok {
		if d, err := toInt(depth); err == nil {
			opts.Depth = &d
		}
	}
	if ids, ok := req.GetArguments()["ids"]; ok {
		if idsStr, ok := ids.(string); ok && idsStr != "" {
			opts.IDs = strings.Split(idsStr, ",")
		}
	}
	if geometry, ok := req.GetArguments()["geometry"]; ok {
		if g, ok := geometry.(string); ok {
			opts.Geometry = g
		}
	}

	file, err := s.figma.Files.GetFile(ctx, fileKey, opts)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get file failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(file, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetFileNodes(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	nodeIDs, err := req.RequireString("node_ids")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	opts := &api.GetFileNodesOptions{
		IDs: strings.Split(nodeIDs, ","),
	}

	if depth, ok := req.GetArguments()["depth"]; ok {
		if d, err := toInt(depth); err == nil {
			opts.Depth = &d
		}
	}

	nodes, err := s.figma.Files.GetFileNodes(ctx, fileKey, opts)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get file nodes failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(nodes, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetImages(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	nodeIDs, err := req.RequireString("node_ids")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	opts := &api.GetImagesOptions{
		IDs: strings.Split(nodeIDs, ","),
	}

	if format, ok := req.GetArguments()["format"]; ok {
		if f, ok := format.(string); ok {
			opts.Format = f
		}
	}
	if scale, ok := req.GetArguments()["scale"]; ok {
		if s, err := toFloat64(scale); err == nil {
			opts.Scale = &s
		}
	}
	if useAbsBounds, ok := req.GetArguments()["use_absolute_bounds"]; ok {
		if b, ok := useAbsBounds.(bool); ok {
			opts.UseAbsoluteBounds = b
		}
	}

	resp, err := s.figma.Files.GetImages(ctx, fileKey, opts)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get images failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetImageFills(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := s.figma.Files.GetImageFills(ctx, fileKey)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get image fills failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetFileVersions(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := s.figma.Files.GetFileVersions(ctx, fileKey, nil)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get file versions failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetFileMetadata(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	resp, err := s.figma.Files.GetFileMetadata(ctx, fileKey)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get file metadata failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleSearchNodes(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	searchName, err := req.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	nodes, err := s.editor.GetFileWithNodeSearch(ctx, fileKey, searchName)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("search nodes failed: %v", err)), nil
	}

	summaries := make([]util.NodeSummary, 0, len(nodes))
	for _, n := range nodes {
		summaries = append(summaries, util.SummarizeNode(n))
	}

	data, _ := json.MarshalIndent(summaries, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetNodesByType(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	nodeType, err := req.RequireString("node_type")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	nodes, err := s.editor.GetNodesByType(ctx, fileKey, nodeType)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get nodes by type failed: %v", err)), nil
	}

	summaries := make([]util.NodeSummary, 0, len(nodes))
	for _, n := range nodes {
		summaries = append(summaries, util.SummarizeNode(n))
	}

	data, _ := json.MarshalIndent(summaries, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleGetFileStructure(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	maxDepth := 3
	if depth, ok := req.GetArguments()["max_depth"]; ok {
		if d, err := toInt(depth); err == nil {
			maxDepth = d
		}
	}

	summaries, err := s.editor.GetFileStructure(ctx, fileKey, maxDepth)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("get file structure failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(summaries, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleExportFrames(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	format := "png"
	if f, ok := req.GetArguments()["format"]; ok {
		if fStr, ok := f.(string); ok {
			format = fStr
		}
	}

	scale := 1.0
	if sc, ok := req.GetArguments()["scale"]; ok {
		if s, err := toFloat64(sc); err == nil {
			scale = s
		}
	}

	images, err := s.editor.ExportFrames(ctx, fileKey, format, scale)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("export frames failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(images, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleExportNodeAsImage(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	nodeID, err := req.RequireString("node_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	format := "png"
	if f, ok := req.GetArguments()["format"]; ok {
		if fStr, ok := f.(string); ok {
			format = fStr
		}
	}

	scale := 1.0
	if sc, ok := req.GetArguments()["scale"]; ok {
		if s, err := toFloat64(sc); err == nil {
			scale = s
		}
	}

	imgURL, err := s.editor.ExportNodeAsImage(ctx, fileKey, nodeID, format, scale)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("export node as image failed: %v", err)), nil
	}

	return mcp.NewToolResultText(imgURL), nil
}

func (s *Server) handleFetchImageFromURL(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	imageURL, err := req.RequireString("url")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	data, format, err := util.FetchImageFromURL(ctx, imageURL)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("fetch image failed: %v", err)), nil
	}

	result := map[string]interface{}{
		"format":     format,
		"size_bytes": len(data),
		"base64":     util.ImageToBase64(data, format),
	}

	jsonData, _ := json.MarshalIndent(result, "", "  ")
	return mcp.NewToolResultText(string(jsonData)), nil
}

func (s *Server) handleUploadImageFromURL(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	imageURL, err := req.RequireString("image_url")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	result, err := s.editor.UploadImageFromURL(ctx, fileKey, imageURL)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("upload image failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(result, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleUploadMultipleImagesFromURLs(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	imageURLsStr, err := req.RequireString("image_urls")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	imageURLs := strings.Split(imageURLsStr, ",")
	for i, u := range imageURLs {
		imageURLs[i] = strings.TrimSpace(u)
	}

	results, errs := s.editor.UploadMultipleImagesFromURLs(ctx, fileKey, imageURLs)

	type uploadResponse struct {
		Results []*util.UploadImageFromURLResult `json:"results"`
		Errors  []string                         `json:"errors,omitempty"`
	}

	resp := uploadResponse{Results: results}
	for _, e := range errs {
		resp.Errors = append(resp.Errors, e.Error())
	}

	data, _ := json.MarshalIndent(resp, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleSearchTextInFile(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	fileKey, err := req.RequireString("file_key")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	searchText, err := req.RequireString("text")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	nodes, err := s.editor.SearchTextInFile(ctx, fileKey, searchText)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("search text failed: %v", err)), nil
	}

	type textResult struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		Characters string `json:"characters"`
	}

	results := make([]textResult, 0, len(nodes))
	for _, n := range nodes {
		results = append(results, textResult{
			ID:         n.ID,
			Name:       n.Name,
			Characters: n.Characters,
		})
	}

	data, _ := json.MarshalIndent(results, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) handleParseFigmaURL(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	figmaURL, err := req.RequireString("url")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	parsed, err := util.ParseFigmaURL(figmaURL)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("parse URL failed: %v", err)), nil
	}

	data, _ := json.MarshalIndent(parsed, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func toInt(v interface{}) (int, error) {
	switch val := v.(type) {
	case int:
		return val, nil
	case int64:
		return int(val), nil
	case float64:
		return int(val), nil
	case string:
		return strconv.Atoi(val)
	default:
		return 0, fmt.Errorf("cannot convert %T to int", v)
	}
}

func toFloat64(v interface{}) (float64, error) {
	switch val := v.(type) {
	case float64:
		return val, nil
	case float32:
		return float64(val), nil
	case int:
		return float64(val), nil
	case int64:
		return float64(val), nil
	case string:
		return strconv.ParseFloat(val, 64)
	default:
		return 0, fmt.Errorf("cannot convert %T to float64", v)
	}
}
