package http

import (
	"urbanstay-api/internal/domain/entity"
	"urbanstay-api/internal/usercase"

	"github.com/gofiber/fiber/v3"
)

type PropertyHandler struct {
	uc *usercase.PropertyUseCase
}

func NewPropertyHandler(uc *usercase.PropertyUseCase) *PropertyHandler {
	return &PropertyHandler{uc: uc}
}

func (h *PropertyHandler) CreateProperty(c fiber.Ctx) error {
	var property entity.Property

	// pegando o body do request
	if err := c.Bind().Body(&property); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	// chamando usercase passando o body
	err := h.uc.ExecuteCreate(c.Context(), &property)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// resposta de sucesso
	return c.SendStatus(fiber.StatusCreated)
}

func (h *PropertyHandler) ListProperties(c fiber.Ctx) error {
	var properties []*entity.Property

	// listando properties
	properties = h.uc.ExecuteList(c.Context())

	return c.Status(fiber.StatusOK).JSON(&properties)
}
