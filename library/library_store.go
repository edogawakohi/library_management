package library

import "main/model"

type Library struct {
	book map[string]model.Book
}

type Borrower struct {
	borrower map[string]model.Borrower
}

func NewLibrary() *Library {

	return &Library{
		book: make(map[string]model.Book),
	}
}
