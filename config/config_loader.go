package config

//
// import (
// 	"log"
//
// 	"gopkg.in/yaml.v2"
// )
//
// const userConfFile string = "/sievert.yml"
//
// type multiSequence []map[string]interface{}
//
// // Config represents user's configuration specified in the sievert.yml file
// type Config struct {
// 	IP       string        `yaml:"ip"`
// 	Services []string      `yaml:"services"`
// 	Folders  multiSequence `yaml:"folders"`
// 	Sites    multiSequence `yaml:"sites"`
// }
//
// // NewConfig creates new config from user's sievert.yml file
// func NewConfig() *Config {
// 	config := &Config{}
// 	config.initializeFromFile()
//
// 	return config
// }
//
// func (config *Config) initializeFromFile() {
// 	err := yaml.Unmarshal(readFile(userConfFile), config)
//
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
