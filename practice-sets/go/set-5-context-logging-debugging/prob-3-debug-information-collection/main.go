package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

type DebugInfo struct {
	Timestamp time.Time              `json:"timestamp"`
	RequestID string                 `json:"request_id"`
	Step      string                 `json:"step"`
	Duration  time.Duration          `json:"duration"`
	Input     map[string]interface{} `json:"input,omitempty"`
	Output    map[string]interface{} `json:"output,omitempty"`
	Error     string                 `json:"error,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type DebugCollector struct {
	requestID string
	steps     []DebugInfo
	startTime time.Time
}

func NewDebugCollector(requestID string) *DebugCollector {
	return &DebugCollector{
		requestID: requestID,
		startTime: time.Now(),
		steps:     make([]DebugInfo, 0),
	}
}

// TODO: Implement RecordStep(step string, input, output map[string]interface{}, err error, metadata map[string]interface{}) that:
// 1. Creates a DebugInfo entry with all provided information
// 2. Calculates duration since last step (or start time for first step)
// 3. Adds the step to the collector
// 4. Logs the step with slog.DebugContext

// TODO: Implement GenerateReport(ctx context.Context, outputDir string) error that:
// 1. Creates the output directory if it doesn't exist
// 2. Marshals all collected debug info to JSON
// 3. Writes to a file named: debug_report_{requestID}_{timestamp}.json
// 4. Logs the report generation with file path
// 5. Returns any errors with proper wrapping

type WorkflowProcessor struct {
	debugEnabled bool
}

func NewWorkflowProcessor(debugEnabled bool) *WorkflowProcessor {
	return &WorkflowProcessor{debugEnabled: debugEnabled}
}

// TODO: Implement ProcessWorkflow(ctx context.Context, requestID string, workflowData map[string]interface{}) error that:
// 1. Creates a debug collector if debugging is enabled
// 2. Simulates a multi-step workflow:
//    - "validate_input": Check if workflowData has required fields
//    - "transform_data": Convert data (simulate processing)
//    - "call_external_service": Simulate API call (can fail)
//    - "save_results": Save processed data
// 3. Records each step in the debug collector
// 4. On completion or error, generates debug report
// 5. Uses proper error handling and logging throughout

func simulateExternalAPICall(data map[string]interface{}) (map[string]interface{}, error) {
	// Simulate API call that sometimes fails
	if data["fail"] == true {
		return nil, errors.New("external service unavailable")
	}

	time.Sleep(100 * time.Millisecond) // Simulate network delay

	return map[string]interface{}{
		"api_result":       "success",
		"processed_count":  len(data),
		"api_timestamp":    time.Now(),
	}, nil
}

func main() {
	processor := NewWorkflowProcessor(true) // Enable debugging

	testCases := []struct {
		name         string
		requestID    string
		workflowData map[string]interface{}
	}{
		{
			name:      "successful workflow",
			requestID: "req-success-1",
			workflowData: map[string]interface{}{
				"user_id": "user123",
				"action":  "process",
				"data":    []string{"item1", "item2"},
			},
		},
		{
			name:      "workflow with API failure",
			requestID: "req-fail-1",
			workflowData: map[string]interface{}{
				"user_id": "user456",
				"action":  "process",
				"fail":    true, // This will cause API call to fail
			},
		},
		{
			name:      "workflow with missing data",
			requestID: "req-invalid-1",
			workflowData: map[string]interface{}{
				"action": "process",
				// Missing user_id
			},
		},
	}

	for _, tc := range testCases {
		fmt.Printf("\n=== Testing: %s ===\n", tc.name)

		ctx := context.Background()
		err := processor.ProcessWorkflow(ctx, tc.requestID, tc.workflowData)

		if err != nil {
			slog.ErrorContext(ctx, "workflow failed",
				"requestID", tc.requestID,
				"error", err)
		} else {
			slog.InfoContext(ctx, "workflow completed successfully",
				"requestID", tc.requestID)
		}
	}

	fmt.Println("\nCheck the generated debug reports in the current directory!")
}

// Answering the questions:
// 1. How does debug information collection help with troubleshooting client issues?
//   -
//
// 2. Why include timing information in debug reports?
//   -
//
// 3. When should debug collection be enabled vs disabled?
//   -
