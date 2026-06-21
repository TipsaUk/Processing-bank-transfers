package inmemory

import (
	"database/sql"
	"processing-bank-transfers/internal/model"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (r *AccountRepository) CreateAccount(account model.BankAccount) error {
	query := `INSERT INTO accounts (id, account_holder, balance, currency) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(query, account.ID, account.AccountHolder, account.Balance, account.Currency)
	return err
}

func (r *AccountRepository) GetBalance(id string) (float64, error) {
	query := `SELECT balance FROM accounts WHERE id = $1`
	var balance float64
	err := r.db.QueryRow(query, id).Scan(balance)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
