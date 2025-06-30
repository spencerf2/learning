package main

import "fmt"

type FileProcessor struct {
	inputDir string
	outputDir string
	maxFiles int
}

func NewFileProcessor(inputDir, outputDir string, maxFiles int) *FileProcessor {
	return &FileProcessor{
		inputDir: inputDir,
		outputDir: outputDir,
		maxFiles: maxFiles,
	}
}

// Note to self: With the above, I know that the return has to have "&" on it. I
//               recal that that means we're returning a memory address, which
//               makes sense, because that's what a pointer is. So, if I print
//               the variable being returned before returning it, then I should
//               see that we print a memory address. Will do in a separate commit
//               for tracking purposes!

func (f FileProcessor) ProcessFiles() {
	fmt.Printf(
		"Processing files from %s to %s, max %d files",
		f.inputDir,
		f.outputDir,
		f.maxFiles,
	)
}

// Note to self: Believe it's fine not to use a pointer for FileProcessor on
//               ProcessFiles() because we don't need to do anything with
//               the struct but access its values.

func main() {
	f := NewFileProcessor("input_directory", "output_directory", 5)
	f.ProcessFiles()
}

// Note to self: Nice, ran it, it works!

// Answering the questions:
// 1. Why do constructors usually return *FileProcessor instead of FileProcessor?
//   - Because the reason a constructor creates something is usually so that it can
//     be used elsewhere--updated. It isn't usually created just to be immediately
//     consumed, though that's certainly possible.
// 
// 2. How is this similar to your work's NewDebugReportGenerator()?
//   - Ooh, good question. Will have to come back to this tomorrow! Need to switch
//     to starting my work day :)
