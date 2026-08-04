package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	Owner        string
	Balance      int
	Transactions []string
}

func (b *BankAccount) Deposit(amount int) error {
	if amount <= 0 {
		return errors.New("invalid amount")
	}
	b.Balance += amount
	message := fmt.Sprintf("Deposit successful! $%v", amount)
	b.Transactions = append(b.Transactions, message)

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
	message := fmt.Sprintf("Withdrawal successful! $%v", amount)
	b.Transactions = append(b.Transactions, message)

	return nil
}

func (b *BankAccount) PrintInfo() {
	fmt.Printf("Account Name: %v\nBalance: $%v\n", b.Owner, b.Balance)
}

func (b *BankAccount) PrintBalance() {
	fmt.Printf("Account Balance: $%v\n", b.Balance)
}

func (b *BankAccount) ShowTransactions() {
	if len(b.Transactions) == 0 {
		fmt.Printf("\nNo Transactions Recorded!\n")
	} else {
		fmt.Printf("\nTransaction History: \n")
		for _, transaction := range b.Transactions {
			fmt.Printf("%v\n", transaction)
		}

	}
}
