package files

import (
	"github.com/fer2d2/sievert/files/filesystem"
	"github.com/fer2d2/sievert/stubs"
)

// FileCommon represents a configuration file to be loaded or written into disk
type FileCommon struct {
	Label string
	Path  string
}

// Exists checks if the specified files exists in the filesystem
func (fileCommon *FileCommon) Exists() bool {
	filesystem.FileExists(fileCommon.Path)
}

// Stub returns the stub for the specified file type through the Label
func (fileCommon *FileCommon) Stub() ([]byte, error) {
	stub := stubs.Get(fileCommon.Label)
	return []byte(stub)
}

// FileContent returns the file content if it exists
func (fileCommon *FileCommon) FileContent() ([]byte, error) {
	if fileCommon.Exists() {
		contentBytes, err = filesystem.ReadFile(fileSievert.Path)
	}

	return contentBytes, err
}

// FileContentOrStub returns the file content if it exists or its stub otherwise
func (fileCommon *FileCommon) FileContentOrStub() ([]byte, error) {
	if fileCommon.Exists() {
		contentBytes, err = fileCommon.FileContent()
	} else {
		contentBytes, err = fileCommon.Stub()
	}

	return contentBytes, err
}
