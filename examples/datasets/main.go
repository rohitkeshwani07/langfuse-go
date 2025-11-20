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

	// Create a dataset
	datasetName := "example-dataset-" + fmt.Sprint(time.Now().Unix())
	dataset, err := c.CreateDataset(ctx, &client.CreateDatasetRequest{
		Name:        datasetName,
		Description: client.String("Example evaluation dataset"),
		Metadata: map[string]interface{}{
			"created_by": "example",
		},
	})
	if err != nil {
		log.Fatalf("Failed to create dataset: %v", err)
	}
	fmt.Printf("Created dataset: %s (ID: %s)\n", dataset.Name, dataset.ID)

	// Create dataset items
	questions := []struct {
		input    string
		expected string
	}{
		{"What is the capital of France?", "Paris"},
		{"What is 2 + 2?", "4"},
		{"What color is the sky?", "Blue"},
	}

	for i, q := range questions {
		item, err := c.CreateDatasetItem(ctx, &client.CreateDatasetItemRequest{
			DatasetName: client.String(datasetName),
			Input: map[string]interface{}{
				"question": q.input,
			},
			ExpectedOutput: map[string]interface{}{
				"answer": q.expected,
			},
			Metadata: map[string]interface{}{
				"index": i,
			},
		})
		if err != nil {
			log.Fatalf("Failed to create dataset item: %v", err)
		}
		fmt.Printf("Created dataset item: %s\n", item.ID)
	}

	// Create a dataset run
	runName := "example-run-" + fmt.Sprint(time.Now().Unix())
	run, err := c.CreateDatasetRun(ctx, &client.CreateDatasetRunRequest{
		Name:        runName,
		DatasetID:   dataset.ID,
		Description: client.String("Example evaluation run"),
		Metadata: map[string]interface{}{
			"model": "gpt-4",
		},
	})
	if err != nil {
		log.Fatalf("Failed to create dataset run: %v", err)
	}
	fmt.Printf("Created dataset run: %s (ID: %s)\n", run.Name, run.ID)

	// List dataset items
	items, err := c.GetDatasetItems(ctx, datasetName, &client.PaginationParams{
		Page:  client.Int(1),
		Limit: client.Int(10),
	})
	if err != nil {
		log.Fatalf("Failed to get dataset items: %v", err)
	}

	fmt.Printf("\nDataset items (%d total):\n", items.Meta.TotalItems)
	for _, item := range items.Data {
		fmt.Printf("  - Item %s: %v\n", item.ID, item.Input)
	}

	// List datasets
	datasets, err := c.GetDatasets(ctx, &client.PaginationParams{
		Page:  client.Int(1),
		Limit: client.Int(10),
	})
	if err != nil {
		log.Fatalf("Failed to get datasets: %v", err)
	}

	fmt.Printf("\nAll datasets (%d total):\n", datasets.Meta.TotalItems)
	for _, ds := range datasets.Data {
		fmt.Printf("  - %s (ID: %s)\n", ds.Name, ds.ID)
	}

	fmt.Println("\nDataset example completed successfully!")
}
