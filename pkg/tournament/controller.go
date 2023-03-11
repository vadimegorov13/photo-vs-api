package tournament

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

/*
The handler type contains the database connection.
*/
type handler struct {
	DB *gorm.DB
}

/*
RegisterRoutes creates a new handler and registers the API routes with the given Fiber app.
@param {*fiber.App} app - The Fiber app instance.
@param {*gorm.DB} db - The GORM database connection.
*/
func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	// Define the API routes
	api := app.Group("/api")
	v1 := api.Group("/v1")
	tournament := v1.Group("/tournament")

	// Register the route handlers
	tournament.Post("/", h.CreateTournament)
	tournament.Get("/:id", h.GetTournamentData)
	tournament.Patch("/:id", h.DelteTournament)
}
