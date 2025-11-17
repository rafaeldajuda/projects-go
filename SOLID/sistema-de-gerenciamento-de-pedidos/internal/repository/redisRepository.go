package repository

import (
	"context"
	"encoding/json"
	"main/internal/domain"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) *RedisRepository {
	return &RedisRepository{
		client: client,
	}
}

func (r *RedisRepository) Create(order domain.Order) (*domain.Order, error) {
	data, _ := json.Marshal(order)
	err := r.client.Set(context.Background(), order.ID, data, 0).Err()
	if err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *RedisRepository) ListAll() ([]domain.Order, error) {
	var orders []domain.Order
	var cursor uint64
	for {
		keys, cursor, err := r.client.Scan(context.Background(), cursor, "*", 1000).Result()
		if err != nil {
			return nil, err
		}

		for _, v := range keys {
			var order domain.Order
			err := r.client.Get(context.Background(), v).Scan(&order)
			if err != nil {
				return nil, err
			}
			orders = append(orders, order)
		}

		if cursor == 0 {
			break
		}
	}

	return orders, nil
}

func (r *RedisRepository) FindByID(id string) (*domain.Order, error) {
	var order domain.Order
	data, err := r.client.Get(context.Background(), id).Bytes()
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &order); err != nil {
		return nil, err
	}
	return &order, nil
}

func (r *RedisRepository) Update(order domain.Order) (*domain.Order, error) {
	data, _ := json.Marshal(order)
	err := r.client.Set(context.Background(), order.ID, data, 0).Err()
	if err != nil {
		return nil, err
	}
	return &order, nil
}
