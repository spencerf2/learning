package main

import "fmt"

type Reader interface {
	Read(id string) ([]byte, error)
}

type Writer interface {
	Write(id string, data []byte) error
}

type ReadWriter interface {
	Reader
	Writer
}

type Database struct {
	ReadWriter
}

// Note to self: Should I actually be doing something like:
/*
				 type Dataase struct {
				 	ReadWriter ReadWriter
				 }
*/
//               Or is this like JavaScript, where just doing it
//               the way I have is the same as `ReadWriter ReadWriter`?

type Cache struct {
	read Reader
}

type Logger struct {
	write Writer
}

// Note to self: If I do:
/*
				 type Cache struct {
				 	read Reader
				 }
*/
//               Then there's a syntax error saying:
//
//               field read is unused (U1000)go-staticcheck
//
//               But if I delete "read" from `read Reader` so the only thing
//               within the struct is `Reader``, then it's fine.
//               I believe this means I must implment the Reader interface
//               so that `read` actually points to something.

func (c *Cache) read(userID string) ([]byte, error) {
	data := []byte(fmt.Sprintf("read_in_%s", userID))
	return data, nil
}

// Note to self: If I implement a `read()` method on Cache, then there's
//               an error message saying:
//
//               field and method with the same name readcompilerDuplicateFieldAndMethod
//               main.go(32, 2): other declaration of read
//
//               Thinking it'll go away if I make Cache struct only contain `Reader`.

func LoadUserProfile(rw ReadWriter, userID string, newData []byte) error {
	// Read current data
	_, err := rw.Read(fmt.Sprintf("profile_%s", userID))
	if err != nil {
		return err
	}

	// Write new data
	return rw.Write(fmt.Sprintf("profile_%s", userID), newData)
}

func main() {
	db := Database{}
	LoadUserProfile(db, "spencer_ID", []byte("cusco_data"))
}