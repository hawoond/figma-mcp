package api

import (
	"context"
	"fmt"

	"github.com/hawoond/figma-mcp/pkg/figma/client"
	"github.com/hawoond/figma-mcp/pkg/figma/types"
)

type WebhooksAPI struct {
	client *client.Client
}

func NewWebhooksAPI(c *client.Client) *WebhooksAPI {
	return &WebhooksAPI{client: c}
}

func (a *WebhooksAPI) CreateWebhook(ctx context.Context, req *types.WebhookRequest) (*types.Webhook, error) {
	var result types.Webhook
	if err := a.client.Post(ctx, "webhooks", req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *WebhooksAPI) GetWebhook(ctx context.Context, webhookID string) (*types.Webhook, error) {
	var result types.Webhook
	if err := a.client.Get(ctx, fmt.Sprintf("webhooks/%s", webhookID), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *WebhooksAPI) UpdateWebhook(ctx context.Context, webhookID string, req *types.WebhookRequest) (*types.Webhook, error) {
	var result types.Webhook
	if err := a.client.Put(ctx, fmt.Sprintf("webhooks/%s", webhookID), req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *WebhooksAPI) DeleteWebhook(ctx context.Context, webhookID string) error {
	return a.client.Delete(ctx, fmt.Sprintf("webhooks/%s", webhookID), nil, nil)
}

func (a *WebhooksAPI) GetTeamWebhooks(ctx context.Context, teamID string) (*types.WebhooksResponse, error) {
	var result types.WebhooksResponse
	if err := a.client.Get(ctx, fmt.Sprintf("teams/%s/webhooks", teamID), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *WebhooksAPI) GetWebhookPayloads(ctx context.Context, webhookID string) (*types.WebhookPayloadsResponse, error) {
	var result types.WebhookPayloadsResponse
	if err := a.client.Get(ctx, fmt.Sprintf("webhooks/%s/requests", webhookID), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
