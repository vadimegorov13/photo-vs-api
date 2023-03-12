package tournament

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vadimegorov13/photo-vs-api/pkg/common/models"
)

func (h handler) GetTournament(c *fiber.Ctx) error {
	id := c.Params("id")

	var tournament models.Tournament

	if result := h.DB.First(&tournament, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.JSON(&tournament)
}
