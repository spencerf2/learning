package main

import (
	"bytes"
	"context"
	"encoding/json"
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

// Answering the questions:
// 1. Why use interfaces for HTTP clients in testing?
//   -
//
// 2. How do you test error conditions in HTTP clients?
//   -
//
// 3. What should you verify when testing HTTP requests?
//   -
