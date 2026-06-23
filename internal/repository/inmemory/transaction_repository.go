package inmemory

import (
	"database/sql"

	"processing-bank-transfers/internal/model"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (r *TransactionRepository) Create(tx model.Transaction) error {
	query := `
		INSERT INTO transactions (
			id,
			from_account,
			to_account,
			amount,
			status,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.Exec(
		query,
		tx.ID,
		tx.FromAccount,
		tx.ToAccount,
		tx.Amount,
		tx.Status,
		tx.Timestamp,
	)

	return err
}

func (r *TransactionRepository) GetTransactionsByAccount(
	accountID string,
) ([]model.Transaction, error) {

	query := `
		SELECT
			id,
			from_account,
			to_account,
			amount,
			created_at,
			status
		FROM transactions
		WHERE from_account = $1
		   OR to_account = $1
		ORDER BY created_at DESC
	`

	rows, err := r.db.Query(query, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []model.Transaction

	for rows.Next() {
		var tx model.Transaction

		err := rows.Scan(
			&tx.ID,
			&tx.FromAccount,
			&tx.ToAccount,
			&tx.Amount,
			&tx.Timestamp,
			&tx.Status,
		)
		if err != nil {
			return nil, err
		}

		transactions = append(transactions, tx)
	}

	return transactions, nil
}
