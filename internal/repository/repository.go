package repository

import (
	"context"

	"processing-bank-transfers/internal/model"
)

// AccountRepository declares account persistence contract.
type AccountRepository interface {
	Create(ctx context.Context, account model.BankAccount) (string, error)
	GetByID(ctx context.Context, accountID string) (model.BankAccount, error)
	UpdateBalance(ctx context.Context, accountID string, newBalance float64) error
}

// TransactionRepository declares transaction persistence contract.
type TransactionRepository interface {
	Create(ctx context.Context, tx model.Transaction) (string, error)
	ListByAccountID(ctx context.Context, accountID string) ([]model.Transaction, error)
}
