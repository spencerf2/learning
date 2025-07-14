package main

import "fmt"

type Writer interface {
	Write(data string) error
}

type FileWriter struct {
	fileName string
}

func (f *FileWriter) Write(somethingOtherThanData string) error {
	fmt.Printf("Writing '[%s]' to file: [%s]\n", somethingOtherThanData, f.fileName)
	return nil
}

type ConsoleWriter struct {}

func (c *ConsoleWriter) Write(data string) error {
	fmt.Printf("Console output: [%s]\n", data)
	return nil
}

func ProcessData(w Writer, data string) error {
	return w.Write(data)
}

func main() {
	f := &FileWriter{fileName: "spencer's file"}
	c := &ConsoleWriter{}

	ProcessData(f, "FileWriter's data")
	ProcessData(c, "ConsoleWriter's data")
}

// Note to self: When I first wrote this up, I didn't add "&" to the beginning
//               of the FileWriter and ConsoleWriter instances. That resulted
//               in the following error on both ProcessData lines:
//
//               cannot use f (variable of struct type FileWriter) as Writer
//               value in argument to ProcessData: FileWriter does not
//               implement Writer (method Write has pointer
//               receiver) compilerInvalidIfaceAssign
//
//               I notice that if I make FileWriter's Write() a value receiver,
//               the error message goes away. Experimenting, I know that it's
//               ok to pass the pointer into ProcessData, even if Write is a
//               value receiver.


// Answering the questions:
// 1. How does Go know that FileWriter and ConsoleWriter implement Writer?
//   - Because we pass them to ProcessData which is a function that's first
//     argument is of type Writer--the interface. I recall that an interface
//     says what the requirements are to be an interface. So, in this case,
//     to be a Writer, all that FileWriter and ConsoleWriter need to do is
//     have a Write method that follows the Writer's Write method signature
//     and thereby take in a string and return an error. From playing, I know
//     that the variable name of the Write method isn't important.
//     "somethingOtherThanData" works just as well as "data" for the Write
//     signatures.
//     Also, from experimenting, I know that the Write methods do need to take
//     in a string and return an error. Without that we get error:
//
//     cannot use f (variable of type *FileWriter) as Writer value in argument
//     to ProcessData: *FileWriter does not implement Writer (wrong type for
//     method Write) compilerInvalidIfaceAssign
//
// 2. What happens if you remove the Write method from one of your structs?
//   - We get an error on the instance we pass into ProcessData. For ex. if we
//     comment out ConsoleWriter's Write method we get an error on "c" in:
//
//     ProcessData(c, "ConsoleWriter's data")
//
//     Which says:
//
//     cannot use c (variable of type *ConsoleWriter) as Writer value in
//     argument to ProcessData: *ConsoleWriter does not implement Writer
//     (missing method Write) compilerInvalidIfaceAssign
//
//     In other words, the struct no longer satisfies the Writer interface,
//     and can thus no longer be used as one.
//
// 3. Why doesn't Go require explicit interface declarations?
//   - They've opted to allow for implicit interfaces. I recall that that
//     has some advantages, especially when it comes to third party libraries.
//     Normally, in a language like Python, if you wanted to use a method that
//     comes from a third party library, you'd have to subclass the class of
//     the library that has the method.
//     That creates issues when the library's creator does not make it easy
//     to subclass the class with the method. You end up having to do things
//     like re-implement the entire class just to get the method, or overwrite
//     the method on the class--copying code and editing it. Things like this.
//     When we have implicit interfaces, we end up being able to use third
//     party library methods much more easily, and it saves on a lot of other
//     code that would otherwise be needed just to utilize a method.
//     Basically, we get the benefits of inheritence and subclassing, without
//     as much boilerplate code. And anything that "walks and quacks like a
//     duck, is considered a duck".
