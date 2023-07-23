package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
    // Create a new logger instance
    log = logrus.New()

    // Set logger formatter to JSON format
    log.SetFormatter(&logrus.JSONFormatter{})

    // Set the output to standard output (you can also log to a file, etc.)
    log.SetOutput(os.Stdout)

    // Set the log level. You can set it to log only errors, warnings, info, or everything (debug)
    log.SetLevel(logrus.InfoLevel)
}

// SetLogLevel allows you to set the log level dynamically
func SetLogLevel(level logrus.Level) {
    log.SetLevel(level)
}

// LogError logs an error message
func LogError(message string, err error) {
    log.WithFields(logrus.Fields{
        "error": err.Error(),
    }).Error(message)
}

// LogInfo logs an informational message
func LogInfo(message string) {
    log.Info(message)
}

// LogWarning logs a warning message
func LogWarning(message string) {
    log.Warning(message)
}

// LogDebug logs a debug message
func LogDebug(message string) {
    log.Debug(message)
}
