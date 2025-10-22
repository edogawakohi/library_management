package model

import (
	"fmt"
)

type Borrower struct {
	Id    string
	Name  string
	Email string
}

func (borrower Borrower) GetInfo() string {
	shortId := borrower.Id[:4] + "..." + borrower.Id[len(borrower.Id)-4:]

	return fmt.Sprintf("%-10s|%-10s|%-10s\n", shortId, borrower.Name, borrower.Email)

}
