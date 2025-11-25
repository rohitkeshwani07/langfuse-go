package traces

import (
	"encoding/json"
	"sort"
	"time"

	"github.com/rohitkeshwani07/langfuse-go/observations"
)

// Trace represents a trace in Langfuse
type Trace struct {
	ID           string                      `json:"id"`
	Name         string                      `json:"name,omitempty"`
	UserID       string                      `json:"userId,omitempty"`
	SessionID    string                      `json:"sessionId,omitempty"`
	Release      *string                     `json:"release,omitempty"`
	Version      *string                     `json:"version,omitempty"`
	Metadata     map[string]interface{}      `json:"metadata,omitempty"`
	Tags         []string                    `json:"tags,omitempty"`
	Input        json.RawMessage             `json:"input,omitempty"`
	Output       json.RawMessage             `json:"output,omitempty"`
	Timestamp    time.Time                   `json:"timestamp"`
	CreatedAt    time.Time                   `json:"createdAt"`
	UpdatedAt    time.Time                   `json:"updatedAt"`
	Public       bool                        `json:"public,omitempty"`
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

// ObservationNode represents a node in the trace tree structure.
// This is a compact representation that excludes input/output fields to reduce memory allocation
// when working with large trace trees.
type ObservationNode struct {
	*observations.Observation
	Children []*ObservationNode `json:"children,omitempty"`
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
	Input       json.RawMessage        `json:"input,omitempty"`
	Output      json.RawMessage        `json:"output,omitempty"`
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
func buildObservationTree(observations []*observations.Observation) []*ObservationNode {
	// First pass: build a set of all valid observation IDs
	validIDs := make(map[string]bool)
	for _, observation := range observations {
		validIDs[observation.ID] = true
	}

	idToNodes := make(map[string][]*ObservationNode)
	for _, observation := range observations {
		node := &ObservationNode{Observation: observation}

		// Treat as root if ParentObservationID is empty OR if the parent doesn't exist in the observation set
		if observation.ParentObservationID == "" || !validIDs[observation.ParentObservationID] {
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
func (t *Trace) ToTraceTree() *TraceTree {
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
