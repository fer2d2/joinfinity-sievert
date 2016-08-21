package util

import log "github.com/Sirupsen/logrus"

var Logger = log.New()

func init() {
	Logger.Level = log.DebugLevel
}

func LogFileCreated(path string) {
	Logger.WithFields(log.Fields{
		"path": path,
	}).Info("Successfully created")
}

func LogRunningCommand(command string, args []string) {
	Logger.WithFields(log.Fields{
		"command": command,
		"args":    args,
	}).Info("Executing command")
}
