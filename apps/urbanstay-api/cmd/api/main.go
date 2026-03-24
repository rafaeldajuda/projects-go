package main

import (
	"log"
	"urbanstay-api/internal/delivery/http"
	"urbanstay-api/internal/repository/sql"
	"urbanstay-api/internal/usecase"

	"github.com/gofiber/fiber/v3"
)

func init() {

}

func main() {
	app := fiber.New()

	// repo := memory.NewMemoryRepository()
	repo, err := sql.NewSqlProperty("sqlite3", "./data/teste.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer repo.Conn.Close()

	uc := usecase.NewPropertyUseCase(repo)
	handler := http.NewPropertyHandler(uc)

	app.Post("/properties", handler.CreateProperty)
	app.Get("/properties", handler.ListProperties)

	if err := app.Listen(":8000"); err != nil {
		log.Fatal("server error:", err.Error())
	}
}
