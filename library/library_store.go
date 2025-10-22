package library

import "main/model"

type Library struct {
	book        map[string]model.Book
	borrower    map[string]model.Borrower
	transaction map[string]model.Transaction
}

func NewLibrary() *Library {

	return &Library{
		book:        make(map[string]model.Book),
		borrower:    make(map[string]model.Borrower),
		transaction: make(map[string]model.Transaction),
	}

}
