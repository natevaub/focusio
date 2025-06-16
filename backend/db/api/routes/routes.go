package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natevaub/focus-companion/backend/db/api/handlers"
)

// RegisterRoutes sets up all the API routes
func RegisterRoutes(app *fiber.App, userHandler *handlers.UserHandler) {
	// User routes
	users := app.Group("/api/users")
	users.Post("/", userHandler.CreateUser)
}
