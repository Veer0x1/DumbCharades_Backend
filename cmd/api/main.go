package main

import (
	"log"

	"github.com/Veer0x1/DumbCharades/config"
	"github.com/Veer0x1/DumbCharades/internal/handler"
	"github.com/Veer0x1/DumbCharades/pkg/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.LoadConfig()

	database, err := db.ConnectPostgres(config.Config.DB_URL)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	app := fiber.New()

	handler.SetupRoutes(app, database)
	log.Fatalf("%v", app.Listen(":"+config.Config.PORT))
}
