package handler

import (
	"hexagonal_v2/internal/core/domain"
	"hexagonal_v2/internal/core/port"
	"strconv"

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

func (a *airpodHandler) GetAirpodByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID format"})
	}

	airpod, err := a.airpod.GetAirpodByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(airpod)
}

func (a *airpodHandler) GetAirpods(c *fiber.Ctx) error {
	airpods, err := a.airpod.GetAirpods()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(airpods)
}

func (a *airpodHandler) GetAirpodByUserID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID format"})
	}

	airpods, err := a.airpod.GetAirpodByUserID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(airpods)
}

func (a *airpodHandler) UpdateAirpod(c *fiber.Ctx) error {
	var airpod domain.Airpod
	if err := c.BodyParser(&airpod); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := a.airpod.UpdateAirpod(&airpod); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(airpod)
}

func (a *airpodHandler) DeleteAirpod(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID format"})
	}

	if err := a.airpod.DeleteAirpod(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
