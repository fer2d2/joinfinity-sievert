package filesystem

import (
	"io/ioutil"
	"os"
	"os/user"
)

// ReadFile reads the content of a given file
func ReadFile(path string) ([]byte, error) {
	configFileContent, err := ioutil.ReadFile(path)
	return configFileContent, err
}

// WriteFile writes the content of a given string to a file in the specified path
func WriteFile(content string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	file.Sync()

	return nil
}

func FileExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	if fileInfo.IsDir() {
		return false
	}

	return true
}

func DirExists(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	if !fileInfo.IsDir() {
		return false
	}

	return true
}

// func GetContentFromFileOrDefault(filePath string, defaultValue string) []byte {
// 	byteContent := []byte(defaultValue)
// 	if FileExists(filePath) {
// 		byteContent = ReadFile(filePath)
// 	}
//
// 	return byteContent
// }

func GetConfigDir() (string, error) {
	confDir := os.Getenv("SIEVERT_CONF")

	if confDir == "" {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}

		confDir = usr.HomeDir + "/.sievert"
	}

	return confDir, nil
}
