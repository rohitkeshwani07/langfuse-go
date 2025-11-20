// Package core provides the base HTTP client for the Langfuse API.
package core

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/bytedance/sonic"
)

const (
	// DefaultBaseURL is the default Langfuse API base URL
	DefaultBaseURL = "https://cloud.langfuse.com"
	// DefaultTimeout is the default HTTP client timeout
	DefaultTimeout = 30 * time.Second
)

// HTTPClient provides the base HTTP client functionality
type HTTPClient struct {
	BaseURL    string
	PublicKey  string
	SecretKey  string
	HTTPClient *http.Client
}

// Option is a functional option for configuring the HTTPClient
type Option func(*HTTPClient)

// WithBaseURL sets a custom base URL for the client
func WithBaseURL(baseURL string) Option {
	return func(c *HTTPClient) {
		c.BaseURL = baseURL
	}
}

// WithHTTPClient sets a custom HTTP client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *HTTPClient) {
		c.HTTPClient = httpClient
	}
}

// WithTimeout sets a custom timeout for the HTTP client
func WithTimeout(timeout time.Duration) Option {
	return func(c *HTTPClient) {
		c.HTTPClient.Timeout = timeout
	}
}

// NewHTTPClient creates a new base HTTP client
func NewHTTPClient(publicKey, secretKey string, opts ...Option) *HTTPClient {
	client := &HTTPClient{
		BaseURL:   DefaultBaseURL,
		PublicKey: publicKey,
		SecretKey: secretKey,
		HTTPClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}

// DoRequest performs an HTTP request with authentication
func (c *HTTPClient) DoRequest(ctx context.Context, method, path string, body interface{}, result interface{}) error {
	var reqBody io.Reader
	if body != nil {
		jsonData, err := sonic.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequestWithContext(ctx, method, c.BaseURL+path, reqBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// Add Basic Auth
	auth := base64.StdEncoding.EncodeToString([]byte(c.PublicKey + ":" + c.SecretKey))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to perform request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	if result != nil && resp.StatusCode != http.StatusNoContent {
		if err := sonic.ConfigDefault.NewDecoder(resp.Body).Decode(result); err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}
