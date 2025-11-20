// Package models provides a client for managing model configurations in Langfuse.
package models

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/langfuse/langfuse-go/core"
	"github.com/langfuse/langfuse-go/types"
)

// Client provides methods for model operations
type Client struct {
	httpClient *core.HTTPClient
}

// NewClient creates a new model client
func NewClient(httpClient *core.HTTPClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// Create creates a new model configuration
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*Model, error) {
	var response Model
	if err := c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/models", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Get retrieves a model by ID
func (c *Client) Get(ctx context.Context, modelID string) (*Model, error) {
	var response Model
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, "/api/public/models/"+url.PathEscape(modelID), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// List retrieves all models with pagination
func (c *Client) List(ctx context.Context, params *types.PaginationParams) (*ListResponse, error) {
	path := "/api/public/models"
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

// Delete deletes a model by ID
func (c *Client) Delete(ctx context.Context, modelID string) error {
	return c.httpClient.DoRequest(ctx, http.MethodDelete, "/api/public/models/"+url.PathEscape(modelID), nil, nil)
}
