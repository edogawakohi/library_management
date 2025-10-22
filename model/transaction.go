package model

import "time"

type Transaction struct {
	TransactionId string
	BookId        string
	BorrowId      string
	DateBorrow    time.Time
	DateReturn    time.Time
}
