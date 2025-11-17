package service

import (
	"errors"
	"main/internal/domain"
	"main/internal/repository"
)

type OrderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order *domain.Order) error {
	if order.TotalValue <= 0 {
		return errors.New("order total must be greater than zero")
	}

	// Aqui podem entrar várias regras:
	// - Validar se o cliente existe
	// - Calcular descontos
	// - Gerar número de pedido
	// - Aplicar promoções

	if order.Customer.Name == "" || order.Customer.Email == "" {
		return errors.New("invalid customer data")
	}

	_, err := s.repo.Create(*order)
	if err != nil {
		return err
	}

	return nil
}

func (s *OrderService) ListAllOrders() ([]domain.Order, error) {
	return s.repo.ListAll()
}

func (s *OrderService) GetOrderByID(id string) (*domain.Order, error) {
	if id == "" {
		return nil, errors.New("invalid id")
	}

	order, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) UpdateOrder(order domain.Order) (*domain.Order, error) {
	if order.ID == "" {
		return nil, errors.New("invalid id")
	} else if order.Customer.Name == "" || order.Customer.Email == "" {
		return nil, errors.New("invalid customer data")
	}

	_, err := s.repo.Update(order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
