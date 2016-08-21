package files

import (
	"log"

	"github.com/fer2d2/sievert/stubs"
	"github.com/fer2d2/sievert/util"
	yaml "gopkg.in/yaml.v2"
)

type multiSequence []map[string]interface{}

// SievertContent represents user's configuration specified in the sievert.yml
// file
type SievertContent struct {
	IP       string        `yaml:"ip"`
	Services []string      `yaml:"services"`
	Folders  multiSequence `yaml:"folders"`
	Sites    multiSequence `yaml:"sites"`
}

type FileSievert struct {
	FileCommon
	Content SievertContent
}

func NewFileSievert() *FileSievert {
	filePath := util.GetConfigDir() + "/" + "sievert.yml"
	fileStub := stubs.SievertYmlTpl

	fileSievert := new(FileSievert)
	fileSievert.FileCommon = FileCommon{filePath, fileStub}

	fileSievert.InitFromYaml()

	return fileSievert
}

func (fileSievert *FileSievert) InitFromYaml() {
	sievertContent := new(SievertContent)
	err := yaml.Unmarshal([]byte(fileSievert.Stub), sievertContent)
	if err != nil {
		log.Fatal(err)
	}

	fileSievert.Content = *sievertContent
}
