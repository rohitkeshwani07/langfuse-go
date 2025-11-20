package sessions

import (
	"time"

	"github.com/rohitkeshwani07/langfuse-go/traces"
	"github.com/rohitkeshwani07/langfuse-go/types"
)

// Session represents a session
type Session struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	ProjectID string    `json:"projectId"`
}

// WithTraces represents a session with its traces
type WithTraces struct {
	Session
	Traces []traces.Trace `json:"traces"`
}

// ListResponse represents a paginated list of sessions
type ListResponse struct {
	Data []Session          `json:"data"`
	Meta types.MetaResponse `json:"meta"`
}
