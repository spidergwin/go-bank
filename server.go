package main

import (
	"net/http"
	"strings"
)

func InitiateServer(port string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Welcome)

	mux.HandleFunc("/health", Healthz)
	mux.HandleFunc("/about", About)

	mux.HandleFunc("/accounts", CreateAccount)
	mux.HandleFunc("/accounts/{owner}", GetAccountDetails)
	mux.HandleFunc("/accounts/{owner}/deposit", DepositHandler)
	mux.HandleFunc("/accounts/{owner}/withdraw", WithdrawalHandler)
	mux.HandleFunc("/accounts/{owner}/transactions", GetTransactions)

	port, _ = strings.CutPrefix(port, ":")

	if err := http.ListenAndServe(port, mux); err != nil {
		return err
	}

	return nil
}
