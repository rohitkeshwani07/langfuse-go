package types

import "time"

// MetaResponse represents common metadata for pagination
type MetaResponse struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}

// PaginationParams represents common pagination parameters
type PaginationParams struct {
	Page  *int
	Limit *int
}

// HealthResponse represents the health status of the API
type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version,omitempty"`
}

// Project represents a Langfuse project
type Project struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Projects represents a list of projects
type Projects struct {
	Data []Project `json:"data"`
}

// Usage represents token usage information
type Usage struct {
	PromptTokens     *int     `json:"promptTokens,omitempty"`
	CompletionTokens *int     `json:"completionTokens,omitempty"`
	TotalTokens      *int     `json:"totalTokens,omitempty"`
	Unit             *string  `json:"unit,omitempty"`
	Input            *int     `json:"input,omitempty"`
	Output           *int     `json:"output,omitempty"`
	Total            *int     `json:"total,omitempty"`
	InputCost        *float64 `json:"inputCost,omitempty"`
	OutputCost       *float64 `json:"outputCost,omitempty"`
	TotalCost        *float64 `json:"totalCost,omitempty"`
}

// OpenAIUsage represents OpenAI-specific usage information
type OpenAIUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

// Sort represents sorting parameters
type Sort struct {
	Column string `json:"column"`
	Order  string `json:"order"`
}
