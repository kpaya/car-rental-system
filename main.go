package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
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
		var input dto.InputCreateANewUserDTO
		// if err := c.BodyParser(&input); err != nil {
		if err := json.Unmarshal(c.Body(), &input); err != nil {
			code := fiber.StatusBadRequest
			return c.Status(code).JSON(fiber.Map{
				"msg":  fmt.Errorf(`error: %s`, err).Error(),
				"code": code,
			})
		}
		repository := repository.NewUserRepository(Db)
		usecase := usecase.NewCreateANewUserUseCase(repository)
		output, err := usecase.Execute(input)
		if err != nil {
			code := fiber.StatusBadRequest
			return c.Status(code).JSON(fiber.Map{
				"msg":  fmt.Errorf(`error: %s`, err).Error(),
				"code": code,
			})
		}
		c.Status(fiber.StatusCreated).JSON(output)
		return nil
	})

	app.Get("/user/:id<guid>", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			code := fiber.StatusBadRequest
			return c.Status(code).JSON(fiber.Map{
				"msg":  fmt.Errorf(`error: %s`, errors.New("you must provide a valide id")).Error(),
				"code": code,
			})
		}
		repository := repository.NewUserRepository(Db)
		usecase := usecase.NewFindAUserByIdUseCase(repository)
		output, err := usecase.Execute(dto.InputFindAUserByIdDTO{Id: id})
		if err != nil {
			code := fiber.StatusBadRequest
			return c.Status(code).JSON(fiber.Map{
				"msg":  fmt.Errorf(`error: %s`, err).Error(),
				"code": code,
			})
		}
		c.Status(fiber.StatusOK).JSON(output)
		return nil
	})

	app.Get("/user", func(c *fiber.Ctx) error {
		repository := repository.NewUserRepository(Db)
		usecase := usecase.NewListUserUseCase(repository)

		output, err := usecase.Execute()
		if err != nil {
			code := fiber.StatusInternalServerError
			return c.Status(code).JSON(fiber.Map{
				"msg":  fmt.Errorf(`error: %s`, err.Error()),
				"code": code,
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"result": output,
			"IP":     c.IP(),
		})
	})

	app.Post("/vehicle", func(c *fiber.Ctx) error {
		var inputDto dto.InputCreateAVehicleDTO
		repository := repository.NewVehicleRepository(Db)
		usecase := usecase.NewCreateVehicleUseCase(repository)

		err := json.Unmarshal(c.Body(), &inputDto)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg":  err.Error(),
				"code": fiber.StatusBadRequest,
			})
		}
		output := usecase.Execute(inputDto)

		return c.JSON(output)

	})

	log.Panic(app.Listen(":8081"))

}
