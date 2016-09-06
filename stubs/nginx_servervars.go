package stubs

// ServerVars represents the values to replace when parsing an NGINX Server
// template
type ServerVars struct {
	Port       uint16
	SslPort    uint16
	ServerName string
	ServerRoot string
}
