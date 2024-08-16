package handler

import (
	"hexagonal_v2/internal/core/domain"
	"hexagonal_v2/internal/core/port"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type locationHandler struct {
	location port.LocationService
}

func NewLocationHandler(location port.LocationService) *locationHandler {
	return &locationHandler{location: location}
}

func (l *locationHandler) CreateLocation(c *fiber.Ctx) error {
	var location domain.Location
	if err := c.BodyParser(&location); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := l.location.CreateLocation(&location); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(location)

}

func (l *locationHandler) GetLocationByID(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID format"})
	}

	location, err := l.location.GetLocationByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(location)
}

func (l *locationHandler) GetLocations(c *fiber.Ctx) error {
	locations, err := l.location.GetLocations()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(locations)
}

func (l *locationHandler) UpdateLocation(c *fiber.Ctx) error {
	var location domain.Location
	if err := c.BodyParser(&location); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := l.location.UpdateLocation(&location); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(location)
}

func (l *locationHandler) DeleteLocation(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid ID format"})
	}

	if err := l.location.DeleteLocation(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
