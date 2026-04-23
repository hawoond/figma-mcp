package util

import (
	"context"
	"fmt"
	"strings"

	"github.com/hawoond/figma-mcp/pkg/figma/api"
	"github.com/hawoond/figma-mcp/pkg/figma/types"
)

type Editor struct {
	files     *api.FilesAPI
	variables *api.VariablesAPI
}

func NewEditor(files *api.FilesAPI, variables *api.VariablesAPI) *Editor {
	return &Editor{
		files:     files,
		variables: variables,
	}
}

type UploadImageFromURLResult struct {
	ImageRef  string `json:"image_ref"`
	Format    string `json:"format"`
	SizeBytes int    `json:"size_bytes"`
}

func (e *Editor) UploadImageFromURL(ctx context.Context, fileKey string, imageURL string) (*UploadImageFromURLResult, error) {
	data, format, err := FetchImageFromURL(ctx, imageURL)
	if err != nil {
		return nil, fmt.Errorf("fetch image from URL: %w", err)
	}

	mimeType := FormatToMIMEType(format)
	resp, err := e.files.UploadImage(ctx, fileKey, data, mimeType)
	if err != nil {
		return nil, fmt.Errorf("upload image to figma: %w", err)
	}

	return &UploadImageFromURLResult{
		ImageRef:  resp.ImageRef,
		Format:    format,
		SizeBytes: len(data),
	}, nil
}

type UploadAndApplyResult struct {
	ImageRef  string `json:"image_ref"`
	NodeID    string `json:"node_id"`
	ScaleMode string `json:"scale_mode"`
	Format    string `json:"format"`
	SizeBytes int    `json:"size_bytes"`
}

func (e *Editor) UploadImageFromURLAndGetRef(ctx context.Context, fileKey string, imageURL string) (*UploadImageFromURLResult, error) {
	return e.UploadImageFromURL(ctx, fileKey, imageURL)
}

func (e *Editor) UploadMultipleImagesFromURLs(ctx context.Context, fileKey string, imageURLs []string) ([]*UploadImageFromURLResult, []error) {
	results := make([]*UploadImageFromURLResult, 0, len(imageURLs))
	var errs []error

	for _, imgURL := range imageURLs {
		result, err := e.UploadImageFromURL(ctx, fileKey, imgURL)
		if err != nil {
			errs = append(errs, fmt.Errorf("URL %s: %w", imgURL, err))
			continue
		}
		results = append(results, result)
	}

	return results, errs
}

func (e *Editor) GetFileWithNodeSearch(ctx context.Context, fileKey string, searchName string) ([]*types.Node, error) {
	file, err := e.files.GetFile(ctx, fileKey, nil)
	if err != nil {
		return nil, fmt.Errorf("get file: %w", err)
	}
	return FindNodesByName(&file.Document, searchName, false), nil
}

func (e *Editor) GetNodesByType(ctx context.Context, fileKey string, nodeType string) ([]*types.Node, error) {
	file, err := e.files.GetFile(ctx, fileKey, nil)
	if err != nil {
		return nil, fmt.Errorf("get file: %w", err)
	}
	return FindNodesByType(&file.Document, nodeType), nil
}

func (e *Editor) ExportAllImages(ctx context.Context, fileKey string, format string, scale float64) (map[string]string, error) {
	file, err := e.files.GetFile(ctx, fileKey, nil)
	if err != nil {
		return nil, fmt.Errorf("get file: %w", err)
	}

	imageNodes := CollectImageNodes(&file.Document)
	if len(imageNodes) == 0 {
		return map[string]string{}, nil
	}

	nodeIDs := make([]string, 0, len(imageNodes))
	for _, n := range imageNodes {
		nodeIDs = append(nodeIDs, n.ID)
	}

	opts := &api.GetImagesOptions{
		IDs:    nodeIDs,
		Format: format,
	}
	if scale > 0 {
		opts.Scale = &scale
	}

	resp, err := e.files.GetImages(ctx, fileKey, opts)
	if err != nil {
		return nil, fmt.Errorf("get images: %w", err)
	}

	return resp.Images, nil
}

func (e *Editor) ExportFrames(ctx context.Context, fileKey string, format string, scale float64) (map[string]string, error) {
	file, err := e.files.GetFile(ctx, fileKey, nil)
	if err != nil {
		return nil, fmt.Errorf("get file: %w", err)
	}

	var frameIDs []string
	WalkNodes(&file.Document, func(node *types.Node, depth int) {
		if node.Type == "FRAME" || node.Type == "COMPONENT" || node.Type == "COMPONENT_SET" {
			frameIDs = append(frameIDs, node.ID)
		}
	})

	if len(frameIDs) == 0 {
		return map[string]string{}, nil
	}

	opts := &api.GetImagesOptions{
		IDs:    frameIDs,
		Format: format,
	}
	if scale > 0 {
		opts.Scale = &scale
	}

	resp, err := e.files.GetImages(ctx, fileKey, opts)
	if err != nil {
		return nil, fmt.Errorf("get images: %w", err)
	}

	return resp.Images, nil
}

func (e *Editor) GetAllImageFillURLs(ctx context.Context, fileKey string) (map[string]string, error) {
	resp, err := e.files.GetImageFills(ctx, fileKey)
	if err != nil {
		return nil, fmt.Errorf("get image fills: %w", err)
	}
	return resp.Meta.Images, nil
}

func (e *Editor) DownloadAllImageFills(ctx context.Context, fileKey string) ([]DownloadedAsset, error) {
	urls, err := e.GetAllImageFillURLs(ctx, fileKey)
	if err != nil {
		return nil, err
	}

	assets, errs := DownloadAssets(ctx, urls)
	if len(errs) > 0 {
		errMsgs := make([]string, 0, len(errs))
		for _, e := range errs {
			errMsgs = append(errMsgs, e.Error())
		}
		return assets, fmt.Errorf("partial download errors: %s", strings.Join(errMsgs, "; "))
	}

	return assets, nil
}

type VariableSummary struct {
	Collections []CollectionSummary `json:"collections"`
}

type CollectionSummary struct {
	ID        string         `json:"id"`
	Name      string         `json:"name"`
	Modes     []string       `json:"modes"`
	Variables []VariableInfo `json:"variables"`
}

type VariableInfo struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	Type         string      `json:"type"`
	DefaultValue interface{} `json:"default_value"`
}

func (e *Editor) GetVariableSummary(ctx context.Context, fileKey string) (*VariableSummary, error) {
	resp, err := e.variables.GetLocalVariables(ctx, fileKey)
	if err != nil {
		return nil, fmt.Errorf("get local variables: %w", err)
	}

	summary := &VariableSummary{}

	for _, col := range resp.Meta.VariableCollections {
		cs := CollectionSummary{
			ID:   col.ID,
			Name: col.Name,
		}
		for _, mode := range col.Modes {
			cs.Modes = append(cs.Modes, mode.Name)
		}

		for _, varID := range col.VariableIDs {
			variable, ok := resp.Meta.Variables[varID]
			if !ok {
				continue
			}

			var defaultValue interface{}
			if len(col.Modes) > 0 {
				defaultValue = variable.ValuesByMode[col.DefaultModeID]
			}

			cs.Variables = append(cs.Variables, VariableInfo{
				ID:           variable.ID,
				Name:         variable.Name,
				Type:         variable.ResolvedType,
				DefaultValue: defaultValue,
			})
		}

		summary.Collections = append(summary.Collections, cs)
	}

	return summary, nil
}

func (e *Editor) SearchTextInFile(ctx context.Context, fileKey string, searchText string) ([]*types.Node, error) {
	file, err := e.files.GetFile(ctx, fileKey, nil)
	if err != nil {
		return nil, fmt.Errorf("get file: %w", err)
	}

	var results []*types.Node
	WalkNodes(&file.Document, func(node *types.Node, depth int) {
		if node.Type == "TEXT" && strings.Contains(node.Characters, searchText) {
			results = append(results, node)
		}
	})

	return results, nil
}

func (e *Editor) GetFileStructure(ctx context.Context, fileKey string, maxDepth int) ([]NodeSummary, error) {
	depth := maxDepth
	opts := &api.GetFileOptions{
		Depth: &depth,
	}

	file, err := e.files.GetFile(ctx, fileKey, opts)
	if err != nil {
		return nil, fmt.Errorf("get file: %w", err)
	}

	var summaries []NodeSummary
	WalkNodes(&file.Document, func(node *types.Node, d int) {
		if d <= maxDepth {
			summaries = append(summaries, SummarizeNode(node))
		}
	})

	return summaries, nil
}

func (e *Editor) GetNodeDetails(ctx context.Context, fileKey string, nodeIDs []string) (map[string]*types.Node, error) {
	resp, err := e.files.GetFileNodes(ctx, fileKey, &api.GetFileNodesOptions{
		IDs: nodeIDs,
	})
	if err != nil {
		return nil, fmt.Errorf("get file nodes: %w", err)
	}

	result := make(map[string]*types.Node, len(resp.Nodes))
	for id, fileNode := range resp.Nodes {
		node := fileNode.Document
		result[id] = &node
	}
	return result, nil
}

func (e *Editor) ExportNodeAsImage(ctx context.Context, fileKey, nodeID, format string, scale float64) (string, error) {
	opts := &api.GetImagesOptions{
		IDs:    []string{nodeID},
		Format: format,
	}
	if scale > 0 {
		opts.Scale = &scale
	}

	resp, err := e.files.GetImages(ctx, fileKey, opts)
	if err != nil {
		return "", fmt.Errorf("get image: %w", err)
	}

	imgURL, ok := resp.Images[nodeID]
	if !ok {
		return "", fmt.Errorf("no image URL returned for node %s", nodeID)
	}

	return imgURL, nil
}

func (e *Editor) FetchAndEncodeNodeImage(ctx context.Context, fileKey, nodeID, format string, scale float64) (string, error) {
	imgURL, err := e.ExportNodeAsImage(ctx, fileKey, nodeID, format, scale)
	if err != nil {
		return "", err
	}

	data, imgFormat, err := FetchImageFromURL(ctx, imgURL)
	if err != nil {
		return "", fmt.Errorf("fetch image: %w", err)
	}

	return ImageToBase64(data, imgFormat), nil
}
