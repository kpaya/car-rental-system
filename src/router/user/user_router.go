package user

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kpaya/car-rental-system/src/infra/middleware"
	"github.com/kpaya/car-rental-system/src/repository"
	router_dto "github.com/kpaya/car-rental-system/src/router"
	usecase "github.com/kpaya/car-rental-system/src/usecase/users"
	"github.com/kpaya/car-rental-system/src/usecase/users/dto"
)

func UserRouterInitializer(commons *router_dto.CommonsBundle) error {

	userGroupRouter := commons.FiberApp.Group("/user")

	createUserRouterGroup := userGroupRouter.Group("/create")

	createUserRouterGroup.Post("/member", func(c *fiber.Ctx) error {
		var input dto.InputCreateANewUserDTO
		if err := c.BodyParser(&input); err != nil {
			code := fiber.StatusBadRequest
			return c.Status(code).JSON(fiber.Map{
				"msg":  err.Error(),
				"code": code,
			})
		}

		input.Type = "M"
		repository := repository.NewUserRepository(commons.Db)
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

	createUserRouterGroup.Post("/receptionist", middleware.UserValidationMiddleware, func(c *fiber.Ctx) error {

		var input dto.InputCreateANewUserDTO
		if err := c.BodyParser(&input); err != nil {
			code := fiber.StatusBadRequest
			return c.Status(code).JSON(fiber.Map{
				"msg":  err.Error(),
				"code": code,
			})
		}
		input.Type = "R"
		repository := repository.NewUserRepository(commons.Db)
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

	userGroupRouter.Get("/:id<guid>", middleware.UserValidationMiddleware, func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			code := fiber.StatusBadRequest
			return c.Status(code).JSON(fiber.Map{
				"msg":  fmt.Errorf(`error: %s`, errors.New("you must provide a valid id")).Error(),
				"code": code,
			})
		}
		repository := repository.NewUserRepository(commons.Db)
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

	userGroupRouter.Get("/list", middleware.UserValidationMiddleware, func(c *fiber.Ctx) error {
		repository := repository.NewUserRepository(commons.Db)
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
	return nil
}
