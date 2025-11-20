# Langfuse Go Client

[![Go Reference](https://pkg.go.dev/badge/github.com/langfuse/langfuse-go.svg)](https://pkg.go.dev/github.com/langfuse/langfuse-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/langfuse/langfuse-go)](https://goreportcard.com/report/github.com/langfuse/langfuse-go)

Official Go client library for [Langfuse](https://langfuse.com) - the open source LLM engineering platform.

## Installation

```bash
go get github.com/langfuse/langfuse-go
```

## Quick Start

```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/langfuse/langfuse-go/client"
)

func main() {
	// Initialize the client
	c := client.NewClient(
		"your-public-key",
		"your-secret-key",
	)

	// Create a trace
	ctx := context.Background()
	err := c.CreateTrace(ctx, &client.CreateTraceBody{
		Name:   client.String("my-trace"),
		UserID: client.String("user-123"),
		Input:  map[string]interface{}{"query": "What is Langfuse?"},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Trace created successfully!")
}
```

## Configuration

### Custom Base URL

For self-hosted instances:

```go
c := client.NewClient(
	"your-public-key",
	"your-secret-key",
	client.WithBaseURL("https://your-instance.com"),
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

c := client.NewClient(
	"your-public-key",
	"your-secret-key",
	client.WithHTTPClient(httpClient),
)
```

### Custom Timeout

```go
c := client.NewClient(
	"your-public-key",
	"your-secret-key",
	client.WithTimeout(60 * time.Second),
)
```

## Usage Examples

### Traces

#### Create a Trace

```go
err := c.CreateTrace(ctx, &client.CreateTraceBody{
	ID:        client.String("trace-123"),
	Name:      client.String("my-llm-app"),
	UserID:    client.String("user-456"),
	SessionID: client.String("session-789"),
	Metadata: map[string]interface{}{
		"environment": "production",
		"version":     "1.0.0",
	},
	Tags:   []string{"production", "important"},
	Input:  map[string]interface{}{"query": "What is Langfuse?"},
	Output: map[string]interface{}{"response": "Langfuse is an LLM engineering platform"},
})
```

#### Get a Trace

```go
trace, err := c.GetTrace(ctx, "trace-123")
if err != nil {
	log.Fatal(err)
}
fmt.Printf("Trace: %+v\n", trace)
```

#### Update a Trace

```go
err := c.UpdateTrace(ctx, "trace-123", &client.UpdateTraceBody{
	Name:   client.String("updated-name"),
	Public: client.Bool(true),
	Output: map[string]interface{}{"response": "Updated response"},
})
```

### Observations

#### Create a Generation

```go
err := c.CreateGeneration(ctx, &client.CreateGenerationBody{
	ID:      client.String("gen-123"),
	TraceID: client.String("trace-123"),
	Name:    client.String("openai-completion"),
	Model:   client.String("gpt-4"),
	Input: []map[string]interface{}{
		{"role": "user", "content": "What is Langfuse?"},
	},
	Output: map[string]interface{}{
		"role":    "assistant",
		"content": "Langfuse is an LLM engineering platform",
	},
	Usage: &client.Usage{
		PromptTokens:     client.Int(10),
		CompletionTokens: client.Int(20),
		TotalTokens:      client.Int(30),
	},
	ModelParameters: map[string]interface{}{
		"temperature": 0.7,
		"max_tokens":  100,
	},
})
```

#### Create a Span

```go
import "time"

err := c.CreateSpan(ctx, &client.CreateSpanBody{
	ID:                  client.String("span-123"),
	TraceID:             client.String("trace-123"),
	ParentObservationID: client.String("parent-span-456"),
	Name:                client.String("database-query"),
	StartTime:           client.Time(time.Now()),
	EndTime:             client.Time(time.Now().Add(100 * time.Millisecond)),
	Input: map[string]interface{}{
		"query": "SELECT * FROM users WHERE id = ?",
		"params": []interface{}{123},
	},
	Output: map[string]interface{}{
		"rows": 1,
	},
	Metadata: map[string]interface{}{
		"database": "postgres",
	},
})
```

#### Create an Event

```go
err := c.CreateEvent(ctx, &client.CreateEventBody{
	ID:      client.String("event-123"),
	TraceID: client.String("trace-123"),
	Name:    client.String("user-feedback"),
	Input: map[string]interface{}{
		"rating":  5,
		"comment": "Great response!",
	},
	Level: (*client.ObservationLevel)(client.String("DEFAULT")),
})
```

#### Get Observations

```go
observations, err := c.GetObservations(ctx, &client.GetObservationsParams{
	Page:    client.Int(1),
	Limit:   client.Int(50),
	TraceID: client.String("trace-123"),
})
if err != nil {
	log.Fatal(err)
}

for _, obs := range observations.Data {
	fmt.Printf("Observation: %s (Type: %s)\n", obs.ID, obs.Type)
}
```

### Scores

#### Create a Score

```go
// Numeric score
score, err := c.CreateScore(ctx, &client.CreateScoreRequest{
	Name:    "accuracy",
	Value:   0.95,
	TraceID: "trace-123",
	Comment: client.String("High accuracy response"),
})

// Categorical score
score, err := c.CreateScore(ctx, &client.CreateScoreRequest{
	Name:     "sentiment",
	Value:    "positive",
	DataType: (*client.ScoreDataType)(client.String("CATEGORICAL")),
	TraceID:  "trace-123",
})

// Boolean score
score, err := c.CreateScore(ctx, &client.CreateScoreRequest{
	Name:     "is_correct",
	Value:    true,
	DataType: (*client.ScoreDataType)(client.String("BOOLEAN")),
	TraceID:  "trace-123",
})
```

#### Get Scores

```go
scores, err := c.GetScores(ctx, &client.GetScoresParams{
	Page:    client.Int(1),
	Limit:   client.Int(50),
	TraceID: client.String("trace-123"),
})
if err != nil {
	log.Fatal(err)
}

for _, score := range scores.Data {
	fmt.Printf("Score: %s = %v\n", score.Name, score.Value)
}
```

### Datasets

#### Create a Dataset

```go
dataset, err := c.CreateDataset(ctx, &client.CreateDatasetRequest{
	Name:        "my-eval-dataset",
	Description: client.String("Evaluation dataset for my LLM app"),
	Metadata: map[string]interface{}{
		"purpose": "testing",
	},
})
```

#### Create Dataset Items

```go
item, err := c.CreateDatasetItem(ctx, &client.CreateDatasetItemRequest{
	DatasetName: client.String("my-eval-dataset"),
	Input: map[string]interface{}{
		"query": "What is the capital of France?",
	},
	ExpectedOutput: map[string]interface{}{
		"answer": "Paris",
	},
	Metadata: map[string]interface{}{
		"category": "geography",
	},
})
```

#### Create Dataset Runs

```go
run, err := c.CreateDatasetRun(ctx, &client.CreateDatasetRunRequest{
	Name:        "run-2024-01-15",
	DatasetID:   "dataset-123",
	Description: client.String("Evaluation run for GPT-4"),
	Metadata: map[string]interface{}{
		"model": "gpt-4",
	},
})
```

#### Link Traces to Dataset Runs

```go
runItem, err := c.CreateDatasetRunItem(ctx, &client.CreateDatasetRunItemRequest{
	RunName:       "run-2024-01-15",
	DatasetItemID: "item-123",
	TraceID:       "trace-123",
	Metadata: map[string]interface{}{
		"runtime_ms": 250,
	},
})
```

### Sessions

#### Get a Session

```go
session, err := c.GetSession(ctx, "session-123")
if err != nil {
	log.Fatal(err)
}

fmt.Printf("Session has %d traces\n", len(session.Traces))
```

#### List Sessions

```go
sessions, err := c.GetSessions(ctx, &client.PaginationParams{
	Page:  client.Int(1),
	Limit: client.Int(50),
})
```

### Models

#### Create Model Configuration

```go
model, err := c.CreateModel(ctx, &client.CreateModelRequest{
	ModelName:    "gpt-4-custom",
	MatchPattern: "gpt-4*",
	InputPrice:   client.Float64(0.03),
	OutputPrice:  client.Float64(0.06),
	Unit:         (*client.ModelUsageUnit)(client.String("TOKENS")),
})
```

#### List Models

```go
models, err := c.GetModels(ctx, &client.PaginationParams{
	Page:  client.Int(1),
	Limit: client.Int(50),
})
```

### Prompts

#### Create a Chat Prompt

```go
prompt, err := c.CreateChatPrompt(ctx, &client.CreateChatPromptRequest{
	Name: "customer-support-prompt",
	Prompt: []client.ChatMessage{
		{Role: "system", Content: "You are a helpful customer support agent."},
		{Role: "user", Content: "{{user_question}}"},
	},
	Tags:   []string{"production", "customer-support"},
	Labels: []string{"v1"},
	Config: map[string]interface{}{
		"temperature": 0.7,
		"max_tokens":  500,
	},
})
```

#### Create a Text Prompt

```go
prompt, err := c.CreateTextPrompt(ctx, &client.CreateTextPromptRequest{
	Name:   "completion-prompt",
	Prompt: "Answer the following question: {{question}}",
	Tags:   []string{"production"},
	Config: map[string]interface{}{
		"temperature": 0.5,
	},
})
```

#### Get a Prompt

```go
// Get latest version
prompt, err := c.GetPrompt(ctx, "customer-support-prompt", nil, nil)

// Get specific version
prompt, err := c.GetPrompt(ctx, "customer-support-prompt", client.Int(2), nil)

// Get by label
prompt, err := c.GetPrompt(ctx, "customer-support-prompt", nil, client.String("production"))
```

### Comments

#### Create a Comment

```go
comment, err := c.CreateComment(ctx, &client.CreateCommentRequest{
	Content:    "This trace needs review",
	ObjectType: client.CommentObjectTypeTrace,
	ObjectID:   "trace-123",
})
```

#### Get Comments

```go
comments, err := c.GetComments(ctx, &client.GetCommentsParams{
	ObjectType: (*client.CommentObjectType)(client.String("TRACE")),
	ObjectID:   client.String("trace-123"),
})
```

### Media

#### Upload Media

```go
// Get upload URL
uploadResp, err := c.GetMediaUploadURL(ctx, &client.GetMediaUploadUrlRequest{
	ContentType:   client.MediaContentTypeImagePng,
	ContentLength: 1024,
	TraceID:       client.String("trace-123"),
	Field:         client.String("input"),
})

// Upload file to uploadResp.UploadURL using standard HTTP PUT
// Then associate the media with the trace
err = c.PatchMedia(ctx, uploadResp.MediaID, &client.PatchMediaBody{
	TraceID: client.String("trace-123"),
	Field:   client.String("input"),
})
```

### Metrics

#### Get Daily Metrics

```go
import "time"

metrics, err := c.GetDailyMetrics(ctx, &client.GetDailyMetricsParams{
	FromTimestamp: client.Time(time.Now().AddDate(0, 0, -7)),
	ToTimestamp:   client.Time(time.Now()),
	Tags:          []string{"production"},
})

for _, metric := range metrics {
	fmt.Printf("Date: %s, Traces: %d, Cost: %.4f\n",
		metric.Date, metric.CountTraces, metric.TotalCost)
}
```

### Annotation Queues

#### Create an Annotation Queue

```go
queue, err := c.CreateAnnotationQueue(ctx, &client.CreateAnnotationQueueRequest{
	Name:           "manual-review-queue",
	Description:    client.String("Queue for traces requiring manual review"),
	ScoreConfigIDs: []string{"config-1", "config-2"},
})
```

#### Add Items to Queue

```go
item, err := c.CreateAnnotationQueueItem(ctx, "queue-123", &client.CreateAnnotationQueueItemRequest{
	ObjectID:   "trace-123",
	ObjectType: client.AnnotationQueueObjectTypeTrace,
})
```

#### Update Queue Item Status

```go
status := client.AnnotationQueueStatusCompleted
item, err := c.UpdateAnnotationQueueItem(ctx, "queue-123", "item-123", &client.UpdateAnnotationQueueItemRequest{
	Status: &status,
})
```

## Helper Functions

The library provides convenient pointer helper functions for optional fields:

```go
// Primitives
client.Bool(true)
client.String("value")
client.Int(42)
client.Int64(42)
client.Float64(3.14)

// Time
client.Time(time.Now())

// UUID
client.UUID(uuid.New())

// Dates
client.MustParseDate("2024-01-15")
client.MustParseDateTime("2024-01-15T10:30:00Z")
```

## Error Handling

All client methods return errors that should be checked:

```go
trace, err := c.GetTrace(ctx, "trace-123")
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

err := c.CreateTrace(ctx, &client.CreateTraceBody{
	Name: client.String("my-trace"),
})
```

## License

MIT

## Contributing

Contributions are welcome! Please see the [GitHub repository](https://github.com/langfuse/langfuse-go) for more information.

## Support

- [Documentation](https://langfuse.com/docs)
- [Discord Community](https://discord.gg/langfuse)
- [GitHub Issues](https://github.com/langfuse/langfuse-go/issues)
