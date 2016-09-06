package files

import (
	"github.com/fer2d2/sievert/stubs"
	"github.com/fer2d2/sievert/util"
	yaml "gopkg.in/yaml.v2"
)

type multiSequence []map[string]interface{}

// SievertContent represents the user's configuration specified in the
// sievert.yml parsed as an struct
type SievertContent struct {
	IP       string        `yaml:"ip"`
	Services []string      `yaml:"services"`
	Folders  multiSequence `yaml:"folders"`
	Sites    multiSequence `yaml:"sites"`
}

// FileSievert represents a project file containing the user's configuration
type FileSievert struct {
	FileCommon
	Content *SievertContent
}

func NewFileSievert() (*FileSievert, error) {
	fileSievert := new(FileSievert)

	fileSievert.initFileCommon()
	err := fileSievert.initContent()

	return fileSievert, err
}

func (fileSievert *FileSievert) initFileCommon() {
	path := util.GetConfigDir() + "/" + "sievert.yml"
	label := stubs.SievertYml

	fileSievert.FileCommon = FileCommon{label, path}
}

func (fileSievert *FileSievert) initContent() error {
	content := new(SievertContent)
	err := yaml.Unmarshal(fileSievert.GetFileContentOrStub(), content)

	fileSievert.Content = content

	return err
}
