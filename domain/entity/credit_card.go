package entity

import (
	"errors"
	"regexp"
	"time"
)

type CardCredit struct {
	Number          string
	Name            string
	ExpirationMonth uint8
	ExpirationYear  int
	Cvv             string
}

func NewCardCredit(number string, name string, expirationMonth uint8, expirationYear int, cvv string) (*CardCredit, error) {
	card := &CardCredit{
		Number:          number,
		Name:            name,
		ExpirationMonth: expirationMonth,
		ExpirationYear:  expirationYear,
		Cvv:             cvv,
	}

	err := card.IsValid()

	if err != nil {
		return nil, err
	}

	return card, nil
}

func (card *CardCredit) IsValid() error {
	err := card.isNumberValid()

	if err != nil {
		return err
	}

	err = card.isMonthValid()

	if err != nil {
		return err
	}

	err = card.isYearValid()

	if err != nil {
		return err
	}

	return nil
}

func (card *CardCredit) isNumberValid() error {
	// validate credit card number
	re := regexp.MustCompile(`^((5(([1-2]|[4-5])[0-9]{8}|0((1|6)([0-9]{7}))|3(0(4((0|[2-9])[0-9]{5})|([0-3]|[5-9])[0-9]{6})|[1-9][0-9]{7})))|((508116)\\d{4,10})|((502121)\\d{4,10})|((589916)\\d{4,10})|(2[0-9]{15})|(67[0-9]{14})|(506387)\\d{4,10})`)

	if !re.MatchString(card.Number) {
		return errors.New("invalid credit card number")
	}

	return nil
}

func (card *CardCredit) isMonthValid() error {
	// validate expiration month
	if card.ExpirationMonth < 1 || card.ExpirationMonth > 12 {
		return errors.New("invalid expiration month")
	}

	return nil
}

func (card *CardCredit) isYearValid() error {
	if card.ExpirationYear >= time.Now().Year() {
		return nil
	}

	return errors.New("invalid expiration year")
}
