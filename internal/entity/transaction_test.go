package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	clientFrom, _ := NewClient("John Doe", "L4uQK@example.com")
	accountFrom := NewAccount(clientFrom)
	accountFrom.Credit(1000.00)

	clientTo, _ := NewClient("Jane Doe", "onpen@huitiho.lv")
	accountTo := NewAccount(clientTo)
	accountTo.Credit(1000.00)

	transaction, err := NewTransaction(accountFrom, accountTo, 100.00)

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, 100.00, transaction.Amount)
	assert.Equal(t, transaction.AccountFrom.ID, accountFrom.ID)
	assert.Equal(t, 900.00, accountFrom.Balance)
	assert.Equal(t, transaction.AccountTo.ID, accountTo.ID)
	assert.Equal(t, 1100.00, accountTo.Balance)
}

func TestCreateTransactionWithZeroAmount(t *testing.T) {
	clientFrom, _ := NewClient("John Doe", "L4uQK@example.com")
	accountFrom := NewAccount(clientFrom)
	accountFrom.Credit(1000.00)

	clientTo, _ := NewClient("Jane Doe", "onpen@huitiho.lv")
	accountTo := NewAccount(clientTo)
	accountTo.Credit(1000.00)

	transaction, err := NewTransaction(accountFrom, accountTo, 0.00)

	assert.NotNil(t, err)
	assert.Error(t, err, "Amount must be greater than 0")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.00, accountFrom.Balance)
	assert.Equal(t, 1000.00, accountTo.Balance)
}

func TestCreateTransactionWithNegativeAmount(t *testing.T) {
	clientFrom, _ := NewClient("John Doe", "L4uQK@example.com")
	accountFrom := NewAccount(clientFrom)
	accountFrom.Credit(1000.00)

	clientTo, _ := NewClient("Jane Doe", "onpen@huitiho.lv")
	accountTo := NewAccount(clientTo)
	accountTo.Credit(1000.00)

	transaction, err := NewTransaction(accountFrom, accountTo, -1000.00)

	assert.NotNil(t, err)
	assert.Error(t, err, "Amount must be greater than 0")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.00, accountFrom.Balance)
	assert.Equal(t, 1000.00, accountTo.Balance)
}

func TestCreateTransactionWithInsufficientBalance(t *testing.T) {
	clientFrom, _ := NewClient("John Doe", "L4uQK@example.com")
	accountFrom := NewAccount(clientFrom)
	accountFrom.Credit(1000.00)

	clientTo, _ := NewClient("Jane Doe", "onpen@huitiho.lv")
	accountTo := NewAccount(clientTo)
	accountTo.Credit(1000.00)

	transaction, err := NewTransaction(accountFrom, accountTo, 10000.00)

	assert.NotNil(t, err)
	assert.Error(t, err, "Account balance is not enough")
	assert.Nil(t, transaction)
	assert.Equal(t, 1000.00, accountFrom.Balance)
	assert.Equal(t, 1000.00, accountTo.Balance)
}
