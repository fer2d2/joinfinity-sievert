package config

import (
	"os"
	"os/user"
)

// ProjectRoot returns the directory where the current sievert's project lives in
func ProjectRoot() (string, error) {
	confDir := os.Getenv("SIEVERT_CONF")

	if confDir == "" {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}

		confDir = usr.HomeDir + "/.sievert"
	}

	return confDir, nil
}
