package model

import "time"

type Borrower struct {
	Id    string
	Name  string
	Email string
	Date  time.Time
}
