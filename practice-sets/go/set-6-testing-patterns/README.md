# Go Practice Set 6: Testing Patterns for Production Code

Create a new directory: go-learning/testing-patterns/ and work through these problems in order.

## Problem 1: Basic Unit Testing and Table-Driven Tests

**Goal**: Master Go's testing fundamentals and the table-driven pattern

```go
// File: calculator.go
package main

import (
    "errors"
    "fmt"
)

type Calculator struct {
    precision int
}

func NewCalculator(precision int) *Calculator {
    return &Calculator{precision: precision}
}

// TODO: Implement Add(a, b float64) float64
// TODO: Implement Divide(a, b float64) (float64, error) - return error if b is 0
// TODO: Implement Percentage(value, percent float64) float64

// File: calculator_test.go
package main

import (
    "math"
    "testing"
)

// TODO: Write TestCalculator_Add that uses table-driven tests with these cases:
// {name: "positive numbers", a: 2.5, b: 3.7, want: 6.2}
// {name: "negative numbers", a: -1.5, b: -2.3, want: -3.8}
// {name: "zero values", a: 0, b: 5.5, want: 5.5}
// {name: "large numbers", a: 1000000.1, b: 2000000.2, want: 3000000.3}

// TODO: Write TestCalculator_Divide that tests:
// - Normal division
// - Division by zero (should return error)
// - Division with precision handling
// Use table-driven tests with struct: {name, a, b, want, wantErr}

// TODO: Write TestCalculator_Percentage that tests:
// - 50% of 100 = 50
// - 25% of 200 = 50  
// - 100% of 75 = 75
// - 0% of anything = 0

func TestMain(m *testing.M) {
    // TODO: Add setup/teardown if needed
    // Run tests and exit with result
}

// Helper function for comparing floats
func floatEquals(a, b, tolerance float64) bool {
    return math.Abs(a-b) < tolerance
}
```

### Questions to answer

- Why use table-driven tests instead of separate test functions?
- How do you handle floating-point comparison in tests?
- When should you use t.Error() vs t.Fatal()?

## Problem 2: Testing HTTP Clients and Services (Your Work Pattern)

**Goal**: Test the HTTP service communication patterns from your codebase

```go
// File: api_client.go
package main

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type APIClient struct {
    baseURL string
    client  HTTPClientInterface
}

// Interface for testing - allows mocking
type HTTPClientInterface interface {
    Do(req *http.Request) (*http.Response, error)
}

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

type APIResponse struct {
    Success bool `json:"success"`
    Data    User `json:"data,omitempty"`
    Error   string `json:"error,omitempty"`
}

func NewAPIClient(baseURL string, client HTTPClientInterface) *APIClient {
    if client == nil {
        client = &http.Client{Timeout: 30 * time.Second}
    }
    return &APIClient{
        baseURL: baseURL,
        client:  client,
    }
}

// TODO: Implement GetUser(ctx context.Context, userID int) (*User, error)
// TODO: Implement CreateUser(ctx context.Context, user User) (*User, error)

// File: api_client_test.go
package main

import (
    "bytes"
    "context"
    "encoding/json"
    "errors"
    "io"
    "net/http"
    "strings"
    "testing"
)

// Mock HTTP client for testing
type MockHTTPClient struct {
    DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
    if m.DoFunc != nil {
        return m.DoFunc(req)
    }
    return &http.Response{}, nil
}

// TODO: Write TestAPIClient_GetUser that tests:
// - Successful user retrieval
// - User not found (404)
// - Server error (500)
// - Network error
// - Invalid JSON response
// Use the MockHTTPClient to simulate different responses

// TODO: Write TestAPIClient_CreateUser that tests:
// - Successful user creation
// - Validation error (400)
// - Server error (500)
// - Request timeout
// Check that the request body contains correct JSON

// Helper function to create mock responses
func createMockResponse(statusCode int, body interface{}) *http.Response {
    var bodyReader io.Reader
    
    if body != nil {
        jsonBody, _ := json.Marshal(body)
        bodyReader = bytes.NewReader(jsonBody)
    } else {
        bodyReader = strings.NewReader("")
    }
    
    return &http.Response{
        StatusCode: statusCode,
        Body:       io.NopCloser(bodyReader),
        Header:     make(http.Header),
    }
}

// TODO: Write benchmark test BenchmarkAPIClient_GetUser
// TODO: Write example test ExampleAPIClient_GetUser
```

### Questions to answer

- Why use interfaces for HTTP clients in testing?
- How do you test error conditions in HTTP clients?
- What should you verify when testing HTTP requests?

## Problem 3: Testing with Dependencies and Mocking (Service Integration)

**Goal**: Test complex service interactions like your Go ↔ Python orchestration

```go
// File: orchestrator.go
package main

import (
    "context"
    "encoding/json"
    "fmt"
    "log/slog"
)

type ProcessingRequest struct {
    ID   string                 `json:"id"`
    Data map[string]interface{} `json:"data"`
}

type ProcessingResult struct {
    ID      string                 `json:"id"`
    Success bool                   `json:"success"`
    Output  map[string]interface{} `json:"output,omitempty"`
    Error   string                 `json:"error,omitempty"`
}

// Interfaces for dependencies
type NLPService interface {
    ProcessText(ctx context.Context, text string) (*ProcessingResult, error)
}

type DatabaseService interface {
    SaveResult(ctx context.Context, result ProcessingResult) error
    GetRequest(ctx context.Context, requestID string) (*ProcessingRequest, error)
}

type Logger interface {
    InfoContext(ctx context.Context, msg string, args ...any)
    ErrorContext(ctx context.Context, msg string, args ...any)
}

type Orchestrator struct {
    nlpService NLPService
    dbService  DatabaseService
    logger     Logger
}

func NewOrchestrator(nlp NLPService, db DatabaseService, logger Logger) *Orchestrator {
    return &Orchestrator{
        nlpService: nlp,
        dbService:  db,
        logger:     logger,
    }
}

// TODO: Implement ProcessRequest(ctx context.Context, requestID string) error that:
// 1. Gets request from database
// 2. Extracts text from request data
// 3. Calls NLP service to process text
// 4. Saves result to database
// 5. Logs each step
// 6. Returns any errors with proper wrapping

// File: orchestrator_test.go
package main

import (
    "context"
    "errors"
    "testing"
)

// Mock implementations
type MockNLPService struct {
    ProcessTextFunc func(ctx context.Context, text string) (*ProcessingResult, error)
    CallCount       int
}

func (m *MockNLPService) ProcessText(ctx context.Context, text string) (*ProcessingResult, error) {
    m.CallCount++
    if m.ProcessTextFunc != nil {
        return m.ProcessTextFunc(ctx, text)
    }
    return &ProcessingResult{Success: true}, nil
}

type MockDatabaseService struct {
    SaveResultFunc func(ctx context.Context, result ProcessingResult) error
    GetRequestFunc func(ctx context.Context, requestID string) (*ProcessingRequest, error)
    SaveCallCount  int
    GetCallCount   int
}

func (m *MockDatabaseService) SaveResult(ctx context.Context, result ProcessingResult) error {
    m.SaveCallCount++
    if m.SaveResultFunc != nil {
        return m.SaveResultFunc(ctx, result)
    }
    return nil
}

func (m *MockDatabaseService) GetRequest(ctx context.Context, requestID string) (*ProcessingRequest, error) {
    m.GetCallCount++
    if m.GetRequestFunc != nil {
        return m.GetRequestFunc(ctx, requestID)
    }
    return &ProcessingRequest{ID: requestID}, nil
}

type MockLogger struct {
    InfoCalls  []LogCall
    ErrorCalls []LogCall
}

type LogCall struct {
    Message string
    Args    []any
}

func (m *MockLogger) InfoContext(ctx context.Context, msg string, args ...any) {
    m.InfoCalls = append(m.InfoCalls, LogCall{Message: msg, Args: args})
}

func (m *MockLogger) ErrorContext(ctx context.Context, msg string, args ...any) {
    m.ErrorCalls = append(m.ErrorCalls, LogCall{Message: msg, Args: args})
}

// TODO: Write TestOrchestrator_ProcessRequest_Success that:
// - Sets up mocks for successful flow
// - Verifies all services are called with correct parameters
// - Checks that result is saved with expected data
// - Verifies appropriate logging occurred

// TODO: Write TestOrchestrator_ProcessRequest_DatabaseError that:
// - Mocks database GetRequest to return error
// - Verifies error is properly handled and logged
// - Ensures NLP service is not called
// - Checks error message contains proper context

// TODO: Write TestOrchestrator_ProcessRequest_NLPError that:
// - Mocks NLP service to return error
// - Verifies error handling and logging
// - Ensures database SaveResult is not called

// TODO: Write TestOrchestrator_ProcessRequest_SaveError that:
// - Mocks SaveResult to return error
// - Verifies all upstream calls succeeded
// - Checks error handling for save failure

func TestOrchestrator_ProcessRequest_CallCounts(t *testing.T) {
    // TODO: Test that services are called the expected number of times
    // Verify no unexpected calls are made
}
```

### Questions to answer

- Why use interfaces for dependencies instead of concrete types?
- How do you verify that mocked functions are called with correct parameters?
- When should you test call counts vs just behavior?

## Problem 4: Integration Testing and Test Helpers

**Goal**: Test complete workflows and create reusable test utilities

```go
// File: integration_test.go
package main

import (
    "context"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "os"
    "testing"
    "time"
)

// Test server setup
type TestServer struct {
    server     *httptest.Server
    responses  map[string]TestResponse
    callCounts map[string]int
}

type TestResponse struct {
    StatusCode int
    Body       interface{}
    Delay      time.Duration
}

func NewTestServer() *TestServer {
    ts := &TestServer{
        responses:  make(map[string]TestResponse),
        callCounts: make(map[string]int),
    }
    
    ts.server = httptest.NewServer(http.HandlerFunc(ts.handleRequest))
    return ts
}

func (ts *TestServer) handleRequest(w http.ResponseWriter, r *http.Request) {
    // TODO: Implement request handler that:
    // 1. Tracks call counts by endpoint
    // 2. Returns configured responses
    // 3. Simulates delays if configured
    // 4. Logs requests for debugging
}

func (ts *TestServer) SetResponse(endpoint string, response TestResponse) {
    ts.responses[endpoint] = response
}

func (ts *TestServer) GetCallCount(endpoint string) int {
    return ts.callCounts[endpoint]
}

func (ts *TestServer) Close() {
    ts.server.Close()
}

func (ts *TestServer) URL() string {
    return ts.server.URL
}

// Test database setup
type TestDatabase struct {
    requests map[string]ProcessingRequest
    results  map[string]ProcessingResult
}

func NewTestDatabase() *TestDatabase {
    return &TestDatabase{
        requests: make(map[string]ProcessingRequest),
        results:  make(map[string]ProcessingResult),
    }
}

func (db *TestDatabase) SaveResult(ctx context.Context, result ProcessingResult) error {
    db.results[result.ID] = result
    return nil
}

func (db *TestDatabase) GetRequest(ctx context.Context, requestID string) (*ProcessingRequest, error) {
    if req, exists := db.requests[requestID]; exists {
        return &req, nil
    }
    return nil, errors.New("request not found")
}

func (db *TestDatabase) AddTestRequest(req ProcessingRequest) {
    db.requests[req.ID] = req
}

func (db *TestDatabase) GetSavedResult(requestID string) (ProcessingResult, bool) {
    result, exists := db.results[requestID]
    return result, exists
}

// Integration test
func TestOrchestrator_Integration(t *testing.T) {
    // TODO: Set up test server and database
    // TODO: Create real orchestrator with test dependencies
    // TODO: Test complete workflow:
    //   - Add test request to database
    //   - Configure NLP service response
    //   - Process request
    //   - Verify result is saved correctly
    //   - Check all service interactions
}

// TODO: Write TestOrchestrator_Integration_SlowNLP that:
// - Configures NLP service with long delay
// - Tests timeout handling
// - Verifies graceful failure

// TODO: Write TestOrchestrator_Integration_NetworkFailure that:
// - Simulates network failures at different points
// - Tests retry logic if implemented
// - Verifies error handling

// Test helpers
func setupTestEnvironment(t *testing.T) (*TestServer, *TestDatabase, func()) {
    // TODO: Set up test server and database
    // TODO: Return cleanup function
    // TODO: Log setup for debugging
}

func createTestRequest(id string, text string) ProcessingRequest {
    return ProcessingRequest{
        ID: id,
        Data: map[string]interface{}{
            "text": text,
            "timestamp": time.Now(),
        },
    }
}

func assertResultSaved(t *testing.T, db *TestDatabase, requestID string, expectSuccess bool) {
    result, exists := db.GetSavedResult(requestID)
    if !exists {
        t.Fatalf("Expected result to be saved for request %s", requestID)
    }
    
    if result.Success != expectSuccess {
        t.Errorf("Expected success=%v, got success=%v", expectSuccess, result.Success)
    }
}

// Performance test
func TestOrchestrator_Performance(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping performance test in short mode")
    }
    
    // TODO: Test processing many requests
    // TODO: Measure timing and memory usage
    // TODO: Verify no memory leaks
    // TODO: Check concurrent processing
}

// TODO: Write TestMain that:
// - Sets up test database
// - Configures test logging
// - Runs tests
// - Cleans up resources
```

### Questions to answer

- When should you use integration tests vs unit tests?
- How do you make integration tests reliable and fast?
- Why use test helpers and what should they provide?
