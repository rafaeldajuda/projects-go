package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"urbanstay-api/internal/delivery/http"
	"urbanstay-api/internal/repository/sql"
	"urbanstay-api/internal/usecase"
	lg "urbanstay-api/pkg/logger"

	"github.com/gofiber/fiber/v3"

	"github.com/joho/godotenv"
)

var (
	host, port, postRoute, getRoute string
	driverDB, dataSourceDB          string
)

func init() {
	os.Setenv("LOG_LEVEL", "DEBUG")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	host = os.Getenv("HOST")
	port = os.Getenv("PORT")
	postRoute = os.Getenv("POST_ROUTE")
	getRoute = os.Getenv("GET_ROUTE")
	driverDB = os.Getenv("DRIVER_DB")
	dataSourceDB = os.Getenv("DATA_SOURCE_DB")
}

func main() {
	// start log
	lg.StartZapLog()
	defer lg.Sync()

	// start fiber
	app := fiber.New()

	// repo := memory.NewMemoryRepository()
	repo, err := sql.NewSqlProperty(driverDB, dataSourceDB)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer repo.Conn.Close()

	uc := usecase.NewPropertyUseCase(repo)
	handler := http.NewPropertyHandler(uc)

	// routes
	app.Post(postRoute, handler.CreateProperty)
	app.Get(getRoute, handler.ListProperties)

	// start server
	go func() {
		lg.Info(fmt.Sprintf("start server port: %s", port))
		if err := app.Listen(fmt.Sprintf("%s:%s", host, port)); err != nil {
			lg.Error(err.Error())
		}
	}()

	// graceful-shutdown
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	// Your cleanup tasks go here
	// Ex: db.Close()
	// fmt.Println("Running cleanup tasks...")

	fmt.Println("Fiber was successful shutdown.")
}
