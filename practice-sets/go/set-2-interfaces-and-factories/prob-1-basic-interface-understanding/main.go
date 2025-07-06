package main

type Writer interface {
	Write(data string) error
}

type FileWriter struct {
	fileName string
}

func main() {

}

// Answering the questions:
