package http

import (
	"fmt"
	"urbanstay-api/internal/domain/entity"
	"urbanstay-api/internal/usecase"
	lg "urbanstay-api/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

type PropertyHandler struct {
	uc *usecase.PropertyUseCase
}

func NewPropertyHandler(uc *usecase.PropertyUseCase) *PropertyHandler {
	return &PropertyHandler{uc: uc}
}

func (h *PropertyHandler) CreateProperty(c fiber.Ctx) error {
	lg.Info("start CreateProperty")
	lg.Info(c.FullPath())
	lg.Debug(fmt.Sprintf("body request: %v", string(c.Body())))

	// pegando o body do request
	var property entity.Property
	if err := c.Bind().Body(&property); err != nil {
		lg.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// chamando usecase passando o body
	err := h.uc.ExecuteCreate(c.Context(), &property)
	if err != nil {
		lg.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// resposta de sucesso
	status := fiber.StatusCreated
	lg.Debug("response:")
	lg.Info(fmt.Sprintf("status: %d", status))
	return c.SendStatus(status)
}

func (h *PropertyHandler) ListProperties(c fiber.Ctx) error {
	var properties []*entity.Property

	// listando properties
	properties, err := h.uc.ExecuteList(c.Context())
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&properties)
}
