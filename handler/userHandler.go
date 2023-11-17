package handler

import (
	"user-product-management/app/auth"
	"user-product-management/app/users"
	"user-product-management/helper"

	"github.com/gofiber/fiber/v2"
)

type userHandler struct {
	userService users.UserService
	authService auth.Service
}

func NewUserHandler(userService users.UserService, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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

	user, err := h.userService.Register(input)
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

	user, err := h.userService.Login(input)
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	// generate token
	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	user.Token = token

	return helper.Response(c, 200, "login success", user)
}

func (h *userHandler) FindAll(c *fiber.Ctx) error {

	// fmt.Println(c.Locals("userID"))
	users, err := h.userService.FindAll()
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}
	return helper.Response(c, 200, "find all user", users)
}
