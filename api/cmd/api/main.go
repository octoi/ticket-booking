package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/octoi/ticket-booking/config"
	"github.com/octoi/ticket-booking/db"
	"github.com/octoi/ticket-booking/handlers"
	"github.com/octoi/ticket-booking/repositories"
	"github.com/octoi/ticket-booking/services"
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
	ticketRepository := repositories.NewTicketRepository(db)
	authRepository := repositories.NewAuthRepository(db)

	// Service
	authService := services.NewAuthService(authRepository)

	// Routing
	server := app.Group("/api")

	// Handlers
	handlers.NewEventHandler(server.Group("/event"), eventRepository)
	handlers.NewTicketHandler(server.Group("/event"), ticketRepository)

	app.Listen(fmt.Sprintf(":" + envConfig.ServerPort))
}
