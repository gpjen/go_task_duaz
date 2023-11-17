package main

import (
	"log"
	"user-product-management/app/auth"
	"user-product-management/app/middleware"
	"user-product-management/app/users"
	"user-product-management/db"
	"user-product-management/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowMethods: "GET, POST, PUT, DELETE",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	db.InitMysqlDB()

	// repository
	userRepository := users.NewUserRepository(db.DB)

	// service
	userServices := users.NewUserService(userRepository)
	authServices := auth.NewService()

	// handler
	userhandler := handler.NewUserHandler(userServices, authServices)

	// route
	apiV1 := app.Group("v1")

	apiV1.Post("/user", userhandler.Register)
	apiV1.Get("/users", middleware.AuthMiddleware(userServices, authServices), userhandler.FindAll)
	apiV1.Get("/user/:id", middleware.AuthMiddleware(userServices, authServices), userhandler.FindByID)
	apiV1.Put("/user/:id", middleware.AuthMiddleware(userServices, authServices), userhandler.UpdateUser)
	apiV1.Post("/login", userhandler.Login)

	app.Listen(":3000")
}
