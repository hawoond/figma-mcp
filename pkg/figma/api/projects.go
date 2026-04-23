package api

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/hawoond/figma-mcp/pkg/figma/client"
	"github.com/hawoond/figma-mcp/pkg/figma/types"
)

type ProjectsAPI struct {
	client *client.Client
}

func NewProjectsAPI(c *client.Client) *ProjectsAPI {
	return &ProjectsAPI{client: c}
}

func (a *ProjectsAPI) GetTeamProjects(ctx context.Context, teamID string) (*types.TeamProjectsResponse, error) {
	var result types.TeamProjectsResponse
	if err := a.client.Get(ctx, fmt.Sprintf("teams/%s/projects", teamID), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type GetProjectFilesOptions struct {
	BranchData bool
}

func (a *ProjectsAPI) GetProjectFiles(ctx context.Context, projectID string, opts *GetProjectFilesOptions) (*types.ProjectFilesResponse, error) {
	params := url.Values{}
	if opts != nil && opts.BranchData {
		params.Set("branch_data", "true")
	}

	var result types.ProjectFilesResponse
	if err := a.client.Get(ctx, fmt.Sprintf("projects/%s/files", projectID), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type GetTeamComponentsOptions struct {
	PageSize *int
	After    *int
	Before   *int
}

func (a *ProjectsAPI) GetTeamComponents(ctx context.Context, teamID string, opts *GetTeamComponentsOptions) (*types.TeamComponentsResponse, error) {
	params := url.Values{}
	if opts != nil {
		if opts.PageSize != nil {
			params.Set("page_size", strconv.Itoa(*opts.PageSize))
		}
		if opts.After != nil {
			params.Set("after", strconv.Itoa(*opts.After))
		}
		if opts.Before != nil {
			params.Set("before", strconv.Itoa(*opts.Before))
		}
	}

	var result types.TeamComponentsResponse
	if err := a.client.Get(ctx, fmt.Sprintf("teams/%s/components", teamID), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *ProjectsAPI) GetFileComponents(ctx context.Context, fileKey string) (*types.TeamComponentsResponse, error) {
	var result types.TeamComponentsResponse
	if err := a.client.Get(ctx, fmt.Sprintf("files/%s/components", fileKey), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *ProjectsAPI) GetComponent(ctx context.Context, key string) (*types.TeamComponent, error) {
	var result types.TeamComponent
	if err := a.client.Get(ctx, fmt.Sprintf("components/%s", key), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *ProjectsAPI) GetTeamComponentSets(ctx context.Context, teamID string, opts *GetTeamComponentsOptions) (*types.TeamComponentSetsResponse, error) {
	params := url.Values{}
	if opts != nil {
		if opts.PageSize != nil {
			params.Set("page_size", strconv.Itoa(*opts.PageSize))
		}
		if opts.After != nil {
			params.Set("after", strconv.Itoa(*opts.After))
		}
		if opts.Before != nil {
			params.Set("before", strconv.Itoa(*opts.Before))
		}
	}

	var result types.TeamComponentSetsResponse
	if err := a.client.Get(ctx, fmt.Sprintf("teams/%s/component_sets", teamID), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *ProjectsAPI) GetFileComponentSets(ctx context.Context, fileKey string) (*types.TeamComponentSetsResponse, error) {
	var result types.TeamComponentSetsResponse
	if err := a.client.Get(ctx, fmt.Sprintf("files/%s/component_sets", fileKey), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *ProjectsAPI) GetComponentSet(ctx context.Context, key string) (*types.TeamComponentSet, error) {
	var result types.TeamComponentSet
	if err := a.client.Get(ctx, fmt.Sprintf("component_sets/%s", key), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *ProjectsAPI) GetTeamStyles(ctx context.Context, teamID string, opts *GetTeamComponentsOptions) (*types.TeamStylesResponse, error) {
	params := url.Values{}
	if opts != nil {
		if opts.PageSize != nil {
			params.Set("page_size", strconv.Itoa(*opts.PageSize))
		}
		if opts.After != nil {
			params.Set("after", strconv.Itoa(*opts.After))
		}
		if opts.Before != nil {
			params.Set("before", strconv.Itoa(*opts.Before))
		}
	}

	var result types.TeamStylesResponse
	if err := a.client.Get(ctx, fmt.Sprintf("teams/%s/styles", teamID), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *ProjectsAPI) GetFileStyles(ctx context.Context, fileKey string) (*types.TeamStylesResponse, error) {
	var result types.TeamStylesResponse
	if err := a.client.Get(ctx, fmt.Sprintf("files/%s/styles", fileKey), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *ProjectsAPI) GetStyle(ctx context.Context, key string) (*types.TeamStyle, error) {
	var result types.TeamStyle
	if err := a.client.Get(ctx, fmt.Sprintf("styles/%s", key), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
