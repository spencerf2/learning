# Problem 1: Basic HTTP Client and JSON Handling

**Goal**: Understand Go's HTTP client patterns and JSON marshaling

```go
import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

type APIResponse struct {
    Success bool   `json:"success"`
    Message string `json:"message"`
    Data    User   `json:"data,omitempty"`
}

type HTTPClient struct {
    baseURL string
    client  *http.Client
}

func NewHTTPClient(baseURL string) *HTTPClient {
    return &HTTPClient{
        baseURL: baseURL,
        client: &http.Client{
            Timeout: 30 * time.Second,
        },
    }
}

// TODO: Implement GetUser(ctx context.Context, userID int) (*User, error) that:
// 1. Creates a GET request to {baseURL}/users/{userID}
// 2. Sets the request context
// 3. Makes the HTTP call
// 4. Reads and unmarshals the JSON response into APIResponse
// 5. Returns the User data or an error with proper context

// TODO: Implement CreateUser(ctx context.Context, user User) (*User, error) that:
// 1. Marshals the user to JSON
// 2. Creates a POST request to {baseURL}/users with JSON body
// 3. Sets Content-Type header to "application/json"
// 4. Makes the HTTP call and handles the response

func main() {
    client := NewHTTPClient("https://jsonplaceholder.typicode.com")
    ctx := context.Background()

    // TODO: Test GetUser with ID 1
    // TODO: Test CreateUser with a new user
    // Handle and print any errors with full context
}
```

## Questions to answer

- Why use context.Context with HTTP requests?
- What's the purpose of struct tags like json:"id"?
- Why set a timeout on the HTTP client?
