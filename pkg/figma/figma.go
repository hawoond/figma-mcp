package figma

import (
	"github.com/hawoond/figma-mcp/pkg/figma/api"
	"github.com/hawoond/figma-mcp/pkg/figma/client"
)

type Client struct {
	Files        *api.FilesAPI
	Comments     *api.CommentsAPI
	Users        *api.UsersAPI
	Projects     *api.ProjectsAPI
	Webhooks     *api.WebhooksAPI
	Variables    *api.VariablesAPI
	DevResources *api.DevResourcesAPI
	Analytics    *api.AnalyticsAPI
}

func New(token string, opts ...client.Option) *Client {
	c := client.New(token, opts...)
	return &Client{
		Files:        api.NewFilesAPI(c),
		Comments:     api.NewCommentsAPI(c),
		Users:        api.NewUsersAPI(c),
		Projects:     api.NewProjectsAPI(c),
		Webhooks:     api.NewWebhooksAPI(c),
		Variables:    api.NewVariablesAPI(c),
		DevResources: api.NewDevResourcesAPI(c),
		Analytics:    api.NewAnalyticsAPI(c),
	}
}
