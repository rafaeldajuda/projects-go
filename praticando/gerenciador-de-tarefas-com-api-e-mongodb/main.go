package main

import (
	"context"

	"github.com/gofiber/fiber/v2"

	db "main/db"
	"main/handler"
)

func init() {

}

func main() {
	// start fiber
	app := fiber.New()
	h, err := handler.NewHandler(app)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := h.Mongdb.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	db.ConfigDB(h.Mongdb)

	// add routes v1
	h.Routes()

	// start server
	if err := app.Listen("127.0.0.1:8080"); err != nil {
		panic(err)
	}
}
