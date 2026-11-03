package main

import (
	"errors"
	"time"
)

type BankAccount struct {
	Owner        string
	Balance      int
	Transactions []Transaction
}

func (b *BankAccount) Deposit(amount int) error {
	if amount <= 0 {
		err := errors.New("invalid amount")
		b.SaveTrx(DepositTransaction, Failed, amount, err)
		return err
	}
	b.Balance += amount

	// Save Transaction
	b.SaveTrx(DepositTransaction, Success, amount, nil)
	return nil
}

func (b *BankAccount) Withdraw(amount int) error {
	if amount <= 0 {
		err := errors.New("invalid amount")
		b.SaveTrx(WithdrawalTransaction, Failed, amount, err)
		return err
	}

	if b.Balance < amount {
		err := errors.New("insufficient funds")
		b.SaveTrx(WithdrawalTransaction, Failed, amount, err)
		return err
	}
	b.Balance -= amount

	// Save Transaction
	b.SaveTrx(WithdrawalTransaction, Success, amount, nil)
	return nil
}

func (b *BankAccount) SaveTrx(trxType string, status string, amt int, err error) {
	transaction := Transaction{
		Type:   trxType,
		Status: status,
		Amount: amt,
		Err:    err,
		Time:   time.Now(),
	}

	b.Transactions = append(b.Transactions, transaction)
}
