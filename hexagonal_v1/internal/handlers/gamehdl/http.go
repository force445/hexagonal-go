package gamehdl

import (
	"hexagonal/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type HTTPHandler struct {
	gameService ports.GameService
}

func NewHTTPHandler(gameService ports.GameService) *HTTPHandler {
	return &HTTPHandler{
		gameService: gameService,
	}
}

func (h *HTTPHandler) Get(c *fiber.Ctx) error {
	game, err := h.gameService.Get(c.Params("id"))
	if err != nil {
		return c.JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(game)
}

// func (hdl *HTTPHandler) RevealCell(c *fiber.Ctx) {
// 	body := BodyRevealCell{}
// 	c.BodyParser(&body)

// 	game, err := hdl.gameService.Reveal(c.Params("id"), body.Row, body.Col)
// 	if err != nil {
// 		c.JSON(fiber.Map{"message": err.Error()})
// 		return
// 	}

// 	c.JSON(BuildResponseRevealCell(game))
// }
