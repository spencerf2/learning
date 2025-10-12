package main

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

// Answering the questions:
// 1. What's the difference between %v and %w in fmt.Errorf?
//   - 
//
// 2. How does error wrapping help with debugging?
//   - 
//
// 3. What information gets lost if you use %v instead of %w?
//   - 
