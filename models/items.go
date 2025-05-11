package models

import "time"

type Item struct {
	ID        string    `json:"id" faker:"uuid_hyphenated"`
	Name      string    `json:"name" faker:"name"`
	Price     float64   `json:"price" faker:"oneof: 15, 56, 43"`
	Quantity  int64     `json:"quantity" faker:"oneof: 10, 20, 30"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
