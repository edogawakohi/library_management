package model

import (
	"fmt"
	"time"
)

type Borrower struct {
	Id         string
	Name       string
	Email      string
	DateBorrow time.Time
	DateReturn time.Time
}

func (borrower Borrower) GetInfo() string {
	shortId := borrower.Id[:4] + "..." + borrower.Id[len(borrower.Id)-4:]

	// Định dạng thời gian
	dateBorrow := borrower.DateBorrow.Format("2006-01-02 15:04")
	dateReturn := borrower.DateReturn.Format("2006-01-02 15:04")

	return fmt.Sprintf("%-10s|%-10s|%-10s|%-15s|%-15s\n", shortId, borrower.Name, borrower.Email, dateBorrow, dateReturn)

}
