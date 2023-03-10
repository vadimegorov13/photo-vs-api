package tournament

import (
	"github.com/gofiber/fiber/v2"
)

func (h handler) GetTournamentData(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(id)
}
