package api

import (
	"main/internal/service"

	"github.com/gofiber/fiber/v2"
)

type OrderHandler struct {
	App     *fiber.App
	Service service.OrderServiceInterface
}

func (h *OrderHandler) NewOrderHanlder(app *fiber.App, s service.OrderServiceInterface) *OrderHandler {
	return &OrderHandler{
		App:     app,
		Service: s,
	}
}

func (h *OrderHandler) Routes() {

}
