package filesystem

import yaml "gopkg.in/yaml.v2"

func WriteYaml(content interface{}, path string) error {
	bytesContent, err := yaml.Marshal(content)

	if err != nil {
		return err
	}

	return WriteFile(string(bytesContent[:]), path)
}
