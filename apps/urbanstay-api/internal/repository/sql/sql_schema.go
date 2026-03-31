package sql

import "time"

var PropertySchema = `
CREATE TABLE IF NOT EXISTS property (
id VARCHAR(64) NOT NULL PRIMARY KEY,
name  VARCHAR(100),
description VARCHAR(255),
price_per_night DECIMAL(10, 2),
address VARCHAR(200),
created_at TIMESTAMP NOT NULL
);
`

type Property struct {
	ID            string    `db:"id"`
	Name          string    `db:"name"`
	Description   string    `db:"description"`
	PricePerNight float64   `db:"price_per_night"`
	Address       string    `db:"address"`
	CreatedAt     time.Time `db:"created_at"`
}
