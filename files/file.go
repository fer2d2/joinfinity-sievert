package files

import (
	"github.com/fer2d2/sievert/files/filesystem"
	"github.com/fer2d2/sievert/stubs"
)

// GenericYaml represents the type of a YAML file which doesn't match any fixed
// struct
type GenericYaml map[string]interface{}

type sequenceElement map[string]string

type multiSequence []sequenceElement

// FileCommon represents a configuration file to be loaded or written into disk
type FileCommon struct {
	Label string
	Path  string
}

// Exists checks if the specified files exists in the filesystem
func (fileCommon *FileCommon) Exists() bool {
	return filesystem.FileExists(fileCommon.Path)
}

// Stub returns the stub for the specified file type through the Label
func (fileCommon *FileCommon) Stub() ([]byte, error) {
	stub, err := stubs.Get(fileCommon.Label)
	return []byte(stub), err
}

// FileContent returns the file content if it exists
func (fileCommon *FileCommon) FileContent() ([]byte, error) {
	return filesystem.ReadFile(fileCommon.Path)
}

// FileContentOrStub returns the file content if it exists or its stub otherwise
func (fileCommon *FileCommon) FileContentOrStub() ([]byte, error) {
	var contentBytes []byte
	var err error

	if fileCommon.Exists() {
		contentBytes, err = fileCommon.FileContent()
	} else {
		contentBytes, err = fileCommon.Stub()
	}

	return contentBytes, err
}
