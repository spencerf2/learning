# Go Practice Set 4: HTTP Clients and Service Communication

Create a new directory: go-learning/http-service-communication/ and work through these problems in order.

## Problem 1: Basic HTTP Client and JSON Handling

**Goal**: Understand Go's HTTP client patterns and JSON marshaling

```go
import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

type APIResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Data    User   `json:"data,omitempty"`
}

type HTTPClient struct {
    baseURL string
    client  *http.Client
}

func NewHTTPClient(baseURL string) *HTTPClient {
    return &HTTPClient{
        baseURL: baseURL,
        client: &http.Client{
            Timeout: 30 * time.Second,
        },
    }
}

// TODO: Implement GetUser(ctx context.Context, userID int) (*User, error) that:
// 1. Creates a GET request to {baseURL}/users/{userID}
// 2. Sets the request context
// 3. Makes the HTTP call
// 4. Reads and unmarshals the JSON response into APIResponse
// 5. Returns the User data or an error with proper context

// TODO: Implement CreateUser(ctx context.Context, user User) (*User, error) that:
// 1. Marshals the user to JSON
// 2. Creates a POST request to {baseURL}/users with JSON body
// 3. Sets Content-Type header to "application/json"
// 4. Makes the HTTP call and handles the response

func main() {
    client := NewHTTPClient("https://jsonplaceholder.typicode.com")
    ctx := context.Background()
    
    // TODO: Test GetUser with ID 1
    // TODO: Test CreateUser with a new user
    // Handle and print any errors with full context
}
```

### Questions to answer

- Why use context.Context with HTTP requests?
- What's the purpose of struct tags like json:"id"?
- Why set a timeout on the HTTP client?

## Problem 2: Service-to-Service Communication (Go ↔ Python Pattern)

**Goal**: Mirror your work's orchestration pattern

```go
import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "log/slog"
    "net/http"
    "time"
)

// Represents a request to your Python NLP service
type NLPRequest struct {
    Text     string            `json:"text"`
    Model    string            `json:"model"`
    Options  map[string]string `json:"options,omitempty"`
}

// Represents response from Python service
type NLPResponse struct {
    ProcessedText string  `json:"processed_text"`
    Confidence    float64 `json:"confidence"`
    Metadata      map[string]interface{} `json:"metadata,omitempty"`
    Error         string  `json:"error,omitempty"`
}

type PythonNLPClient struct {
    baseURL string
    client  *http.Client
}

func NewPythonNLPClient(baseURL string) *PythonNLPClient {
    return &PythonNLPClient{
        baseURL: baseURL,
        client: &http.Client{
            Timeout: 60 * time.Second, // NLP can be slow
        },
    }
}

// TODO: Implement ProcessText(ctx context.Context, req NLPRequest) (*NLPResponse, error) that:
// 1. Logs the request with slog.InfoContext
// 2. Marshals the request to JSON
// 3. Makes POST request to {baseURL}/process
// 4. Handles both success and error responses
// 5. If the response has an Error field, return that as a Go error
// 6. Logs the successful response
// 7. Wraps any errors with meaningful context

// Orchestrator that coordinates multiple service calls
type Orchestrator struct {
    nlpClient *PythonNLPClient
}

func NewOrchestrator(nlpServiceURL string) *Orchestrator {
    return &Orchestrator{
        nlpClient: NewPythonNLPClient(nlpServiceURL),
    }
}

// TODO: Implement ProcessUserContent(ctx context.Context, userText string) error that:
// 1. Logs the start of processing
// 2. Calls the NLP service with different models: "sentiment" and "summarization"
// 3. Logs each service response
// 4. Returns any errors with proper wrapping
// 5. Logs successful completion

func main() {
    orch := NewOrchestrator("http://localhost:8080")
    ctx := context.Background()
    
    testTexts := []string{
        "This is a great product!",
        "I'm not sure about this...",
        "", // Test empty text
    }
    
    for _, text := range testTexts {
        fmt.Printf("\n--- Processing: '%s' ---\n", text)
        if err := orch.ProcessUserContent(ctx, text); err != nil {
            slog.ErrorContext(ctx, "processing failed", "text", text, "error", err)
        }
    }
}
```

### Questions to answer

- Why use longer timeouts for NLP service calls?
- How do you handle service errors vs HTTP errors?
- Why log both requests and responses in service communication?

## Problem 3: Retry Logic and Circuit Breaker Pattern

**Goal**: Handle unreliable service communication (production reality)

```go
import (
    "context"
    "encoding/json"
    "fmt"
    "log/slog"
    "math/rand"
    "net/http"
    "time"
)

type RetryableClient struct {
    baseURL    string
    client     *http.Client
    maxRetries int
    baseDelay  time.Duration
}

func NewRetryableClient(baseURL string, maxRetries int) *RetryableClient {
    return &RetryableClient{
        baseURL:    baseURL,
        maxRetries: maxRetries,
        baseDelay:  100 * time.Millisecond,
        client: &http.Client{
            Timeout: 10 * time.Second,
        },
    }
}

type ServiceRequest struct {
    Data string `json:"data"`
}

type ServiceResponse struct {
    Result string `json:"result"`
    Status string `json:"status"`
}

// TODO: Implement shouldRetry(err error, statusCode int) bool that returns true for:
// - Network errors (err != nil)
// - 5xx status codes
// - 429 (rate limited)
// Returns false for 4xx errors (except 429)

// TODO: Implement CallServiceWithRetry(ctx context.Context, endpoint string, req ServiceRequest) (*ServiceResponse, error) that:
// 1. Tries the request up to maxRetries times
// 2. Uses exponential backoff: delay = baseDelay * (2^attempt) + jitter
// 3. Logs each retry attempt with attempt number
// 4. Checks if the error/status code is retryable
// 5. Returns the final error if all retries fail
// 6. Uses proper error wrapping

func (c *RetryableClient) calculateDelay(attempt int) time.Duration {
    // Exponential backoff with jitter
    delay := c.baseDelay * time.Duration(1<<attempt) // 2^attempt
    jitter := time.Duration(rand.Intn(100)) * time.Millisecond
    return delay + jitter
}

// Mock a flaky service for testing
type FlakeyService struct {
    successRate float64
    callCount   int
}

func (f *FlakeyService) HandleRequest() (int, ServiceResponse) {
    f.callCount++
    
    if rand.Float64() < f.successRate {
        return http.StatusOK, ServiceResponse{
            Result: fmt.Sprintf("Success on call %d", f.callCount),
            Status: "ok",
        }
    }
    
    // Randomly return different error types
    errorTypes := []int{http.StatusInternalServerError, http.StatusBadGateway, http.StatusTooManyRequests}
    statusCode := errorTypes[rand.Intn(len(errorTypes))]
    
    return statusCode, ServiceResponse{
        Result: "",
        Status: "error",
    }
}

func main() {
    client := NewRetryableClient("http://mock-service", 3)
    ctx := context.Background()
    
    // Test with different success rates
    successRates := []float64{0.0, 0.3, 0.8, 1.0}
    
    for _, rate := range successRates {
        fmt.Printf("\n--- Testing with %d%% success rate ---\n", int(rate*100))
        
        req := ServiceRequest{Data: "test data"}
        resp, err := client.CallServiceWithRetry(ctx, "/process", req)
        
        if err != nil {
            fmt.Printf("Final error: %v\n", err)
        } else {
            fmt.Printf("Success: %v\n", resp)
        }
    }
}
```

### Questions to answer

- Why use exponential backoff instead of fixed delays?
- Which HTTP status codes should trigger retries vs immediate failure?
- How does context cancellation interact with retry logic?

## Problem 4: Concurrent Service Calls and Coordination

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

### Questions to answer

- When should you use parallel vs sequential service calls?
- How do you handle partial failures in orchestration workflows?
- Why pass context.Context through the entire call chain?
