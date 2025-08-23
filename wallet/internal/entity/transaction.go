package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	ID          string
	FromAccount *Account
	ToAccount   *Account
	Amount      float64
	CreatedAt   time.Time
}

func NewTransaction(fromAccount, toAccount *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		ID:          uuid.New().String(),
		FromAccount: fromAccount,
		ToAccount:   toAccount,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}
	err := transaction.Validate()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *Transaction) Validate() error {
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if t.FromAccount.Balance < t.Amount {
		return errors.New("insufficient funds")
	}
	return nil
}

func (t *Transaction) Commit() {
	t.FromAccount.Debit(t.Amount)
	t.ToAccount.Credit(t.Amount)
}
