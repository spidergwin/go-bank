package main

import (
	"fmt"
	"strings"
	"time"
)

type Transaction struct {
	Type   string
	Status string
	Amount int
	Time   time.Time
	Err    error
}

const (
	DepositTransaction    = "deposit"
	WithdrawalTransaction = "withdrawal"
	Pending               = "pending"
	Success               = "successful"
	Failed                = "failed"
)

func ParseTrxMsg(trx Transaction) string {
	var trxType string

	if trx.Type == DepositTransaction {
		trxType = strings.Replace(trx.Type, "d", "D", 1)
	} else {
		trxType = strings.Replace(trx.Type, "w", "W", 1)
	}

	var trxErr string
	if trx.Err != nil {
		trxErr = fmt.Sprintf("%s", trx.Err)
	}

	trxMessage := fmt.Sprintf("%s %s - $%v %s || %s ", trxType, trx.Status, trx.Amount, trxErr, trx.Time.Format(time.RFC1123))
	return trxMessage
}
