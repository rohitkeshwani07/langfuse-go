package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/langfuse/langfuse-go/client"
)

func main() {
	// Get credentials from environment variables
	publicKey := os.Getenv("LANGFUSE_PUBLIC_KEY")
	secretKey := os.Getenv("LANGFUSE_SECRET_KEY")

	if publicKey == "" || secretKey == "" {
		log.Fatal("LANGFUSE_PUBLIC_KEY and LANGFUSE_SECRET_KEY must be set")
	}

	// Initialize the Langfuse client
	c := client.NewClient(publicKey, secretKey)

	ctx := context.Background()

	// Check API health
	health, err := c.Health(ctx)
	if err != nil {
		log.Fatalf("Failed to check health: %v", err)
	}
	fmt.Printf("API Status: %s\n", health.Status)

	// Create a trace
	traceID := "example-trace-" + fmt.Sprint(time.Now().Unix())
	err = c.CreateTrace(ctx, &client.CreateTraceBody{
		ID:     client.String(traceID),
		Name:   client.String("example-trace"),
		UserID: client.String("user-123"),
		Input: map[string]interface{}{
			"query": "What is Langfuse?",
		},
		Metadata: map[string]interface{}{
			"example": true,
		},
	})
	if err != nil {
		log.Fatalf("Failed to create trace: %v", err)
	}
	fmt.Printf("Created trace: %s\n", traceID)

	// Create a generation
	generationID := "example-gen-" + fmt.Sprint(time.Now().Unix())
	err = c.CreateGeneration(ctx, &client.CreateGenerationBody{
		ID:      client.String(generationID),
		TraceID: client.String(traceID),
		Name:    client.String("gpt-4-completion"),
		Model:   client.String("gpt-4"),
		Input: []map[string]interface{}{
			{
				"role":    "user",
				"content": "What is Langfuse?",
			},
		},
		Output: map[string]interface{}{
			"role":    "assistant",
			"content": "Langfuse is an open source LLM engineering platform.",
		},
		Usage: &client.Usage{
			PromptTokens:     client.Int(15),
			CompletionTokens: client.Int(25),
			TotalTokens:      client.Int(40),
		},
		ModelParameters: map[string]interface{}{
			"temperature": 0.7,
			"max_tokens":  500,
		},
	})
	if err != nil {
		log.Fatalf("Failed to create generation: %v", err)
	}
	fmt.Printf("Created generation: %s\n", generationID)

	// Create a score
	scoreResp, err := c.CreateScore(ctx, &client.CreateScoreRequest{
		Name:    "accuracy",
		Value:   0.95,
		TraceID: traceID,
		Comment: client.String("High quality response"),
	})
	if err != nil {
		log.Fatalf("Failed to create score: %v", err)
	}
	fmt.Printf("Created score: %s\n", scoreResp.ID)

	// Wait a bit for data to be available
	time.Sleep(2 * time.Second)

	// Get the trace
	trace, err := c.GetTrace(ctx, traceID)
	if err != nil {
		log.Fatalf("Failed to get trace: %v", err)
	}
	fmt.Printf("Retrieved trace: %s (Name: %s)\n", trace.ID, *trace.Name)

	// Get scores for the trace
	scores, err := c.GetScores(ctx, &client.GetScoresParams{
		TraceID: client.String(traceID),
		Limit:   client.Int(10),
	})
	if err != nil {
		log.Fatalf("Failed to get scores: %v", err)
	}
	fmt.Printf("Found %d scores for trace\n", len(scores.Data))
	for _, score := range scores.Data {
		fmt.Printf("  - %s: %v\n", score.Name, score.Value)
	}

	fmt.Println("\nExample completed successfully!")
}
