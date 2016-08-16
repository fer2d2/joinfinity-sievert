package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fer2d2/sievert/cmd"
	"github.com/fer2d2/sievert/stubs"
)

func main() {
	fmt.Println(stubs.AppLogoTpl)

	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(-1)
	}
}

// func main() {
// 	userConfig := config.NewConfig()
//
// 	dockerCompose := config.NewDockerCompose(userConfig)
//
// 	fmt.Printf("content: %v\n", dockerCompose)
//
// 	for _, site := range userConfig.Sites {
// 		fmt.Printf("map: %v\n", site["map"])
// 		fmt.Printf("to: %v\n", site["to"])
// 		fmt.Printf("template: %v\n", site["template"])
// 	}
//
// 	if userConfig.Folders == nil {
// 		fmt.Println("Folders is nil")
// 	}
//
// 	for _, folder := range userConfig.Folders {
// 		fmt.Printf("map: %v\n", folder["map"])
// 		fmt.Printf("to: %v\n", folder["to"])
// 	}
// }
