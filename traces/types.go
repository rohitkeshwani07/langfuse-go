package traces

import (
	"sort"
	"time"

	"github.com/langfuse/langfuse-go/observations"
)

// Trace represents a trace in Langfuse
type Trace struct {
	ID           string                      `json:"id"`
	Name         *string                     `json:"name,omitempty"`
	UserID       *string                     `json:"userId,omitempty"`
	SessionID    *string                     `json:"sessionId,omitempty"`
	Release      *string                     `json:"release,omitempty"`
	Version      *string                     `json:"version,omitempty"`
	Metadata     map[string]interface{}      `json:"metadata,omitempty"`
	Tags         []string                    `json:"tags,omitempty"`
	Input        interface{}                 `json:"input,omitempty"`
	Output       interface{}                 `json:"output,omitempty"`
	Timestamp    time.Time                   `json:"timestamp"`
	CreatedAt    time.Time                   `json:"createdAt"`
	UpdatedAt    time.Time                   `json:"updatedAt"`
	Public       *bool                       `json:"public,omitempty"`
	ProjectID    string                      `json:"projectId"`
	Environment  string                      `json:"environment,omitempty"`
	Bookmarked   bool                        `json:"bookmarked,omitempty"`
	ExternalID   interface{}                 `json:"externalId,omitempty"`
	Latency      float64                     `json:"latency,omitempty"`
	TotalCost    float64                     `json:"totalCost,omitempty"`
	HtmlPath     string                      `json:"htmlPath,omitempty"`
	Scores       []interface{}               `json:"scores,omitempty"`
	Observations []*observations.Observation `json:"observations,omitempty"`
}

// CreateTraceRequest represents the request body for creating a trace
type CreateTraceRequest struct {
	ID        *string                `json:"id,omitempty"`
	Name      *string                `json:"name,omitempty"`
	UserID    *string                `json:"userId,omitempty"`
	SessionID *string                `json:"sessionId,omitempty"`
	Release   *string                `json:"release,omitempty"`
	Version   *string                `json:"version,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Tags      []string               `json:"tags,omitempty"`
	Input     interface{}            `json:"input,omitempty"`
	Output    interface{}            `json:"output,omitempty"`
	Timestamp *time.Time             `json:"timestamp,omitempty"`
	Public    *bool                  `json:"public,omitempty"`
}

// UpdateTraceRequest represents the request body for updating a trace
type UpdateTraceRequest struct {
	Name      *string                `json:"name,omitempty"`
	UserID    *string                `json:"userId,omitempty"`
	SessionID *string                `json:"sessionId,omitempty"`
	Release   *string                `json:"release,omitempty"`
	Version   *string                `json:"version,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Tags      []string               `json:"tags,omitempty"`
	Input     interface{}            `json:"input,omitempty"`
	Output    interface{}            `json:"output,omitempty"`
	Public    *bool                  `json:"public,omitempty"`
}

// CompactObservation represents a single observation in the trace tree response.
// This is a compact representation that excludes input/output fields to reduce memory allocation
// when fetching traces with large payloads.
type CompactObservation struct {
	ID                   string                 `json:"id"`
	TraceID              *string                `json:"traceId,omitempty"`
	ProjectID            string                 `json:"projectId"`
	Type                 string                 `json:"type"`
	ParentObservationID  string                 `json:"parentObservationId,omitempty"`
	StartTime            time.Time              `json:"startTime"`
	EndTime              *time.Time             `json:"endTime,omitempty"`
	Name                 string                 `json:"name"`
	Environment          string                 `json:"environment,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	Level                string                 `json:"level,omitempty"`
	StatusMessage        *string                `json:"statusMessage,omitempty"`
	Version              *string                `json:"version,omitempty"`
	ModelParameters      map[string]interface{} `json:"modelParameters,omitempty"`
	CompletionStartTime  *time.Time             `json:"completionStartTime,omitempty"`
	CreatedAt            time.Time              `json:"createdAt"`
	UpdatedAt            time.Time              `json:"updatedAt"`
	UsageDetails         map[string]interface{} `json:"usageDetails,omitempty"`
	CostDetails          map[string]interface{} `json:"costDetails,omitempty"`
	Model                *string                `json:"model,omitempty"`
	PromptID             *string                `json:"promptId,omitempty"`
	PromptName           *string                `json:"promptName,omitempty"`
	PromptVersion        *int                   `json:"promptVersion,omitempty"`
	Latency              float64                `json:"latency,omitempty"`
	TimeToFirstToken     *float64               `json:"timeToFirstToken,omitempty"`
	InputPrice           float64                `json:"inputPrice,omitempty"`
	OutputPrice          float64                `json:"outputPrice,omitempty"`
	TotalPrice           float64                `json:"totalPrice,omitempty"`
	CalculatedInputCost  *float64               `json:"calculatedInputCost,omitempty"`
	CalculatedOutputCost *float64               `json:"calculatedOutputCost,omitempty"`
	CalculatedTotalCost  float64                `json:"calculatedTotalCost,omitempty"`
	Unit                 string                 `json:"unit,omitempty"`
	PromptTokens         *int                   `json:"promptTokens,omitempty"`
	CompletionTokens     *int                   `json:"completionTokens,omitempty"`
	TotalTokens          *int                   `json:"totalTokens,omitempty"`
	ModelID              *string                `json:"modelId,omitempty"`
	Usage                map[string]interface{} `json:"usage,omitempty"`
}

// CompactTrace represents the JSON payload with flat observations.
// Uses CompactObservation to reduce memory allocation for large trace payloads.
type CompactTrace struct {
	ID           string                  `json:"id"`
	ProjectID    string                  `json:"projectId"`
	Name         string                  `json:"name"`
	Timestamp    time.Time               `json:"timestamp"`
	Environment  string                  `json:"environment,omitempty"`
	Tags         []string                `json:"tags,omitempty"`
	Bookmarked   bool                    `json:"bookmarked,omitempty"`
	Release      *string                 `json:"release,omitempty"`
	Version      *string                 `json:"version,omitempty"`
	UserID       string                  `json:"userId,omitempty"`
	SessionID    string                  `json:"sessionId,omitempty"`
	Public       bool                    `json:"public,omitempty"`
	Input        interface{}             `json:"input,omitempty"`
	Output       interface{}             `json:"output,omitempty"`
	Metadata     map[string]interface{}  `json:"metadata,omitempty"`
	CreatedAt    time.Time               `json:"createdAt"`
	UpdatedAt    time.Time               `json:"updatedAt"`
	ExternalID   interface{}             `json:"externalId,omitempty"`
	Scores       []interface{}           `json:"scores,omitempty"`
	Latency      float64                 `json:"latency,omitempty"`
	Observations []*CompactObservation   `json:"observations,omitempty"`
	HtmlPath     string                  `json:"htmlPath,omitempty"`
	TotalCost    float64                 `json:"totalCost,omitempty"`
}

// ObservationNode represents a node in the trace tree structure.
// This is a compact representation that excludes input/output fields to reduce memory allocation
// when working with large trace trees.
type ObservationNode struct {
	ID                   string                 `json:"id"`
	TraceID              *string                `json:"traceId,omitempty"`
	ProjectID            string                 `json:"projectId"`
	Name                 string                 `json:"name"`
	Type                 string                 `json:"type"`
	Environment          string                 `json:"environment,omitempty"`
	StartTime            time.Time              `json:"startTime"`
	EndTime              *time.Time             `json:"endTime,omitempty"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	Level                string                 `json:"level,omitempty"`
	StatusMessage        *string                `json:"statusMessage,omitempty"`
	Version              *string                `json:"version,omitempty"`
	Latency              float64                `json:"latency,omitempty"`
	Children             []*ObservationNode     `json:"children,omitempty"`
	ModelParameters      map[string]interface{} `json:"modelParameters,omitempty"`
	CompletionStartTime  *time.Time             `json:"completionStartTime,omitempty"`
	CreatedAt            time.Time              `json:"createdAt"`
	UpdatedAt            time.Time              `json:"updatedAt"`
	UsageDetails         map[string]interface{} `json:"usageDetails,omitempty"`
	CostDetails          map[string]interface{} `json:"costDetails,omitempty"`
	Model                *string                `json:"model,omitempty"`
	PromptID             *string                `json:"promptId,omitempty"`
	PromptName           *string                `json:"promptName,omitempty"`
	PromptVersion        *int                   `json:"promptVersion,omitempty"`
	TimeToFirstToken     *float64               `json:"timeToFirstToken,omitempty"`
	InputPrice           float64                `json:"inputPrice,omitempty"`
	OutputPrice          float64                `json:"outputPrice,omitempty"`
	TotalPrice           float64                `json:"totalPrice,omitempty"`
	CalculatedInputCost  *float64               `json:"calculatedInputCost,omitempty"`
	CalculatedOutputCost *float64               `json:"calculatedOutputCost,omitempty"`
	CalculatedTotalCost  float64                `json:"calculatedTotalCost,omitempty"`
	Unit                 string                 `json:"unit,omitempty"`
	PromptTokens         *int                   `json:"promptTokens,omitempty"`
	CompletionTokens     *int                   `json:"completionTokens,omitempty"`
	TotalTokens          *int                   `json:"totalTokens,omitempty"`
	ModelID              *string                `json:"modelId,omitempty"`
	Usage                map[string]interface{} `json:"usage,omitempty"`
}

// TraceTree represents the trace with tree-structured observations
type TraceTree struct {
	ID          string                 `json:"id"`
	ProjectID   string                 `json:"projectId"`
	Name        string                 `json:"name"`
	Timestamp   time.Time              `json:"timestamp"`
	Environment string                 `json:"environment,omitempty"`
	Tags        []string               `json:"tags,omitempty"`
	Bookmarked  bool                   `json:"bookmarked,omitempty"`
	Release     *string                `json:"release,omitempty"`
	Version     *string                `json:"version,omitempty"`
	UserID      string                 `json:"userId,omitempty"`
	SessionID   string                 `json:"sessionId,omitempty"`
	Public      bool                   `json:"public,omitempty"`
	Input       interface{}            `json:"input,omitempty"`
	Output      interface{}            `json:"output,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
	CreatedAt   time.Time              `json:"createdAt"`
	UpdatedAt   time.Time              `json:"updatedAt"`
	ExternalID  interface{}            `json:"externalId,omitempty"`
	Scores      []interface{}          `json:"scores,omitempty"`
	Latency     float64                `json:"latency,omitempty"`
	RootNode    []*ObservationNode     `json:"rootNode"`
	HtmlPath    string                 `json:"htmlPath,omitempty"`
	TotalCost   float64                `json:"totalCost,omitempty"`
}

// buildObservationTree converts flat observations to a tree structure
func buildObservationTree(observations []*CompactObservation) []*ObservationNode {
	idToNodes := make(map[string][]*ObservationNode)
	for _, observation := range observations {
		node := &ObservationNode{
			ID:                   observation.ID,
			TraceID:              observation.TraceID,
			ProjectID:            observation.ProjectID,
			Type:                 observation.Type,
			Environment:          observation.Environment,
			StartTime:            observation.StartTime,
			EndTime:              observation.EndTime,
			Name:                 observation.Name,
			Metadata:             observation.Metadata,
			Level:                observation.Level,
			StatusMessage:        observation.StatusMessage,
			Version:              observation.Version,
			ModelParameters:      observation.ModelParameters,
			CostDetails:          observation.CostDetails,
			Model:                observation.Model,
			PromptID:             observation.PromptID,
			PromptName:           observation.PromptName,
			PromptVersion:        observation.PromptVersion,
			Latency:              observation.Latency,
			TimeToFirstToken:     observation.TimeToFirstToken,
			InputPrice:           observation.InputPrice,
			OutputPrice:          observation.OutputPrice,
			TotalPrice:           observation.TotalPrice,
			CalculatedInputCost:  observation.CalculatedInputCost,
			CalculatedOutputCost: observation.CalculatedOutputCost,
			CalculatedTotalCost:  observation.CalculatedTotalCost,
			Unit:                 observation.Unit,
			PromptTokens:         observation.PromptTokens,
			CompletionTokens:     observation.CompletionTokens,
			TotalTokens:          observation.TotalTokens,
			ModelID:              observation.ModelID,
			Usage:                observation.Usage,
			CreatedAt:            observation.CreatedAt,
			UpdatedAt:            observation.UpdatedAt,
			UsageDetails:         observation.UsageDetails,
			CompletionStartTime:  observation.CompletionStartTime,
		}

		if observation.ParentObservationID == "" {
			idToNodes["root"] = append(idToNodes["root"], node)
		} else {
			idToNodes[observation.ParentObservationID] = append(idToNodes[observation.ParentObservationID], node)
		}
	}

	// Helper function to build tree recursively
	var buildTree func(node *ObservationNode)
	buildTree = func(node *ObservationNode) {
		if children, ok := idToNodes[node.ID]; ok {
			// sort the array of children by the startTime
			sort.Slice(children, func(i, j int) bool {
				return children[i].StartTime.Before(children[j].StartTime)
			})

			node.Children = children

			// build the tree recursively
			for i := range node.Children {
				buildTree(node.Children[i])
			}
		}

		// no children, return
	}

	// Start from the root nodes and build the tree recursively
	rootResults := idToNodes["root"]

	sort.Slice(rootResults, func(i, j int) bool {
		return rootResults[i].StartTime.Before(rootResults[j].StartTime)
	})

	for i := range rootResults {
		buildTree(rootResults[i])
	}

	return rootResults
}

// ToTraceTree converts CompactTrace to TraceTree
func (t *CompactTrace) ToTraceTree() *TraceTree {
	nodes := buildObservationTree(t.Observations)
	return &TraceTree{
		ID:          t.ID,
		ProjectID:   t.ProjectID,
		Name:        t.Name,
		Timestamp:   t.Timestamp,
		Environment: t.Environment,
		Tags:        t.Tags,
		Bookmarked:  t.Bookmarked,
		Release:     t.Release,
		Version:     t.Version,
		UserID:      t.UserID,
		SessionID:   t.SessionID,
		Public:      t.Public,
		Input:       t.Input,
		Output:      t.Output,
		Metadata:    t.Metadata,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
		ExternalID:  t.ExternalID,
		Scores:      t.Scores,
		Latency:     t.Latency,
		RootNode:    nodes,
		HtmlPath:    t.HtmlPath,
		TotalCost:   t.TotalCost,
	}
}
