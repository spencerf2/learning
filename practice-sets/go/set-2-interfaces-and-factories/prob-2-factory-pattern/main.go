package main

import (
	"fmt"
	"log/slog"
)

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
	fmt.Printf("Saving [%s] locally to [%s]\n", filename, ls.basePath)
	return nil
}

func (ls *LocalStorage) Load(filename string) ([]byte, error) {
	fmt.Printf("Loading [%s] from [%s]\n", filename, ls.basePath)
	return []byte("local data"), nil
}

type CloudStorage struct {
	region string
}

func (cs *CloudStorage) Save(filename string, data []byte) error {
	fmt.Printf("Uploading [%s] to cloud in [%s]\n", filename, cs.region)
	return nil
}

func (cs *CloudStorage) Load(filename string) ([]byte, error) {
	fmt.Printf("Downloading [%s] from cloud in [%s]\n", filename, cs.region)
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
	ls, err := NewStorage(StorageType_Local, "localPathDir")
	if err != nil {
		slog.Error("unable to initialize local storage")
	}
	cs, err := NewStorage(StorageType_Cloud, "cuscoPeru")
	if err != nil {
		slog.Error("unable to initialize cloud storage")
	}

	localFileName := "airbnb_file"
	newLocalFileName := "edited_airbnb_file"
	cloudFileName := "cusco_file"
	newCloudFileName := "edited_cusco_file"

	data, err := ls.Load(localFileName)
	if err != nil {
		slog.Error("unable to load", "localFileName", localFileName)
	}

	err = ls.Save(newLocalFileName, data)
	if err != nil {
		slog.Error("unable to save", "newLocalFileName", newLocalFileName, "data", data, "err", err)
	}

	data, err = cs.Load("cusco_file")
	if err != nil {
		slog.Error("unable to load", "cloudFileName", cloudFileName)
	}

	err = cs.Save(newCloudFileName, data)
	if err != nil {
		slog.Error("unable to save", "newCloudFileName", newCloudFileName, "data", data, "err", err)
	}
}

// Note to self: I'm supposed to notice that the caller doesn't know which implementation
//               they got. While the question could create a clearer depiction of this,
//               by actually consuming ls and cs rather than having me invoke them
//               manually, I can see that both of the functions are used the exact
//               same way.
//               The only difference is that in one case, we're using ls and in the
//               other, we're using cs. Since their signatures are defined the same
//               way, and they're both Storage interfaces, if there was a function
//               that consumed them, that function woudn't need to know whether it
//               was looking at the CloudStorage or LocalStorage struct, it'd use
//               either in the same way.

// Answering the questions:
// 1. How is this similar to the NewUnifileFromEnv pattern you found at work?
//   - Now that I've looked at NewUnifileFromEnv in the codebase from work, it's clear
//     how this question demonstrates that the caller doesn't need to know
//     which implementation they got, clarifying my "Note to Self" above.
//     In the work codebase, there's a switch case similar to the NewStorage factory
//     constructor above.
//     NewStorage ensures that there's one easy way to create Storage, regardless of
//     what type of storage it is. We can write code the exact same way throughout
//     the project, regardless of whether we're dealing with local or cloud storage.
//     We can use an env. var. to determine whether we're operating in local or cloud
//     mode, and when writing code we can just focus on "storage" rather than
//     seperating and handling both cloud and local storage.
//
// 2. Why is the factory function useful instead of creating structs directly?
//   - Rather than having to create and invoke a cloud or local storage struct
//     manually, we can rely on the factory and simply tell it the type that we
//     want. This makes the code more maintainable and simplifies instantiation
//     of the different kind of structs. Also, if we fail to instantiate properly,
//     we get an error message rather than having to remember what types of Storage
//     are available in the codebase.
//
// 3. What are the benefits of returning the interface type instead of concrete types?
//   - Rather than having to write code for the different types that could be returned
//     (for example, a factory that creates LocalStorage, and a factory that creates
//     CloudStorage), we can create one factory that conditionally returns each type
//     of storage defined in the code base.
//   - It's clear what the storage options are for the code base.
//   - We have one function for creating storage.
//   - We know that anything the factory returns satisfies the Storage interface, which
//     means can rely on each storage type to have the same methods, and be used in a
//     consistent way.
//     - That gives us the benefits described above, but it also allows us to write
//       code just for "Storage", rather than writing separate bits of code for the
//       various types of storage, and then worrying about what kind of storage we
//       may be dealing with where (for example, we don't need to remember how to
//       load cloud storage as something different from loading local storage--they
//       both work the same way!)
