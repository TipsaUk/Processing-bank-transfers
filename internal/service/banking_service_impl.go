package service

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"processing-bank-transfers/internal/model"
	"processing-bank-transfers/internal/repository"
)

var (
	ErrInvalidAmount      = errors.New("amount must be greater than zero")
	ErrInsufficientFunds  = errors.New("insufficient funds")
	ErrSourceEqualsTarget = errors.New("source and target accounts must be different")
)

type BankingServiceImpl struct {
	accounts     repository.AccountRepository
	transactions repository.TransactionRepository
	sequence     atomic.Uint64
}

func NewBankingService(accounts repository.AccountRepository, transactions repository.TransactionRepository) *BankingServiceImpl {
	return &BankingServiceImpl{
		accounts:     accounts,
		transactions: transactions,
	}
}

func (s *BankingServiceImpl) CreateAccount(ctx context.Context, accountHolder, currency string) (string, error) {
	id := s.nextID("acc")
	account := model.BankAccount{
		ID:            id,
		AccountHolder: accountHolder,
		Balance:       0,
		Currency:      currency,
	}

	if _, err := s.accounts.Create(ctx, account); err != nil {
		return "", err
	}

	return id, nil
}

func (s *BankingServiceImpl) GetBalance(ctx context.Context, accountID string) (float64, error) {
	account, err := s.accounts.GetByID(ctx, accountID)
	if err != nil {
		return 0, err
	}

	return account.Balance, nil
}

func (s *BankingServiceImpl) Transfer(ctx context.Context, fromAccount, toAccount string, amount float64) (TransferResult, error) {
	if amount <= 0 {
		return TransferResult{}, ErrInvalidAmount
	}
	if fromAccount == toAccount {
		return TransferResult{}, ErrSourceEqualsTarget
	}

	from, err := s.accounts.GetByID(ctx, fromAccount)
	if err != nil {
		return TransferResult{}, err
	}

	to, err := s.accounts.GetByID(ctx, toAccount)
	if err != nil {
		return TransferResult{}, err
	}

	if from.Balance < amount {
		return TransferResult{}, ErrInsufficientFunds
	}

	if err := s.accounts.UpdateBalance(ctx, fromAccount, from.Balance-amount); err != nil {
		return TransferResult{}, err
	}
	if err := s.accounts.UpdateBalance(ctx, toAccount, to.Balance+amount); err != nil {
		return TransferResult{}, err
	}

	tx := model.Transaction{
		ID:          s.nextID("tx"),
		FromAccount: fromAccount,
		ToAccount:   toAccount,
		Amount:      amount,
		Timestamp:   time.Now().UTC(),
		Status:      "completed",
	}

	txID, err := s.transactions.Create(ctx, tx)
	if err != nil {
		return TransferResult{}, err
	}

	return TransferResult{TransactionID: txID, Status: tx.Status}, nil
}

func (s *BankingServiceImpl) GetTransactionHistory(ctx context.Context, accountID string) ([]model.Transaction, error) {
	if _, err := s.accounts.GetByID(ctx, accountID); err != nil {
		return nil, err
	}

	return s.transactions.ListByAccountID(ctx, accountID)
}

func (s *BankingServiceImpl) nextID(prefix string) string {
	id := s.sequence.Add(1)
	return fmt.Sprintf("%s-%d", prefix, id)
}
