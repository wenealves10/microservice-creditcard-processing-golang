package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCardCredit_IsNumberValid(t *testing.T) {
	_, err := NewCardCredit("40000000000000000", "Wene Alves", 12, 2023, "123")
	assert.Equal(t, "invalid credit card number", err.Error())

	_, err = NewCardCredit("5164114171298912", "Wene Alves", 12, 2023, "123")
	assert.Nil(t, err)
}

func TestCardCredit_IsMonthValid(t *testing.T) {
	_, err := NewCardCredit("5164114171298912", "Wene Alves", 13, 2023, "123")
	assert.Equal(t, "invalid expiration month", err.Error())

	_, err = NewCardCredit("5164114171298912", "Wene Alves", 12, 2023, "123")
	assert.Nil(t, err)
}

func TestCardCredit_IsYearValid(t *testing.T) {

	lastYear := time.Now().AddDate(-1, 0, 0).Year()
	yearActual := time.Now().Year()

	_, err := NewCardCredit("5164114171298912", "Wene Alves", 12, lastYear, "123")
	assert.Equal(t, "invalid expiration year", err.Error())

	_, err = NewCardCredit("5164114171298912", "Wene Alves", 12, yearActual, "123")
	assert.Nil(t, err)

	_, err = NewCardCredit("5164114171298912", "Wene Alves", 12, yearActual+1, "123")
	assert.Nil(t, err)
}
