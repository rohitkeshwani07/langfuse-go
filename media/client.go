// Package media provides a client for managing media in Langfuse.
package media

import (
	"context"
	"net/http"
	"net/url"

	"github.com/langfuse/langfuse-go/core"
)

// Client provides methods for media operations
type Client struct {
	httpClient *core.HTTPClient
}

// NewClient creates a new media client
func NewClient(httpClient *core.HTTPClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// Get retrieves media metadata by ID
func (c *Client) Get(ctx context.Context, mediaID string) (*Response, error) {
	var response Response
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, "/api/public/media/"+url.PathEscape(mediaID), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetUploadURL retrieves an upload URL for media
func (c *Client) GetUploadURL(ctx context.Context, req *UploadURLRequest) (*UploadURLResponse, error) {
	var response UploadURLResponse
	if err := c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/media", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Patch updates media metadata
func (c *Client) Patch(ctx context.Context, mediaID string, req *PatchRequest) error {
	return c.httpClient.DoRequest(ctx, http.MethodPatch, "/api/public/media/"+url.PathEscape(mediaID), req, nil)
}
