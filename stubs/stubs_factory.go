package stubs

import "errors"

// Exported constants to retrieve the correct stub from the factory
const (
	ConfPrefix = "conf."

	Logo       = ConfPrefix + "logo"
	SievertYml = ConfPrefix + "sievert"
	ComposeYml = ConfPrefix + "compose"

	// Nginx servers

	ServersPrefix = "nginx.servers."

	NginxLaravel = ServersPrefix + "laravel"
	NginxSymfony = ServersPrefix + "symfony"

	// Docker compose partials

	ComposeServicesPrefix = "compose.services."

	Mysql         = ComposeServicesPrefix + "mysql"
	Sqlite        = ComposeServicesPrefix + "sqlite"
	Postgresql    = ComposeServicesPrefix + "postgresql"
	Elasticsearch = ComposeServicesPrefix + "elasticsearch"
	Redis         = ComposeServicesPrefix + "redis"
	Memcached     = ComposeServicesPrefix + "memcached"
	Mongodb       = ComposeServicesPrefix + "mongodb"
)

func init() {
	register(Logo, appLogoTpl)
	register(SievertYml, sievertYmlTpl)
	register(ComposeYml, composeYmlTpl)

	register(NginxLaravel, nginxLaravelTpl)
	register(NginxSymfony, nginxSymfonyTpl)

}

var stubs = make(map[string]string)

func register(name string, content string) error {
	_, registered := stubs[name]
	if registered {
		return errors.New("stubs: template already registered for key " + name)
	}

	stubs[name] = content

	return nil
}

// Get returns the stub for a specified registered text template
func Get(name string) (string, error) {
	content, registered := stubs[name]

	if !registered {
		return "", errors.New("stubs: template not found for key " + name)
	}

	return content, nil
}
