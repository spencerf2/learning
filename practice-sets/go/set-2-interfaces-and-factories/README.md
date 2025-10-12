# Go Practice Set 2: Interfaces and Factory Patterns

Create a new directory: go-learning/interfaces-and-factories/ and work through these problems in order.

## Problem 1: Basic Interface Understanding

**Goal**: Understand how Go interfaces work (implicit satisfaction)

```go
// Define a simple interface
type Writer interface {
    Write(data string) error
}

// TODO: Create a FileWriter struct with a fileName field
// TODO: Add a Write method to FileWriter that prints: "Writing '[data]' to file: [fileName]"
// TODO: Create a ConsoleWriter struct (no fields needed)
// TODO: Add a Write method to ConsoleWriter that prints: "Console output: [data]"

func ProcessData(w Writer, data string) error {
    return w.Write(data)
}

func main() {
    // TODO: Create instances of both FileWriter and ConsoleWriter
    // TODO: Call ProcessData with each one
    // Notice: No explicit "implements" declaration needed!
}
```

### Questions to answer

- How does Go know that FileWriter and ConsoleWriter implement Writer?
- What happens if you remove the Write method from one of your structs?
- Why doesn't Go require explicit interface declarations?

## Problem 2: The Factory Pattern (Like Your Work)

**Goal**: Recreate the NewUnifileFromEnv pattern you discovered

```go
type Storage interface {
    Save(filename string, data []byte) error
    Load(filename string) ([]byte, error)
}

type StorageType string
const (
    StorageType_Local StorageType = "local"
    StorageType_Cloud StorageType = "cloud"
)

// TODO: Create a LocalStorage struct with a basePath field
// TODO: Implement Save method that prints: "Saving [filename] locally to [basePath]"
// TODO: Implement Load method that prints: "Loading [filename] from [basePath]" and returns []byte("local data"), nil

// TODO: Create a CloudStorage struct with a region field
// TODO: Implement Save method that prints: "Uploading [filename] to cloud in [region]"
// TODO: Implement Load method that prints: "Downloading [filename] from cloud in [region]" and returns []byte("cloud data"), nil

// TODO: Create a factory function NewStorage(storageType StorageType, config string) (Storage, error)
// It should return LocalStorage for local type, CloudStorage for cloud type
// Use config as basePath for local, region for cloud
// Return an error for unknown types

func main() {
    // TODO: Create both storage types using your factory
    // TODO: Call Save and Load on each
    // Notice: The caller doesn't know which implementation they got!
}
```

### Questions to answer

- How is this similar to the NewUnifileFromEnv pattern you found at work?
- Why is the factory function useful instead of creating structs directly?
- What are the benefits of returning the interface type instead of concrete types?

## Problem 3: Reading Interface-Based Code

**Goal**: Practice the pattern you encountered with UniFile

```go
// This mimics the pattern you found in your work codebase
type FileSystem interface {
    CreateFolder(path string) error
    WriteFile(path string, data []byte) error
    ReadFile(path string) ([]byte, error)
}

type Environment string
const (
    Env_Development Environment = "dev"
    Env_Production  Environment = "prod"
)

// TODO: Create DevFileSystem struct - implement all three methods with simple print statements
// TODO: Create ProdFileSystem struct - implement all three methods with different print statements

func NewFileSystemFromEnv(env Environment) (FileSystem, error) {
    switch env {
    case Env_Development:
        return &DevFileSystem{}, nil
    case Env_Production:
        return &ProdFileSystem{}, nil
    default:
        return nil, fmt.Errorf("unsupported environment: %s", env)
    }
}

// This simulates code that uses the filesystem without knowing which one
func BackupUserData(fs FileSystem, userID string, data []byte) error {
    folderPath := fmt.Sprintf("/backups/user_%s", userID)
    if err := fs.CreateFolder(folderPath); err != nil {
        return fmt.Errorf("failed to create backup folder: %w", err)
    }

    filePath := fmt.Sprintf("%s/data.json", folderPath)
    if err := fs.WriteFile(filePath, data); err != nil {
        return fmt.Errorf("failed to write backup file: %w", err)
    }

    return nil
}

func main() {
    // TODO: Test BackupUserData with both dev and prod filesystems
    // TODO: Show that the same function works with different implementations
}
```

### Questions to answer

- How does this mirror your work's unifile.go, local.go, and s3.go files?
- Why can BackupUserData work with any FileSystem implementation?
- If you needed to add a new storage type (like database), what would you need to change?

## Problem 4: Interface Composition (Advanced)

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

### Questions to answer

- Why can't you pass Cache to UpdateUserProfile?
- How does interface composition help with flexible function parameters?
- What's the benefit of accepting the smallest interface possible in functions?
