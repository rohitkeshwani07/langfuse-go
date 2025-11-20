package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/langfuse/langfuse-go/client"
	"github.com/langfuse/langfuse-go/traces"
)

func main() {
	// Get credentials from environment variables
	publicKey := os.Getenv("LANGFUSE_PUBLIC_KEY")
	secretKey := os.Getenv("LANGFUSE_SECRET_KEY")

	if publicKey == "" || secretKey == "" {
		log.Fatal("LANGFUSE_PUBLIC_KEY and LANGFUSE_SECRET_KEY must be set")
	}

	// Initialize the Langfuse client
	c := client.New(publicKey, secretKey)

	ctx := context.Background()

	// Test trace ID and observation ID
	traceID := "37ae885d46abc96bde952bcc387304b7"
	observationID := "db5343cd53694930"

	fmt.Println("=== GetTree Example ===")
	fmt.Printf("Fetching trace tree for: %s\n\n", traceID)

	// Get trace tree with compact observations
	tree, err := c.Traces.GetTree(ctx, traceID)
	if err != nil {
		log.Fatalf("Failed to get trace tree: %v", err)
	}

	// Pretty print the tree
	treeJSON, err := json.MarshalIndent(tree, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal tree: %v", err)
	}
	fmt.Println("Trace Tree (CompactTrace with nested observations):")
	fmt.Println(string(treeJSON))

	fmt.Printf("\n\nTrace Summary:\n")
	fmt.Printf("  ID: %s\n", tree.ID)
	fmt.Printf("  Name: %s\n", tree.Name)
	fmt.Printf("  Total Cost: %.4f\n", tree.TotalCost)
	fmt.Printf("  Latency: %.2fms\n", tree.Latency)
	fmt.Printf("  Root Observation Nodes: %d\n", len(tree.RootNode))

	// Print observation tree structure
	fmt.Printf("\n\nObservation Tree Structure:\n")
	printObservationTree(tree.RootNode, 0)

	fmt.Println("\n\n=== Observation.Get Example ===")
	fmt.Printf("Fetching observation: %s\n\n", observationID)

	// Get single observation with full details (includes input/output)
	observation, err := c.Observations.Get(ctx, observationID)
	if err != nil {
		log.Fatalf("Failed to get observation: %v", err)
	}

	// Pretty print the observation
	obsJSON, err := json.MarshalIndent(observation, "", "  ")
	if err != nil {
		log.Fatalf("Failed to marshal observation: %v", err)
	}
	fmt.Println("Observation Details (includes input/output):")
	fmt.Println(string(obsJSON))

	fmt.Println("\n\nExample completed successfully!")
}

// printObservationTree recursively prints the observation tree structure
func printObservationTree(observations []*traces.ObservationNode, depth int) {
	indent := ""
	for i := 0; i < depth; i++ {
		indent += "  "
	}

	for _, obs := range observations {
		fmt.Printf("%s- %s (%s) [%s]\n", indent, obs.Name, obs.Type, obs.ID)
		if len(obs.Children) > 0 {
			printObservationTree(obs.Children, depth+1)
		}
	}
}
