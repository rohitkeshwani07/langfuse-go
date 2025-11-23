// Package client provides a unified client for the Langfuse API.
// It combines all domain-specific clients into a single interface.
package client

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/rohitkeshwani07/langfuse-go/annotations"
	"github.com/rohitkeshwani07/langfuse-go/comments"
	"github.com/rohitkeshwani07/langfuse-go/core"
	"github.com/rohitkeshwani07/langfuse-go/datasets"
	"github.com/rohitkeshwani07/langfuse-go/ingestion"
	"github.com/rohitkeshwani07/langfuse-go/media"
	"github.com/rohitkeshwani07/langfuse-go/metrics"
	"github.com/rohitkeshwani07/langfuse-go/models"
	"github.com/rohitkeshwani07/langfuse-go/observations"
	"github.com/rohitkeshwani07/langfuse-go/prompts"
	"github.com/rohitkeshwani07/langfuse-go/scores"
	"github.com/rohitkeshwani07/langfuse-go/sessions"
	"github.com/rohitkeshwani07/langfuse-go/traces"
	"github.com/rohitkeshwani07/langfuse-go/types"
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

// CreateTraceID generates a unique trace ID for use with Langfuse.
//
// This function generates a unique trace ID for use with various Langfuse APIs.
// It can either generate a random ID or create a deterministic ID based on
// a seed string.
//
// Trace IDs must be 32 lowercase hexadecimal characters, representing 16 bytes.
// This function ensures the generated ID meets this requirement. If you need to
// correlate an external ID with a Langfuse trace ID, use the external ID as the
// seed to get a valid, deterministic Langfuse trace ID.
//
// Parameters:
//   - seed: Optional string to use as a seed for deterministic ID generation.
//     If provided, the same seed will always produce the same ID.
//     If empty, a random ID will be generated.
//
// Returns:
//   - A 32-character lowercase hexadecimal string representing the Langfuse trace ID.
//   - An error if random ID generation fails (extremely unlikely).
//
// Example:
//
//	// Generate a random trace ID
//	traceID, err := client.CreateTraceID("")
//	if err != nil {
//	    // handle error
//	}
//
//	// Generate a deterministic ID based on a seed
//	sessionTraceID, _ := client.CreateTraceID("session-456")
//
//	// Correlate an external ID with a Langfuse trace ID
//	externalID := "external-system-123456"
//	correlatedTraceID, _ := client.CreateTraceID(externalID)
func (c *Client) CreateTraceID(seed string) (string, error) {
	if seed == "" {
		// Generate a random 16-byte trace ID
		traceIDBytes := make([]byte, 16)
		_, err := rand.Read(traceIDBytes)
		if err != nil {
			return "", fmt.Errorf("failed to generate random trace ID: %w", err)
		}
		return hex.EncodeToString(traceIDBytes), nil
	}

	// Generate deterministic ID based on seed
	hash := sha256.Sum256([]byte(seed))
	return hex.EncodeToString(hash[:16]), nil
}
