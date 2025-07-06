package main

import "fmt"

type FileProcessor struct {
	inputDir string
	outputDir string
	maxFiles int
}

func NewFileProcessor(inputDir, outputDir string, maxFiles int) *FileProcessor {
	f := &FileProcessor{
		inputDir: inputDir,
		outputDir: outputDir,
		maxFiles: maxFiles,
	}
	fmt.Println("FileProcessor pointer address", f)
	// Note to self: Above didn't print memory address.
	//               Need to learn more about fmt package.
	//               Will return to this later.

	// UPDATE:
	// ============================================================================
	// EXTENDED EXERCISE: fmt Package Exploration
	// ============================================================================
	// - Try these different formatting approaches and observe the differences:
	// 
	// fmt.Println("Method 1:", f) // What I've been doing
	// fmt.Printf("Method 2: %v\n", f) // Removes new line hence needing \n
	// fmt.Printf("Method 3: %p\n", f) // Here's actual pointer address
	// fmt.Printf("Method 4: %#v\n", f) // This one is interesting. Want to look it
	//                                     up later.
	//                                     &main.FileProcessor{inputDir:"input_directory", outputDir:"output_directory", maxFiles:5}
	//
	// Update 2025-07-06:
	// %#v is "Go-syntax representation". You can copy and paste it to create the
	// exact same struct.

	// Questions to answer after experimenting:
	// 1. What's the difference between %v, %p, and %#v format verbs?
	//   - %v is value, which is the default Println seems to do.
	//   - %p is the pointer memory address
	//   - %#v seems to be the actual function with its args being printed.
	// 
	// 2. Why might Go choose to format pointers differently in Println vs Printf?
	//   - I think Println is your more standard "just print this line", while
	//     Printf is for getting specific with what you're looking to output.
	// 
	// 3. Which method actually shows you the memory address?
	//   - Printf with %p
	// 
	// 4. When would you use each formatting approach in debugging?
	//   - Println for when you just want to see the structure of what's printed.
	//   - Printf for when you want more granular info.

	// BONUS: Try these additional experiments:
	// 1. What happens if you print the address of f itself? (&f)
	// 
	// fmt.Println("&f:", &f)
	/*
		Method 3: 0x140001080c0
		&f: 0x1400005c038

		You get the memory address of the variable f rather than what f
		itself points to.
	*/
	// 
	// 2. Compare printing a pointer vs the value it points to (*f)
	// 
	// fmt.Println("&f:", &f)
	// fmt.Println("*f:", *f)
	/*
		Method 3: 0x140001140c0
		&f: 0x1400010e030
		*f: {input_directory output_directory 5}

		When you print the pointer, you get the memory address. When
		you print the value it points to, you get the call signature
		with variables. Want to double check this.
	*/
	// 
	// 3. Try %T to see the type information
	// 
	// fmt.Printf("Address of f variable: %p\n", &f)
	// fmt.Printf("Value f points to: %v\n", *f)
	// fmt.Printf("Type of f: %T\n", f)
	/*
		Method 3: 0x140001120c0
		Address of f variable: 0x1400010c030
		Value f points to: {input_directory output_directory 5}
		Type of f: *main.FileProcessor

		Very interesting. So that means that %#v that we saw earlier included
		the type info. Though, in that case we saw &main.FileProcessor rather
		than *main.FileProcessor. Again, want to research this a little to
		solidify understanding.
	*/
	// ============================================================================
	// END EXTENDED EXERCISE
	// ============================================================================

	return f
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
//
// Update 2025-07-06: ^ is exactly correct.

func main() {
	f := NewFileProcessor("input_directory", "output_directory", 5)
	f.ProcessFiles()

	// UPDATE 2025-07-06:
	// ============================================================================
	// GOING DEEPER: & vs * Understanding
	// ============================================================================
	// To solidify understanding of & and *, try this experiment:
	//
	val := FileProcessor{inputDir: "test", outputDir: "test", maxFiles: 1}
	ptr := &val
	//
	fmt.Printf("\nval with %%#v: %#v\n", val)
	fmt.Printf("ptr with %%#v: %#v\n", ptr)
	//
	// Questions to answer:
	// 1. What does val show with %#v? (no & or * in the output)
	// 2. What does ptr show with %#v? (should show & in the output)
	// 3. Why does one show & and the other doesn't?
	// 4. How does this connect to the %T output you saw earlier?
	// ============================================================================
	// END GOING DEEPER
	// ============================================================================
}

// Note to self: Nice, ran it, it works!

// Answering the questions:
// 1. Why do constructors usually return *FileProcessor instead of FileProcessor?
//   - Because the reason a constructor creates something is usually so that it can
//     be used elsewhere / updated. It isn't usually created just to be immediately
//     consumed and discarded, though that's certainly possible.
//
//     Update 2025-07-06: Also, passing pointers around is more memory efficient
//                        than passing copies of structs.
// 
// 2. How is this similar to your work's NewDebugReportGenerator()?
//   - NewDebugReportGenerator is a factory constructor, that creates an instance of
//     DebugReportGenerator which holds onto two fields:
//     i. One that keeps track of the location where the DebugReport should be written.
//     ii. One that holds onto the file system abstraction allowing the application
//         to read, write, list files, or create a folder on either AWS S3 or the
//         local file system, depending on which env var is set.
