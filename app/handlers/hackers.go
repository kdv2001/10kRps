package handlers

import (
	"10kRps/app/usecases"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

type HackersHandler struct {
	useCases usecases.HackersUseCases
}

func CreateHackersHandler(cases usecases.HackersUseCases) HackersHandler {
	return HackersHandler{useCases: cases}
}

func (h *HackersHandler) Get(c *fiber.Ctx) error {
	res, err := h.useCases.GetAllHackers(c.Params("group"))
	if err != nil {
		return err
	}
	b, err := json.Marshal(res)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	c.Set("Content-type", "application/json; charset=utf-8")
	return c.Send(b)
}
