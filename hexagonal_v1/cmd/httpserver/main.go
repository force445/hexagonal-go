package main

import (
	"hexagonal/internal/core/services"
	"hexagonal/internal/handlers/gamehdl"
	"hexagonal/internal/repositories/gamesrepo"
	"hexagonal/pkg/uidgen"

	"github.com/gofiber/fiber/v2"
)

func main() {

	gameRepository := gamesrepo.NewMemKVS()
	gameService := services.New(gameRepository, uidgen.New())
	gameHandler := gamehdl.NewHTTPHandler(gameService)

	router := fiber.New()
	router.Get("/games/:id", gameHandler.Get)

	router.Listen(":3000")
}
