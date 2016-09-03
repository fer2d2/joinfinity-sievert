package stubs

import "github.com/fer2d2/sievert/util"

// Exported constatns to retrieve the correct stub from the factory
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

	// Register()
	// Register()

}

var stubs = make(map[string]string)

func register(name string, content string) {
	_, registered := stubs[name]
	if registered {
		util.Logger.Fatalf("Stub already registered for key '%s'", name)
	}

	stubs[name] = content
}

// Get returns the stub for a specified registered text template
func Get(name string) string {
	content, registered := stubs[name]
	if !registered {
		util.Logger.Fatalf("Stub not found for key '%s'", name)
	}

	return content
}
