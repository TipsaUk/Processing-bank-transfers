package inmemory

import (
	"context"
	"database/sql"
	"processing-bank-transfers/internal/model"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) Create(account model.BankAccount) error {
	query := `INSERT INTO accounts (id, account_holder, balance, currency) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, account.ID, account.AccountHolder, account.Balance, account.Currency)
	return err
}

func (r *AccountRepository) GetByID(ctx context.Context, id string) (model.BankAccount, error) {
	query := `SELECT id, account_holder, balance, currency FROM accounts WHERE id = $1`
	var account model.BankAccount
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&account.ID,
		&account.AccountHolder,
		&account.Balance,
		&account.Currency,
	)
	if err != nil {
		return model.BankAccount{}, err
	}
	return account, nil
}
