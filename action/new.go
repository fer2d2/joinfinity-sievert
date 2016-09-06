package action

// import (
// 	"os"
//
// 	"github.com/fer2d2/sievert/files"
// 	"github.com/fer2d2/sievert/util"
// )
//
// func RequireConfigDirNotCreated() {
// 	if util.DirExists(confDir) {
// 		util.Logger.Fatal("Sievert config directory exists. Check your " +
// 			"$SIEVERT_CONF value.")
// 	}
// }
//
// func CreateConfigDir() {
// 	err := os.Mkdir(confDir, 0744)
// 	if err != nil {
// 		util.Logger.Fatal("Error creating " + confDir + ": " + err.Error())
// 	} else {
// 		util.LogFileCreated(confDir)
// 	}
// }
//
// func CreateSievertYmlFile() {
// 	sievertFile := files.NewFileSievert()
// 	sievertFile.WriteStubFile()
//
// 	util.LogFileCreated(sievertFile.Path)
// }
