package service

import (
	"context"
	"testing"

	"processing-bank-transfers/internal/repository/inmemory"
)

func createTestService(t *testing.T) *BankingServiceImpl {
	t.Helper()
	accounts := inmemory.NewAccountRepository()
	transactions := inmemory.NewTransactionRepository()
	return NewBankingService(accounts, transactions)
}

func TestCreateAccountAndGetBalance(t *testing.T) {
	svc := createTestService(t)
	ctx := context.Background()

	accountID, err := svc.CreateAccount(ctx, "Alice", "USD")
	if err != nil {
		t.Fatalf("CreateAccount() error = %v", err)
	}

	balance, err := svc.GetBalance(ctx, accountID)
	if err != nil {
		t.Fatalf("GetBalance() error = %v", err)
	}

	if balance != 0 {
		t.Fatalf("expected zero balance, got %v", balance)
	}
}

func TestTransferSuccess(t *testing.T) {
	svc := createTestService(t)
	ctx := context.Background()

	fromID, _ := svc.CreateAccount(ctx, "Alice", "USD")
	toID, _ := svc.CreateAccount(ctx, "Bob", "USD")

	if err := svc.accounts.UpdateBalance(ctx, fromID, 150); err != nil {
		t.Fatalf("seed balance error = %v", err)
	}

	result, err := svc.Transfer(ctx, fromID, toID, 100)
	if err != nil {
		t.Fatalf("Transfer() error = %v", err)
	}

	if result.Status != "completed" {
		t.Fatalf("expected completed status, got %s", result.Status)
	}

	fromBalance, _ := svc.GetBalance(ctx, fromID)
	toBalance, _ := svc.GetBalance(ctx, toID)

	if fromBalance != 50 {
		t.Fatalf("expected sender balance 50, got %v", fromBalance)
	}
	if toBalance != 100 {
		t.Fatalf("expected recipient balance 100, got %v", toBalance)
	}
}

func TestTransferInsufficientFunds(t *testing.T) {
	svc := createTestService(t)
	ctx := context.Background()

	fromID, _ := svc.CreateAccount(ctx, "Alice", "USD")
	toID, _ := svc.CreateAccount(ctx, "Bob", "USD")

	_, err := svc.Transfer(ctx, fromID, toID, 10)
	if err == nil {
		t.Fatalf("expected insufficient funds error, got nil")
	}
	if err != ErrInsufficientFunds {
		t.Fatalf("expected ErrInsufficientFunds, got %v", err)
	}
}

func TestGetTransactionHistory(t *testing.T) {
	svc := createTestService(t)
	ctx := context.Background()

	fromID, _ := svc.CreateAccount(ctx, "Alice", "USD")
	toID, _ := svc.CreateAccount(ctx, "Bob", "USD")
	thirdID, _ := svc.CreateAccount(ctx, "Charlie", "USD")

	if err := svc.accounts.UpdateBalance(ctx, fromID, 200); err != nil {
		t.Fatalf("seed balance error = %v", err)
	}

	_, _ = svc.Transfer(ctx, fromID, toID, 50)
	_, _ = svc.Transfer(ctx, fromID, thirdID, 30)

	history, err := svc.GetTransactionHistory(ctx, fromID)
	if err != nil {
		t.Fatalf("GetTransactionHistory() error = %v", err)
	}

	if len(history) != 2 {
		t.Fatalf("expected 2 transactions, got %d", len(history))
	}

	for _, tx := range history {
		if tx.Status != "completed" {
			t.Fatalf("expected completed tx, got %s", tx.Status)
		}
	}
}

func TestGetTransactionHistoryForUnknownAccount(t *testing.T) {
	svc := createTestService(t)
	ctx := context.Background()

	_, err := svc.GetTransactionHistory(ctx, "missing-account")
	if err == nil {
		t.Fatalf("expected account not found error, got nil")
	}
}

func TestTransferRequiresDifferentAccounts(t *testing.T) {
	svc := createTestService(t)
	ctx := context.Background()

	id, _ := svc.CreateAccount(ctx, "Alice", "USD")
	if err := svc.accounts.UpdateBalance(ctx, id, 100); err != nil {
		t.Fatalf("seed balance error = %v", err)
	}

	_, err := svc.Transfer(ctx, id, id, 10)
	if err == nil {
		t.Fatalf("expected source equals target error, got nil")
	}
	if err != ErrSourceEqualsTarget {
		t.Fatalf("expected ErrSourceEqualsTarget, got %v", err)
	}
}

func TestTransferInvalidAmount(t *testing.T) {
	svc := createTestService(t)
	ctx := context.Background()

	fromID, _ := svc.CreateAccount(ctx, "Alice", "USD")
	toID, _ := svc.CreateAccount(ctx, "Bob", "USD")

	_, err := svc.Transfer(ctx, fromID, toID, 0)
	if err == nil {
		t.Fatalf("expected invalid amount error, got nil")
	}
	if err != ErrInvalidAmount {
		t.Fatalf("expected ErrInvalidAmount, got %v", err)
	}
}
