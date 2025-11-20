// Package sessions provides a client for managing sessions in Langfuse.
package sessions

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/langfuse/langfuse-go/core"
	"github.com/langfuse/langfuse-go/types"
)

// Client provides methods for session operations
type Client struct {
	httpClient *core.HTTPClient
}

// NewClient creates a new session client
func NewClient(httpClient *core.HTTPClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// Get retrieves a session by ID
func (c *Client) Get(ctx context.Context, sessionID string) (*WithTraces, error) {
	var response WithTraces
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, "/api/public/sessions/"+url.PathEscape(sessionID), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// List retrieves all sessions with pagination
func (c *Client) List(ctx context.Context, params *types.PaginationParams) (*ListResponse, error) {
	path := "/api/public/sessions"
	if params != nil {
		query := url.Values{}
		if params.Page != nil {
			query.Set("page", strconv.Itoa(*params.Page))
		}
		if params.Limit != nil {
			query.Set("limit", strconv.Itoa(*params.Limit))
		}
		if len(query) > 0 {
			path += "?" + query.Encode()
		}
	}

	var response ListResponse
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
