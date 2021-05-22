package database

import (
	"encoding/json"

	"github.com/syndtr/goleveldb/leveldb"
)

type Transaction struct {
	FirstObjectInput   []byte `json:"FirstObjectInput"`
	FirstObjectOutput  []byte `json:"FirstObjectOutput"`
	SecondObjectInput  []byte `json:"SecondObjectInput"`
	SecondObjectOutput []byte `json:"SecondObjectOutput"`
	ShiftInfo          []byte `json:"ShiftInfo"`
	HashLink           []byte `json:"HashLink"`
}

var transactionData, _ = leveldb.OpenFile("database/transactionData", nil)

func (transaction Transaction) WriteTransaction() {
	concatenatedTransaction := [][]byte{
		transaction.FirstObjectInput,
		transaction.FirstObjectOutput,
		transaction.SecondObjectInput,
		transaction.SecondObjectOutput,
		transaction.ShiftInfo,
		transaction.HashLink,
	}
	transactionInfo := calc.ConcatenateMessage(concatenatedTransaction)
	transactionHash := calc.Hash(transactionInfo)
	transactionAsBytes, _ := json.Marshal(transaction)
	transactionData.Put(transactionHash, transactionAsBytes, nil)
}
