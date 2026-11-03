package main

import (
	"encoding/json"
	"log"
	"os"
)

func LoadAccounts(source string) {
	fileContent, err := os.ReadFile(source)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(fileContent, &Accounts)
	if err != nil {
		log.Fatalln(err)
	}
}

func SaveAccounts() error {
	AccountsJSON, err := json.MarshalIndent(Accounts, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile("accounts.json", AccountsJSON, 0o644)
	if err != nil {
		return err
	}

	return nil
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
