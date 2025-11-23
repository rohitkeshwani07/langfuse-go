// Package langfuse provides a Go client for the Langfuse API.
package langfuse

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// CreateTraceID generates a unique trace ID for use with Langfuse.
//
// This function generates a unique trace ID for use with various Langfuse APIs.
// It can either generate a random ID or create a deterministic ID based on
// a seed string.
//
// Trace IDs must be 32 lowercase hexadecimal characters, representing 16 bytes.
// This function ensures the generated ID meets this requirement. If you need to
// correlate an external ID with a Langfuse trace ID, use the external ID as the
// seed to get a valid, deterministic Langfuse trace ID.
//
// Parameters:
//   - seed: Optional string to use as a seed for deterministic ID generation.
//     If provided, the same seed will always produce the same ID.
//     If empty, a random ID will be generated.
//
// Returns:
//   - A 32-character lowercase hexadecimal string representing the Langfuse trace ID.
//   - An error if random ID generation fails (extremely unlikely).
//
// Example:
//
//	// Generate a random trace ID
//	traceID, err := langfuse.CreateTraceID("")
//	if err != nil {
//	    // handle error
//	}
//
//	// Generate a deterministic ID based on a seed
//	sessionTraceID, _ := langfuse.CreateTraceID("session-456")
//
//	// Correlate an external ID with a Langfuse trace ID
//	externalID := "external-system-123456"
//	correlatedTraceID, _ := langfuse.CreateTraceID(externalID)
func CreateTraceID(seed string) (string, error) {
	if seed == "" {
		// Generate a random 16-byte trace ID
		traceIDBytes := make([]byte, 16)
		_, err := rand.Read(traceIDBytes)
		if err != nil {
			return "", fmt.Errorf("failed to generate random trace ID: %w", err)
		}
		return hex.EncodeToString(traceIDBytes), nil
	}

	// Generate deterministic ID based on seed
	hash := sha256.Sum256([]byte(seed))
	return hex.EncodeToString(hash[:16]), nil
}
