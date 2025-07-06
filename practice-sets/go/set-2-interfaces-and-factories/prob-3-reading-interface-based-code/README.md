# Problem 3: Reading Interface-Based Code

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

## Questions to answer

- How does this mirror your work's unifile.go, local.go, and s3.go files?
- Why can BackupUserData work with any FileSystem implementation?
- If you needed to add a new storage type (like database), what would you need to change?
