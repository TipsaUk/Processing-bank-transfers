package inmemory

import (
	"context"
	"sync"

	"processing-bank-transfers/internal/model"
)

type TransactionRepository struct {
	mu           sync.RWMutex
	transactions []model.Transaction
}

func NewTransactionRepository() *TransactionRepository {
	return &TransactionRepository{transactions: make([]model.Transaction, 0)}
}

func (r *TransactionRepository) Create(_ context.Context, tx model.Transaction) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.transactions = append(r.transactions, tx)
	return tx.ID, nil
}

func (r *TransactionRepository) ListByAccountID(_ context.Context, accountID string) ([]model.Transaction, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]model.Transaction, 0)
	for _, tx := range r.transactions {
		if tx.FromAccount == accountID || tx.ToAccount == accountID {
			result = append(result, tx)
		}
	}

	return result, nil
}
