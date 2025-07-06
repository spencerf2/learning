# Problem 4: Interface Composition (Advanced)

**Goal**: Understand how interfaces can be combined (common in production Go)

```go
type Reader interface {
    Read(id string) ([]byte, error)
}

type Writer interface {
    Write(id string, data []byte) error
}

// TODO: Create a ReadWriter interface that embeds both Reader and Writer
// TODO: Create a Database struct that implements ReadWriter
// TODO: Create a Cache struct that implements only Reader
// TODO: Create a Logger struct that implements only Writer

// Function that only needs reading capability
func LoadUserProfile(r Reader, userID string) ([]byte, error) {
    return r.Read(fmt.Sprintf("profile_%s", userID))
}

// Function that needs both reading and writing
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
    // TODO: Show that Database can be used with both functions
    // TODO: Show that Cache can only be used with LoadUserProfile
    // TODO: Try to use Cache with UpdateUserProfile - what happens?
}
```

## Questions to answer

- Why can't you pass Cache to UpdateUserProfile?
- How does interface composition help with flexible function parameters?
- What's the benefit of accepting the smallest interface possible in functions?
