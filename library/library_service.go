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
	fmt.Printf("Id new book: %s\n", id)
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

	if _, err := lib.borrower[id]; err {
		return fmt.Errorf("borrower id %s already exist", id)
	}

	lib.borrower[id] = model.Borrower{
		Id:    id,
		Name:  name,
		Email: email,
	}
	fmt.Printf("Id new borrower: %s\n", id)
	fmt.Println("Added Successfully!!!")
	return nil
}

func ListBorrower(lib *Library) error {
	if len(lib.borrower) == 0 {
		fmt.Println("The library currently has no borrower!")
		return nil
	}

	fmt.Printf("%-11s|%-15s|%15s|\n", "ID", "NAME", "EMAIL")
	fmt.Println(strings.Repeat("-", 50))

	for _, borrower := range lib.borrower {
		fmt.Print(borrower.GetInfo())
	}

	return nil
}

func BorrowBook(lib *Library) error {
	fmt.Println("DEBUG: Current borrowers in library:")
	for id, b := range lib.borrower {
		fmt.Printf("ID: %s | Name: %s | Email: %s\n", id, b.Name, b.Email)
	}

	fmt.Println("DEBUG: Current book in library:")
	for id := range lib.book {
		fmt.Printf("ID: %s \n", id)
	}
	for {
		bookId := utils.CheckStringId("Input book Id: ")
		borrowId := utils.CheckStringId("Input borrowId: ")

		book, exist := lib.book[bookId]

		if !exist {
			fmt.Println("This book does not exist!")
			continue
		}

		if book.IsBorrow {
			fmt.Println("This book is already borrowed!")
			continue
		}

		_, borrower := lib.borrower[borrowId]

		if !borrower {
			fmt.Println("This borrower does not exist!")
			continue
		}

		//update in map
		book.IsBorrow = true
		lib.book[bookId] = book

		transactionId := utils.GenerateId()

		lib.transaction[transactionId] = model.Transaction{
			TransactionId: transactionId,
			BookId:        bookId,
			BorrowId:      borrowId,
			DateBorrow:    time.Now(),
		}

		fmt.Println("✅ Successful book borrowing!")
		return nil
	}
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
