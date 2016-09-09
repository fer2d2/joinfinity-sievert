package logger

import logrus "github.com/Sirupsen/logrus"

// Log is an instance of the logrus logger to be used across the application
var Log = log.New()

func init() {
	Log.Level = log.DebugLevel
}

// FileCreated logs the successfull creation of a file
func FileCreated(path string) {
	Log.WithFields(logrus.Fields{
		"path": path,
	}).Info("Successfully created")
}

// CommandCall logs the execution of a cli command
func CommandCall(command string, args []string) {
	Log.WithFields(logrus.Fields{
		"command": command,
		"args":    args,
	}).Info("Executing command")
}
