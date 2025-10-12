# Problem 2: Testing HTTP Clients and Services (Your Work Pattern)

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

## Questions to answer

- Why use interfaces for HTTP clients in testing?
- How do you test error conditions in HTTP clients?
- What should you verify when testing HTTP requests?
