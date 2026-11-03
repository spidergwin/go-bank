package main

type Bank struct {
	Application string `json:"application"`
	Version     string `json:"version"`
}

type CreateAccReqFmt struct {
	Owner string `json:"owner"`
}

type DepositWithdrawalReqFmt struct {
	Amount int `json:"amount"`
}

type DepositWithdrawalRespFmt struct {
	Message string `json:"message"`
	Balance int    `json:"balance"`
}
