package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/kpaya/car-rental-system/src/infra/database"
	router "github.com/kpaya/car-rental-system/src/router"
	access_router "github.com/kpaya/car-rental-system/src/router/access"
	user_router "github.com/kpaya/car-rental-system/src/router/user"
	vehicle_router "github.com/kpaya/car-rental-system/src/router/vehicle"
)

var Db *sql.DB

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic(err.Error())
	}

	Db = database.NewDb()
}

func main() {
	app := fiber.New()

	dataCommon := &router.CommonsBundle{
		Db:       Db,
		FiberApp: app,
	}

	user_router.UserRouterInitializer(dataCommon)
	vehicle_router.VehicleRouterInitializer(dataCommon)
	access_router.AccessRouterInitializer(dataCommon)

	log.Panic(app.Listen(":8081"))

}
