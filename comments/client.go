// Package comments provides a client for managing comments in Langfuse.
package comments

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/langfuse/langfuse-go/core"
)

// Client provides methods for comment operations
type Client struct {
	httpClient *core.HTTPClient
}

// NewClient creates a new comment client
func NewClient(httpClient *core.HTTPClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// Create creates a new comment
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	var response CreateResponse
	if err := c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/comments", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Get retrieves a comment by ID
func (c *Client) Get(ctx context.Context, commentID string) (*Comment, error) {
	var response Comment
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, "/api/public/comments/"+url.PathEscape(commentID), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// List retrieves comments with optional filtering
func (c *Client) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/api/public/comments"
	if params != nil {
		query := url.Values{}
		if params.Page != nil {
			query.Set("page", strconv.Itoa(*params.Page))
		}
		if params.Limit != nil {
			query.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.ObjectType != nil {
			query.Set("objectType", string(*params.ObjectType))
		}
		if params.ObjectID != nil {
			query.Set("objectId", *params.ObjectID)
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
