package files

type FileDockerCompose struct {
	FileCommon
	Content GenericYaml
}

func NewFileDockerCompose() *FileDockerCompose {
	// filePath, _ := config.ProjectRoot()
	// filePath += "/" + "docker-compose.yml"
	// fileStub := stubs.Get(stubs.ComposeYml)

	fileDockerCompose := new(FileDockerCompose)
	// fileDockerCompose.FileCommon = FileCommon{filePath, fileStub}
	//
	// fileDockerCompose.InitFromYaml()

	return fileDockerCompose
}

func (fileDockerCompose *FileDockerCompose) InitFromYaml() {
	// dockerComposeContent := new(GenericYaml)
	// err := yaml.Unmarshal([]byte(fileDockerCompose.Stub), dockerComposeContent)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fileDockerCompose.Content = *dockerComposeContent
}
