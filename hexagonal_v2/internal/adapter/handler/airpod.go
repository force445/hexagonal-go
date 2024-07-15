package handler

import (
	"hexagonal_v2/internal/core/domain"
	"hexagonal_v2/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type airpodHandler struct {
	airpod port.AirpodService
}

func NewAirpodHandler(airpod port.AirpodService) *airpodHandler {
	return &airpodHandler{airpod: airpod}
}

func (a *airpodHandler) CreateAirpod(c *fiber.Ctx) error {
	var airpod domain.Airpod
	if err := c.BodyParser(&airpod); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := a.airpod.CreateAirpod(&airpod); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(airpod)
}
