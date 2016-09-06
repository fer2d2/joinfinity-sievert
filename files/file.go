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
func (fileCommon *FileCommon) Stub() (string, error) {
	return stubs.Get(fileCommon.Label)
}

func (fileCommon *FileCommon) GetFileContentOrStub() ([]byte, error) {

	contentBytes, err := []byte(fileSievert.Stub())

	if fileCommon.Exists() {
		contentBytes, err = filesystem.ReadFile(fileSievert.Path)
	}

	return contentBytes, err
}
