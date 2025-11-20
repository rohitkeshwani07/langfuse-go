package datasets

import (
	"time"

	"github.com/rohitkeshwani07/langfuse-go/types"
)

// Dataset represents a dataset
type Dataset struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	Status      *string   `json:"status,omitempty"`
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
	ProjectID   string                 `json:"projectId"`
}

// CreateRequest represents the request body for creating a dataset
type CreateRequest struct {
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// ListResponse represents a paginated list of datasets
type ListResponse struct {
	Data []Dataset          `json:"data"`
	Meta types.MetaResponse `json:"meta"`
}

// Item represents an item in a dataset
type Item struct {
	ID                  string                 `json:"id"`
	DatasetID           string                 `json:"datasetId"`
	DatasetName         string                 `json:"datasetName"`
	Input               interface{}            `json:"input"`
	ExpectedOutput      interface{}            `json:"expectedOutput,omitempty"`
	Metadata            map[string]interface{} `json:"metadata,omitempty"`
	SourceTraceID       *string                `json:"sourceTraceId,omitempty"`
	SourceObservationID *string                `json:"sourceObservationId,omitempty"`
	Status              string                 `json:"status"`
	CreatedAt           time.Time              `json:"createdAt"`
	UpdatedAt           time.Time              `json:"updatedAt"`
}

// CreateItemRequest represents the request body for creating a dataset item
type CreateItemRequest struct {
	DatasetName         *string                `json:"datasetName,omitempty"`
	Input               interface{}            `json:"input"`
	ExpectedOutput      interface{}            `json:"expectedOutput,omitempty"`
	Metadata            map[string]interface{} `json:"metadata,omitempty"`
	SourceTraceID       *string                `json:"sourceTraceId,omitempty"`
	SourceObservationID *string                `json:"sourceObservationId,omitempty"`
}

// ItemListResponse represents a paginated list of dataset items
type ItemListResponse struct {
	Data []Item             `json:"data"`
	Meta types.MetaResponse `json:"meta"`
}

// Run represents a run of a dataset
type Run struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	DatasetID   string                 `json:"datasetId"`
	DatasetName string                 `json:"datasetName"`
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
}

// CreateRunRequest represents the request body for creating a dataset run
type CreateRunRequest struct {
	Name        string                 `json:"name"`
	Description *string                `json:"description,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	DatasetID   string                 `json:"datasetId"`
}

// RunListResponse represents a paginated list of dataset runs
type RunListResponse struct {
	Data []Run              `json:"data"`
	Meta types.MetaResponse `json:"meta"`
}

// RunItem represents an item in a dataset run
type RunItem struct {
	ID             string                 `json:"id"`
	DatasetRunID   string                 `json:"datasetRunId"`
	DatasetRunName string                 `json:"datasetRunName"`
	DatasetItemID  string                 `json:"datasetItemId"`
	TraceID        string                 `json:"traceId"`
	ObservationID  *string                `json:"observationId,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt      time.Time              `json:"createdAt"`
	UpdatedAt      time.Time              `json:"updatedAt"`
}

// CreateRunItemRequest represents the request body for creating a dataset run item
type CreateRunItemRequest struct {
	RunName        string                 `json:"runName"`
	RunDescription *string                `json:"runDescription,omitempty"`
	RunMetadata    map[string]interface{} `json:"runMetadata,omitempty"`
	DatasetItemID  string                 `json:"datasetItemId"`
	TraceID        string                 `json:"traceId"`
	ObservationID  *string                `json:"observationId,omitempty"`
	Metadata       map[string]interface{} `json:"metadata,omitempty"`
}
