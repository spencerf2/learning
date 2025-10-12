# Problem 2: Structured Logging with slog (Your Work Pattern)

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

## Questions to answer

- Why use structured logging instead of printf-style logging?
- How do default attributes (like processorID) help with debugging?
- When should you use InfoContext vs WarnContext vs ErrorContext?
