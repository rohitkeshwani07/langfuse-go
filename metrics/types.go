package metrics

import (
	"time"

	"github.com/rohitkeshwani07/langfuse-go/types"
)

// Daily represents daily metrics
type Daily struct {
	Date              string       `json:"date"`
	CountTraces       int          `json:"countTraces"`
	CountObservations int          `json:"countObservations"`
	TotalCost         float64      `json:"totalCost"`
	Usage             []types.Usage `json:"usage"`
}

// DailyParams represents query parameters for getting daily metrics
type DailyParams struct {
	TraceName     *string
	UserID        *string
	Tags          []string
	FromTimestamp *time.Time
	ToTimestamp   *time.Time
}
