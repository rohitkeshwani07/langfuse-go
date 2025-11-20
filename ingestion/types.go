package ingestion

// Request represents a request to the ingestion API
type Request struct {
	Batch    []interface{}          `json:"batch"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// Response represents a response from the ingestion API
type Response struct {
	Successes []string `json:"successes"`
	Errors    []struct {
		ID      string `json:"id"`
		Status  int    `json:"status"`
		Message string `json:"message"`
		Error   string `json:"error"`
	} `json:"errors"`
}
