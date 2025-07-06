# Problem 3: Constructor Pattern (Like Your Work Code)

**Goal**: Understand the NewSomething() pattern you see everywhere

```go
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

## Questions to answer:

- Why do constructors usually return *FileProcessor instead of FileProcessor?
- How is this similar to your work's NewDebugReportGenerator()?
