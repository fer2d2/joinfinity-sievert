package config

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

const composeConfFile string = "/stubs/docker-compose/docker-compose.yml.stub"

// DockerCompose represents the configuration specified in the
// docker-compose.yml file
type DockerCompose map[string]interface{}

// NewDockerCompose creates new DockerCompose struct representing a
// docker-compose.yml file
func NewDockerCompose(config *Config) *DockerCompose {
	dockerCompose := &DockerCompose{}

	dockerCompose.initializeFromFile()
	dockerCompose.createFile()

	return dockerCompose
}

func (dockerCompose *DockerCompose) initializeFromFile() {
	err := yaml.Unmarshal(readFile(composeConfFile), dockerCompose)

	if err != nil {
		log.Fatal(err)
	}
}

func (dockerCompose *DockerCompose) createFile() {
	dockerComposeText, err := yaml.Marshal(dockerCompose)
	fmt.Printf("--- docker-compose.yml content:\n%s\n", string(dockerComposeText))

	if err != nil {
		log.Fatal(err)
	}
}
