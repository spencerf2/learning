# Problem 2: The Factory Pattern (Like Your Work)

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

## Questions to answer

- How is this similar to the NewUnifileFromEnv pattern you found at work?
- Why is the factory function useful instead of creating structs directly?
- What are the benefits of returning the interface type instead of concrete types?
