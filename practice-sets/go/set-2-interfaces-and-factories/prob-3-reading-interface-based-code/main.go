package main

import (
	"fmt"
	"log/slog"
)

type FileSystem interface {
	CreateFolder(path string) error
	WriteFile(path string, data []byte) error
	ReadFile(path string) ([]byte, error)
}

type Environment string
const (
	Env_Development Environment = "dev"
	Env_Production Environment = "prod"
)

type DevFileSystem struct {}
type ProdFileSystem struct {}

func (dfs *DevFileSystem) CreateFolder(path string) error {
	fmt.Println("Dev folder created at", path)
	return nil
}

func (dfs *DevFileSystem) WriteFile(path string, data []byte) error {
	fmt.Println("Wrote dev file to", path)
	return nil
}

func (dfs *DevFileSystem) ReadFile(path string) ([]byte, error) {
	fmt.Println("Read dev file at", path)
	return []byte("dev file data"), nil
}

func (pfs *ProdFileSystem) CreateFolder(path string) error {
	fmt.Println("Prod folder created at", path)
	return nil
}

func (pfs *ProdFileSystem) WriteFile(path string, data []byte) error {
	fmt.Println("Wrote prod file to", path)
	return nil
}

func (pfs *ProdFileSystem) ReadFile(path string) ([]byte, error) {
	fmt.Println("Read prod file at", path)
	return []byte("prod file data"), nil
}

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
	userID := "spence"
	devFileSystem, err := NewFileSystemFromEnv(Env_Development)
	if err != nil {
		slog.Error("unable to instantiate DevFileSystem")
	}
	prodFileSystem, err := NewFileSystemFromEnv(Env_Production)
	if err != nil {
		slog.Error("unable to instantiate ProdFileSystem")
	}

	devFilePath := "path_to_dev_file"
	devFileData, err := devFileSystem.ReadFile(devFilePath)
	if err != nil {
		slog.Error("failed to read in dev file", "devFilePath", devFilePath)
	}

	prodFilePath := "path_to_prod_file"
	prodFileData, err := prodFileSystem.ReadFile(prodFilePath)
	if err != nil {
		slog.Error("failed to read in prod file", "prodFilePath", prodFilePath)
	}

	if err := BackupUserData(devFileSystem, userID, devFileData); err != nil {
		slog.Error("unable to backup dev data", "userID", userID, "err", err)
	}
	if err := BackupUserData(prodFileSystem, userID, prodFileData); err != nil {
		slog.Error("unable to backup prod data", "userID", userID, "err", err)
	}
}

// Answering the questions:
// 1. How does this mirror your work's unifile.go, local.go, and s3.go files?
//   - unifile.go contains the interface and factory function.
//   - local.go and s3.go both:
//     - contain a struct for their respective file systems
//       and a factory function to create that struct.
//     - implement the methods necessary to be considered a UniFile interface
//       - s3.go contains additional functions required for the interface methods
//         to work.
//   - This shows the power of Go's interface and struct system. The implementation
//     details for the local file system in local.go are far simpler than what's
//     required to do the same thing with S3, however, once we've written these
//     methods for the different file systems, the rest of the code base doesn't
//     care whether we're storing in S3 or locally--we're just using the "file system"
//     or unifile in the case of the work codebase. Super cool!
//
// 2. Why can BackupUserData work with any FileSystem implementation?
//   - Because both file systems follow the same contract--that established by
//     the FileSystem interface. As far as the caller (BackupUserData) is
//     concerned, they both have Read, Write, and CreateFolder methods with the
//     same signature and return details.
//
// 3. If you needed to add a new storage type (like database), what would you need
//    to change?
//   - We'd need to add a new struct representing the database, and for that struct to
//     implement the FileSystem interface methods. Additional work may be
//     required to implement those methods (such as in the case of S3--where we have
//     to utilize the s3 library, our config getter, and instantiate the
//     s3 client), but once they're implemented no other changes are needed for the
//     app to store to the database instead of s3 or locally.
//     (With this said, it should be noted that we have a different mechanism present
//     in the code base for the database. The db is not represented by UniFile, which is
//     for unstructured data (files/blobs) storage, whether stored locally or in s3. The
//     db is handled differently since it's for structured data storage.)
