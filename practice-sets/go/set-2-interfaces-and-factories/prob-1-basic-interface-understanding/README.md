# Problem 1: Basic Interface Understanding

**Goal**: Understand how Go interfaces work (implicit satisfaction)

```go
// Define a simple interface
type Writer interface {
    Write(data string) error
}

// TODO: Create a FileWriter struct with a fileName field
// TODO: Add a Write method to FileWriter that prints: "Writing '[data]' to file: [fileName]"
// TODO: Create a ConsoleWriter struct (no fields needed)
// TODO: Add a Write method to ConsoleWriter that prints: "Console output: [data]"

func ProcessData(w Writer, data string) error {
    return w.Write(data)
}

func main() {
    // TODO: Create instances of both FileWriter and ConsoleWriter
    // TODO: Call ProcessData with each one
    // Notice: No explicit "implements" declaration needed!
}
```

## Questions to answer

- How does Go know that FileWriter and ConsoleWriter implement Writer?
- What happens if you remove the Write method from one of your structs?
- Why doesn't Go require explicit interface declarations?
