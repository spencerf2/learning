package main

import (
	"context"
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
	Success bool   `json:"success"`
	Data    User   `json:"data,omitempty"`
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
