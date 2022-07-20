package database

import (
	"github.com/disgoorg/snowflake/v2"
	"strconv"
	"time"
)

type Transaction struct {
	Id        string `json:"id"`
	UserId    string `json:"userId"`
	Amount    int64  `json:"amount"`
	Timestamp int64  `json:"timestamp"`
	Type      string `json:"type"`
	From      string `json:"from"`
}

func generateTransactionId() string {
	return strconv.FormatUint(uint64(snowflake.New(time.Now())), 10)
}

func newTransaction(userId string, amount int64, type_ string, from string) Transaction {
	return Transaction{
		Id:        generateTransactionId(),
		UserId:    userId,
		Amount:    amount,
		Timestamp: time.Now().Unix(),
		Type:      type_,
		From:      from,
	}
}

func writeToLedger(txn Transaction) error {
	c := getClient()

	_, err := c.JSONSet(formatTransactionKey(txn.Id), ".", txn)

	if err != nil {
		return err
	}

	return nil
}
