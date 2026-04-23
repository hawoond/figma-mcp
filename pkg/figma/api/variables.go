package api

import (
	"context"
	"fmt"

	"github.com/hawoond/figma-mcp/pkg/figma/client"
	"github.com/hawoond/figma-mcp/pkg/figma/types"
)

type VariablesAPI struct {
	client *client.Client
}

func NewVariablesAPI(c *client.Client) *VariablesAPI {
	return &VariablesAPI{client: c}
}

func (a *VariablesAPI) GetLocalVariables(ctx context.Context, fileKey string) (*types.LocalVariablesResponse, error) {
	var result types.LocalVariablesResponse
	if err := a.client.Get(ctx, fmt.Sprintf("files/%s/variables/local", fileKey), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *VariablesAPI) GetPublishedVariables(ctx context.Context, fileKey string) (*types.PublishedVariablesResponse, error) {
	var result types.PublishedVariablesResponse
	if err := a.client.Get(ctx, fmt.Sprintf("files/%s/variables/published", fileKey), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *VariablesAPI) PostVariables(ctx context.Context, fileKey string, req *types.PostVariablesRequest) (*types.PostVariablesResponse, error) {
	var result types.PostVariablesResponse
	if err := a.client.Post(ctx, fmt.Sprintf("files/%s/variables", fileKey), req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
