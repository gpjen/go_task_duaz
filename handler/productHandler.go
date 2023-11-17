package handler

import (
	"strconv"
	"user-product-management/app/products"
	"user-product-management/app/users"
	"user-product-management/helper"

	"github.com/gofiber/fiber/v2"
)

type productHandler struct {
	productService products.Service
}

func NewProductHandler(productService products.Service) *productHandler {
	return &productHandler{productService}
}

func (h *productHandler) CreateNewProduct(c *fiber.Ctx) error {
	input := new(products.ProductRegister)

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

	product, err := h.productService.Create(*input, user.ID)
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	return helper.Response(c, 200, "create new product", product)
}

func (h *productHandler) FindAll(c *fiber.Ctx) error {
	products, err := h.productService.FindAll()
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	return helper.Response(c, 200, "get all products", products)
}

func (h *productHandler) FindByID(c *fiber.Ctx) error {

	stringId := c.Params("id")

	dataId, err := strconv.Atoi(stringId)
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	product, err := h.productService.FindByID(uint(dataId))
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	return helper.Response(c, 200, "get product by id", product)

}

func (h *productHandler) UpdateProduct(c *fiber.Ctx) error {

	stringId := c.Params("id")

	dataId, err := strconv.Atoi(stringId)
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	input := new(products.ProductUpdate)

	if err := c.BodyParser(input); err != nil {
		return helper.Response(c, 400, "Invalid JSON", err.Error())
	}

	validateList := helper.ValidateInput(input)
	if validateList != nil {
		return helper.Response(c, 422, "validation error", validateList)
	}

	userContex, ok := c.Locals("user").(users.UserContex)
	if !ok {
		return helper.Response(c, 500, "internal server error", userContex)
	}

	productUpdate, err := h.productService.UpdateProduct(*input, uint(dataId), userContex)
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	return helper.Response(c, 200, "update product", productUpdate)
}

func (h *productHandler) DeleteProduct(c *fiber.Ctx) error {

	stringId := c.Params("id")

	dataId, err := strconv.Atoi(stringId)
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	userContex, ok := c.Locals("user").(users.UserContex)
	if !ok {
		return helper.Response(c, 500, "internal server error", userContex)
	}

	productDelete, err := h.productService.DeleteProduct(uint(dataId), userContex)
	if err != nil {
		return helper.Response(c, 400, err.Error(), nil)
	}

	return helper.Response(c, 200, "delete product", productDelete)
}
