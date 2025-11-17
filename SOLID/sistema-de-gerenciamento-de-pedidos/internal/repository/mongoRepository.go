package repository

import (
	"context"
	"main/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoOrderRepository struct {
	collection *mongo.Collection
}

func NewMongoOrderRepository(db *mongo.Database) *MongoOrderRepository {
	return &MongoOrderRepository{
		collection: db.Collection("orders"),
	}
}

func (m *MongoOrderRepository) Create(order domain.Order) (*domain.Order, error) {
	_, err := m.collection.InsertOne(context.Background(), order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (m *MongoOrderRepository) ListAll() ([]domain.Order, error) {
	cur, err := m.collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	var orders []domain.Order
	if err := cur.All(context.Background(), &orders); err != nil {
		return nil, err
	}
	return orders, nil
}

func (m *MongoOrderRepository) FindByID(id string) (*domain.Order, error) {
	var order domain.Order
	err := m.collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&order)
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (m *MongoOrderRepository) Update(order domain.Order) (*domain.Order, error) {
	_, err := m.collection.UpdateOne(
		context.Background(),
		bson.M{"id": order.ID},
		bson.M{"$set": order},
	)
	if err != nil {
		return nil, err
	}
	return &order, nil
}
