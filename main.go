package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type BankAccount struct {
	Owner   string
	Balance int
}

func (b BankAccount) PrintInfo() {
	fmt.Printf("Account Name: %v\nBalance: $%v\n", b.Owner, b.Balance)
}

func (b BankAccount) PrintBalance() {
	fmt.Printf("Account Balance: $%v\n", b.Balance)
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

	welcome() // Display the welcome message including the user name.

	for {
		displayMenu() // Display the transactional options for the user.

		var menuChoice int
		if choice, err := readInt("Choose"); err != nil {
			fmt.Println(err)
		} else {
			menuChoice = choice
		}

		switch menuChoice {
		case 1:
			handleDeposit(&account)
		case 2:
			handleWithdrawal(&account)
		case 3:
			account.PrintBalance()
		case 4:
			account.PrintInfo()
		case 5:
			fmt.Println("Thank you for banking with us!")
			os.Exit(0)
		default:
			fmt.Println("\nWrong input! Please enter 1, 2, 3, 4 or 5!")
			continue
		}
	}
}

func handleDeposit(account *BankAccount) {
	var amount int
	for {
		if amt, err := readInt("Enter amount"); err != nil {
			fmt.Printf("\nInvalid amount: %v\n", err)
		} else {
			amount = amt
			break
		}
	}

	if err := account.Deposit(amount); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Deposit successful! $%v\n", amount)
	}
}

func handleWithdrawal(account *BankAccount) {
	var amount int
	for {
		if amt, err := readInt("Enter amount"); err != nil {
			fmt.Printf("\nInvalid amount: %v\n", err)
		} else {
			amount = amt
			break
		}
	}
	var choice string
	for {
		fmt.Printf("\nWithdraw $%v? (y/n) ", amount)

		if _, err := fmt.Scan(&choice); err != nil {
			fmt.Printf("\nInvalid choice: %v", err)
		} else {
			break
		}
	}

	choice = strings.ToLower(choice)
	switch choice {
	case "y", "yes":
		if err := account.Withdraw(amount); err != nil {
			fmt.Println(err)
			return
		}
	case "n", "no":
		fmt.Println("Withdrawal request cancelled!")
		return
	}
	fmt.Printf("\nWithdrawal successful! $%v", amount)
}

func welcome() {
	welcome := `
************************
                       *
      GO BANK          *
                       *
************************
Welcome Joseph!
`

	fmt.Printf("%v", welcome)
}

func displayMenu() {
	menu := `
========================
1. Deposit
2. Withdraw
3. Check Balance
4. Account Info
5. Exit
========================
`

	fmt.Printf("%v", menu)
}

func readInt(prompt string) (int, error) {
	fmt.Print(prompt + ": ")

	text, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	amount, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		return 0, err
	}

	return amount, nil
}
