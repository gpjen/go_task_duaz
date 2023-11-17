package handler

import (
	"strconv"
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

func (h *userHandler) FindByID(c *fiber.Ctx) error {

	stringId := c.Params("id")

	dataId, err := strconv.Atoi(stringId)
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	user, err := h.userService.FindByID(uint(dataId))
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}
	return helper.Response(c, 200, "find user by id", user)

}

func (h *userHandler) UpdateUser(c *fiber.Ctx) error {
	stringId := c.Params("id")

	dataId, err := strconv.Atoi(stringId)
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	input := new(users.UserUpdate)
	if err := c.BodyParser(input); err != nil {
		return helper.Response(c, 400, "Invalid JSON", err.Error())
	}

	validateList := helper.ValidateInput(input)
	if validateList != nil {
		return helper.Response(c, 422, "validation error", validateList)
	}

	user, ok := c.Locals("user").(users.UserContex)
	if !ok {
		return helper.Response(c, 500, "internal server error", user)
	}

	if (uint(dataId) != user.ID) && (user.Role != "admin") {
		return helper.Response(c, 401, "only the owner can modify this user's data", nil)
	}

	userUpdate, err := h.userService.UpdateUser(*input, uint(dataId))
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	return helper.Response(c, 200, "update user success", userUpdate)
}

func (h *userHandler) DeleteUser(c *fiber.Ctx) error {

	stringId := c.Params("id")

	dataId, err := strconv.Atoi(stringId)
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	user, ok := c.Locals("user").(users.UserContex)
	if !ok {
		return helper.Response(c, 500, "internal server error", user)
	}

	if (uint(dataId) != user.ID) && (user.Role != "admin") {
		return helper.Response(c, 401, "only the owner can delete this user's data", nil)
	}

	userDelete, err := h.userService.DeleteUser(uint(dataId))
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	return helper.Response(c, 200, "delete user success", userDelete)
}
