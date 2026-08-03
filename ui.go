package main

import (
	"fmt"
)

func welcome() {
	welcome := `
************************
                       *
      GO BANK          *
                       *
************************
`

	fmt.Printf("%v", welcome)
}

func displayMenu(text string) {
	menu := `

Current Account: %v

========================
1. Create Account
2. Switch Account
3. Deposit
4. Withdraw
5. Check Balance
6. Account Info
7. View Transactions
8. Exit
========================`

	fmt.Println(fmt.Sprintf(menu, text))
}
