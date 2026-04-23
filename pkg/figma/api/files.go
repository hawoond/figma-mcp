package api

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/hawoond/figma-mcp/pkg/figma/client"
	"github.com/hawoond/figma-mcp/pkg/figma/types"
)

type FilesAPI struct {
	client *client.Client
}

func NewFilesAPI(c *client.Client) *FilesAPI {
	return &FilesAPI{client: c}
}

type GetFileOptions struct {
	Version    string
	IDs        []string
	Depth      *int
	Geometry   string
	PluginData string
	BranchData bool
}

func (a *FilesAPI) GetFile(ctx context.Context, fileKey string, opts *GetFileOptions) (*types.File, error) {
	params := url.Values{}
	if opts != nil {
		if opts.Version != "" {
			params.Set("version", opts.Version)
		}
		if len(opts.IDs) > 0 {
			params.Set("ids", strings.Join(opts.IDs, ","))
		}
		if opts.Depth != nil {
			params.Set("depth", strconv.Itoa(*opts.Depth))
		}
		if opts.Geometry != "" {
			params.Set("geometry", opts.Geometry)
		}
		if opts.PluginData != "" {
			params.Set("plugin_data", opts.PluginData)
		}
		if opts.BranchData {
			params.Set("branch_data", "true")
		}
	}

	var result types.File
	if err := a.client.Get(ctx, fmt.Sprintf("files/%s", fileKey), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type GetFileNodesOptions struct {
	IDs        []string
	Version    string
	Depth      *int
	Geometry   string
	PluginData string
}

func (a *FilesAPI) GetFileNodes(ctx context.Context, fileKey string, opts *GetFileNodesOptions) (*types.FileNodes, error) {
	params := url.Values{}
	if opts != nil {
		if len(opts.IDs) > 0 {
			params.Set("ids", strings.Join(opts.IDs, ","))
		}
		if opts.Version != "" {
			params.Set("version", opts.Version)
		}
		if opts.Depth != nil {
			params.Set("depth", strconv.Itoa(*opts.Depth))
		}
		if opts.Geometry != "" {
			params.Set("geometry", opts.Geometry)
		}
		if opts.PluginData != "" {
			params.Set("plugin_data", opts.PluginData)
		}
	}

	var result types.FileNodes
	if err := a.client.Get(ctx, fmt.Sprintf("files/%s/nodes", fileKey), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type GetImagesOptions struct {
	IDs     []string
	Scale   *float64
	Format  string
	SVGIncludeID bool
	SVGSimplifyStroke bool
	UseAbsoluteBounds bool
	Version string
}

func (a *FilesAPI) GetImages(ctx context.Context, fileKey string, opts *GetImagesOptions) (*types.ImageResponse, error) {
	params := url.Values{}
	if opts != nil {
		if len(opts.IDs) > 0 {
			params.Set("ids", strings.Join(opts.IDs, ","))
		}
		if opts.Scale != nil {
			params.Set("scale", strconv.FormatFloat(*opts.Scale, 'f', -1, 64))
		}
		if opts.Format != "" {
			params.Set("format", opts.Format)
		}
		if opts.SVGIncludeID {
			params.Set("svg_include_id", "true")
		}
		if opts.SVGSimplifyStroke {
			params.Set("svg_simplify_stroke", "true")
		}
		if opts.UseAbsoluteBounds {
			params.Set("use_absolute_bounds", "true")
		}
		if opts.Version != "" {
			params.Set("version", opts.Version)
		}
	}

	var result types.ImageResponse
	if err := a.client.Get(ctx, fmt.Sprintf("images/%s", fileKey), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *FilesAPI) GetImageFills(ctx context.Context, fileKey string) (*types.ImageFillsResponse, error) {
	var result types.ImageFillsResponse
	if err := a.client.Get(ctx, fmt.Sprintf("files/%s/images", fileKey), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *FilesAPI) GetFileMetadata(ctx context.Context, fileKey string) (*types.FileMetadata, error) {
	var result types.FileMetadata
	if err := a.client.Get(ctx, fmt.Sprintf("files/%s/meta", fileKey), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type GetFileVersionsOptions struct {
	PageSize *int
	Before   *int
	After    *int
}

func (a *FilesAPI) GetFileVersions(ctx context.Context, fileKey string, opts *GetFileVersionsOptions) (*types.FileVersionsResponse, error) {
	params := url.Values{}
	if opts != nil {
		if opts.PageSize != nil {
			params.Set("page_size", strconv.Itoa(*opts.PageSize))
		}
		if opts.Before != nil {
			params.Set("before", strconv.Itoa(*opts.Before))
		}
		if opts.After != nil {
			params.Set("after", strconv.Itoa(*opts.After))
		}
	}

	var result types.FileVersionsResponse
	if err := a.client.Get(ctx, fmt.Sprintf("files/%s/versions", fileKey), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
