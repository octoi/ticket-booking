package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/octoi/ticket-booking/handlers"
	"github.com/octoi/ticket-booking/repositories"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "TicketBooking",
		ServerHeader: "Fiber",
	})

	// Repositories
	eventRepository := repositories.NewEventRepository(nil)

	// Repositories
	server := app.Group("/api")

	// Repositories
	handlers.NewEventHandler(server.Group("/event"), eventRepository)

	app.Listen(":3000")
}
