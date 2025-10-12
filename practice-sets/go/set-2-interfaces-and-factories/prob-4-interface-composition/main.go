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
	records map[int]string
}

// Note to self: Should I actually be doing something like:
/*
				 type Dataase struct {
				 	ReadWriter ReadWriter
				 }
*/
//               Or is this like JavaScript, where just doing it
//               the way I have is the same as `ReadWriter ReadWriter`?
//
//               Answer (2025-10-12): Neither! The struct needs to
//                                    implement the interface, not embed
//                                    it!

func (d *Database) Read(id string) ([]byte, error) {
	return []byte(id), nil
}

func (d *Database) Write(id string, data []byte) error {
	fmt.Printf("created/updated id=%s using data=%s", id, data)
	return nil
}

type Cache struct {}

type Logger struct {}

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
//
//               Upadte (2025-10-12): When I wrote the above, I had confused
//                                    myself. The struct needs to implement
//                                    the the interface, not embed it. When I
//                                    created a field "read" of type Reader
//                                    interface, of course there was an error
//                                    reporting I needed to utilize it.

func (c *Cache) Read(userID string) ([]byte, error) {
	data := []byte(fmt.Appendf([]byte("starting_byte"), "read_in_%s", userID))
	return data, nil
}

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
//
//               Update (2025-10-12): The above was kind of on the right track. It makes
//                                    sense that you can't have a field and a method with
//                                    the same name as that'd create namespace collisions.
//                                    We didn't need to embed the interface, we needed to
//                                    implement it! With that said, a struct could embed
//                                    an interface--it'd make sense to do so when you want
//                                    to wrap an interface with additional info. You can
//                                    store variables on the struct, thereby associating
//                                    them with the interface. You can also override the
//                                    interfaces methods.

func (l *Logger) Write(id string, data []byte) error {
	fmt.Printf("logging id=%s data=%s", id, data)
	return nil
}

func LoadUserProfile(rw ReadWriter, userID string) ([]byte, error) {
	return rw.Read(fmt.Sprintf("profile_%s", userID))
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
	db := &Database{make(map[int]string)}
	db.records[1] = "first_record"

	userID := "spencer_ID"

	data, err := LoadUserProfile(db, userID)
	if err != nil {
		slog.Error("could not load user using database")
	}
	slog.Info("loaded user profile from db", "data", data)

	updatedData := fmt.Append(data, "added_data")
	err = UpdateUserProfile(db, userID, updatedData)
	if err != nil {
		slog.Error("could not update user using database")
	}

	// The above proves that the Database can be used with both functions.

	cache := &Cache{}

	data, err = LoadUserProfile(cache, userID)
	if err != nil {
		slog.Error("could not load user using cache")
	}
	slog.Info("loaded user profile from cache", "data", data)

	// Attempting to use cache with UpdateUserProfile
	/*
		updatedData2 := fmt.Append(data, "added_data_again")
		err = UpdateUserProfile(cache, userID, updatedData2)
		if err != nil {
			slog.Error("could not update user using database")
		}
	*/
	// The above results in the following error on `cache`:
	//
	// cannot use cache (variable of type *Cache) as ReadWriter value in argument to
	// UpdateUserProfile: *Cache does not implement ReadWriter (missing method Write
	// compilerInvalidIfaceAssign
	//
	// This makes sense because UpdateUserProfile requires an update a.k.a. write
}

// Answering the questions:
// 1. Why can't you pass Cache to UpdateUserProfile?
//   - It requires a write operation to make an update.
//   - Using the interface type in the UpdateUserProfile function definition
//     results in automatic enforcement. It's instantly clear that Cache cannot
//     be used because it is not an implementation of ReadWriter due to not having
//     a Write method.
//
// 2. How does interface composition help with flexible function parameters?
//   - Because the function requires a ReadWriter interface, we can separately
//     define a Reader and a Writer, and then bring them together for use in this
//     function. Without composition we'd have to define a separate interface or
//     "class" that allows both reading and writing. Thanks to Go's flexibility, we're
//     able to utilize these separate interfaces together for the purpose of
//     updating user profiles.
//
// 3. What's the benefit of accepting the smallest interface possible in functions?
//   - It keeps the function light
//   - It keeps what's absolutely necessary clear.
//   - It means a greater number of interfaces can be used with the function, so the
//     function can be defined one time but used in many different scenarios. Rather
//     than having to create a new function in the event that we want to use a
//     different interface (or class) with it, as in many other languages, in Go, we
//     can reuse the same function.
//   - Thanks to composition, we can have many small single purpose interfaces that
//     are then composed together when the need arises.
//   - If we were passing a ReadWriter into a function for some reason, even though
//     only reading was required, and then later, we wanted to reuse the function, but
//     with an interface that only implemented a reader--we could do that. We could
//     define the function as only requiring a Read, and then use both interfaces that
//     only support reading (Cache) and ones that support reading and writing
//     (ReadWriter). We wouldn't have to add a fake Write method to Cache, just so it
//     could satisfy a ReadWriter requirement the function may have originally been
//     written as requiring.
