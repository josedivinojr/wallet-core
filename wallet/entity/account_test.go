package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, _ := NewClient("John", "j@j.com")
	account := NewAccount(client)

	assert.NotNil(t, account)
	assert.Equal(t, client, account.Client)
	assert.NotNil(t, account.ID)
	assert.Equal(t, 0.0, account.Balance)
}

func TestCreateNewAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)

	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John", "j@j.com")
	account := NewAccount(client)
	account.Credit(100)

	assert.Equal(t, 100.0, account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John", "j@j.com")
	account := NewAccount(client)
	account.Credit(100)
	account.Debit(50)

	assert.Equal(t, 50.0, account.Balance)
}
