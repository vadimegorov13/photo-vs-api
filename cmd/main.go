package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/vadimegorov13/photo-vs-api/pkg/common/config"
	"github.com/vadimegorov13/photo-vs-api/pkg/common/db"
	"github.com/vadimegorov13/photo-vs-api/pkg/routes/user"
)

func main() {
	// Initialize the database connection
	dbHandler := db.InitDB()

	firebaseAuth := config.SetupFirebase()

	// Set up Fiber configuration
	fiberConf := config.FiberConfig()

	// Set up logger configuration and open the log file
	logConf, file := config.LoggerConfig()
	defer file.Close()

	// Create a new Fiber app with the specified configuration
	app := fiber.New(fiberConf)

	// Middleware function to set db and firebase auth to Fiber context.
	app.Use(func(c *fiber.Ctx) error {
		// Set the database connection and Firebase auth object to the context.
		c.Locals("firebaseAuth", firebaseAuth)

		// Call the next middleware function in the chain.
		return c.Next()
	})

	corsConf := cors.Config{
		AllowOrigins: "http://127.0.0.1:5174",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	}

	app.Use(cors.New(corsConf))

	// Use the logger middleware with the specified configuration
	app.Use(logger.New(logConf))

	// Register routes for the Post handler with the initialized database connection
	user.RegisterRoutes(app, dbHandler)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
