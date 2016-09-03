package files

import (
	"log"

	"github.com/fer2d2/sievert/stubs"
	"github.com/fer2d2/sievert/util"
	yaml "gopkg.in/yaml.v2"
)

type GenericYaml map[string]interface{}

type FileDockerCompose struct {
	FileCommon
	Content GenericYaml
}

func NewFileDockerCompose() *FileDockerCompose {
	filePath := util.GetConfigDir() + "/" + "docker-compose.yml"
	fileStub := stubs.Get(stubs.ComposeYml)

	fileDockerCompose := new(FileDockerCompose)
	fileDockerCompose.FileCommon = FileCommon{filePath, fileStub}

	fileDockerCompose.InitFromYaml()

	return fileDockerCompose
}

func (fileDockerCompose *FileDockerCompose) InitFromYaml() {
	dockerComposeContent := new(GenericYaml)
	err := yaml.Unmarshal([]byte(fileDockerCompose.Stub), dockerComposeContent)
	if err != nil {
		log.Fatal(err)
	}

	fileDockerCompose.Content = *dockerComposeContent
}
