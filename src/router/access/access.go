package access

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/kpaya/car-rental-system/src/repository"
	router_dto "github.com/kpaya/car-rental-system/src/router"
	"github.com/kpaya/car-rental-system/src/service"
	usecase "github.com/kpaya/car-rental-system/src/usecase/users"
	"github.com/kpaya/car-rental-system/src/usecase/users/dto"
)

func AccessRouterInitializer(commons *router_dto.CommonsBundle) error {
	commons.FiberApp.Post("/token", func(c *fiber.Ctx) error {
		userRepository := repository.NewUserRepository(commons.Db)
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

		token, err := service.CreateJWTToAccess(userFound.ID, userFound.Name, userFound.Email)

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

	return nil
}
