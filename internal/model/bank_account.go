package model

// BankAccount describes a bank account domain model.
type BankAccount struct {
	ID            string  `json:"id" db:"id"`
	AccountHolder string  `json:"accountHolder" db:"account_holder"`
	Balance       float64 `json:"balance" db:"balance"`
	Currency      string  `json:"currency" db:"currency"`
}
