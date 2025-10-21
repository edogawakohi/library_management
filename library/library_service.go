package library

import (
	"fmt"
	"main/model"
	"main/utils"
	"strings"
	"time"
)

func AddBook(lib *Library) error {
	id := utils.GenerateId()
	title := utils.CheckString("Input title book: ")
	author := utils.CheckString("Input author: ")

	if _, err := lib.book[id]; err {
		return fmt.Errorf("book with id: %s already exist", id)
	}

	lib.book[id] = model.Book{
		Id:       id,
		Title:    title,
		Author:   author,
		IsBorrow: false,
	}
	fmt.Println("Added Successfully!!!")
	return nil
}

func ListBook(lib *Library) error {

	if len(lib.book) == 0 {
		fmt.Println("The library currently has no books!")
		return nil
	}

	fmt.Printf("%-11s|%-15s|%-10s|%-5s\n", "ID", "TITLE", "AUTHOR", "STATUS")
	fmt.Println(strings.Repeat("-", 40))

	for _, book := range lib.book {
		fmt.Println(book.GetInfoBook())
	}
	return nil
}

func AddBorrower(lib *Library) error {

	id := utils.GenerateId()
	name := utils.CheckString("Input name: ")
	email := utils.CheckEmail("Input email: ")
	dateBorrow := time.Time{}
	dateReturn := time.Time{}

	if _, err := lib.borrower[id]; err {
		return fmt.Errorf("borrower id %s already exist", id)
	}

	lib.borrower[id] = model.Borrower{
		Id:         id,
		Name:       name,
		Email:      email,
		DateBorrow: dateBorrow,
		DateReturn: dateReturn,
	}
	fmt.Println("Added Successfully!!!")
	return nil
}

func ListBorrower(lib *Library) error {
	if len(lib.borrower) == 0 {
		fmt.Println("The library currently has no borrower!")
		return nil
	}

	fmt.Printf("%-11s|%-15s|%15s|%-12s|%-10s\n", "ID", "NAME", "EMAIL", "DATE BORROW", "DATE RETURN")
	fmt.Println(strings.Repeat("-", 50))

	for _, borrower := range lib.borrower {
		fmt.Print(borrower.GetInfo())
	}

	return nil
}

func BorrowBook() error {
	return nil
}

func ListHistory() error {
	return nil
}

func ReturnBook() error {
	return nil
}

func SearchBook() error {
	return nil
}
