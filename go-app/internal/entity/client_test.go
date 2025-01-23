package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "L4uQK@example.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "L4uQK@example.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "L4uQK@example.com")
	err := client.Update("Jane Doe Updated", "onpen@huitiho.lv")
	assert.Nil(t, err)
	assert.Equal(t, "Jane Doe Updated", client.Name)
	assert.Equal(t, "onpen@huitiho.lv", client.Email)
}

func TestUpdateClientWhenArgsAreInvalid(t *testing.T) {
	client, _ := NewClient("John Doe", "L4uQK@example.com")
	err := client.Update("", "L4uQK@example.com")
	assert.NotNil(t, err)
	assert.Error(t, err, "Name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "L4uQK@example.com")
	account := NewAccount(client)
	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))
}
