package dto

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type CommonsBundle struct {
	Db       *sql.DB
	FiberApp *fiber.App
}
