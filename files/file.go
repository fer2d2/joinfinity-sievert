package files

import (
	"log"
	"os"
)

// FileCommon represents a configuration file to be loaded or written into
// disk
type FileCommon struct {
	Path string
	Stub string
}

type ConfigFile interface {
	FileExists() bool
	WriteStubFile()
}

type YamlFile interface {
	InitFromYaml()
}

func (fileCommon *FileCommon) WriteStubFile() {
	file, err := os.Create(fileCommon.Path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	_, err = file.WriteString(fileCommon.Stub)
	if err != nil {
		log.Fatal(err)
	}

	file.Sync()
}
