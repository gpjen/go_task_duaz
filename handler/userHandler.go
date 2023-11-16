package handler

import (
	"user-product-management/app/users"
	"user-product-management/helper"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	UserService users.UserService
}

func NewUserHandler(userService users.UserService) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) Register(c *fiber.Ctx) error {

	input := new(users.UserRegister)

	if err := c.BodyParser(input); err != nil {
		return helper.Response(c, 400, "Invalid JSON", err.Error())
	}

	validateList := helper.ValidateInput(input)
	if validateList != nil {
		return helper.Response(c, fiber.StatusUnprocessableEntity, "validation error", validateList)
	}

	user, err := h.UserService.Register(input)
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	return helper.Response(c, 200, "register new user", user)
}

func (h *userHandler) Login(c *fiber.Ctx) error {

	input := new(users.UserLogin)

	if err := c.BodyParser(input); err != nil {
		return helper.Response(c, 400, "Invalid JSON", err.Error())
	}

	validateList := helper.ValidateInput(input)
	if validateList != nil {
		return helper.Response(c, fiber.StatusUnprocessableEntity, "validation error", validateList)
	}

	user, err := h.UserService.Login(input)
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	return helper.Response(c, 200, "login success", user)
}
