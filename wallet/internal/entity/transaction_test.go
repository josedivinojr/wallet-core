package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTransaction(t *testing.T) {
	client1, _ := NewClient("John", "j@j.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Marie", "m@m.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 100)

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, account1.ID, transaction.FromAccount.ID)
	assert.Equal(t, account2.ID, transaction.ToAccount.ID)
	assert.Equal(t, float64(100), transaction.Amount)
}

func TestNewTransactionWithInsufficientBalance(t *testing.T) {
	client1, _ := NewClient("John", "j@j.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Marie", "m@m.com")
	account2 := NewAccount(client2)

	account1.Credit(100)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 200)

	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Equal(t, "insufficient funds", err.Error())
}

func TestCommit(t *testing.T) {
	client1, _ := NewClient("John", "j@j.com")
	account1 := NewAccount(client1)
	client2, _ := NewClient("Marie", "m@m.com")
	account2 := NewAccount(client2)

	account1.Credit(1000)
	account2.Credit(1000)

	transaction, err := NewTransaction(account1, account2, 100)

	assert.Nil(t, err)

	transaction.Commit()

	assert.Equal(t, float64(900), account1.Balance)
	assert.Equal(t, float64(1100), account2.Balance)
}
