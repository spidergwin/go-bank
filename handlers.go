package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Welcome
func Welcome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		_, err := fmt.Fprintln(w, "Welcome to Go Banking!")
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Not found", http.StatusNotFound)
	}
}

// Health check
func Healthz(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_, err := fmt.Fprintln(w, "OK")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// About
func About(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	res := Bank{
		Application: "Go Bank",
		Version:     "v5.0",
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Creating Account
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody CreateAccReqFmt

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	key := strings.ToLower(reqBody.Owner)

	var account *BankAccount
	if _, exists := Accounts[key]; exists {
		http.Error(w, "Account already exists", http.StatusConflict)
		return
	}

	account = &BankAccount{
		Owner:        reqBody.Owner,
		Balance:      0,
		Transactions: []Transaction{},
	}
	Accounts[key] = account
	if err := SaveAccounts(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	} // Save accounts to JSON file

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(account); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Retrieving account details
func GetAccountDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	owner := r.PathValue("owner")

	key := strings.ToLower(owner)
	var account *BankAccount
	if acc, exists := Accounts[key]; !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	} else {
		account = acc
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Deposit handler
func DepositHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody DepositWithdrawalReqFmt

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	owner := r.PathValue("owner")
	account, err := GetAccount(owner)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusNotFound)
		return
	}

	if err := account.Deposit(reqBody.Amount); err != nil {
		error := fmt.Sprintf("%s", err)
		http.Error(w, error, http.StatusBadRequest)
		return
	} else {

		if err := SaveAccounts(); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		} // Save accounts to JSON file

		response := DepositWithdrawalRespFmt{
			Message: "Deposit Successful",
			Balance: account.Balance,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

// Withdrawal handler
func WithdrawalHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reqBody DepositWithdrawalReqFmt

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	owner := r.PathValue("owner")

	account, err := GetAccount(owner)
	if err != nil {
		http.Error(w, fmt.Sprintf("%s", err), http.StatusNotFound)
		return
	}

	if err := account.Withdraw(reqBody.Amount); err != nil {
		error := fmt.Sprintf("%s", err)
		http.Error(w, error, http.StatusBadRequest)
		return
	} else {

		if err := SaveAccounts(); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		} // Save accounts to JSON file

		response := DepositWithdrawalRespFmt{
			Message: "Withdrawal Successful",
			Balance: account.Balance,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

// Get transactions
func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	owner := r.PathValue("owner")

	key := strings.ToLower(owner)
	var account *BankAccount
	if acc, exists := Accounts[key]; !exists {
		http.Error(w, "Account not found", http.StatusNotFound)
		return
	} else {
		account = acc
	}

	w.Header().Set("Content-Type", "application/json")
	transactions := account.Transactions
	if err := json.NewEncoder(w).Encode(transactions); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
