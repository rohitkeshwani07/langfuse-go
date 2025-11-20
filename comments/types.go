package comments

import (
	"time"

	"github.com/langfuse/langfuse-go/types"
)

// Comment represents a comment on a trace or observation
type Comment struct {
	ID           string                  `json:"id"`
	Content      string                  `json:"content"`
	ObjectType   string `json:"objectType"`
	ObjectID     string                  `json:"objectId"`
	ProjectID    string                  `json:"projectId"`
	AuthorUserID string                  `json:"authorUserId"`
	CreatedAt    time.Time               `json:"createdAt"`
	UpdatedAt    time.Time               `json:"updatedAt"`
}

// CreateRequest represents the request body for creating a comment
type CreateRequest struct {
	Content    string                  `json:"content"`
	ObjectType string `json:"objectType"`
	ObjectID   string                  `json:"objectId"`
}

// CreateResponse represents the response for creating a comment
type CreateResponse struct {
	ID string `json:"id"`
}

// ListParams represents query parameters for listing comments
type ListParams struct {
	Page       *int
	Limit      *int
	ObjectType *string
	ObjectID   *string
}

// ListResponse represents a paginated list of comments
type ListResponse struct {
	Data []Comment          `json:"data"`
	Meta types.MetaResponse `json:"meta"`
}
