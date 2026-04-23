package api

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/hawoond/figma-mcp/pkg/figma/client"
	"github.com/hawoond/figma-mcp/pkg/figma/types"
)

type AnalyticsAPI struct {
	client *client.Client
}

func NewAnalyticsAPI(c *client.Client) *AnalyticsAPI {
	return &AnalyticsAPI{client: c}
}

type GetActivityLogsOptions struct {
	Events    []string
	StartTime *int64
	EndTime   *int64
	Limit     *int
	Cursor    string
	Order     string
}

func (a *AnalyticsAPI) GetActivityLogs(ctx context.Context, opts *GetActivityLogsOptions) (*types.ActivityLogsResponse, error) {
	params := url.Values{}
	if opts != nil {
		for _, e := range opts.Events {
			params.Add("events", e)
		}
		if opts.StartTime != nil {
			params.Set("start_time", strconv.FormatInt(*opts.StartTime, 10))
		}
		if opts.EndTime != nil {
			params.Set("end_time", strconv.FormatInt(*opts.EndTime, 10))
		}
		if opts.Limit != nil {
			params.Set("limit", strconv.Itoa(*opts.Limit))
		}
		if opts.Cursor != "" {
			params.Set("cursor", opts.Cursor)
		}
		if opts.Order != "" {
			params.Set("order", opts.Order)
		}
	}

	var result types.ActivityLogsResponse
	if err := a.client.Get(ctx, "activity_logs", params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type GetLibraryAnalyticsOptions struct {
	Cursor    string
	GroupBy   string
	StartDate string
	EndDate   string
	Order     string
	PageSize  *int
}

func (a *AnalyticsAPI) GetLibraryAnalyticsComponents(ctx context.Context, fileKey string, opts *GetLibraryAnalyticsOptions) (*types.LibraryAnalyticsComponentsResponse, error) {
	params := url.Values{}
	if opts != nil {
		if opts.Cursor != "" {
			params.Set("cursor", opts.Cursor)
		}
		if opts.GroupBy != "" {
			params.Set("group_by", opts.GroupBy)
		}
		if opts.StartDate != "" {
			params.Set("start_date", opts.StartDate)
		}
		if opts.EndDate != "" {
			params.Set("end_date", opts.EndDate)
		}
		if opts.Order != "" {
			params.Set("order", opts.Order)
		}
		if opts.PageSize != nil {
			params.Set("page_size", strconv.Itoa(*opts.PageSize))
		}
	}

	var result types.LibraryAnalyticsComponentsResponse
	if err := a.client.Get(ctx, fmt.Sprintf("analytics/libraries/%s/component", fileKey), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *AnalyticsAPI) GetLibraryAnalyticsStyles(ctx context.Context, fileKey string, opts *GetLibraryAnalyticsOptions) (*types.LibraryAnalyticsStylesResponse, error) {
	params := url.Values{}
	if opts != nil {
		if opts.Cursor != "" {
			params.Set("cursor", opts.Cursor)
		}
		if opts.GroupBy != "" {
			params.Set("group_by", opts.GroupBy)
		}
		if opts.StartDate != "" {
			params.Set("start_date", opts.StartDate)
		}
		if opts.EndDate != "" {
			params.Set("end_date", opts.EndDate)
		}
		if opts.Order != "" {
			params.Set("order", opts.Order)
		}
		if opts.PageSize != nil {
			params.Set("page_size", strconv.Itoa(*opts.PageSize))
		}
	}

	var result types.LibraryAnalyticsStylesResponse
	if err := a.client.Get(ctx, fmt.Sprintf("analytics/libraries/%s/style", fileKey), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (a *AnalyticsAPI) GetLibraryAnalyticsVariables(ctx context.Context, fileKey string, opts *GetLibraryAnalyticsOptions) (*types.LibraryAnalyticsVariablesResponse, error) {
	params := url.Values{}
	if opts != nil {
		if opts.Cursor != "" {
			params.Set("cursor", opts.Cursor)
		}
		if opts.GroupBy != "" {
			params.Set("group_by", opts.GroupBy)
		}
		if opts.StartDate != "" {
			params.Set("start_date", opts.StartDate)
		}
		if opts.EndDate != "" {
			params.Set("end_date", opts.EndDate)
		}
		if opts.Order != "" {
			params.Set("order", opts.Order)
		}
		if opts.PageSize != nil {
			params.Set("page_size", strconv.Itoa(*opts.PageSize))
		}
	}

	var result types.LibraryAnalyticsVariablesResponse
	if err := a.client.Get(ctx, fmt.Sprintf("analytics/libraries/%s/variable", fileKey), params, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
