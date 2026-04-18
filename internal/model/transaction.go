package model

import "time"

// Transaction describes money transfer operation between two accounts.
type Transaction struct {
	ID          string    `json:"id" db:"id"`
	FromAccount string    `json:"fromAccount" db:"from_account"`
	ToAccount   string    `json:"toAccount" db:"to_account"`
	Amount      float64   `json:"amount" db:"amount"`
	Timestamp   time.Time `json:"timestamp" db:"timestamp"`
	Status      string    `json:"status" db:"status"`
}
