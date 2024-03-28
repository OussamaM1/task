// Package services provides functionality related to configuration and managing todos file.
package services

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var enableLogging bool

const (
	// white color reset
	colorReset = "\033[0m"
	// green color
	colorGreen = "\033[32m"
	// yellow color
	colorYellow = "\033[33m"
	// blue color
	colorBlue = "\033[34m"
	// gray color
	colorGray = "\033[90m"
)

// init service - loads environment variables from a .env file and sets the env variable.
func init() {
	if err := godotenv.Load(); err != nil {
		// If the .env file is not found, it's not considered an error
		var pathError *os.PathError
		if !errors.As(err, &pathError) {
			log.Fatalln("Error loading the .env file: ", err)
		}
	}
	env := os.Getenv("environment")
	enableLogging = env == "development"
}

// LogInfo func - logs info message if logging is enabled
func LogInfo(format string, args ...interface{}) {
	format = fmt.Sprintf("%s%s%s\n", colorGreen, format, colorReset)
	if enableLogging {
		log.Printf(format, args...)
	}
}

// LogWarn func - logs Warning message if logging is enabled
func LogWarn(format string, args ...interface{}) {
	format = fmt.Sprintf("%s%s%s\n", colorYellow, format, colorReset)
	if enableLogging {
		log.Printf(format, args...)
	}
}
