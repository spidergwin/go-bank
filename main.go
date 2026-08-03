package main

import (
	"fmt"
	"os"
)

var (
	Accounts       = map[string]*BankAccount{}
	CurrentAccount *BankAccount
)

func main() {
	welcome() // Display the welcome message including the user name.

	for {
		if CurrentAccount == nil {
			displayMenu("None, please select account or create a new account!!!")
		} else {
			displayMenu(CurrentAccount.Owner)
		}
		var menuChoice int
		if choice, err := ReadInt("Choose"); err != nil {
			fmt.Println(err)
		} else {
			menuChoice = choice
		}

		switch menuChoice {
		case 1:
			CreateAccount()
		case 2:
			SwitchAccount()
		case 3:
			if CurrentAccount == nil {
				fmt.Println("No account selected.")
				continue
			}

			HandleDeposit(CurrentAccount)
		case 4:
			if CurrentAccount == nil {
				fmt.Println("No account selected.")
				continue
			}

			HandleWithdrawal(CurrentAccount)
		case 5:
			if CurrentAccount == nil {
				fmt.Println("No Account Selected! please select an account or create a new account!")
				return
			}

			CurrentAccount.PrintBalance()
		case 6:
			if CurrentAccount == nil {
				fmt.Println("No Account Selected! please select an account or create a new account!")
				return
			}

			CurrentAccount.PrintInfo()
		case 7:
			if CurrentAccount == nil {
				fmt.Println("No Account Selected! please select an account or create a new account!")
				return
			}

			CurrentAccount.ShowTransactions()
		case 8:
			fmt.Println("Thank you for banking with us!")
			os.Exit(0)
		default:
			fmt.Println("\nWrong input! Please enter 1, 2, 3, 4 or 5!")
			continue
		}
	}
}
