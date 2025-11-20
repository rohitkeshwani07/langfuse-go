package scores

import (
	"time"

	"github.com/langfuse/langfuse-go/types"
)

// Score represents a score for a trace or observation
type Score struct {
	ID            string              `json:"id"`
	Name          string              `json:"name"`
	Value         interface{}         `json:"value"`
	DataType      types.ScoreDataType `json:"dataType"`
	Source        types.ScoreSource   `json:"source"`
	Comment       *string             `json:"comment,omitempty"`
	TraceID       string              `json:"traceId"`
	ObservationID *string             `json:"observationId,omitempty"`
	ConfigID      *string             `json:"configId,omitempty"`
	CreatedAt     time.Time           `json:"createdAt"`
	UpdatedAt     time.Time           `json:"updatedAt"`
	AuthorUserID  *string             `json:"authorUserId,omitempty"`
}

// CreateRequest represents the request body for creating a score
type CreateRequest struct {
	ID            *string              `json:"id,omitempty"`
	Name          string               `json:"name"`
	Value         interface{}          `json:"value"`
	DataType      *types.ScoreDataType `json:"dataType,omitempty"`
	Comment       *string              `json:"comment,omitempty"`
	TraceID       string               `json:"traceId"`
	ObservationID *string              `json:"observationId,omitempty"`
	ConfigID      *string              `json:"configId,omitempty"`
}

// CreateResponse represents the response for creating a score
type CreateResponse struct {
	ID string `json:"id"`
}

// ListParams represents query parameters for listing scores
type ListParams struct {
	Page    *int
	Limit   *int
	TraceID *string
	UserID  *string
}

// ListResponse represents a paginated list of scores
type ListResponse struct {
	Data []Score            `json:"data"`
	Meta types.MetaResponse `json:"meta"`
}

// Config represents a score configuration
type Config struct {
	ID          string              `json:"id"`
	Name        string              `json:"name"`
	DataType    types.ScoreDataType `json:"dataType"`
	Description *string             `json:"description,omitempty"`
	MinValue    *float64            `json:"minValue,omitempty"`
	MaxValue    *float64            `json:"maxValue,omitempty"`
	Categories  []string            `json:"categories,omitempty"`
	CreatedAt   time.Time           `json:"createdAt"`
	UpdatedAt   time.Time           `json:"updatedAt"`
}
