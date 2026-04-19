package inmemory

import (
	"context"
	"errors"
	"sync"

	"processing-bank-transfers/internal/model"
)

var ErrAccountNotFound = errors.New("account not found")

type AccountRepository struct {
	mu       sync.RWMutex
	accounts map[string]model.BankAccount
}

func NewAccountRepository() *AccountRepository {
	return &AccountRepository{accounts: make(map[string]model.BankAccount)}
}

func (r *AccountRepository) Create(_ context.Context, account model.BankAccount) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.accounts[account.ID] = account
	return account.ID, nil
}

func (r *AccountRepository) GetByID(_ context.Context, accountID string) (model.BankAccount, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	account, ok := r.accounts[accountID]
	if !ok {
		return model.BankAccount{}, ErrAccountNotFound
	}

	return account, nil
}

func (r *AccountRepository) UpdateBalance(_ context.Context, accountID string, newBalance float64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	account, ok := r.accounts[accountID]
	if !ok {
		return ErrAccountNotFound
	}

	account.Balance = newBalance
	r.accounts[accountID] = account
	return nil
}
