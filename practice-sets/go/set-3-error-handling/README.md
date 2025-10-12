# Go Practice Set 3: Error Handling and Wrapping

Create a new directory: go-learning/error-handling/ and work through these problems in order.

## Problem 1: Basic Error Creation and Checking

**Goal**: Understand Go's error fundamentals

```go
import (
    "errors"
    "fmt"
)

type User struct {
    ID   int
    Name string
    Age  int
}

// TODO: Create a function ValidateUser(user User) error that returns:
// - An error if user.Name is empty (use errors.New())
// - An error if user.Age is less than 0 or greater than 150 (use fmt.Errorf())
// - nil if the user is valid

// TODO: Create a function SaveUser(user User) error that:
// 1. First calls ValidateUser
// 2. If validation fails, return the validation error
// 3. If validation passes, print "Saving user: [name]" and return nil

func main() {
    users := []User{
        {ID: 1, Name: "Alice", Age: 30},
        {ID: 2, Name: "", Age: 25},
        {ID: 3, Name: "Bob", Age: -5},
    }

    // TODO: Try to save each user and print any errors
    for _, user := range users {
        if err := SaveUser(user); err != nil {
            fmt.Printf("Error saving user %d: %v\n", user.ID, err)
        }
    }
}
```

### Questions to answer

- When should you use errors.New() vs fmt.Errorf()?
- Why does Go use explicit error returns instead of exceptions?
- What's the difference between %v and %s when printing errors?

## Problem 2: Error Wrapping with %w (Your Work Pattern)

**Goal**: Practice the exact pattern you encountered: fmt.Errorf("context: %w", err)

```go
import (
    "errors"
    "fmt"
    "os"
)

type FileProcessor struct {
    outputDir string
}

func NewFileProcessor(outputDir string) *FileProcessor {
    return &FileProcessor{outputDir: outputDir}
}

// TODO: Implement CreateOutputDirectory() error that:
// 1. Tries to create the directory using os.MkdirAll(f.outputDir, 0755)
// 2. If it fails, wrap the error with context: "failed to create output directory: %w"
// 3. Return nil if successful

// TODO: Implement WriteFile(filename string, data []byte) error that:
// 1. First calls CreateOutputDirectory()
// 2. If that fails, wrap with context: "failed to prepare output: %w"
// 3. Then tries to write file using os.WriteFile(filepath.Join(f.outputDir, filename), data, 0644)
// 4. If that fails, wrap with context: "failed to write file: %w"

func (f *FileProcessor) ProcessFile(filename string, data []byte) error {
    // TODO: Call WriteFile and wrap any error with: "failed to process file %s: %w"
    return nil
}

func main() {
    // Test with a directory that will cause permission errors
    processor := NewFileProcessor("/root/forbidden")

    err := processor.ProcessFile("test.txt", []byte("test data"))
    if err != nil {
        fmt.Printf("Full error: %v\n", err)

        // TODO: Use errors.Unwrap() to see the original error
        // TODO: Use errors.Is() to check if it's a permission error
    }
}
```

### Questions to answer

- What's the difference between %v and %w in fmt.Errorf?
- How does error wrapping help with debugging?
- What information gets lost if you use %v instead of %w?

## Problem 3: Error Chain Analysis and Unwrapping

**Goal**: Understand how to work with wrapped errors (debugging production issues)

```go
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
```

### Questions to answer

- How do you check for specific error types in a wrapped error chain?
- Why is errors.Is() better than string comparison for error checking?
- When would you use errors.As() vs errors.Is()?

## Problem 4: Production Error Patterns (Context + Logging)

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

### Questions to answer

- Why use slog.ErrorContext() instead of just fmt.Printf() for errors?
- How does context help with error logging and tracing?
- When should you log an error vs just return it?
