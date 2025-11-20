// Package traces provides a client for managing traces in Langfuse.
package traces

import (
	"context"
	"net/http"
	"net/url"

	"github.com/langfuse/langfuse-go/core"
)

// Client provides methods for trace operations
type Client struct {
	httpClient *core.HTTPClient
}

// NewClient creates a new trace client
func NewClient(httpClient *core.HTTPClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// Create creates a new trace
func (c *Client) Create(ctx context.Context, req *CreateTraceRequest) error {
	return c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/traces", req, nil)
}

// Get retrieves a trace by ID
func (c *Client) Get(ctx context.Context, traceID string) (*Trace, error) {
	var response Trace
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, "/api/public/traces/"+url.PathEscape(traceID), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetIn retrieves a trace by ID into the provided output variable.
// This is an optimization to allow reusing allocated memory and avoid allocations.
func (c *Client) GetIn(ctx context.Context, traceID string, out interface{}) error {
	return c.httpClient.DoRequest(ctx, http.MethodGet, "/api/public/traces/"+url.PathEscape(traceID), nil, out)
}

// Update updates a trace
func (c *Client) Update(ctx context.Context, traceID string, req *UpdateTraceRequest) error {
	return c.httpClient.DoRequest(ctx, http.MethodPatch, "/api/public/traces/"+url.PathEscape(traceID), req, nil)
}

// GetTree retrieves a trace by ID with observations in a tree structure
func (c *Client) GetTree(ctx context.Context, traceID string) (*TraceTree, error) {
	var response CompactTrace
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, "/api/public/traces/"+url.PathEscape(traceID), nil, &response); err != nil {
		return nil, err
	}
	return response.ToTraceTree(), nil
}
