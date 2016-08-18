package config

// TODO en la carpeta STUBS dejamos las cadenas de texto tal cual, y en la
// carpeta config las cargamos en sus respectivas structs

// AbstractFile represents a configuration file to be loaded or written into
// disk
type AbstractFile struct {
	Filename string
	Content  string
}

type ConfigFile interface {
	FileExists() bool
	WriteStubFile()
	LoadFromDisk()
}

type YamlFile interface {
	Parse()
}

// TODO mover todo esto

// NewAppLogo creates a new instance of AppLogo
func NewAppLogo() *AppLogo {
	appLogo := &AppLogo{AbstractFile{filename, appLogoTpl}}

	return appLogo
}

func (appLogo *AppLogo) fileExists() bool {
	return false
}
