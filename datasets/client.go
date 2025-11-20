// Package datasets provides a client for managing datasets in Langfuse.
package datasets

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/rohitkeshwani07/langfuse-go/core"
	"github.com/rohitkeshwani07/langfuse-go/types"
)

// Client provides methods for dataset operations
type Client struct {
	httpClient *core.HTTPClient
}

// NewClient creates a new dataset client
func NewClient(httpClient *core.HTTPClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// Create creates a new dataset
func (c *Client) Create(ctx context.Context, req *CreateRequest) (*Dataset, error) {
	var response Dataset
	if err := c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/datasets", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// Get retrieves a dataset by name
func (c *Client) Get(ctx context.Context, datasetName string) (*Dataset, error) {
	var response Dataset
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, "/api/public/datasets/"+url.PathEscape(datasetName), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// List retrieves all datasets with pagination
func (c *Client) List(ctx context.Context, params *types.PaginationParams) (*ListResponse, error) {
	path := "/api/public/datasets"
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

// CreateItem creates a new dataset item
func (c *Client) CreateItem(ctx context.Context, req *CreateItemRequest) (*Item, error) {
	var response Item
	if err := c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/dataset-items", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetItem retrieves a dataset item by ID
func (c *Client) GetItem(ctx context.Context, itemID string) (*Item, error) {
	var response Item
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, "/api/public/dataset-items/"+url.PathEscape(itemID), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListItems retrieves items for a dataset
func (c *Client) ListItems(ctx context.Context, datasetName string, params *types.PaginationParams) (*ItemListResponse, error) {
	path := "/api/public/datasets/" + url.PathEscape(datasetName) + "/items"
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

	var response ItemListResponse
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateRun creates a new dataset run
func (c *Client) CreateRun(ctx context.Context, req *CreateRunRequest) (*Run, error) {
	var response Run
	if err := c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/dataset-runs", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetRun retrieves a dataset run by ID
func (c *Client) GetRun(ctx context.Context, runID string) (*Run, error) {
	var response Run
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, "/api/public/dataset-runs/"+url.PathEscape(runID), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListRuns retrieves runs for a dataset
func (c *Client) ListRuns(ctx context.Context, datasetName string, params *types.PaginationParams) (*RunListResponse, error) {
	path := "/api/public/datasets/" + url.PathEscape(datasetName) + "/runs"
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

	var response RunListResponse
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateRunItem creates a new dataset run item
func (c *Client) CreateRunItem(ctx context.Context, req *CreateRunItemRequest) (*RunItem, error) {
	var response RunItem
	if err := c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/dataset-run-items", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
