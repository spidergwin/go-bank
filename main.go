package main

import "log"

var Accounts = map[string]*BankAccount{}

func main() {
	if exists, _ := FileExists("accounts.json"); exists {
		LoadAccounts("accounts.json")
	}
	log.Fatalln(InitiateServer("3000"))
}
