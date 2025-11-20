package annotations

import (
	"time"

	"github.com/langfuse/langfuse-go/types"
)

// Queue represents an annotation queue
type Queue struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	Description    *string   `json:"description,omitempty"`
	ScoreConfigIDs []string  `json:"scoreConfigIds"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// QueueItem represents an item in an annotation queue
type QueueItem struct {
	ID          string                          `json:"id"`
	QueueID     string                          `json:"queueId"`
	ObjectID    string                          `json:"objectId"`
	ObjectType  string `json:"objectType"`
	Status      string     `json:"status"`
	CompletedAt *time.Time                      `json:"completedAt,omitempty"`
	CreatedAt   time.Time                       `json:"createdAt"`
	UpdatedAt   time.Time                       `json:"updatedAt"`
}

// CreateQueueRequest represents the request body for creating an annotation queue
type CreateQueueRequest struct {
	Name           string   `json:"name"`
	Description    *string  `json:"description,omitempty"`
	ScoreConfigIDs []string `json:"scoreConfigIds"`
}

// CreateQueueItemRequest represents the request body for creating an annotation queue item
type CreateQueueItemRequest struct {
	ObjectID   string                          `json:"objectId"`
	ObjectType string `json:"objectType"`
	Status     *string    `json:"status,omitempty"`
}

// UpdateQueueItemRequest represents the request body for updating an annotation queue item
type UpdateQueueItemRequest struct {
	Status *string `json:"status,omitempty"`
}

// DeleteQueueItemResponse represents the response for deleting an annotation queue item
type DeleteQueueItemResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// CreateAssignmentResponse represents the response for creating a queue assignment
type CreateAssignmentResponse struct {
	UserID    string `json:"userId"`
	QueueID   string `json:"queueId"`
	ProjectID string `json:"projectId"`
}

// AssignmentRequest represents the request body for creating a queue assignment
type AssignmentRequest struct {
	UserID string `json:"userId"`
}

// DeleteAssignmentResponse represents the response for deleting a queue assignment
type DeleteAssignmentResponse struct {
	Success bool `json:"success"`
}

// ListQueuesResponse represents a paginated list of annotation queues
type ListQueuesResponse struct {
	Data []Queue            `json:"data"`
	Meta types.MetaResponse `json:"meta"`
}

// ListQueueItemsParams represents parameters for listing queue items
type ListQueueItemsParams struct {
	Page   *int
	Limit  *int
	Status *string
}

// ListQueueItemsResponse represents a paginated list of annotation queue items
type ListQueueItemsResponse struct {
	Data []QueueItem        `json:"data"`
	Meta types.MetaResponse `json:"meta"`
}
