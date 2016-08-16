package config

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

func readFile(fileName string) []byte {
	configFileContent, err := ioutil.ReadFile(getConfigDir() + fileName)
	if err != nil {
		log.Fatal(err)
	}

	return configFileContent
}

func getConfigDir() string {
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
