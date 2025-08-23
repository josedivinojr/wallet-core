package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John", "john@john.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John", client.Name)
	assert.Equal(t, "john@john.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")

	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {

	client, _ := NewClient("John", "john@john.com")

	err := client.Update("Updated John", "updated@john.com")

	assert.Nil(t, err)
	assert.Equal(t, "Updated John", client.Name)
	assert.Equal(t, "updated@john.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John", "john@john.com")

	err := client.Update("", "updated@john.com")
	assert.Error(t, err, "name is required")

	err = client.Update("Updated John", "")
	assert.Error(t, err, "email is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)

	err := client.AddAccount(account)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
