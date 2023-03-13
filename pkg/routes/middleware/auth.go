package middleware

import (
	"context"
	"net/http"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/gofiber/fiber/v2"
)

// AuthMiddleware: to verify all authorized operations
func AuthMiddleware(c *fiber.Ctx) error {
	firebaseAuth := c.Locals("firebaseAuth").(*auth.Client)
	authorizationToken := c.Get("Authorization")
	idToken := strings.TrimSpace(strings.Replace(authorizationToken, "Bearer", "", 1))

	if idToken == "" {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Id token not available"})
	}

	// Verify token
	token, err := firebaseAuth.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid token"})
	}

	// Get user data from token
	userID := token.UID
	userDisplayName := token.Claims["name"].(string)
	userEmail := token.Claims["email"].(string)

	// Set user data in fiber context locals
	c.Locals("userID", userID)
	c.Locals("userDisplayName", userDisplayName)
	c.Locals("userEmail", userEmail)
	c.Locals("idToken", idToken)

	return c.Next()
}
