// Package scores provides a client for managing scores in Langfuse.
package scores

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rohitkeshwani07/langfuse-go/core"
)

// Client provides methods for score operations
type Client struct {
	httpClient *core.HTTPClient
}

// NewClient creates a new score client
func NewClient(httpClient *core.HTTPClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// Create creates a new score
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var response CreateResponse
	if err := c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/scores", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Get retrieves a score by ID
func (c *Client) Get(ctx context.Context, scoreID string) (*Score, error) {
	var response Score
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, "/api/public/scores/"+url.PathEscape(scoreID), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// List retrieves scores with optional filtering
func (c *Client) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/api/public/scores"
	if params != nil {
		query := url.Values{}
		if params.Page != nil {
			query.Set("page", strconv.Itoa(*params.Page))
		}
		if params.Limit != nil {
			query.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.TraceID != nil {
			query.Set("traceId", *params.TraceID)
		}
		if params.UserID != nil {
			query.Set("userId", *params.UserID)
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

// Delete deletes a score by ID
func (c *Client) Delete(ctx context.Context, scoreID string) error {
	return c.httpClient.DoRequest(ctx, http.MethodDelete, "/api/public/scores/"+url.PathEscape(scoreID), nil, nil)
}
