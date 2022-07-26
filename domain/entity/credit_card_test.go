package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreditCard_IsNumberValid(t *testing.T) {
	_, err := NewCreditCard("40000000000000000", "Wene Alves", 12, 2027, "123")
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCreditCard("5164114171298912", "Wene Alves", 12, 2027, "123")
	assert.Nil(t, err)
}

func TestCreditCard_IsMonthValid(t *testing.T) {
	_, err := NewCreditCard("5164114171298912", "Wene Alves", 13, 2027, "123")
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCreditCard("5164114171298912", "Wene Alves", 12, 2027, "123")
	assert.Nil(t, err)
}

func TestCreditCard_IsYearValid(t *testing.T) {

	lastYear := time.Now().AddDate(-1, 0, 0).Year()
	yearActual := time.Now().Year()

	_, err := NewCreditCard("5164114171298912", "Wene Alves", 12, lastYear, "123")
	assert.Equal(t, "invalid expiration year", err.Error())

	_, err = NewCreditCard("5164114171298912", "Wene Alves", 12, yearActual, "123")
	assert.Nil(t, err)

	_, err = NewCreditCard("5164114171298912", "Wene Alves", 12, yearActual+1, "123")
	assert.Nil(t, err)
}

func TestCreditCard_IsCvvValid(t *testing.T) {
	_, err := NewCreditCard("5164114171298912", "Wene Alves", 12, 2027, "12")
	assert.Equal(t, "invalid cvv", err.Error())

	_, err = NewCreditCard("5164114171298912", "Wene Alves", 12, 2027, "123")
	assert.Nil(t, err)
}
