# Problem 2: Service-to-Service Communication (Go ↔ Python Pattern)

**Goal**: Mirror your work's orchestration pattern

```go
import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "log/slog"
    "net/http"
    "time"
)

// Represents a request to your Python NLP service
type NLPRequest struct {
    Text     string            `json:"text"`
    Model    string            `json:"model"`
    Options  map[string]string `json:"options,omitempty"`
}

// Represents response from Python service
type NLPResponse struct {
    ProcessedText string  `json:"processed_text"`
    Confidence    float64 `json:"confidence"`
    Metadata      map[string]interface{} `json:"metadata,omitempty"`
    Error         string  `json:"error,omitempty"`
}

type PythonNLPClient struct {
    baseURL string
    client  *http.Client
}

func NewPythonNLPClient(baseURL string) *PythonNLPClient {
    return &PythonNLPClient{
        baseURL: baseURL,
        client: &http.Client{
            Timeout: 60 * time.Second, // NLP can be slow
        },
    }
}

// TODO: Implement ProcessText(ctx context.Context, req NLPRequest) (*NLPResponse, error) that:
// 1. Logs the request with slog.InfoContext
// 2. Marshals the request to JSON
// 3. Makes POST request to {baseURL}/process
// 4. Handles both success and error responses
// 5. If the response has an Error field, return that as a Go error
// 6. Logs the successful response
// 7. Wraps any errors with meaningful context

// Orchestrator that coordinates multiple service calls
type Orchestrator struct {
    nlpClient *PythonNLPClient
}

func NewOrchestrator(nlpServiceURL string) *Orchestrator {
    return &Orchestrator{
        nlpClient: NewPythonNLPClient(nlpServiceURL),
    }
}

// TODO: Implement ProcessUserContent(ctx context.Context, userText string) error that:
// 1. Logs the start of processing
// 2. Calls the NLP service with different models: "sentiment" and "summarization"
// 3. Logs each service response
// 4. Returns any errors with proper wrapping
// 5. Logs successful completion

func main() {
    orch := NewOrchestrator("http://localhost:8080")
    ctx := context.Background()

    testTexts := []string{
        "This is a great product!",
        "I'm not sure about this...",
        "", // Test empty text
    }

    for _, text := range testTexts {
        fmt.Printf("\n--- Processing: '%s' ---\n", text)
        if err := orch.ProcessUserContent(ctx, text); err != nil {
            slog.ErrorContext(ctx, "processing failed", "text", text, "error", err)
        }
    }
}

```

## Questions to answer

- Why use longer timeouts for NLP service calls?
- How do you handle service errors vs HTTP errors?
- Why log both requests and responses in service communication?
