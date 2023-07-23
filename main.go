package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
	"github.com/kpaya/car-rental-system/src/infra/database"
	"github.com/kpaya/car-rental-system/src/repository"
	"github.com/kpaya/car-rental-system/src/service"
	usecase "github.com/kpaya/car-rental-system/src/usecase/users"
	"github.com/kpaya/car-rental-system/src/usecase/users/dto"
	_ "github.com/lib/pq"
)

var Db *sql.DB
var JwtService *service.JWTService

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panic(err.Error())
	}

	Db = database.NewDb()

	JwtService = service.NewJWTService(jwt.RegisteredClaims{
		Issuer:   "RentalCarSystem",
		Subject:  "CreateJWTToRouterAccess",
		Audience: []string{"Users"},
	})
}

func main() {
	app := fiber.New()

	app.Post("/token", func(c *fiber.Ctx) error {
		userRepository := repository.NewUserRepository(Db)
		findUserUseCase := usecase.NewFindUserByEmailAndPasswordUseCase(userRepository)
		var findUserInput dto.InputFindUserByEmailAndPasswordDTO

		json.Unmarshal(c.Body(), &findUserInput)

		userFound, err := findUserUseCase.Execute(findUserInput)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg":  err.Error(),
				"code": fiber.StatusBadRequest,
			})
		}

		token, err := JwtService.CreateJWTToAccess(userFound.ID, userFound.Name, userFound.Email)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"msg":  err.Error(),
				"code": fiber.StatusBadRequest,
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"code": token,
		})
	})

	app.Post("/user/create", func(c *fiber.Ctx) error {
		var input dto.InputCreateANewUserDTO
		if err := c.BodyParser(&input); err != nil {
			code := fiber.StatusBadRequest
			return c.Status(code).JSON(fiber.Map{
				"msg":  err.Error(),
				"code": code,
			})
		}
		repository := repository.NewUserRepository(Db)
		usecase := usecase.NewCreateANewUserUseCase(repository)
		output, err := usecase.Execute(input)
		if err != nil {
			code := fiber.StatusBadRequest
			return c.Status(code).JSON(fiber.Map{
				"msg":  err.Error(),
				"code": code,
			})
		}
		c.Status(fiber.StatusCreated).JSON(output)
		return nil
	})

	userGroupRouter := app.Group("/user", JwtService.ValidateJWTToAccess)

	userGroupRouter.Get("/:id<guid>", func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			code := fiber.StatusBadRequest
			return c.Status(code).JSON(fiber.Map{
				"msg":  fmt.Errorf(`error: %s`, errors.New("you must provide a valid id")).Error(),
				"code": code,
			})
		}
		repository := repository.NewUserRepository(Db)
		usecase := usecase.NewFindAUserByIdUseCase(repository)
		output, err := usecase.Execute(dto.InputFindAUserByIdDTO{Id: id})
		if err != nil {
			code := fiber.StatusBadRequest
			return c.Status(code).JSON(fiber.Map{
				"msg":  err.Error(),
				"code": code,
			})
		}
		c.Status(fiber.StatusOK).JSON(output)
		return nil
	})

	userGroupRouter.Get("/list", func(c *fiber.Ctx) error {
		repository := repository.NewUserRepository(Db)
		usecase := usecase.NewListUserUseCase(repository)

		output, err := usecase.Execute()
		if err != nil {
			code := fiber.StatusInternalServerError
			return c.Status(code).JSON(fiber.Map{
				"msg":  err.Error(),
				"code": code,
			})
		}
		return c.Status(fiber.StatusOK).JSON(output)
	})

	vehicleRouterGroup := app.Group("/vehicle", JwtService.ValidateJWTToAccess)

	vehicleRouterGroup.Post("/create", func(c *fiber.Ctx) error {
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
