package main

import (
	"errors"
	"strings"
)

func GetAccount(owner string) (*BankAccount, error) {
	key := strings.ToLower(owner)
	if acc, exists := Accounts[key]; !exists {
		return nil, errors.New("account not found")
	} else {
		return acc, nil
	}
}
