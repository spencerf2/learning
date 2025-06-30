# Problem 1: Basic Struct Methods

**Goal**: Understand the difference between methods and functions

```go
// Create a simple Counter struct
type Counter struct {
    value int
}

// TODO: Add a method called GetValue() that returns the current value
// TODO: Add a method called SetValue(newValue int) that updates the value
// TODO: Create a regular function (not a method) called CreateCounter(initialValue int) that returns a new Counter

func main() {
    // Test your code:
    c := CreateCounter(5)
    fmt.Println(c.GetValue()) // Should print 5
    c.SetValue(10)
    fmt.Println(c.GetValue()) // Should print 10
}
```

## Questions to answer

- What's the difference between func (c Counter) GetValue() and func GetValue(c Counter)?
- Why do we use methods instead of just functions?

