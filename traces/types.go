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

// TreeObservation represents a single observation in the trace tree response
type TreeObservation struct {
	ID                   string                 `json:"id"`
	Type                 string                 `json:"type"`
	ParentObservationID  string                 `json:"parentObservationId"`
	StartTime            time.Time              `json:"startTime"`
	EndTime              time.Time              `json:"endTime"`
	Name                 string                 `json:"name"`
	Metadata             map[string]interface{} `json:"metadata"`
	ModelParameters      map[string]interface{} `json:"modelParameters"`
	CompletionStartTime  string                 `json:"completionStartTime"`
	CreatedAt            string                 `json:"createdAt"`
	UpdatedAt            string                 `json:"updatedAt"`
	UsageDetails         map[string]interface{} `json:"usageDetails"`
	CostDetails          map[string]interface{} `json:"costDetails"`
	Model                string                 `json:"model"`
	Latency              float64                `json:"latency"`
	InputPrice           float64                `json:"inputPrice"`
	OutputPrice          float64                `json:"outputPrice"`
	TotalPrice           float64                `json:"totalPrice"`
	CalculatedInputCost  float64                `json:"calculatedInputCost"`
	CalculatedOutputCost float64                `json:"calculatedOutputCost"`
	CalculatedTotalCost  float64                `json:"calculatedTotalCost"`
	Unit                 string                 `json:"unit"`
	PromptTokens         int64                  `json:"promptTokens"`
	CompletionTokens     int64                  `json:"completionTokens"`
	TotalTokens          int64                  `json:"totalTokens"`
	ModelID              string                 `json:"modelId"`
	Usage                map[string]interface{} `json:"usage"`
}

// TraceWithObservations represents the JSON payload with flat observations
type TraceWithObservations struct {
	ID           string                 `json:"id"`
	Name         string                 `json:"name"`
	Timestamp    time.Time              `json:"timestamp"`
	UserID       string                 `json:"userId"`
	SessionID    string                 `json:"sessionId"`
	Input        map[string]interface{} `json:"input"`
	Output       map[string]interface{} `json:"output"`
	Metadata     map[string]interface{} `json:"metadata"`
	CreatedAt    string                 `json:"createdAt"`
	UpdatedAt    string                 `json:"updatedAt"`
	ExternalID   string                 `json:"externalId"`
	Latency      float64                `json:"latency"`
	Observations []*TreeObservation     `json:"observations"`
	TotalCost    float64                `json:"totalCost"`
}

// ObservationNode represents a node in the trace tree structure
type ObservationNode struct {
	ID                   string                 `json:"id"`
	Name                 string                 `json:"name"`
	Type                 string                 `json:"type"`
	StartTime            time.Time              `json:"startTime"`
	EndTime              time.Time              `json:"endTime"`
	Metadata             map[string]interface{} `json:"metadata,omitempty"`
	Latency              float64                `json:"latency"`
	Children             []*ObservationNode     `json:"children,omitempty"`
	ModelParameters      map[string]interface{} `json:"modelParameters,omitempty"`
	CompletionStartTime  string                 `json:"completionStartTime"`
	CreatedAt            string                 `json:"createdAt"`
	UpdatedAt            string                 `json:"updatedAt"`
	UsageDetails         map[string]interface{} `json:"usageDetails"`
	CostDetails          map[string]interface{} `json:"costDetails"`
	Model                string                 `json:"model"`
	InputPrice           float64                `json:"inputPrice"`
	OutputPrice          float64                `json:"outputPrice"`
	TotalPrice           float64                `json:"totalPrice"`
	CalculatedInputCost  float64                `json:"calculatedInputCost"`
	CalculatedOutputCost float64                `json:"calculatedOutputCost"`
	CalculatedTotalCost  float64                `json:"calculatedTotalCost"`
	Unit                 string                 `json:"unit"`
	PromptTokens         int64                  `json:"promptTokens"`
	CompletionTokens     int64                  `json:"completionTokens"`
	TotalTokens          int64                  `json:"totalTokens"`
	ModelID              string                 `json:"modelId"`
	Usage                map[string]interface{} `json:"usage"`
}

// TraceTree represents the trace with tree-structured observations
type TraceTree struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Timestamp  time.Time              `json:"timestamp"`
	UserID     string                 `json:"userId"`
	SessionID  string                 `json:"sessionId"`
	Input      map[string]interface{} `json:"input"`
	Output     map[string]interface{} `json:"output"`
	Metadata   map[string]interface{} `json:"metadata"`
	CreatedAt  string                 `json:"createdAt"`
	UpdatedAt  string                 `json:"updatedAt"`
	ExternalID interface{}            `json:"externalId"`
	Latency    float64                `json:"latency"`
	RootNode   []*ObservationNode     `json:"rootNode"`
	TotalCost  float64                `json:"totalCost"`
}

// buildObservationTree converts flat observations to a tree structure
func buildObservationTree(observations []*TreeObservation) []*ObservationNode {
	idToNodes := make(map[string][]*ObservationNode)
	for _, observation := range observations {
		node := &ObservationNode{
			ID:                   observation.ID,
			Type:                 observation.Type,
			StartTime:            observation.StartTime,
			EndTime:              observation.EndTime,
			Name:                 observation.Name,
			Metadata:             observation.Metadata,
			ModelParameters:      observation.ModelParameters,
			CostDetails:          observation.CostDetails,
			Model:                observation.Model,
			Latency:              observation.Latency,
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

// ToTraceTree converts TraceWithObservations to TraceTree
func (t *TraceWithObservations) ToTraceTree() *TraceTree {
	nodes := buildObservationTree(t.Observations)
	return &TraceTree{
		ID:         t.ID,
		Name:       t.Name,
		Timestamp:  t.Timestamp,
		UserID:     t.UserID,
		SessionID:  t.SessionID,
		Input:      t.Input,
		Output:     t.Output,
		Metadata:   t.Metadata,
		CreatedAt:  t.CreatedAt,
		UpdatedAt:  t.UpdatedAt,
		ExternalID: t.ExternalID,
		RootNode:   nodes,
		Latency:    t.Latency,
		TotalCost:  t.TotalCost,
	}
}
