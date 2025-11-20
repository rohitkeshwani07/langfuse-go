// Package annotations provides a client for managing annotation queues in Langfuse.
package annotations

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/langfuse/langfuse-go/core"
	"github.com/langfuse/langfuse-go/types"
)

// Client provides methods for annotation queue operations
type Client struct {
	httpClient *core.HTTPClient
}

// NewClient creates a new annotations client
func NewClient(httpClient *core.HTTPClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// ListQueues retrieves all annotation queues
func (c *Client) ListQueues(ctx context.Context, params *types.PaginationParams) (*ListQueuesResponse, error) {
	path := "/api/public/annotation-queues"
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

	var response ListQueuesResponse
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateQueue creates a new annotation queue
func (c *Client) CreateQueue(ctx context.Context, req *CreateQueueRequest) (*Queue, error) {
	var response Queue
	if err := c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/annotation-queues", req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetQueue retrieves an annotation queue by ID
func (c *Client) GetQueue(ctx context.Context, queueID string) (*Queue, error) {
	var response Queue
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, "/api/public/annotation-queues/"+url.PathEscape(queueID), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// ListQueueItems retrieves items for an annotation queue
func (c *Client) ListQueueItems(ctx context.Context, queueID string, params *ListQueueItemsParams) (*ListQueueItemsResponse, error) {
	path := "/api/public/annotation-queues/" + url.PathEscape(queueID) + "/items"
	if params != nil {
		query := url.Values{}
		if params.Page != nil {
			query.Set("page", strconv.Itoa(*params.Page))
		}
		if params.Limit != nil {
			query.Set("limit", strconv.Itoa(*params.Limit))
		}
		if params.Status != nil {
			query.Set("status", string(*params.Status))
		}
		if len(query) > 0 {
			path += "?" + query.Encode()
		}
	}

	var response ListQueueItemsResponse
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateQueueItem creates a new annotation queue item
func (c *Client) CreateQueueItem(ctx context.Context, queueID string, req *CreateQueueItemRequest) (*QueueItem, error) {
	var response QueueItem
	path := "/api/public/annotation-queues/" + url.PathEscape(queueID) + "/items"
	if err := c.httpClient.DoRequest(ctx, http.MethodPost, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetQueueItem retrieves an annotation queue item
func (c *Client) GetQueueItem(ctx context.Context, queueID, itemID string) (*QueueItem, error) {
	var response QueueItem
	path := "/api/public/annotation-queues/" + url.PathEscape(queueID) + "/items/" + url.PathEscape(itemID)
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// UpdateQueueItem updates an annotation queue item
func (c *Client) UpdateQueueItem(ctx context.Context, queueID, itemID string, req *UpdateQueueItemRequest) (*QueueItem, error) {
	var response QueueItem
	path := "/api/public/annotation-queues/" + url.PathEscape(queueID) + "/items/" + url.PathEscape(itemID)
	if err := c.httpClient.DoRequest(ctx, http.MethodPatch, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteQueueItem deletes an annotation queue item
func (c *Client) DeleteQueueItem(ctx context.Context, queueID, itemID string) (*DeleteQueueItemResponse, error) {
	var response DeleteQueueItemResponse
	path := "/api/public/annotation-queues/" + url.PathEscape(queueID) + "/items/" + url.PathEscape(itemID)
	if err := c.httpClient.DoRequest(ctx, http.MethodDelete, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// CreateAssignment creates an assignment for a user to an annotation queue
func (c *Client) CreateAssignment(ctx context.Context, queueID string, req *AssignmentRequest) (*CreateAssignmentResponse, error) {
	var response CreateAssignmentResponse
	path := "/api/public/annotation-queues/" + url.PathEscape(queueID) + "/assignments"
	if err := c.httpClient.DoRequest(ctx, http.MethodPost, path, req, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// DeleteAssignment deletes an assignment from an annotation queue
func (c *Client) DeleteAssignment(ctx context.Context, queueID, userID string) (*DeleteAssignmentResponse, error) {
	var response DeleteAssignmentResponse
	path := "/api/public/annotation-queues/" + url.PathEscape(queueID) + "/assignments/" + url.PathEscape(userID)
	if err := c.httpClient.DoRequest(ctx, http.MethodDelete, path, nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
