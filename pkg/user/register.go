package user

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/vadimegorov13/photo-vs-api/pkg/common/models"
)

func (h handler) RegisterUser(c *fiber.Ctx) error {
	body := c.Body()

	var user models.User

	if err := json.Unmarshal(body, &user); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if result := h.DB.Create(&user); result.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, result.Error.Error())
	}

	return c.JSON(&user)
}
