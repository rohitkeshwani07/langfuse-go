package observations

import (
	"time"

	"github.com/langfuse/langfuse-go/types"
)

// Observation represents an observation (event, span, or generation)
type Observation struct {
	ID                  string                  `json:"id"`
	TraceID             *string                 `json:"traceId,omitempty"`
	ParentObservationID *string                 `json:"parentObservationId,omitempty"`
	Type                types.ObservationType   `json:"type"`
	Name                *string                 `json:"name,omitempty"`
	StartTime           time.Time               `json:"startTime"`
	EndTime             *time.Time              `json:"endTime,omitempty"`
	CompletionStartTime *time.Time              `json:"completionStartTime,omitempty"`
	Model               *string                 `json:"model,omitempty"`
	ModelParameters     map[string]interface{}  `json:"modelParameters,omitempty"`
	Input               interface{}             `json:"input,omitempty"`
	Output              interface{}             `json:"output,omitempty"`
	Metadata            map[string]interface{}  `json:"metadata,omitempty"`
	Level               *types.ObservationLevel `json:"level,omitempty"`
	StatusMessage       *string                 `json:"statusMessage,omitempty"`
	Version             *string                 `json:"version,omitempty"`
	Usage               *types.Usage            `json:"usage,omitempty"`
	PromptTokens        *int                    `json:"promptTokens,omitempty"`
	CompletionTokens    *int                    `json:"completionTokens,omitempty"`
	TotalTokens         *int                    `json:"totalTokens,omitempty"`
	CreatedAt           time.Time               `json:"createdAt"`
	UpdatedAt           time.Time               `json:"updatedAt"`
}

// CreateEventRequest represents the request body for creating an event
type CreateEventRequest struct {
	ID                  *string                 `json:"id,omitempty"`
	TraceID             *string                 `json:"traceId,omitempty"`
	ParentObservationID *string                 `json:"parentObservationId,omitempty"`
	Name                *string                 `json:"name,omitempty"`
	StartTime           *time.Time              `json:"startTime,omitempty"`
	Metadata            map[string]interface{}  `json:"metadata,omitempty"`
	Input               interface{}             `json:"input,omitempty"`
	Output              interface{}             `json:"output,omitempty"`
	Level               *types.ObservationLevel `json:"level,omitempty"`
	StatusMessage       *string                 `json:"statusMessage,omitempty"`
	Version             *string                 `json:"version,omitempty"`
}

// UpdateEventRequest represents the request body for updating an event
type UpdateEventRequest struct {
	Metadata      map[string]interface{}  `json:"metadata,omitempty"`
	Input         interface{}             `json:"input,omitempty"`
	Output        interface{}             `json:"output,omitempty"`
	Level         *types.ObservationLevel `json:"level,omitempty"`
	StatusMessage *string                 `json:"statusMessage,omitempty"`
}

// CreateSpanRequest represents the request body for creating a span
type CreateSpanRequest struct {
	ID                  *string                 `json:"id,omitempty"`
	TraceID             *string                 `json:"traceId,omitempty"`
	ParentObservationID *string                 `json:"parentObservationId,omitempty"`
	Name                *string                 `json:"name,omitempty"`
	StartTime           *time.Time              `json:"startTime,omitempty"`
	EndTime             *time.Time              `json:"endTime,omitempty"`
	Metadata            map[string]interface{}  `json:"metadata,omitempty"`
	Input               interface{}             `json:"input,omitempty"`
	Output              interface{}             `json:"output,omitempty"`
	Level               *types.ObservationLevel `json:"level,omitempty"`
	StatusMessage       *string                 `json:"statusMessage,omitempty"`
	Version             *string                 `json:"version,omitempty"`
}

// UpdateSpanRequest represents the request body for updating a span
type UpdateSpanRequest struct {
	EndTime       *time.Time              `json:"endTime,omitempty"`
	Metadata      map[string]interface{}  `json:"metadata,omitempty"`
	Input         interface{}             `json:"input,omitempty"`
	Output        interface{}             `json:"output,omitempty"`
	Level         *types.ObservationLevel `json:"level,omitempty"`
	StatusMessage *string                 `json:"statusMessage,omitempty"`
}

// CreateGenerationRequest represents the request body for creating a generation
type CreateGenerationRequest struct {
	ID                  *string                 `json:"id,omitempty"`
	TraceID             *string                 `json:"traceId,omitempty"`
	ParentObservationID *string                 `json:"parentObservationId,omitempty"`
	Name                *string                 `json:"name,omitempty"`
	StartTime           *time.Time              `json:"startTime,omitempty"`
	EndTime             *time.Time              `json:"endTime,omitempty"`
	CompletionStartTime *time.Time              `json:"completionStartTime,omitempty"`
	Model               *string                 `json:"model,omitempty"`
	ModelParameters     map[string]interface{}  `json:"modelParameters,omitempty"`
	Metadata            map[string]interface{}  `json:"metadata,omitempty"`
	Input               interface{}             `json:"input,omitempty"`
	Output              interface{}             `json:"output,omitempty"`
	Usage               *types.Usage            `json:"usage,omitempty"`
	PromptName          *string                 `json:"promptName,omitempty"`
	PromptVersion       *int                    `json:"promptVersion,omitempty"`
	Level               *types.ObservationLevel `json:"level,omitempty"`
	StatusMessage       *string                 `json:"statusMessage,omitempty"`
	Version             *string                 `json:"version,omitempty"`
}

// UpdateGenerationRequest represents the request body for updating a generation
type UpdateGenerationRequest struct {
	EndTime             *time.Time              `json:"endTime,omitempty"`
	CompletionStartTime *time.Time              `json:"completionStartTime,omitempty"`
	Metadata            map[string]interface{}  `json:"metadata,omitempty"`
	Input               interface{}             `json:"input,omitempty"`
	Output              interface{}             `json:"output,omitempty"`
	Usage               *types.Usage            `json:"usage,omitempty"`
	Level               *types.ObservationLevel `json:"level,omitempty"`
	StatusMessage       *string                 `json:"statusMessage,omitempty"`
}

// ListParams represents query parameters for listing observations
type ListParams struct {
	Page    *int
	Limit   *int
	Name    *string
	UserID  *string
	Type    *types.ObservationType
	TraceID *string
}

// ListResponse represents a paginated list of observations
type ListResponse struct {
	Data []Observation      `json:"data"`
	Meta types.MetaResponse `json:"meta"`
}
