package repository

import "main/internal/domain"

type OrderRepository interface {
	Create(order domain.Order) (*domain.Order, error)
	ListAll() ([]domain.Order, error)
	FindByID(id string) (*domain.Order, error)
	Update(order domain.Order) (*domain.Order, error)
}
