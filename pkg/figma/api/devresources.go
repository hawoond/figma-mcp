package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hawoond/figma-mcp/pkg/figma/client"
	"github.com/hawoond/figma-mcp/pkg/figma/types"
)

type DevResourcesAPI struct {
	client *client.Client
}

func NewDevResourcesAPI(c *client.Client) *DevResourcesAPI {
	return &DevResourcesAPI{client: c}
}

func (a *DevResourcesAPI) GetDevResources(ctx context.Context, fileKey string, nodeIDs []string) (*types.DevResourcesResponse, error) {
	params := url.Values{}
	if len(nodeIDs) > 0 {
		for _, id := range nodeIDs {
			params.Add("node_ids", id)
		}
	}

	var result types.DevResourcesResponse
	if err := a.client.Get(ctx, fmt.Sprintf("files/%s/dev_resources", fileKey), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *DevResourcesAPI) CreateDevResources(ctx context.Context, req *types.CreateDevResourcesRequest) (*types.CreateDevResourcesResponse, error) {
	var result types.CreateDevResourcesResponse
	if err := a.client.Post(ctx, "dev_resources", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *DevResourcesAPI) UpdateDevResources(ctx context.Context, req *types.UpdateDevResourcesRequest) (*types.DevResourcesResponse, error) {
	var result types.DevResourcesResponse
	if err := a.client.Put(ctx, "dev_resources", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *DevResourcesAPI) DeleteDevResource(ctx context.Context, fileKey, devResourceID string) error {
	return a.client.Delete(ctx, fmt.Sprintf("files/%s/dev_resources/%s", fileKey, devResourceID), nil, nil)
}
