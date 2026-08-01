package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Owner   string
	Balance int
}

func (b BankAccount) PrintBalance() {
	fmt.Printf("Account Name: %v\nBalance: $%v\n", b.Owner, b.Balance)
}

func (b *BankAccount) Deposit(amount int) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}
	b.Balance += amount
	return nil
}

func (b *BankAccount) Withdraw(amount int) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}

	if b.Balance < amount {
		return errors.New("insufficient funds")
	}
	b.Balance -= amount
	return nil
}

func main() {
	account := BankAccount{
		Owner:   "Joseph",
		Balance: 1000,
	}

	account.PrintBalance()

	if err := account.Deposit(500); err != nil {
		fmt.Println(err)
	}

	if err := account.Withdraw(2000); err != nil {
		fmt.Println(err)
	}

	account.PrintBalance()
}
