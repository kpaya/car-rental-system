package dto

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/kpaya/car-rental-system/src/service"
)

type CommonsBundle struct {
	Db       *sql.DB
	Jwt      *service.JWTService
	FiberApp *fiber.App
}
