package models

import "time"

type Transaction struct {
	Id        string    `json:"id"`
	Merchant  Merchant  `json:"merchant"`
	Customer  Customer  `json:"customer"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
