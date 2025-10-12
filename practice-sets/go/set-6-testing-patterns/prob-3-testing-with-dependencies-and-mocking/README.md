# Problem 3: Testing with Dependencies and Mocking (Service Integration)

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

## Questions to answer

- Why use interfaces for dependencies instead of concrete types?
- How do you verify that mocked functions are called with correct parameters?
- When should you test call counts vs just behavior?
