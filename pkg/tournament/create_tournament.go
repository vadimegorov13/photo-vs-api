package tournament

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/vadimegorov13/photo-vs-api/pkg/common/models"
)

func (h handler) CreateTournament(c *fiber.Ctx) error {
	body := c.Body()

	var lead models.Tournament

	if err := json.Unmarshal(body, &lead); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if result := h.DB.Create(&lead); result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}

	return c.JSON(&lead)
}
