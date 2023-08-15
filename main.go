package main

import (
	"context"
	"golang-clean/application/repository"
	"golang-clean/application/usecase"
	"golang-clean/infrastructure"
	"golang-clean/presentation"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx := context.Background()
	mongodb := infrastructure.NewMongodb(ctx, "mongodb://localhost:27017", "user", "users")
	user_repository := repository.NewRepository(ctx, mongodb)
	user_usecase := usecase.NewUserUsecase(ctx, user_repository)
	user_api := presentation.NewUserApi(user_usecase)
	app := user_api.Router(ctx, fiber.New())
	log.Fatal(app.Listen(":8080"))
}
