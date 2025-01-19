package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "L4uQK@example.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.ID, account.Client.ID)
}

func TestNewAccountWithBlankClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestAccountCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "L4uQK@example.com")
	account := NewAccount(client)
	account.Credit(100.00)
	assert.Equal(t, 100.00, account.Balance)
}

func TestAccountDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "L4uQK@example.com")
	account := NewAccount(client)
	account.Credit(200.00)
	account.Debit(50.00)
	assert.Equal(t, 150.00, account.Balance)
}
