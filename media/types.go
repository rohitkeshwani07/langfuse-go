package media

import (
	"time"
)

// Response represents media metadata
type Response struct {
	MediaID       string                 `json:"mediaId"`
	ContentType   string `json:"contentType"`
	ContentLength int                    `json:"contentLength"`
	UploadedAt    time.Time              `json:"uploadedAt"`
}

// UploadURLRequest represents the request for getting a media upload URL
type UploadURLRequest struct {
	ContentType   string `json:"contentType"`
	ContentLength int                    `json:"contentLength"`
	TraceID       *string                `json:"traceId,omitempty"`
	ObservationID *string                `json:"observationId,omitempty"`
	Field         *string                `json:"field,omitempty"`
}

// UploadURLResponse represents the response for getting a media upload URL
type UploadURLResponse struct {
	MediaID   string `json:"mediaId"`
	UploadURL string `json:"uploadUrl"`
}

// PatchRequest represents the request body for patching media
type PatchRequest struct {
	TraceID       *string `json:"traceId,omitempty"`
	ObservationID *string `json:"observationId,omitempty"`
	Field         *string `json:"field,omitempty"`
}
