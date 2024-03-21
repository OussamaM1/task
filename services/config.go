package services

import (
	"errors"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var enableLogging bool

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

// Logf logs the message if logging is enabled
func Logf(format string, args ...interface{}) {
	if enableLogging {
		log.Printf(format, args...)
	}
}
