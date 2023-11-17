package main

import (
	"log"
	"user-product-management/app/auth"
	"user-product-management/app/middleware"
	"user-product-management/app/users"
	"user-product-management/db"
	"user-product-management/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	db.InitMysqlDB()

	apiV1 := app.Group("v1")

	userRepository := users.NewUserRepository(db.DB)

	userServices := users.NewUserService(userRepository)
	authServices := auth.NewService()

	userhandler := handler.NewUserHandler(userServices, authServices)

	apiV1.Post("/user", userhandler.Register)
	apiV1.Get("/users", middleware.AuthMiddleware(userServices, authServices), userhandler.FindAll)
	apiV1.Put("/user/:id", middleware.AuthMiddleware(userServices, authServices), userhandler.UpdateUser)
	apiV1.Post("/login", userhandler.Login)

	app.Listen(":3000")
}
