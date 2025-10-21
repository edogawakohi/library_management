package library

import "main/model"

type Library struct {
	book     map[string]model.Book
	borrower map[string]model.Borrower
}

func NewLibrary() *Library {

	return &Library{
		book: make(map[string]model.Book),
	}

}

func NewBorrower() *Library {
	return &Library{
		borrower: make(map[string]model.Borrower),
	}
}
