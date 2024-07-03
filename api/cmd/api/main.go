package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/octoi/ticket-booking/config"
	"github.com/octoi/ticket-booking/db"
	"github.com/octoi/ticket-booking/handlers"
	"github.com/octoi/ticket-booking/repositories"
)

func main() {
	envConfig := config.NewEnvConfig()

	db := db.Init(envConfig, db.DBMigrator)

	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(db)

	// Repositories
	server := app.Group("/api")

	// Repositories
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
