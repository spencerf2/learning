# Problem 4: Concurrent Service Calls and Coordination

**Goal**: Handle multiple service calls efficiently (like your orchestration work)

```go
import (
    "context"
    "encoding/json"
    "fmt"
    "log/slog"
    "sync"
    "time"
)

type ServiceResult struct {
    ServiceName string
    Data        interface{}
    Error       error
    Duration    time.Duration
}

type MultiServiceClient struct {
    services map[string]string // service name -> URL
    client   *http.Client
}

func NewMultiServiceClient(services map[string]string) *MultiServiceClient {
    return &MultiServiceClient{
        services: services,
        client:   &http.Client{Timeout: 30 * time.Second},
    }
}

// TODO: Implement callService(ctx context.Context, serviceName, endpoint string, data interface{}) ServiceResult that:
// 1. Records start time
// 2. Makes the HTTP call to the service
// 3. Returns ServiceResult with timing info and any errors
// 4. Includes the service name for identification

// TODO: Implement CallServicesParallel(ctx context.Context, requests map[string]interface{}) map[string]ServiceResult that:
// 1. Calls multiple services concurrently using goroutines
// 2. Uses a WaitGroup to coordinate completion
// 3. Collects all results in a thread-safe way
// 4. Respects context cancellation
// 5. Logs the total operation time

// TODO: Implement CallServicesSequential(ctx context.Context, serviceOrder []string, data interface{}) (map[string]ServiceResult, error) that:
// 1. Calls services in order, stopping on first error
// 2. Each call can use data from previous calls
// 3. Returns partial results even if later calls fail
// 4. Logs progress through the sequence

type OrchestrationPipeline struct {
    client *MultiServiceClient
}

func NewOrchestrationPipeline() *OrchestrationPipeline {
    services := map[string]string{
        "validator":   "http://localhost:8001",
        "nlp":         "http://localhost:8002",
        "summarizer":  "http://localhost:8003",
        "classifier":  "http://localhost:8004",
    }

    return &OrchestrationPipeline{
        client: NewMultiServiceClient(services),
    }
}

// TODO: Implement ProcessDocument(ctx context.Context, document string) error that:
// 1. First validates the document
// 2. Then calls NLP, summarizer, and classifier in parallel
// 3. Logs timing for each phase
// 4. Handles partial failures gracefully
// 5. Returns detailed error info about which services failed

func main() {
    pipeline := NewOrchestrationPipeline()
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
    defer cancel()

    documents := []string{
        "This is a sample document for processing.",
        "Another document with different content.",
        "", // Test empty document
    }

    for i, doc := range documents {
        fmt.Printf("\n--- Processing Document %d ---\n", i+1)
        start := time.Now()

        err := pipeline.ProcessDocument(ctx, doc)
        duration := time.Since(start)

        if err != nil {
            slog.ErrorContext(ctx, "document processing failed",
                "document", doc,
                "duration", duration,
                "error", err)
        } else {
            slog.InfoContext(ctx, "document processing completed",
                "document", doc,
                "duration", duration)
        }
    }
}

```

## Questions to answer

- When should you use parallel vs sequential service calls?
- How do you handle partial failures in orchestration workflows?
- Why pass context.Context through the entire call chain?
