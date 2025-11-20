// Package prompts provides a client for managing prompts in Langfuse.
package prompts

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/langfuse/langfuse-go/core"
)

// Client provides methods for prompt operations
type Client struct {
	httpClient *core.HTTPClient
}

// NewClient creates a new prompt client
func NewClient(httpClient *core.HTTPClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// Get retrieves a prompt by name and version or label
func (c *Client) Get(ctx context.Context, promptName string, params *GetParams) (interface{}, error) {
	path := "/api/public/prompts/" + url.PathEscape(promptName)
	if params != nil {
		query := url.Values{}
		if params.Version != nil {
			query.Set("version", strconv.Itoa(*params.Version))
		}
		if params.Label != nil {
			query.Set("label", *params.Label)
		}
		if len(query) > 0 {
			path += "?" + query.Encode()
		}
	}

	var response interface{}
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}
	return response, nil
}

// List retrieves all prompts with pagination
func (c *Client) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/api/public/prompts"
	if params != nil {
		query := url.Values{}
		if params.Page != nil {
			query.Set("page", strconv.Itoa(*params.Page))
		}
		if params.Limit != nil {
			query.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Name != nil {
			query.Set("name", *params.Name)
		}
		if params.Label != nil {
			query.Set("label", *params.Label)
		}
		if params.Tag != nil {
			query.Set("tag", *params.Tag)
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

// CreateChat creates a new chat prompt
func (c *Client) CreateChat(ctx context.Context, req *CreateChatRequest) (*Chat, error) {
	var response Chat
	if err := c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/prompts", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateText creates a new text prompt
func (c *Client) CreateText(ctx context.Context, req *CreateTextRequest) (*Text, error) {
	var response Text
	if err := c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/prompts", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
