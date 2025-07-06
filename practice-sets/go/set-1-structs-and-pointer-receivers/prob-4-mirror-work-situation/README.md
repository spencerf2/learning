# Problem 4: Mirror Your Work Confusion

**Goal**: Recreate the exact pattern that confused you

```go
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

## Questions to answer:

- Which approach feels more like your existing codebase patterns?
- What are the pros/cons of each approach?
- When might you choose one over the other?
