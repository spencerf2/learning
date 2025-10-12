package main

import (
	"errors"
	"fmt"
	"os"
)

// Simulate some common errors from your work environment
var (
	ErrInvalidConfig = errors.New("invalid configuration")
	ErrServiceDown   = errors.New("external service unavailable")
	ErrPermission    = errors.New("permission denied")
)

func connectToDatabase(config string) error {
	if config == "" {
		return fmt.Errorf("database connection failed: %w", ErrInvalidConfig)
	}
	if config == "prod" {
		return fmt.Errorf("database connection failed: %w", ErrServiceDown)
	}
	return nil
}

func loadUserData(userID string, config string) error {
	if err := connectToDatabase(config); err != nil {
		return fmt.Errorf("failed to load user %s: %w", userID, err)
	}
	return nil
}

func generateReport(userID string, config string) error {
	if err := loadUserData(userID, config); err != nil {
		return fmt.Errorf("report generation failed: %w", err)
	}
	return nil
}

func main() {
	testCases := []struct {
		userID string
		config string
		name   string
	}{
		{"user123", "", "empty config"},
		{"user456", "prod", "service down"},
		{"user789", "dev", "success case"},
	}

	for _, tc := range testCases {
		fmt.Printf("\nTesting %s:\n", tc.name)
		err := generateReport(tc.userID, tc.config)

		if err != nil {
			fmt.Printf("Error: %v\n", err)

			// TODO: Use errors.Is() to check if the root cause is ErrInvalidConfig
			// TODO: Use errors.Is() to check if the root cause is ErrServiceDown
			// TODO: Unwrap the error chain manually using errors.Unwrap() in a loop
		} else {
			fmt.Println("Success!")
		}
	}
}

// Answering the questions:
// 1. How do you check for specific error types in a wrapped error chain?
//   - 
//
// 2. Why is errors.Is() better than string comparison for error checking?
//   - 
//
// 3. When would you use errors.As() vs errors.Is()?
//   - 
