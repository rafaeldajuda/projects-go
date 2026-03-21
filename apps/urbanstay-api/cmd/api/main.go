package main

import (
	"log"
	"urbanstay-api/internal/delivery/http"
	"urbanstay-api/internal/repository/memory"
	"urbanstay-api/internal/usecase"

	"github.com/gofiber/fiber/v3"
)

func init() {

}

func main() {
	app := fiber.New()

	repo := memory.NewMemoryRepository()
	uc := usecase.NewPropertyUseCase(repo)
	handler := http.NewPropertyHandler(uc)

	app.Post("/properties", handler.CreateProperty)
	app.Get("/properties", handler.ListProperties)

	if err := app.Listen(":8000"); err != nil {
		log.Fatal("server error:", err.Error())
	}
}
