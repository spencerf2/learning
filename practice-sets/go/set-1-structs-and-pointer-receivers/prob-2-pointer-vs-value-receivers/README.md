# Problem 2: Pointer vs Value Receivers

**Goal**: Understand why *Counter vs Counter matters

```go
type Counter struct {
    value int
}

// TODO: Try this version first (value receiver):
func (c Counter) IncrementValue() {
    c.value++
}

// TODO: Then try this version (pointer receiver):
func (c *Counter) IncrementPointer() {
    c.value++
}

func main() {
    c := Counter{value: 0}
    
    c.IncrementValue()
    fmt.Println(c.value) // What prints here? Why?
    
    c.IncrementPointer()
    fmt.Println(c.value) // What prints here? Why?
}
```

## Questions to answer:

- Why does IncrementValue() not change the counter?
- Why does IncrementPointer() change the counter?
- When would you use each approach?
