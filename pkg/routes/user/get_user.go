package user

import (
	"time"

	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/vadimegorov13/photo-vs-api/pkg/common/models"
)

func (h handler) SingInUser(c *fiber.Ctx) error {
	// get the user's UID from the context
	uid, ok := c.Locals("userID").(string)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}

	var user models.User

	err := h.DB.Where("id = ?", uid).FirstOrCreate(
		&user, models.User{
			Email:    c.Locals("userEmail").(string),
			Username: c.Locals("userDisplayName").(string),
			Basics:   models.Basics{Id: uid},
		},
	).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to get user data",
			"message": err.Error(),
		})
	}

	idToken := c.Locals("idToken").(string)
	firebaseAuth := c.Locals("firebaseAuth").(*auth.Client)

	session, err := firebaseAuth.SessionCookie(c.Context(), idToken, time.Duration(14*24*time.Hour))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to get user data",
			"message": err.Error(),
		})
	}

	// set session cookie
	cookie := fiber.Cookie{
		Name:     "session",
		Value:    session,
		HTTPOnly: true,
		SameSite: "Strict",
		Path:     "/",
		MaxAge:   60 * 60 * 24 * 14,
		Expires:  time.Now().Add(time.Duration(14 * 24 * time.Hour)),
	}
	c.Cookie(&cookie)

	return c.JSON(&user)
}
