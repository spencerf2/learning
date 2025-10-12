package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"
)

type DataProcessor struct {
	name string
}

func NewDataProcessor(name string) *DataProcessor {
	return &DataProcessor{name: name}
}

// TODO: Implement ProcessData(ctx context.Context, data []string) error that:
// 1. Logs start of processing with slog.InfoContext
// 2. Processes each item in data with a 500ms delay
// 3. Checks ctx.Done() before each item to handle cancellation
// 4. If cancelled, logs cancellation and returns ctx.Err()
// 5. Logs each processed item with slog.InfoContext
// 6. Logs completion

// TODO: Implement ProcessWithTimeout(data []string, timeout time.Duration) error that:
// 1. Creates a context with timeout
// 2. Calls ProcessData with that context
// 3. Handles timeout errors specifically
// 4. Returns meaningful error messages

// TODO: Implement ProcessWithCancellation(data []string) error that:
// 1. Creates a cancellable context
// 2. Starts ProcessData in a goroutine
// 3. Cancels after 3 items are processed (simulate user cancellation)
// 4. Waits for the goroutine to finish
// 5. Returns any error

func main() {
	processor := NewDataProcessor("main-processor")

	testData := []string{"item1", "item2", "item3", "item4", "item5", "item6"}

	fmt.Println("=== Testing with timeout ===")
	if err := processor.ProcessWithTimeout(testData, 2*time.Second); err != nil {
		fmt.Printf("Timeout test error: %v\n", err)
	}

	fmt.Println("\n=== Testing with cancellation ===")
	if err := processor.ProcessWithCancellation(testData); err != nil {
		fmt.Printf("Cancellation test error: %v\n", err)
	}
}

// Answering the questions:
// 1. Why pass context as the first parameter to functions?
//   -
//
// 2. What's the difference between context.WithTimeout and context.WithDeadline?
//   -
//
// 3. How does checking ctx.Done() help with responsive cancellation?
//   -
