package files

import (
	"github.com/fer2d2/sievert/config"
	"github.com/fer2d2/sievert/files/filesystem"
	"github.com/fer2d2/sievert/stubs"
	yaml "gopkg.in/yaml.v2"
)

// TODO replace simple []string with this struct
type Service struct {
	Name    string `yaml:"Name"`
	Version string `yaml:"Version,omitempty"`
	Memory  string `yaml:"Memory,omitempty"`
}

type Site struct {
	Map      string `yaml:"map"`
	To       string `yaml:"to"`
	Template string `yaml:"template"`
}

type Folder struct {
	Map string `yaml:"map"`
	To  string `yaml:"to"`
}

type Database struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Type     string `yaml:"type"`
}

// SievertContent represents the user's configuration specified in the
// sievert.yml parsed as an struct
type SievertContent struct {
	// IP        string     `yaml:"ip"`
	Authorize string     `yaml:"authorize"`
	Keys      []string   `yaml:"keys"`
	Services  []Service  `yaml:"services"`
	Folders   []Folder   `yaml:"folders"`
	Sites     []Site     `yaml:"sites"`
	Databases []Database `yaml:"databases"`
}

// FileSievert represents a project file containing the user's configuration
type FileSievert struct {
	FileCommon
	Content *SievertContent
}

// NewFileSievert creates a FileSievert initialized from an existent file in
// the disk if it exists or loading its stub otherwise
func NewFileSievert() (*FileSievert, error) {
	fileSievert := new(FileSievert)

	fileSievert.initFileCommon()
	err := fileSievert.initContent()

	return fileSievert, err
}

func (fileSievert *FileSievert) initFileCommon() error {
	projectRoot, err := config.ProjectRoot()

	if err != nil {
		return err
	}

	path := projectRoot + "/sievert.yml"
	label := stubs.SievertYml

	fileSievert.FileCommon = FileCommon{label, path}

	return nil
}

func (fileSievert *FileSievert) initContent() error {
	content := new(SievertContent)

	fileOrStub, err := fileSievert.FileContentOrStub()

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(fileOrStub, content)

	fileSievert.Content = content

	return err
}

func (fileSievert FileSievert) Write() error {
	return filesystem.WriteYaml(fileSievert.Content, fileSievert.Path)
}
