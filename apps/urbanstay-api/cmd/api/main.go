package main

import (
	"log"
	"urbanstay-api/internal/delivery/http"
	"urbanstay-api/internal/repository/sql"
	"urbanstay-api/internal/usecase"
	lg "urbanstay-api/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func init() {

}

func main() {
	// start log
	lg.StartZapLog()
	defer lg.Sync()

	// start fiber
	app := fiber.New()

	// repo := memory.NewMemoryRepository()
	repo, err := sql.NewSqlProperty("sqlite3", "./data/teste.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer repo.Conn.Close()

	uc := usecase.NewPropertyUseCase(repo)
	handler := http.NewPropertyHandler(uc)

	// routes
	app.Post("/properties", handler.CreateProperty)
	app.Get("/properties", handler.ListProperties)

	// server
	lg.Info("start server")
	if err := app.Listen(":8000"); err != nil {
		lg.Error(err.Error())
	}
}
