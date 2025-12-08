package main

import (
	"context"
	"log"
	"main/internal/api"
	"main/internal/repository"
	"main/internal/service"
	"os"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	os.Setenv("MONGO_URI", "")
	os.Setenv("SERVICE_HOST", ":8080")
}

func main() {
	// 1 - Conectar ao MongoDB
	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://admin:admin@localhost:27017"
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}
	collection := client.Database("orderdb").Collection("orders")

	// 2 - Criar repositories
	orderRepository := repository.NewMongoOrderRepository(collection.Database())

	// 3 - Criar Services
	orderService := service.NewOrderService(orderRepository)

	// 4 - Criar API/hanlder
	handler := api.NewHandler(orderService)
	handler.StartHandler(fiber.Config{})
	handler.Routes()

	// 5 - Start do servidor
	err = handler.App.Listen(os.Getenv("SERVICE_HOST"))
	if err != nil {
		log.Fatal(err)
	}
}
