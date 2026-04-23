package api

import (
	"context"
	"fmt"
	"net/url"

	"github.com/hawoond/figma-mcp/pkg/figma/client"
	"github.com/hawoond/figma-mcp/pkg/figma/types"
)

type CommentsAPI struct {
	client *client.Client
}

func NewCommentsAPI(c *client.Client) *CommentsAPI {
	return &CommentsAPI{client: c}
}

func (a *CommentsAPI) GetComments(ctx context.Context, fileKey string, asMarkdown bool) (*types.CommentsResponse, error) {
	params := url.Values{}
	if asMarkdown {
		params.Set("as_md", "true")
	}

	var result types.CommentsResponse
	if err := a.client.Get(ctx, fmt.Sprintf("files/%s/comments", fileKey), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *CommentsAPI) PostComment(ctx context.Context, fileKey string, req *types.PostCommentRequest) (*types.Comment, error) {
	var result types.Comment
	if err := a.client.Post(ctx, fmt.Sprintf("files/%s/comments", fileKey), req, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *CommentsAPI) DeleteComment(ctx context.Context, fileKey, commentID string) error {
	return a.client.Delete(ctx, fmt.Sprintf("files/%s/comments/%s", fileKey, commentID), nil, nil)
}

func (a *CommentsAPI) GetCommentReactions(ctx context.Context, fileKey, commentID string, cursor string) (map[string]interface{}, error) {
	params := url.Values{}
	if cursor != "" {
		params.Set("cursor", cursor)
	}

	var result map[string]interface{}
	if err := a.client.Get(ctx, fmt.Sprintf("files/%s/comments/%s/reactions", fileKey, commentID), params, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (a *CommentsAPI) PostCommentReaction(ctx context.Context, fileKey, commentID, emoji string) error {
	body := map[string]string{"emoji": emoji}
	return a.client.Post(ctx, fmt.Sprintf("files/%s/comments/%s/reactions", fileKey, commentID), body, nil)
}

func (a *CommentsAPI) DeleteCommentReaction(ctx context.Context, fileKey, commentID, emoji string) error {
	params := url.Values{}
	params.Set("emoji", emoji)
	return a.client.Delete(ctx, fmt.Sprintf("files/%s/comments/%s/reactions", fileKey, commentID), params, nil)
}
