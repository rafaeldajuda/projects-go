package handler

import (
	"encoding/json"
	"main/db"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type Handler struct {
	App    *fiber.App
	Mongdb *mongo.Client
}

func NewHandler(app *fiber.App) (*Handler, error) {
	mongdb := db.Connect()

	return &Handler{App: app, Mongdb: mongdb}, nil
}

func (h *Handler) Routes() fiber.Router {
	v1 := h.App.Group("/v1")

	v1.Post("/tasks", func(c *fiber.Ctx) error {
		SetHeaderResponse(c)
		id, err := PostTask(c.Body(), h.Mongdb)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		return c.Status(http.StatusOK).SendString(id)
	})

	v1.Get("/tasks", func(c *fiber.Ctx) error {
		SetHeaderResponse(c)
		tasks, err := GetTasks(h.Mongdb)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		tasksB, err := json.Marshal(tasks)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		return c.Status(http.StatusOK).SendString(string(tasksB))
	})

	v1.Get("/tasks/:id", func(c *fiber.Ctx) error {
		SetHeaderResponse(c)
		id := c.Params("id")
		task, err := GetTask(id, h.Mongdb)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		taskB, err := json.Marshal(task)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		return c.Status(http.StatusOK).SendString(string(taskB))
	})

	v1.Put("/tasks/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := PutTask(id, c.Body(), h.Mongdb)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		return c.SendStatus(http.StatusNoContent)
	})

	v1.Delete("/tasks/:id", func(c *fiber.Ctx) error {
		SetHeaderResponse(c)
		id := c.Params("id")
		err := DeleteTask(id, h.Mongdb)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		return c.Status(http.StatusOK).SendString(id)
	})

	return v1
}

func SetHeaderResponse(c *fiber.Ctx) {
	c.Set("Content-Type", "application/json")
}
