package service

import (
	"context"

	"processing-bank-transfers/internal/model"
)

// TransferResult describes transfer execution output.
type TransferResult struct {
	TransactionID string `json:"transactionId"`
	Status        string `json:"status"`
}

// BankingService declares main application use-cases.
type BankingService interface {
	CreateAccount(ctx context.Context, accountHolder, currency string) (string, error)
	GetBalance(ctx context.Context, accountID string) (float64, error)
	Transfer(ctx context.Context, fromAccount, toAccount string, amount float64) (TransferResult, error)
	GetTransactionHistory(ctx context.Context, accountID string) ([]model.Transaction, error)
}
