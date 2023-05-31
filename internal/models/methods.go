package models

import (
	"math/rand"
	"strconv"
	"time"
)

func (a *Account) GenerateAccountNumber() {
	rand.NewSource(time.Now().UnixNano())

	randomNumber := rand.Intn(999999)

	a.AccountNumber = strconv.Itoa(randomNumber)
}
