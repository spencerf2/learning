# Problem 4: Production Error Patterns (Context + Logging)

**Goal**: Mirror the logging patterns you see in your work codebase

```go
import (
    "context"
    "errors"
    "fmt"
    "log/slog"
    "os"
)

type APIClient struct {
    baseURL string
}

func NewAPIClient(baseURL string) *APIClient {
    return &APIClient{baseURL: baseURL}
}

// TODO: Implement FetchUserData(ctx context.Context, userID string) ([]byte, error) that:
// 1. Simulates an API call (just return nil, error for this exercise)
// 2. If userID is empty, return an error with proper context
// 3. If userID is "error", simulate a network error
// 4. Use slog.ErrorContext() to log failures before returning errors
// 5. Wrap errors with meaningful context using %w

func (a *APIClient) ProcessUserRequest(ctx context.Context, userID string) error {
    slog.InfoContext(ctx, "processing user request", "userID", userID)

    data, err := a.FetchUserData(ctx, userID)
    if err != nil {
        // TODO: Log the error with context and return wrapped error
        return fmt.Errorf("failed to process user request: %w", err)
    }

    // TODO: Simulate processing the data
    slog.InfoContext(ctx, "successfully processed user data",
        "userID", userID,
        "dataSize", len(data))

    return nil
}

func main() {
    // Setup structured logging
    logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
        Level: slog.LevelInfo,
    }))
    slog.SetDefault(logger)

    client := NewAPIClient("https://api.example.com")
    ctx := context.Background()

    testCases := []string{"", "error", "user123"}

    for _, userID := range testCases {
        fmt.Printf("\n--- Testing userID: '%s' ---\n", userID)
        err := client.ProcessUserRequest(ctx, userID)
        if err != nil {
            fmt.Printf("Final error: %v\n", err)
        }
    }
}
```

## Questions to answer

- Why use slog.ErrorContext() instead of just fmt.Printf() for errors?
- How does context help with error logging and tracing?
- When should you log an error vs just return it?
