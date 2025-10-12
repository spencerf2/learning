package main

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

// Answering the questions:
// 1. Why use exponential backoff instead of fixed delays?
//   -
//
// 2. Which HTTP status codes should trigger retries vs immediate failure?
//   -
//
// 3. How does context cancellation interact with retry logic?
//   -
