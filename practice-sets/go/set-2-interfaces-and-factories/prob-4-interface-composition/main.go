package main

import (
	"fmt"
	"log/slog"
)

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
	Reader
}

type Logger struct {
	Writer
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

// func (c *Cache) Read(userID string) ([]byte, error) {
// 	data := []byte(fmt.Appendf([]byte("starting_byte"), "read_in_%s", userID))
// 	return data, nil
// }

// Note to self: If I implement a `read()` method on Cache, then there's
//               an error message saying:
//
//               field and method with the same name readcompilerDuplicateFieldAndMethod
//               main.go(32, 2): other declaration of read
//
//               Thinking it'll go away if I make Cache struct only contain `Reader`.
//               It goes away if I capitalize the `R` on the `read()` method.
//               And also if I delete `read` from Cache.
//               I believe this is because Reader is an interface, so Cache struct
//               doesn't need a field to represent it, although it can have one.
//               Cache can simply embed its methods. What does that make Cache then?
//               A struct that embeds the Reader interface. Perhaps that makes sense
//               because Cache may need the ability to read in data.

func LoadUserProfile(r Reader, userID string) ([]byte, error) {
	return r.Read(fmt.Sprintf("profile_%s", userID))
}

func UpdateUserProfile(rw ReadWriter, userID string, newData []byte) error {
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
	userID := "spencer_ID"

	data, err := LoadUserProfile(db, userID)
	if err != nil {
		slog.Error("could not load user using database")
	}
	slog.Info("loaded user profile from db", "data", data)

	err = UpdateUserProfile(db, userID, fmt.Append(data, "added_data"))
	if err != nil {
		slog.Error("could not update user using database")
	}

	// c := Cache{}
	// cacheData, err := LoadUserProfile(c, userID)
}
// https://claude.ai/chat/8c2fef5d-51a4-4f97-8547-c3bb60c318b1
