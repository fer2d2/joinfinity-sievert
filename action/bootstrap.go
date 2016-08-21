package action

import (
	"fmt"
	"os"

	"github.com/fer2d2/sievert/files"
	"github.com/fer2d2/sievert/util"
)

var confDir = util.GetConfigDir()

func BootstrapAction() {
	if util.DirExists(confDir) {
		util.Logger.Fatal("Sievert config directory exists. Try 'sievert " +
			"bootstrap --rebuild'")
	}

	createConfigDir()
	createSievertYmlFile()
	showNextStepsMessage()
}

func createConfigDir() {
	err := os.Mkdir(confDir, 0744)
	if err != nil {
		util.Logger.Fatal("Error creating " + confDir + ": " + err.Error())
	} else {
		util.LogFileCreated(confDir)
	}
}

func createSievertYmlFile() {
	sievertFile := files.NewFileSievert()
	sievertFile.WriteStubFile()

	util.LogFileCreated(sievertFile.Path)
}

func showNextStepsMessage() {
	fmt.Println("\nWhat to do next? Modify your sievert.yml config file and " +
		"run 'sievert provision'")
}
