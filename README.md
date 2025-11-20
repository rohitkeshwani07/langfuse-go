# Langfuse Go Client

[![Go Reference](https://pkg.go.dev/badge/github.com/langfuse/langfuse-go.svg)](https://pkg.go.dev/github.com/langfuse/langfuse-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/langfuse/langfuse-go)](https://goreportcard.com/report/github.com/langfuse/langfuse-go)

Official Go client library for [Langfuse](https://langfuse.com) - the open source LLM engineering platform.

## Features

- 🎯 **Modular Architecture**: Organized into domain-specific packages for better code organization
- 🔒 **Type-Safe**: Full type safety with comprehensive struct definitions
- 🚀 **Complete API Coverage**: All Langfuse API endpoints supported
- 📦 **Easy to Use**: Simple, intuitive API with sensible defaults
- 🔄 **Context Support**: Full support for context-based cancellation and timeouts
- 🛠️ **Flexible Configuration**: Customizable HTTP client, base URL, and timeouts

## Installation

```bash
go get github.com/langfuse/langfuse-go
```

## Quick Start

```go
package main

import (
	"context"
	"log"

	"github.com/langfuse/langfuse-go/client"
	"github.com/langfuse/langfuse-go/traces"
	"github.com/langfuse/langfuse-go/types"
)

func main() {
	// Initialize the client
	c := client.New(
		"your-public-key",
		"your-secret-key",
	)

	// Create a trace
	ctx := context.Background()
	err := c.Traces.Create(ctx, &traces.CreateTraceRequest{
		Name:   types.String("my-trace"),
		UserID: types.String("user-123"),
		Input:  map[string]interface{}{"query": "What is Langfuse?"},
	})
	if err != nil {
		log.Fatal(err)
	}
}
```

## Package Structure

The library is organized into modular packages:

- **`client`** - Main client combining all sub-clients
- **`core`** - Base HTTP client and shared utilities
- **`types`** - Shared types, enums, and helper functions
- **`traces`** - Trace management
- **`observations`** - Events, spans, and generations
- **`scores`** - Score operations
- **`datasets`** - Dataset management
- **`sessions`** - Session operations
- **`models`** - Model configuration
- **`prompts`** - Prompt management
- **`comments`** - Comment operations
- **`media`** - Media handling
- **`metrics`** - Metrics and analytics
- **`annotations`** - Annotation queues
- **`ingestion`** - Batch ingestion

## Configuration

### Custom Base URL

For self-hosted instances:

```go
import "github.com/langfuse/langfuse-go/core"

c := client.New(
	"your-public-key",
	"your-secret-key",
	core.WithBaseURL("https://your-instance.com"),
)
```

### Custom HTTP Client

```go
import (
	"net/http"
	"time"
)

httpClient := &http.Client{
	Timeout: 60 * time.Second,
}

c := client.New(
	"your-public-key",
	"your-secret-key",
	core.WithHTTPClient(httpClient),
)
```

### Custom Timeout

```go
c := client.New(
	"your-public-key",
	"your-secret-key",
	core.WithTimeout(60 * time.Second),
)
```

## Usage Examples

### Traces

```go
import (
	"github.com/langfuse/langfuse-go/traces"
	"github.com/langfuse/langfuse-go/types"
)

// Create a trace
err := c.Traces.Create(ctx, &traces.CreateTraceRequest{
	ID:        types.String("trace-123"),
	Name:      types.String("my-llm-app"),
	UserID:    types.String("user-456"),
	SessionID: types.String("session-789"),
	Metadata: map[string]interface{}{
		"environment": "production",
	},
	Tags: []string{"production", "important"},
})

// Get a trace
trace, err := c.Traces.Get(ctx, "trace-123")

// Update a trace
err = c.Traces.Update(ctx, "trace-123", &traces.UpdateTraceRequest{
	Name:   types.String("updated-name"),
	Public: types.Bool(true),
})
```

### Observations

```go
import (
	"github.com/langfuse/langfuse-go/observations"
	"github.com/langfuse/langfuse-go/types"
)

// Create a generation
err := c.Observations.CreateGeneration(ctx, &observations.CreateGenerationRequest{
	ID:      types.String("gen-123"),
	TraceID: types.String("trace-123"),
	Name:    types.String("openai-completion"),
	Model:   types.String("gpt-4"),
	Input: []map[string]interface{}{
		{"role": "user", "content": "What is Langfuse?"},
	},
	Usage: &types.Usage{
		PromptTokens:     types.Int(10),
		CompletionTokens: types.Int(20),
		TotalTokens:      types.Int(30),
	},
})

// Create a span
err = c.Observations.CreateSpan(ctx, &observations.CreateSpanRequest{
	ID:      types.String("span-123"),
	TraceID: types.String("trace-123"),
	Name:    types.String("database-query"),
})

// Create an event
err = c.Observations.CreateEvent(ctx, &observations.CreateEventRequest{
	ID:      types.String("event-123"),
	TraceID: types.String("trace-123"),
	Name:    types.String("user-feedback"),
})

// List observations
obs, err := c.Observations.List(ctx, &observations.ListParams{
	TraceID: types.String("trace-123"),
	Limit:   types.Int(50),
})
```

### Scores

```go
import (
	"github.com/langfuse/langfuse-go/scores"
	"github.com/langfuse/langfuse-go/types"
)

// Create a numeric score
scoreResp, err := c.Scores.Create(ctx, &scores.CreateRequest{
	Name:    "accuracy",
	Value:   0.95,
	TraceID: "trace-123",
	Comment: types.String("High accuracy response"),
})

// Get a score
score, err := c.Scores.Get(ctx, "score-123")

// List scores
scores, err := c.Scores.List(ctx, &scores.ListParams{
	TraceID: types.String("trace-123"),
	Limit:   types.Int(50),
})

// Delete a score
err = c.Scores.Delete(ctx, "score-123")
```

### Datasets

```go
import (
	"github.com/langfuse/langfuse-go/datasets"
	"github.com/langfuse/langfuse-go/types"
)

// Create a dataset
dataset, err := c.Datasets.Create(ctx, &datasets.CreateRequest{
	Name:        "my-eval-dataset",
	Description: types.String("Evaluation dataset"),
})

// Create dataset items
item, err := c.Datasets.CreateItem(ctx, &datasets.CreateItemRequest{
	DatasetName: types.String("my-eval-dataset"),
	Input: map[string]interface{}{
		"query": "What is the capital of France?",
	},
	ExpectedOutput: map[string]interface{}{
		"answer": "Paris",
	},
})

// Create a dataset run
run, err := c.Datasets.CreateRun(ctx, &datasets.CreateRunRequest{
	Name:      "run-2024-01-15",
	DatasetID: "dataset-123",
})

// Link trace to dataset run
runItem, err := c.Datasets.CreateRunItem(ctx, &datasets.CreateRunItemRequest{
	RunName:       "run-2024-01-15",
	DatasetItemID: "item-123",
	TraceID:       "trace-123",
})

// List datasets
datasets, err := c.Datasets.List(ctx, &types.PaginationParams{
	Page:  types.Int(1),
	Limit: types.Int(50),
})
```

### Sessions

```go
import "github.com/langfuse/langfuse-go/types"

// Get a session with traces
session, err := c.Sessions.Get(ctx, "session-123")

// List sessions
sessions, err := c.Sessions.List(ctx, &types.PaginationParams{
	Page:  types.Int(1),
	Limit: types.Int(50),
})
```

### Models

```go
import (
	"github.com/langfuse/langfuse-go/models"
	"github.com/langfuse/langfuse-go/types"
)

// Create model configuration
model, err := c.Models.Create(ctx, &models.CreateRequest{
	ModelName:    "gpt-4-custom",
	MatchPattern: "gpt-4*",
	InputPrice:   types.Float64(0.03),
	OutputPrice:  types.Float64(0.06),
	Unit:         (*types.ModelUsageUnit)(types.String("TOKENS")),
})

// Get model
model, err := c.Models.Get(ctx, "model-123")

// List models
models, err := c.Models.List(ctx, &types.PaginationParams{
	Page:  types.Int(1),
	Limit: types.Int(50),
})

// Delete model
err = c.Models.Delete(ctx, "model-123")
```

### Prompts

```go
import (
	"github.com/langfuse/langfuse-go/prompts"
	"github.com/langfuse/langfuse-go/types"
)

// Create a chat prompt
prompt, err := c.Prompts.CreateChat(ctx, &prompts.CreateChatRequest{
	Name: "customer-support-prompt",
	Prompt: []prompts.ChatMessage{
		{Role: "system", Content: "You are a helpful assistant."},
		{Role: "user", Content: "{{user_question}}"},
	},
	Tags: []string{"production"},
})

// Create a text prompt
textPrompt, err := c.Prompts.CreateText(ctx, &prompts.CreateTextRequest{
	Name:   "completion-prompt",
	Prompt: "Answer: {{question}}",
})

// Get a prompt (latest version)
prompt, err := c.Prompts.Get(ctx, "customer-support-prompt", nil)

// Get specific version
prompt, err = c.Prompts.Get(ctx, "customer-support-prompt", &prompts.GetParams{
	Version: types.Int(2),
})

// Get by label
prompt, err = c.Prompts.Get(ctx, "customer-support-prompt", &prompts.GetParams{
	Label: types.String("production"),
})

// List prompts
prompts, err := c.Prompts.List(ctx, &prompts.ListParams{
	Page:  types.Int(1),
	Limit: types.Int(50),
	Tag:   types.String("production"),
})
```

### Comments

```go
import (
	"github.com/langfuse/langfuse-go/comments"
	"github.com/langfuse/langfuse-go/types"
)

// Create a comment
comment, err := c.Comments.Create(ctx, &comments.CreateRequest{
	Content:    "This trace needs review",
	ObjectType: types.CommentObjectTypeTrace,
	ObjectID:   "trace-123",
})

// Get comment
comment, err := c.Comments.Get(ctx, "comment-123")

// List comments
comments, err := c.Comments.List(ctx, &comments.ListParams{
	ObjectType: (*types.CommentObjectType)(types.String("TRACE")),
	ObjectID:   types.String("trace-123"),
})
```

### Media

```go
import (
	"github.com/langfuse/langfuse-go/media"
	"github.com/langfuse/langfuse-go/types"
)

// Get upload URL
uploadResp, err := c.Media.GetUploadURL(ctx, &media.UploadURLRequest{
	ContentType:   types.MediaContentTypeImagePng,
	ContentLength: 1024,
	TraceID:       types.String("trace-123"),
})

// Upload file to uploadResp.UploadURL using standard HTTP PUT

// Associate media with trace
err = c.Media.Patch(ctx, uploadResp.MediaID, &media.PatchRequest{
	TraceID: types.String("trace-123"),
	Field:   types.String("input"),
})

// Get media metadata
mediaInfo, err := c.Media.Get(ctx, "media-123")
```

### Metrics

```go
import (
	"github.com/langfuse/langfuse-go/metrics"
	"github.com/langfuse/langfuse-go/types"
	"time"
)

// Get daily metrics
metrics, err := c.Metrics.GetDaily(ctx, &metrics.DailyParams{
	FromTimestamp: types.Time(time.Now().AddDate(0, 0, -7)),
	ToTimestamp:   types.Time(time.Now()),
	Tags:          []string{"production"},
})

for _, metric := range metrics {
	fmt.Printf("Date: %s, Traces: %d, Cost: %.4f\n",
		metric.Date, metric.CountTraces, metric.TotalCost)
}
```

### Annotation Queues

```go
import (
	"github.com/langfuse/langfuse-go/annotations"
	"github.com/langfuse/langfuse-go/types"
)

// Create annotation queue
queue, err := c.Annotations.CreateQueue(ctx, &annotations.CreateQueueRequest{
	Name:           "manual-review-queue",
	Description:    types.String("Queue for manual review"),
	ScoreConfigIDs: []string{"config-1", "config-2"},
})

// Add item to queue
item, err := c.Annotations.CreateQueueItem(ctx, "queue-123", &annotations.CreateQueueItemRequest{
	ObjectID:   "trace-123",
	ObjectType: types.AnnotationQueueObjectTypeTrace,
})

// Update item status
status := types.AnnotationQueueStatusCompleted
item, err = c.Annotations.UpdateQueueItem(ctx, "queue-123", "item-123", &annotations.UpdateQueueItemRequest{
	Status: &status,
})

// List queues
queues, err := c.Annotations.ListQueues(ctx, &types.PaginationParams{
	Page:  types.Int(1),
	Limit: types.Int(50),
})

// List queue items
items, err := c.Annotations.ListQueueItems(ctx, "queue-123", &annotations.ListQueueItemsParams{
	Status: &types.AnnotationQueueStatusPending,
})
```

### Batch Ingestion

```go
import "github.com/langfuse/langfuse-go/ingestion"

// Send batch of events
response, err := c.Ingestion.Ingest(ctx, &ingestion.Request{
	Batch: []interface{}{
		map[string]interface{}{
			"type": "trace-create",
			"body": map[string]interface{}{
				"id":   "trace-123",
				"name": "batch-trace",
			},
		},
		map[string]interface{}{
			"type": "generation-create",
			"body": map[string]interface{}{
				"id":      "gen-123",
				"traceId": "trace-123",
				"name":    "batch-generation",
			},
		},
	},
})

fmt.Printf("Successes: %d, Errors: %d\n",
	len(response.Successes), len(response.Errors))
```

## Helper Functions

The `types` package provides convenient pointer helper functions:

```go
import "github.com/langfuse/langfuse-go/types"

// Primitives
types.Bool(true)
types.String("value")
types.Int(42)
types.Int64(42)
types.Float64(3.14)

// Time
types.Time(time.Now())

// UUID
types.UUID(uuid.New())

// Dates
types.MustParseDate("2024-01-15")
types.MustParseDateTime("2024-01-15T10:30:00Z")
```

## Error Handling

All client methods return errors that should be checked:

```go
trace, err := c.Traces.Get(ctx, "trace-123")
if err != nil {
	// Handle error
	log.Printf("Failed to get trace: %v", err)
	return
}

// Use trace
fmt.Printf("Trace name: %s\n", *trace.Name)
```

## Context Support

All client methods accept a `context.Context` for cancellation and timeouts:

```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()

err := c.Traces.Create(ctx, &traces.CreateTraceRequest{
	Name: types.String("my-trace"),
})
```

## Examples

See the [examples](./examples) directory for complete working examples:

- [basic](./examples/basic) - Basic usage with traces, generations, and scores
- [datasets](./examples/datasets) - Dataset management and evaluation runs

## Package Documentation

For detailed package documentation, see:

- [pkg.go.dev/github.com/langfuse/langfuse-go](https://pkg.go.dev/github.com/langfuse/langfuse-go)

## License

MIT

## Contributing

Contributions are welcome! Please see the [GitHub repository](https://github.com/langfuse/langfuse-go) for more information.

## Support

- [Documentation](https://langfuse.com/docs)
- [Discord Community](https://discord.gg/langfuse)
- [GitHub Issues](https://github.com/langfuse/langfuse-go/issues)
