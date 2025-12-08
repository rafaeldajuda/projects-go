package api

import (
	"main/internal/domain"
	"main/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Handler struct {
	App     *fiber.App
	Service *service.OrderService
}

func NewHandler(service *service.OrderService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) StartHandler(config fiber.Config) {
	h.App = fiber.New(config)
	h.App.Use(logger.New())
}

func (h *Handler) Routes() {
	g := h.App.Group("/v1/orders")
	g.Post("/", h.CreateOrder)
	g.Get("/", h.ListAllOrders)
	g.Get("/:id", h.GetOrderByID)
	g.Put("/:id", h.UpdateOrder)
}

func (h *Handler) CreateOrder(c *fiber.Ctx) error {
	var req domain.Order
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if req.TotalValue < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "order total must be positive",
		})
	}

	if err := h.Service.CreateOrder(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(req)
}

func (h *Handler) ListAllOrders(c *fiber.Ctx) error {
	orders, err := h.Service.ListAllOrders()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error list orders",
		})
	}

	return c.Status(fiber.StatusOK).JSON(orders)
}

func (h *Handler) GetOrderByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid order id value",
		})
	}

	order, err := h.Service.GetOrderByID(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error get order",
		})
	}

	return c.Status(fiber.StatusOK).JSON(order)
}

func (h *Handler) UpdateOrder(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid order id value",
		})
	}

	var req domain.Order
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid body request",
		})
	}
	req.ID = id

	if req.TotalValue < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "order total must be positive",
		})
	}

	_, err := h.Service.UpdateOrder(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "error update order",
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
