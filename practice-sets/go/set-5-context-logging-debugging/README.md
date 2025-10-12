# Go Practice Set 5: Context, Logging, and Debugging Patterns

Create a new directory: go-learning/context-logging-debugging/ and work through these problems in order.

## Problem 1: Context Fundamentals and Cancellation

**Goal**: Understand how context flows through your application

```go
import (
    "context"
    "fmt"
    "log/slog"
    "time"
)

type DataProcessor struct {
    name string
}

func NewDataProcessor(name string) *DataProcessor {
    return &DataProcessor{name: name}
}

// TODO: Implement ProcessData(ctx context.Context, data []string) error that:
// 1. Logs start of processing with slog.InfoContext
// 2. Processes each item in data with a 500ms delay
// 3. Checks ctx.Done() before each item to handle cancellation
// 4. If cancelled, logs cancellation and returns ctx.Err()
// 5. Logs each processed item with slog.InfoContext
// 6. Logs completion

// TODO: Implement ProcessWithTimeout(data []string, timeout time.Duration) error that:
// 1. Creates a context with timeout
// 2. Calls ProcessData with that context
// 3. Handles timeout errors specifically
// 4. Returns meaningful error messages

// TODO: Implement ProcessWithCancellation(data []string) error that:
// 1. Creates a cancellable context
// 2. Starts ProcessData in a goroutine
// 3. Cancels after 3 items are processed (simulate user cancellation)
// 4. Waits for the goroutine to finish
// 5. Returns any error

func main() {
    processor := NewDataProcessor("main-processor")
    
    testData := []string{"item1", "item2", "item3", "item4", "item5", "item6"}
    
    fmt.Println("=== Testing with timeout ===")
    if err := processor.ProcessWithTimeout(testData, 2*time.Second); err != nil {
        fmt.Printf("Timeout test error: %v\n", err)
    }
    
    fmt.Println("\n=== Testing with cancellation ===")
    if err := processor.ProcessWithCancellation(testData); err != nil {
        fmt.Printf("Cancellation test error: %v\n", err)
    }
}
```

### Questions to answer

- Why pass context as the first parameter to functions?
- What's the difference between context.WithTimeout and context.WithDeadline?
- How does checking ctx.Done() help with responsive cancellation?

## Problem 2: Structured Logging with slog (Your Work Pattern)

**Goal**: Master the slog patterns you see in production code

```go
import (
    "context"
    "errors"
    "fmt"
    "log/slog"
    "os"
    "time"
)

type RequestProcessor struct {
    processorID string
    logger      *slog.Logger
}

func NewRequestProcessor(processorID string) *RequestProcessor {
    // TODO: Create a structured logger with:
    // 1. JSON handler for production-like output
    // 2. Include processorID in all log entries as a default attribute
    // 3. Set log level to Info
    
    return &RequestProcessor{
        processorID: processorID,
        // Set logger field
    }
}

type ProcessingRequest struct {
    RequestID   string
    UserID      string
    Action      string
    Data        map[string]interface{}
    ProcessedAt time.Time
}

// TODO: Implement ProcessRequest(ctx context.Context, req ProcessingRequest) error that:
// 1. Logs request start with: requestID, userID, action
// 2. Simulates processing steps with intermediate logging:
//    - "validating request" with validation duration
//    - "processing data" with data size
//    - "saving results" with save duration
// 3. Uses slog.InfoContext for normal flow
// 4. Uses slog.WarnContext for warnings (like empty data)
// 5. Uses slog.ErrorContext for errors (like invalid userID)
// 6. Includes timing information in logs
// 7. Returns errors for invalid requests

// TODO: Implement ProcessBatch(ctx context.Context, requests []ProcessingRequest) error that:
// 1. Logs batch start with batch size
// 2. Processes each request, continuing on individual failures
// 3. Tracks success/failure counts
// 4. Logs progress every 10 requests
// 5. Logs final batch summary with counts and total duration

// TODO: Add log correlation - create a function AddCorrelationID(ctx context.Context, correlationID string) context.Context
// that stores correlation ID in context for request tracing

func main() {
    processor := NewRequestProcessor("worker-1")
    
    // Create test requests
    requests := []ProcessingRequest{
        {
            RequestID: "req-1",
            UserID:    "user-123",
            Action:    "create",
            Data:      map[string]interface{}{"name": "test", "value": 42},
        },
        {
            RequestID: "req-2", 
            UserID:    "", // Invalid - will cause error
            Action:    "update",
            Data:      map[string]interface{}{},
        },
        {
            RequestID: "req-3",
            UserID:    "user-456", 
            Action:    "delete",
            Data:      nil, // Will cause warning
        },
    }
    
    // TODO: Process individual requests with correlation IDs
    for _, req := range requests {
        ctx := context.Background()
        // Add correlation ID to context
        // Process the request
        // Handle any errors
    }
    
    fmt.Println("\n=== Processing as batch ===")
    ctx := context.Background()
    if err := processor.ProcessBatch(ctx, requests); err != nil {
        slog.ErrorContext(ctx, "batch processing failed", "error", err)
    }
}
```

### Questions to answer

- Why use structured logging instead of printf-style logging?
- How do default attributes (like processorID) help with debugging?
- When should you use InfoContext vs WarnContext vs ErrorContext?

## Problem 3: Debug Information Collection (Like Your Work Feature)

**Goal**: Build debug report generation similar to your current work

```go
import (
    "context"
    "encoding/json"
    "fmt"
    "log/slog"
    "os"
    "path/filepath"
    "time"
)

type DebugInfo struct {
    Timestamp   time.Time              `json:"timestamp"`
    RequestID   string                 `json:"request_id"`
    Step        string                 `json:"step"`
    Duration    time.Duration          `json:"duration"`
    Input       map[string]interface{} `json:"input,omitempty"`
    Output      map[string]interface{} `json:"output,omitempty"`
    Error       string                 `json:"error,omitempty"`
    Metadata    map[string]interface{} `json:"metadata,omitempty"`
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
        "api_result": "success",
        "processed_count": len(data),
        "api_timestamp": time.Now(),
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
```

### Questions to answer

- How does debug information collection help with troubleshooting client issues?
- Why include timing information in debug reports?
- When should debug collection be enabled vs disabled?

## Problem 4: Production Debugging Patterns

**Goal**: Debug complex issues like those you encounter with clients

```go
import (
    "context"
    "fmt"
    "log/slog"
    "os"
    "runtime"
    "sync"
    "time"
)

type SystemMonitor struct {
    metrics map[string]interface{}
    mu      sync.RWMutex
    logger  *slog.Logger
}

func NewSystemMonitor() *SystemMonitor {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelDebug,
    }))
    
    return &SystemMonitor{
        metrics: make(map[string]interface{}),
        logger:  logger,
    }
}

// TODO: Implement RecordMetric(key string, value interface{}) that:
// 1. Thread-safely stores the metric
// 2. Logs the metric update with timestamp
// 3. If value is numeric, also log the change from previous value

// TODO: Implement GetMetrics() map[string]interface{} that:
// 1. Thread-safely returns a copy of all metrics
// 2. Includes system metrics: goroutine count, memory usage

// TODO: Implement DumpSystemState(ctx context.Context, reason string) that:
// 1. Collects all current metrics
// 2. Adds runtime information (goroutines, memory, GC stats)
// 3. Logs everything as a single structured log entry
// 4. Includes the reason for the dump

type Application struct {
    monitor   *SystemMonitor
    isRunning bool
    mu        sync.RWMutex
}

func NewApplication() *Application {
    return &Application{
        monitor:   NewSystemMonitor(),
        isRunning: false,
    }
}

// TODO: Implement ProcessTask(ctx context.Context, taskID string, complexity int) error that:
// 1. Records start time and increments active task count
// 2. Simulates work based on complexity (more complexity = longer processing)
// 3. Records metrics: task duration, memory before/after
// 4. Decrements active task count on completion
// 5. If processing takes longer than 5 seconds, dumps system state
// 6. Uses proper logging throughout

// TODO: Implement RunApplication(ctx context.Context, taskCount int) error that:
// 1. Starts the application
// 2. Processes multiple concurrent tasks
// 3. Monitors system health every 2 seconds
// 4. Dumps system state if any concerning metrics are detected
// 5. Gracefully shuts down on context cancellation

func (a *Application) healthCheck(ctx context.Context) {
    for {
        select {
        case <-ctx.Done():
            return
        case <-time.After(2 * time.Second):
            // TODO: Check metrics and dump state if needed
            // Look for: high memory usage, too many goroutines, slow tasks
        }
    }
}

func simulateMemoryLeak() {
    // Intentionally create a small memory leak for testing
    data := make([]byte, 1024*1024) // 1MB
    _ = data
}

func main() {
    app := NewApplication()
    
    // Test with increasing task complexity to trigger monitoring
    complexities := []int{1, 3, 5, 8, 10} // Last ones will trigger long processing
    
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    // TODO: Start health monitoring in background
    
    // TODO: Process tasks with different complexities
    for i, complexity := range complexities {
        taskID := fmt.Sprintf("task-%d", i+1)
        
        // Simulate memory leak on some tasks
        if i%2 == 0 {
            simulateMemoryLeak()
        }
        
        if err := app.ProcessTask(ctx, taskID, complexity); err != nil {
            slog.ErrorContext(ctx, "task failed", "taskID", taskID, "error", err)
        }
        
        time.Sleep(500 * time.Millisecond) // Brief pause between tasks
    }
    
    // Final system state dump
    app.monitor.DumpSystemState(ctx, "application_shutdown")
    
    fmt.Println("Check the logs for system monitoring and debug information!")
}
```

### Questions to answer

- How do you detect when a system is behaving abnormally?
- Why include runtime metrics (goroutines, memory) in debug information?
- When should you automatically dump system state vs wait for manual triggers?
