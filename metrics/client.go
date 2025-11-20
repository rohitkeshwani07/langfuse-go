// Package metrics provides a client for accessing metrics in Langfuse.
package metrics

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/rohitkeshwani07/langfuse-go/core"
)

// Client provides methods for metrics operations
type Client struct {
	httpClient *core.HTTPClient
}

// NewClient creates a new metrics client
func NewClient(httpClient *core.HTTPClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// GetDaily retrieves daily metrics
func (c *Client) GetDaily(ctx context.Context, params *DailyParams) ([]Daily, error) {
	path := "/api/public/metrics/daily"
	if params != nil {
		query := url.Values{}
		if params.TraceName != nil {
			query.Set("traceName", *params.TraceName)
		}
		if params.UserID != nil {
			query.Set("userId", *params.UserID)
		}
		if params.Tags != nil {
			for _, tag := range params.Tags {
				query.Add("tags", tag)
			}
		}
		if params.FromTimestamp != nil {
			query.Set("fromTimestamp", params.FromTimestamp.Format(time.RFC3339))
		}
		if params.ToTimestamp != nil {
			query.Set("toTimestamp", params.ToTimestamp.Format(time.RFC3339))
		}
		if len(query) > 0 {
			path += "?" + query.Encode()
		}
	}

	var response []Daily
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}
