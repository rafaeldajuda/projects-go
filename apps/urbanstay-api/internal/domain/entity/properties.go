package entity

import (
	"errors"
	"time"
)

type Property struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	PricePerNight float64   `json:"price_per_night"`
	Address       string    `json:"address"`
	CreatedAt     time.Time `json:"created_at"`
}

func (p *Property) Validate() error {
	if p.Name == "" {
		return errors.New("o nome do imóvel é obrigatório")
	}
	if p.PricePerNight <= 0 {
		return errors.New("o preço por noite deve ser maior que zero")
	}

	return nil
}
