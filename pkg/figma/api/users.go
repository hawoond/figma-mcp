package api

import (
	"context"

	"github.com/hawoond/figma-mcp/pkg/figma/client"
	"github.com/hawoond/figma-mcp/pkg/figma/types"
)

type UsersAPI struct {
	client *client.Client
}

func NewUsersAPI(c *client.Client) *UsersAPI {
	return &UsersAPI{client: c}
}

func (a *UsersAPI) GetMe(ctx context.Context) (*types.User, error) {
	var result types.User
	if err := a.client.Get(ctx, "me", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
