// Package ingestion provides a client for batch ingestion in Langfuse.
package ingestion

import (
	"context"
	"net/http"

	"github.com/langfuse/langfuse-go/core"
)

// Client provides methods for ingestion operations
type Client struct {
	httpClient *core.HTTPClient
}

// NewClient creates a new ingestion client
func NewClient(httpClient *core.HTTPClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// Ingest sends a batch of events to the ingestion API
func (c *Client) Ingest(ctx context.Context, req *Request) (*Response, error) {
	var response Response
	if err := c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/ingestion", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
