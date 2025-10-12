package main

import (
	"context"
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
