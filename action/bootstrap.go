package action

import (
	"bytes"
	"fmt"
	"html/template"

	log "github.com/Sirupsen/logrus"
	"github.com/fer2d2/sievert/files"
	"github.com/fer2d2/sievert/stubs"
	"github.com/fer2d2/sievert/util"
)

func GenerateProjectFiles() {
}

func GenerateDockerComposeFile() {
	// sievertFile := files.NewFileSievert()
	// dockerComposeFile := files.NewDockerComposeFile()
	//
	// sievertFile.
}

func GenerateNginxServers() {
	sievertFile := files.NewFileSievert()

	sievertContent := sievertFile.Content

	for _, site := range sievertContent.Sites {
		parsedServer := parseNginxServerTemplate(site)
		fmt.Println(parsedServer)
	}
}

func parseNginxServerTemplate(site map[string]interface{}) string {
	util.Logger.WithFields(log.Fields{
		"map":      site["map"],
		"to":       site["to"],
		"template": site["template"],
	}).Info("Found site")

	doc := new(bytes.Buffer)
	serverVars := stubs.ServerVars{
		Port:       80,
		SslPort:    443,
		ServerName: fmt.Sprint(site["map"]),
		ServerRoot: fmt.Sprint(site["to"]),
	}

	template, err := template.New("NginxServer").
		Parse(stubs.Get(stubs.ServersPrefix + fmt.Sprint(site["template"])))

	if err != nil {
		util.Logger.Fatal(err)
	}

	template.Execute(doc, serverVars)

	return doc.String()
}
