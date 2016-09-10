package action

import (
	"fmt"
	"os"

	"github.com/fer2d2/sievert/config"
	"github.com/fer2d2/sievert/files"
	"github.com/fer2d2/sievert/logger"
)

func CreateConfigDir() {
	confDir, err := config.ProjectRoot()

	if err != nil {
		logger.Log.Fatal(err)
	}

	err = os.Mkdir(confDir, 0744)
	if err != nil {
		logger.Log.Fatal("Error creating " + confDir + ": " + err.Error())
	} else {
		logger.FileCreated(confDir)
	}
}

func CreateSievertYmlFile() {
	sievertFile, _ := files.NewFileSievert()
	fmt.Printf("%v", sievertFile.Content)
	sievertFile.Write()
	// util.LogFileCreated(sievertFile.Path)
}
