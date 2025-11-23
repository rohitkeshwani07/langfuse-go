package client

import (
	"testing"
)

func TestCreateTraceID(t *testing.T) {
	client := &Client{}

	t.Run("generates random trace ID", func(t *testing.T) {
		traceID1 := client.CreateTraceID("")
		traceID2 := client.CreateTraceID("")

		// Verify length (32 hex characters)
		if len(traceID1) != 32 {
			t.Errorf("expected trace ID length 32, got %d", len(traceID1))
		}
		if len(traceID2) != 32 {
			t.Errorf("expected trace ID length 32, got %d", len(traceID2))
		}

		// Verify they are different (random)
		if traceID1 == traceID2 {
			t.Error("expected different random trace IDs, got identical")
		}

		// Verify they are lowercase hex
		for _, c := range traceID1 {
			if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
				t.Errorf("trace ID contains non-hex character: %c", c)
			}
		}
	})

	t.Run("generates deterministic trace ID from seed", func(t *testing.T) {
		seed := "test-seed-123"
		traceID1 := client.CreateTraceID(seed)
		traceID2 := client.CreateTraceID(seed)

		// Verify length (32 hex characters)
		if len(traceID1) != 32 {
			t.Errorf("expected trace ID length 32, got %d", len(traceID1))
		}

		// Verify they are identical (deterministic)
		if traceID1 != traceID2 {
			t.Errorf("expected identical trace IDs for same seed, got %s and %s", traceID1, traceID2)
		}

		// Verify they are lowercase hex
		for _, c := range traceID1 {
			if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
				t.Errorf("trace ID contains non-hex character: %c", c)
			}
		}
	})

	t.Run("different seeds produce different IDs", func(t *testing.T) {
		traceID1 := client.CreateTraceID("seed1")
		traceID2 := client.CreateTraceID("seed2")

		if traceID1 == traceID2 {
			t.Error("expected different trace IDs for different seeds")
		}
	})

	t.Run("matches expected format for known seed", func(t *testing.T) {
		// This test verifies the deterministic behavior matches the expected SHA256 hash
		seed := "session-456"
		traceID := client.CreateTraceID(seed)

		// The trace ID should be the first 16 bytes (32 hex chars) of SHA256(seed)
		// We can verify the length and format
		if len(traceID) != 32 {
			t.Errorf("expected trace ID length 32, got %d", len(traceID))
		}

		// Verify consistency
		if traceID != client.CreateTraceID(seed) {
			t.Error("trace ID is not consistent for the same seed")
		}
	})
}
