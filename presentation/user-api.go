package presentation

import (
	"context"
	"golang-clean/domain"

	"github.com/gofiber/fiber/v2"
)

var app = fiber.New()

type UserApi struct {
	userUsecase domain.UserUsecase
}

func NewUserApi(userUsecase domain.UserUsecase) UserApi {
	return UserApi{
		userUsecase: userUsecase,
	}
}

func (u *UserApi) Router(ctx context.Context, app *fiber.App) *fiber.App {
	app.Get("/user", func(c *fiber.Ctx) error {
		users, err := u.userUsecase.GetAll(ctx)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "An error occurred while getting the user"})
		}

		return c.Status(fiber.StatusOK).JSON(users)
	})

	app.Get("/user/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		users, err := u.userUsecase.Get(ctx, id)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "An error occurred while getting the user"})
		}

		return c.Status(fiber.StatusOK).JSON(users)
	})

	app.Post("/user", func(c *fiber.Ctx) error {
		var userDTO domain.UserDTO
		if err := c.BodyParser(&userDTO); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err})
		}
		err := u.userUsecase.Create(ctx, userDTO)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "created"})
	})

	return app
}
