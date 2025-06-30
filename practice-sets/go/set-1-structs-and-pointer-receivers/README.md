# Go Practice Set 1: Structs and Pointer Receivers

Create a new directory: go-learning/structs-and-methods/ and work through these problems in order.

## Problem 1: Basic Struct Methods

**Goal**: Understand the difference between methods and functions
```
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

### Questions to answer:

- What's the difference between func (c Counter) GetValue() and func GetValue(c Counter)?
- Why do we use methods instead of just functions?

## Problem 2: Pointer vs Value Receivers

**Goal**: Understand why *Counter vs Counter matters
```
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

### Questions to answer:

- Why does IncrementValue() not change the counter?
- Why does IncrementPointer() change the counter?
- When would you use each approach?

## Problem 3: Constructor Pattern (Like Your Work Code)

**Goal**: Understand the NewSomething() pattern you see everywhere
```
type FileProcessor struct {
    inputDir  string
    outputDir string
    maxFiles  int
}

// TODO: Create a constructor function NewFileProcessor(inputDir, outputDir string, maxFiles int) *FileProcessor
// Make sure it returns a pointer to the struct

// TODO: Add a method ProcessFiles() that prints:
// "Processing files from [inputDir] to [outputDir], max [maxFiles] files"

func main() {
    // TODO: Use your constructor to create a FileProcessor
    // Call ProcessFiles() on it
}

```

### Questions to answer:

- Why do constructors usually return *FileProcessor instead of FileProcessor?
- How is this similar to your work's NewDebugReportGenerator()?

Problem 4: Mirror Your Work Confusion

**Goal**: Recreate the exact pattern that confused you
```
type ReportGenerator struct {
    outputPath string
}

func NewReportGenerator(outputPath string) *ReportGenerator {
    return &ReportGenerator{
        outputPath: outputPath,
    }
}

func (r *ReportGenerator) Generate(reportName string) error {
    fmt.Printf("Generating %s report to %s\n", reportName, r.outputPath)
    return nil
}

// Now create a Service that uses ReportGenerator
type Service struct {
    name string
    // TODO: Should you add a ReportGenerator field here? Or create it locally in methods?
}

// TODO: Try both approaches:
// Approach 1: Add reportGen *ReportGenerator to the Service struct and NewService constructor
// Approach 2: Create ReportGenerator locally in the method below

func (s *Service) ProcessData() error {
    // TODO: Either use s.reportGen or create a local ReportGenerator
    // Call Generate("data-analysis") on it
    return nil
}

```

### Questions to answer:

- Which approach feels more like your existing codebase patterns?
- What are the pros/cons of each approach?
- When might you choose one over the other?
