/*
Package config provides a function FiberConfig that returns a Fiber
configuration with predefined settings.
*/
package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

// Returns a logger configuration object and a file pointer for a Fiber application
func LoggerConfig() (logger.Config, *os.File) {
	// Create the logs directory if it doesn't exist
	logsDir := "logs"
	os.Mkdir(logsDir, os.ModePerm) // create the logs directory if it doesn't exist

	// Get the current date to create a new log file
	now := time.Now()
	filename := fmt.Sprintf("%s/%d-%02d-%02d.log", logsDir, now.Year(), now.Month(), now.Day())
	// Create or open the log file with read-write, create and append permissions
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// Configure the logger middleware with the log file as the output
	logConf := logger.Config{
		Output: file,
	}

	return logConf, file
}
