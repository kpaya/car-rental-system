package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kpaya/car-rental-system/src/entity"
	"github.com/kpaya/car-rental-system/src/infra/database"
	"github.com/kpaya/car-rental-system/src/repository"
	usecase "github.com/kpaya/car-rental-system/src/usecase/users"
	"github.com/kpaya/car-rental-system/src/usecase/users/dto"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err.Error())
	}

	Db = database.NewDb()
}

func main() {
	app := fiber.New()

	app.Post("/user", func(c *fiber.Ctx) error {
		user := entity.User{}
		json.Unmarshal(c.Body(), &user)
		if err := user.Validate(); err != nil {
			return c.Status(404).JSON(fiber.Map{
				"msg":  fmt.Errorf(`error: %s`, err).Error(),
				"code": 404,
			})
		}
		repository := repository.NewUserRepository(Db)
		usecase := usecase.NewCreateANewUserUseCase(repository)
		output, err := usecase.Execute(dto.InputCreateANewUserDTO{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			Status:   entity.Active,
			Phone:    user.Phone,
		})
		if err != nil {
			return c.Status(404).JSON(fiber.Map{
				"msg":  fmt.Errorf(`error: %s`, err).Error(),
				"code": 404,
			})
		}
		c.JSON(output)
		return nil
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"code": "Oka",
		})
	})

	app.Listen(":8081")

}
