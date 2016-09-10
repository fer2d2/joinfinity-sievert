package files

type FileNginxServer struct {
	FileCommon
	Content string
}

func NewFileNginxServer(server GenericYaml) *FileNginxServer {
	// domain := fmt.Sprint(server["map"])
	// folder := fmt.Sprint(server["to"])
	// template := fmt.Sprint(server["template"])
	//
	// logger.Log.WithFields(logrus.Fields{
	// 	"map":      domain,
	// 	"to":       folder,
	// 	"template": template,
	// }).Info("Creating new site")
	//
	// filePath := config.ProjectRoot() + "/servers/" + domain + ".conf"
	// fileStub := stubs.Get(stubs.ServersPrefix + template)
	//
	fileNginxServer := new(FileNginxServer)
	// fileNginxServer.FileCommon = FileCommon{filePath, fileStub}
	//
	// fileNginxServer.InitFromTemplate(domain, folder)
	//
	return fileNginxServer
}

func (fileNginxServer *FileNginxServer) InitFromTemplate(domain string, folder string) {
	// doc := new(bytes.Buffer)
	// serverVars := stubs.ServerVars{
	// 	Port:       80,
	// 	SslPort:    443,
	// 	ServerName: domain,
	// 	ServerRoot: folder,
	// }
	//
	// template, err := template.New("NginxServer").
	// 	Parse(fileNginxServer.Stub)
	//
	// if err != nil {
	// 	logger.Log.Fatal(err)
	// }
	//
	// template.Execute(doc, serverVars)
	//
	// fileNginxServer.Content = doc.String()
}
