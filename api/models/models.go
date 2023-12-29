package model

import "time"

type payment struct {
	ID        string    `json:"id"`
	Status    string    `json:"status"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
