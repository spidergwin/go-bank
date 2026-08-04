package main

import (
	"fmt"
	"strings"
)

func CreateAccount() {
	fmt.Println("")
	fmt.Println("Create an Account:")
	fmt.Printf("\t")

	// Grab user registration details

	//--- Get Account Name
	var accountName string
	for {
		if name, err := ReadString("Enter Name"); err != nil {
			fmt.Println("Invalid Name:", err)
		} else {
			accountName = name
			break
		}
	}

	key := strings.ToLower(accountName)

	if _, exists := Accounts[key]; exists {
		fmt.Println("Account already exists.")
		return
	}
	Accounts[key] = &BankAccount{
		Owner:        accountName,
		Balance:      0,
		Transactions: []Transaction{},
	}

	SaveAccounts() // Save accounts to JSON file

	CurrentAccount = Accounts[key]

	fmt.Println("")
	fmt.Println("Account created successfully!!, Enjoy Unparalled Banking!!!")
}

func SwitchAccount() {
	fmt.Println("Sign-in to Account:")
	fmt.Printf("\t")

	//--- Get Account Name
	var accountName string
	for {
		if name, err := ReadString("Enter Name"); err != nil {
			fmt.Println("Invalid Name:", err)
		} else {
			accountName = strings.TrimSpace(strings.ToLower(name))
			break
		}
	}

	if Accounts[accountName] != nil {
		CurrentAccount = Accounts[accountName]
	} else {
		fmt.Println("Account not found! Please create a new account!!!")
	}
}

func HandleDeposit(account *BankAccount) {
	var amount int
	for {
		if amt, err := ReadInt("Enter amount $"); err != nil {
			fmt.Printf("\nInvalid amount: %v\n", err)
		} else {
			amount = amt
			break
		}
	}

	if err := account.Deposit(amount); err != nil {
		fmt.Println(err)
	} else {

		SaveAccounts() // Save accounts to JSON file
		message := fmt.Sprintf("Deposit successful! $%v\n", amount)
		fmt.Println(message)
	}
}

func HandleWithdrawal(account *BankAccount) {
	var amount int
	for {
		if amt, err := ReadInt("Enter amount"); err != nil {
			fmt.Printf("\nInvalid amount: %v\n", err)
		} else {
			amount = amt
			break
		}
	}

	var confirmation bool
	for {
		prompt := fmt.Sprintf("Withdraw $%v?", amount)

		if choice, err := ReadConfirmation(prompt); err != nil {
			fmt.Printf("%v", err)
		} else {
			confirmation = choice
			break
		}
	}

	switch confirmation {
	case true:
		if err := account.Withdraw(amount); err != nil {
			fmt.Println(err)
			return
		} else {

			SaveAccounts() // Save accounts to JSON file
			message := fmt.Sprintf("Withdrawal successful! $%v", amount)
			fmt.Println(message)
		}
	case false:
		fmt.Println("Withdrawal request cancelled!")
		return
	}
}
