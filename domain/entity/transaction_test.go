package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransaction_IsValid(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "123"
	transaction.AccountID = "12333"
	transaction.Amount = 100

	assert.Nil(t, transaction.IsValid())
}

func TestTransaction_IsNotValidWithAmountGreaterThan1000(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "123"
	transaction.AccountID = "12333"
	transaction.Amount = 1001

	err := transaction.IsValid()

	assert.Error(t, err)

	assert.Equal(t, "you dont have limit for this transaction", err.Error())
}

func TestTransaction_IsNotValidWithAmountLessThan1(t *testing.T) {
	transaction := NewTransaction()
	transaction.ID = "123"
	transaction.AccountID = "12333"
	transaction.Amount = 0

	err := transaction.IsValid()

	assert.Error(t, err)

	assert.Equal(t, "the amount must be greater than 1", err.Error())
}
