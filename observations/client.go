// Package observations provides a client for managing observations (events, spans, generations) in Langfuse.
package observations

import (
	"context"
	"net/http"
	"net/url"
	"strconv"

	"github.com/langfuse/langfuse-go/core"
)

// Client provides methods for observation operations
type Client struct {
	httpClient *core.HTTPClient
}

// NewClient creates a new observation client
func NewClient(httpClient *core.HTTPClient) *Client {
	return &Client{
		httpClient: httpClient,
	}
}

// CreateEvent creates a new event observation
func (c *Client) CreateEvent(ctx context.Context, req *CreateEventRequest) error {
	return c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/events", req, nil)
}

// UpdateEvent updates an event observation
func (c *Client) UpdateEvent(ctx context.Context, eventID string, req *UpdateEventRequest) error {
	return c.httpClient.DoRequest(ctx, http.MethodPatch, "/api/public/events/"+url.PathEscape(eventID), req, nil)
}

// CreateSpan creates a new span observation
func (c *Client) CreateSpan(ctx context.Context, req *CreateSpanRequest) error {
	return c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/spans", req, nil)
}

// UpdateSpan updates a span observation
func (c *Client) UpdateSpan(ctx context.Context, spanID string, req *UpdateSpanRequest) error {
	return c.httpClient.DoRequest(ctx, http.MethodPatch, "/api/public/spans/"+url.PathEscape(spanID), req, nil)
}

// CreateGeneration creates a new generation observation
func (c *Client) CreateGeneration(ctx context.Context, req *CreateGenerationRequest) error {
	return c.httpClient.DoRequest(ctx, http.MethodPost, "/api/public/generations", req, nil)
}

// UpdateGeneration updates a generation observation
func (c *Client) UpdateGeneration(ctx context.Context, generationID string, req *UpdateGenerationRequest) error {
	return c.httpClient.DoRequest(ctx, http.MethodPatch, "/api/public/generations/"+url.PathEscape(generationID), req, nil)
}

// Get retrieves an observation by ID
func (c *Client) Get(ctx context.Context, observationID string) (*Observation, error) {
	var response Observation
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, "/api/public/observations/"+url.PathEscape(observationID), nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// List retrieves observations with optional filtering
func (c *Client) List(ctx context.Context, params *ListParams) (*ListResponse, error) {
	path := "/api/public/observations"
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
		if params.UserID != nil {
			query.Set("userId", *params.UserID)
		}
		if params.Type != nil {
			query.Set("type", string(*params.Type))
		}
		if params.TraceID != nil {
			query.Set("traceId", *params.TraceID)
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
