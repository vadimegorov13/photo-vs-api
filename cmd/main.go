package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vadimegorov13/photo-vs-api/pkg/common/config"
	"github.com/vadimegorov13/photo-vs-api/pkg/common/db"
	"github.com/vadimegorov13/photo-vs-api/pkg/tournament"
)

func main() {
	// Initialize the database connection
	dbHandler := db.InitDB()

	// Set up Fiber configuration
	fiberConf := config.FiberConfig()

	// Set up logger configuration and open the log file
	logConf, file := config.LoggerConfig()
	defer file.Close()

	// Create a new Fiber app with the specified configuration
	app := fiber.New(fiberConf)

	// Use the logger middleware with the specified configuration
	app.Use(logger.New(logConf))

	// Register routes for the Post handler with the initialized database connection
	tournament.RegisterRoutes(app, dbHandler)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
