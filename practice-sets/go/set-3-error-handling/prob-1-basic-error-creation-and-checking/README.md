# Problem 1: Basic Error Creation and Checking

**Goal**: Understand Go's error fundamentals

```go
import (
    "errors"
    "fmt"
)

type User struct {
    ID   int
    Name string
    Age  int
}

// TODO: Create a function ValidateUser(user User) error that returns:
// - An error if user.Name is empty (use errors.New())
// - An error if user.Age is less than 0 or greater than 150 (use fmt.Errorf())
// - nil if the user is valid

// TODO: Create a function SaveUser(user User) error that:
// 1. First calls ValidateUser
// 2. If validation fails, return the validation error
// 3. If validation passes, print "Saving user: [name]" and return nil

func main() {
    users := []User{
        {ID: 1, Name: "Alice", Age: 30},
        {ID: 2, Name: "", Age: 25},
        {ID: 3, Name: "Bob", Age: -5},
    }

    // TODO: Try to save each user and print any errors
    for _, user := range users {
        if err := SaveUser(user); err != nil {
            fmt.Printf("Error saving user %d: %v\n", user.ID, err)
        }
    }
}
```

## Questions to answer

- When should you use errors.New() vs fmt.Errorf()?
- Why does Go use explicit error returns instead of exceptions?
- What's the difference between %v and %s when printing errors?
