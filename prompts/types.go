package prompts

import "time"

// ChatMessage represents a chat message in a prompt
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Chat represents a chat-based prompt
type Chat struct {
	Name      string        `json:"name"`
	Version   int           `json:"version"`
	Config    interface{}   `json:"config,omitempty"`
	Prompt    []ChatMessage `json:"prompt"`
	Tags      []string      `json:"tags,omitempty"`
	Labels    []string      `json:"labels,omitempty"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
	CreatedBy string        `json:"createdBy"`
	Type      string        `json:"type"`
}

// Text represents a text-based prompt
type Text struct {
	Name      string      `json:"name"`
	Version   int         `json:"version"`
	Config    interface{} `json:"config,omitempty"`
	Prompt    string      `json:"prompt"`
	Tags      []string    `json:"tags,omitempty"`
	Labels    []string    `json:"labels,omitempty"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	CreatedBy string      `json:"createdBy"`
	Type      string      `json:"type"`
}

// CreateChatRequest represents the request body for creating a chat prompt
type CreateChatRequest struct {
	Name   string        `json:"name"`
	Prompt []ChatMessage `json:"prompt"`
	Config interface{}   `json:"config,omitempty"`
	Tags   []string      `json:"tags,omitempty"`
	Labels []string      `json:"labels,omitempty"`
}

// CreateTextRequest represents the request body for creating a text prompt
type CreateTextRequest struct {
	Name   string      `json:"name"`
	Prompt string      `json:"prompt"`
	Config interface{} `json:"config,omitempty"`
	Tags   []string    `json:"tags,omitempty"`
	Labels []string    `json:"labels,omitempty"`
}

// GetParams represents parameters for getting a prompt
type GetParams struct {
	Version *int
	Label   *string
}

// ListParams represents parameters for listing prompts
type ListParams struct {
	Page  *int
	Limit *int
	Name  *string
	Label *string
	Tag   *string
}

// Meta represents prompt metadata
type Meta struct {
	Name      string    `json:"name"`
	Version   int       `json:"version"`
	Type      string    `json:"type"`
	Tags      []string  `json:"tags"`
	Labels    []string  `json:"labels"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// ListResponse represents a list of prompt metadata
type ListResponse struct {
	Data []Meta `json:"data"`
	Meta struct {
		Page       int `json:"page"`
		Limit      int `json:"limit"`
		TotalItems int `json:"totalItems"`
		TotalPages int `json:"totalPages"`
	} `json:"meta"`
}
