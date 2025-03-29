package models

import (
	"time"

	"github.com/google/uuid"
)

type Customer struct {
	Id        string    `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func NewCustomer(username, password string, balance int) Customer {
	return Customer{
		Id:        uuid.New().String(),
		Username:  username,
		Password:  password,
		Balance:   balance,
		CreatedAt: time.Now(),
	}
}
