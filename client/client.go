// Package client provides a unified client for the Langfuse API.
// It combines all domain-specific clients into a single interface.
package client

import (
	"context"
	"net/http"

	"github.com/langfuse/langfuse-go/annotations"
	"github.com/langfuse/langfuse-go/comments"
	"github.com/langfuse/langfuse-go/core"
	"github.com/langfuse/langfuse-go/datasets"
	"github.com/langfuse/langfuse-go/ingestion"
	"github.com/langfuse/langfuse-go/media"
	"github.com/langfuse/langfuse-go/metrics"
	"github.com/langfuse/langfuse-go/models"
	"github.com/langfuse/langfuse-go/observations"
	"github.com/langfuse/langfuse-go/prompts"
	"github.com/langfuse/langfuse-go/scores"
	"github.com/langfuse/langfuse-go/sessions"
	"github.com/langfuse/langfuse-go/traces"
	"github.com/langfuse/langfuse-go/types"
)

// Client is the main Langfuse API client that provides access to all sub-clients.
type Client struct {
	httpClient *core.HTTPClient

	// Domain-specific clients
	Traces       *traces.Client
	Observations *observations.Client
	Scores       *scores.Client
	Datasets     *datasets.Client
	Sessions     *sessions.Client
	Models       *models.Client
	Prompts      *prompts.Client
	Comments     *comments.Client
	Media        *media.Client
	Metrics      *metrics.Client
	Annotations  *annotations.Client
	Ingestion    *ingestion.Client
}

// New creates a new Langfuse client with the given credentials and options.
func New(publicKey, secretKey string, opts ...core.Option) *Client {
	httpClient := core.NewHTTPClient(publicKey, secretKey, opts...)

	return &Client{
		httpClient:   httpClient,
		Traces:       traces.NewClient(httpClient),
		Observations: observations.NewClient(httpClient),
		Scores:       scores.NewClient(httpClient),
		Datasets:     datasets.NewClient(httpClient),
		Sessions:     sessions.NewClient(httpClient),
		Models:       models.NewClient(httpClient),
		Prompts:      prompts.NewClient(httpClient),
		Comments:     comments.NewClient(httpClient),
		Media:        media.NewClient(httpClient),
		Metrics:      metrics.NewClient(httpClient),
		Annotations:  annotations.NewClient(httpClient),
		Ingestion:    ingestion.NewClient(httpClient),
	}
}

// Health checks the health of the API
func (c *Client) Health(ctx context.Context) (*types.HealthResponse, error) {
	var response types.HealthResponse
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, "/api/public/health", nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}

// GetProjects retrieves all projects
func (c *Client) GetProjects(ctx context.Context) (*types.Projects, error) {
	var response types.Projects
	if err := c.httpClient.DoRequest(ctx, http.MethodGet, "/api/public/projects", nil, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
