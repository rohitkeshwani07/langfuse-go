package models

import (
	"time"

	"github.com/langfuse/langfuse-go/types"
)

// Model represents a model configuration
type Model struct {
	ID              string                 `json:"id"`
	ModelName       string                 `json:"modelName"`
	MatchPattern    string                 `json:"matchPattern"`
	StartDate       *time.Time             `json:"startDate,omitempty"`
	InputPrice      *float64               `json:"inputPrice,omitempty"`
	OutputPrice     *float64               `json:"outputPrice,omitempty"`
	TotalPrice      *float64               `json:"totalPrice,omitempty"`
	Unit            *string  `json:"unit,omitempty"`
	TokenizerID     *string                `json:"tokenizerId,omitempty"`
	TokenizerConfig map[string]interface{} `json:"tokenizerConfig,omitempty"`
	CreatedAt       time.Time              `json:"createdAt"`
	UpdatedAt       time.Time              `json:"updatedAt"`
}

// CreateRequest represents the request body for creating a model
type CreateRequest struct {
	ModelName       string                 `json:"modelName"`
	MatchPattern    string                 `json:"matchPattern"`
	StartDate       *time.Time             `json:"startDate,omitempty"`
	InputPrice      *float64               `json:"inputPrice,omitempty"`
	OutputPrice     *float64               `json:"outputPrice,omitempty"`
	TotalPrice      *float64               `json:"totalPrice,omitempty"`
	Unit            *string  `json:"unit,omitempty"`
	TokenizerID     *string                `json:"tokenizerId,omitempty"`
	TokenizerConfig map[string]interface{} `json:"tokenizerConfig,omitempty"`
}

// ListResponse represents a paginated list of models
type ListResponse struct {
	Data []Model            `json:"data"`
	Meta types.MetaResponse `json:"meta"`
}
