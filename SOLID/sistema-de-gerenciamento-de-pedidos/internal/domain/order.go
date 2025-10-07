package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	StatusCreated    OrderStatus = "CRIADO"
	StatusProcessing OrderStatus = "PROCESSANDO"
	StatusCompleted  OrderStatus = "CONCLUIDO"
)

type Order struct {
	ID          string      `json:"id" bson:"id"`
	Customer    Customer    `json:"customer" bson:"customer"`
	Products    []Products  `json:"products" bson:"products"`
	TotalValue  float64     `json:"total_value" bson:"total_value"`
	Status      OrderStatus `json:"status" bson:"status"`
	CreatedAt   time.Time   `json:"createdAt" bson:"createdAt"`
	LastUpdated time.Time   `json:"lastUpdated" bson:"lastUpdated"`
}

// Montar pedido
func NewOrder(customer Customer, products []Products) *Order {
	order := &Order{
		ID:          string(uuid.New().String()),
		Customer:    customer,
		Products:    products,
		Status:      StatusCreated,
		CreatedAt:   time.Now(),
		LastUpdated: time.Now(),
	}
	order.CalculateTotalValue()
	return order
}

// Calcular valor total
func (o *Order) CalculateTotalValue() {
	total := 0.
	for _, p := range o.Products {
		total += p.UnitPrice
	}
	o.TotalValue = total
}

// Atualizar Status
func (o *Order) UpdateStatus(newStatus string) error {
	if !o.CanUpdateStatus(newStatus) {
		return errors.New("invalid status")
	}
	o.Status = OrderStatus(newStatus)
	o.LastUpdated = time.Now()
	return nil
}

// Verifica se pode alterar o status
func (o *Order) CanUpdateStatus(newStatus string) bool {
	switch o.Status {
	case StatusCreated:
		return newStatus == string(StatusProcessing)
	case StatusProcessing:
		return newStatus == string(StatusCompleted)
	default:
		return false
	}
}
