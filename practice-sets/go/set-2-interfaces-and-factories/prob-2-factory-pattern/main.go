package main

import "fmt"

type Storage interface {
	Save(filename string, data []byte) error
	Load(filename string) ([]byte, error)
}

type StorageType string
const (
	StorageType_Local StorageType = "local"
	StorageType_Cloud StorageType = "cloud"
)

type LocalStorage struct {
	basePath string
}

func (ls *LocalStorage) Save(filename string, data []byte) error {
	fmt.Printf("Saving [%s] locally to [%s]", filename, ls.basePath)
	return nil
}

func (ls *LocalStorage) Load(filename string) ([]byte, error) {
	fmt.Printf("Loading [%s] from [%s]", filename, ls.basePath)
	return []byte("local data"), nil
}

type CloudStorage struct {
	region string
}

func (cs *CloudStorage) Save(filename string, data []byte) error {
	fmt.Printf("Uploading [%s] to cloud in [%s]", filename, cs.region)
	return nil
}

func (cs *CloudStorage) Load(filename string) ([]byte, error) {
	fmt.Printf("Downloading [%s] from cloud in [%s]", filename, cs.region)
	return []byte("cloud data"), nil
}

func NewStorage(storageType StorageType, config string) (Storage, error) {
	switch storageType {
	case StorageType_Local:
		return &LocalStorage{basePath: config}, nil
	case StorageType_Cloud:
		return &CloudStorage{region: config}, nil
	default:
		return nil, fmt.Errorf("unkown type %v", storageType)
	}
}

// Note to self: 1.
//               Interestingly, when I make the Save and Load methods pointer
//               receivers, Go expects that the newly created LocalStorage in
//               the NewStorage factory will be a pointer. If it's just a value,
//               we see error:
//
//               cannot use LocalStorage{…} (value of struct type LocalStorage)
//               as Storage value in return statement: LocalStorage does not
//               implement Storage (method Load has pointer receiver)
//               compilerInvalidIfaceAssign
//
//               However, when Save() and Load() are value receivers, then
//               returning LocalStorage{...} is just fine. Additionally, it's
//               also fine to return the pointer &LocalStorage{...}.
//
//               I want to test this to ensure that when the methods accept a
//               value receiver, and NewStorage returns a pointer, that the
//               methods still print correctly. I know they will from work
//               because we'll be accessing a value on the pointer, but still,
//               I want to test it in action here.
//
//               2.
//               There's a blue syntax error on my if, else if, else above:
//
//               could use tagged switch on storageTypeQF1003default
//
//               Going to create a commit here, and then change to a switch to
//               practice the difference.

func main() {

}

// Answering the questions:
