package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/natevaub/focus-companion/backend/db/api/dto"
	"github.com/natevaub/focus-companion/backend/db/api/service"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var user dto.CreateUserRequest
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	newUser, err := h.userService.CreateUser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(201).JSON(newUser)
}
