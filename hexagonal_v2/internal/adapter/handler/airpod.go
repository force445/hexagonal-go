package handler

import (
	"hexagonal_v2/internal/core/port"

	"github.com/gofiber/fiber/v2"
)

type AirpodHandler interface {
	CreateAirpod(c *fiber.Ctx) error
	GetAirpodByID(c *fiber.Ctx) error
	GetAirpods(c *fiber.Ctx) error
	UpdateAirpod(c *fiber.Ctx) error
	DeleteAirpod(c *fiber.Ctx) error
}

type airpodHandler struct {
	port port.AirpodService
}

func NewAirpodHandler(port port.AirpodService) AirpodHandler {
	return &airpodHandler{port: port}
}

func (a *airpodHandler) CreateAirpod(c *fiber.Ctx) error {
	return nil
}

func (a *airpodHandler) GetAirpodByID(c *fiber.Ctx) error {
	return nil
}

func (a *airpodHandler) GetAirpods(c *fiber.Ctx) error {
	return nil
}

func (a *airpodHandler) UpdateAirpod(c *fiber.Ctx) error {
	return nil
}

func (a *airpodHandler) DeleteAirpod(c *fiber.Ctx) error {
	return nil
}
