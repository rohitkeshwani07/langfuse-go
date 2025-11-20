package traces

import (
	"time"

	"github.com/langfuse/langfuse-go/observations"
)

// Trace represents a trace in Langfuse
type Trace struct {
	ID           string                      `json:"id"`
	Name         *string                     `json:"name,omitempty"`
	UserID       *string                     `json:"userId,omitempty"`
	SessionID    *string                     `json:"sessionId,omitempty"`
	Release      *string                     `json:"release,omitempty"`
	Version      *string                     `json:"version,omitempty"`
	Metadata     map[string]interface{}      `json:"metadata,omitempty"`
	Tags         []string                    `json:"tags,omitempty"`
	Input        interface{}                 `json:"input,omitempty"`
	Output       interface{}                 `json:"output,omitempty"`
	Timestamp    time.Time                   `json:"timestamp"`
	CreatedAt    time.Time                   `json:"createdAt"`
	UpdatedAt    time.Time                   `json:"updatedAt"`
	Public       *bool                       `json:"public,omitempty"`
	ProjectID    string                      `json:"projectId"`
	Observations []*observations.Observation `json:"observations,omitempty"`
}

// CreateTraceRequest represents the request body for creating a trace
type CreateTraceRequest struct {
	ID        *string                `json:"id,omitempty"`
	Name      *string                `json:"name,omitempty"`
	UserID    *string                `json:"userId,omitempty"`
	SessionID *string                `json:"sessionId,omitempty"`
	Release   *string                `json:"release,omitempty"`
	Version   *string                `json:"version,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Tags      []string               `json:"tags,omitempty"`
	Input     interface{}            `json:"input,omitempty"`
	Output    interface{}            `json:"output,omitempty"`
	Timestamp *time.Time             `json:"timestamp,omitempty"`
	Public    *bool                  `json:"public,omitempty"`
}

// UpdateTraceRequest represents the request body for updating a trace
type UpdateTraceRequest struct {
	Name      *string                `json:"name,omitempty"`
	UserID    *string                `json:"userId,omitempty"`
	SessionID *string                `json:"sessionId,omitempty"`
	Release   *string                `json:"release,omitempty"`
	Version   *string                `json:"version,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Tags      []string               `json:"tags,omitempty"`
	Input     interface{}            `json:"input,omitempty"`
	Output    interface{}            `json:"output,omitempty"`
	Public    *bool                  `json:"public,omitempty"`
}
