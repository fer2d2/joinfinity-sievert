package util

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func ReadFile(fileName string) []byte {
	configFileContent, err := ioutil.ReadFile(GetConfigDir() + fileName)
	if err != nil {
		log.Fatal(err)
	}

	return configFileContent
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

func GetContentFromFileOrDefault(filePath string, defaultValue string) []byte {
	byteContent := []byte(defaultValue)
	if FileExists(filePath) {
		byteContent = ReadFile(filePath)
	}

	return byteContent
}

func GetConfigDir() string {
	confDir := os.Getenv("SIEVERT_CONF")
	if confDir == "" {
		usr, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		confDir = usr.HomeDir + "/.sievert"
	}

	return confDir
}
