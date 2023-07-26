package vehicle

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/kpaya/car-rental-system/src/infra/middleware"
	"github.com/kpaya/car-rental-system/src/repository"
	router_dto "github.com/kpaya/car-rental-system/src/router"
	usecase "github.com/kpaya/car-rental-system/src/usecase/users"
	"github.com/kpaya/car-rental-system/src/usecase/users/dto"
)

func VehicleRouterInitializer(commons *router_dto.CommonsBundle) error {

	vehicleRouterGroup := commons.FiberApp.Group("/vehicle", middleware.UserValidationMiddleware)

	vehicleRouterGroup.Post("/create", func(c *fiber.Ctx) error {
		var inputDto dto.InputCreateAVehicleDTO
		repository := repository.NewVehicleRepository(commons.Db)
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

	return nil
}
