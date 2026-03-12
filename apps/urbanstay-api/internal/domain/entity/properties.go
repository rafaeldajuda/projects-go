package entity

import "time"

type Property struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	PricePerNight float64   `json:"price_per_night"`
	Address       string    `json:"address"`
	CreatedAt     time.Time `json:"created_at"`
}
