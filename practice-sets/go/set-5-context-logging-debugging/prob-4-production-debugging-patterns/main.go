package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"sync"
	"time"
)

type SystemMonitor struct {
	metrics map[string]interface{}
	mu      sync.RWMutex
	logger  *slog.Logger
}

func NewSystemMonitor() *SystemMonitor {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	return &SystemMonitor{
		metrics: make(map[string]interface{}),
		logger:  logger,
	}
}

// TODO: Implement RecordMetric(key string, value interface{}) that:
// 1. Thread-safely stores the metric
// 2. Logs the metric update with timestamp
// 3. If value is numeric, also log the change from previous value

// TODO: Implement GetMetrics() map[string]interface{} that:
// 1. Thread-safely returns a copy of all metrics
// 2. Includes system metrics: goroutine count, memory usage

// TODO: Implement DumpSystemState(ctx context.Context, reason string) that:
// 1. Collects all current metrics
// 2. Adds runtime information (goroutines, memory, GC stats)
// 3. Logs everything as a single structured log entry
// 4. Includes the reason for the dump

type Application struct {
	monitor   *SystemMonitor
	isRunning bool
	mu        sync.RWMutex
}

func NewApplication() *Application {
	return &Application{
		monitor:   NewSystemMonitor(),
		isRunning: false,
	}
}

// TODO: Implement ProcessTask(ctx context.Context, taskID string, complexity int) error that:
// 1. Records start time and increments active task count
// 2. Simulates work based on complexity (more complexity = longer processing)
// 3. Records metrics: task duration, memory before/after
// 4. Decrements active task count on completion
// 5. If processing takes longer than 5 seconds, dumps system state
// 6. Uses proper logging throughout

// TODO: Implement RunApplication(ctx context.Context, taskCount int) error that:
// 1. Starts the application
// 2. Processes multiple concurrent tasks
// 3. Monitors system health every 2 seconds
// 4. Dumps system state if any concerning metrics are detected
// 5. Gracefully shuts down on context cancellation

func (a *Application) healthCheck(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(2 * time.Second):
			// TODO: Check metrics and dump state if needed
			// Look for: high memory usage, too many goroutines, slow tasks
		}
	}
}

func simulateMemoryLeak() {
	// Intentionally create a small memory leak for testing
	data := make([]byte, 1024*1024) // 1MB
	_ = data
}

func main() {
	app := NewApplication()

	// Test with increasing task complexity to trigger monitoring
	complexities := []int{1, 3, 5, 8, 10} // Last ones will trigger long processing

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// TODO: Start health monitoring in background

	// TODO: Process tasks with different complexities
	for i, complexity := range complexities {
		taskID := fmt.Sprintf("task-%d", i+1)

		// Simulate memory leak on some tasks
		if i%2 == 0 {
			simulateMemoryLeak()
		}

		if err := app.ProcessTask(ctx, taskID, complexity); err != nil {
			slog.ErrorContext(ctx, "task failed", "taskID", taskID, "error", err)
		}

		time.Sleep(500 * time.Millisecond) // Brief pause between tasks
	}

	// Final system state dump
	app.monitor.DumpSystemState(ctx, "application_shutdown")

	fmt.Println("Check the logs for system monitoring and debug information!")
}

// Answering the questions:
// 1. How do you detect when a system is behaving abnormally?
//   -
//
// 2. Why include runtime metrics (goroutines, memory) in debug information?
//   -
//
// 3. When should you automatically dump system state vs wait for manual triggers?
//   -
