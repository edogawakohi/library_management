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

func ListHistory(lib *Library) error {
	if len(lib.transaction) == 0 {
		fmt.Println("No one borrowed the book")
		return nil
	}
	fmt.Printf("%-11s | %-30s | %-30s | %-20s | %-20s\n", "TRANSACTION ID", "BOOK NAME", "BORROWER", "BORROW DATE", "RETURN DATE")
	fmt.Println(strings.Repeat("-", 150))

	for _, transaction := range lib.transaction {
		var bookName string
		var borrowerName string

		if book, exist := lib.book[transaction.BookId]; exist {
			bookName = book.Title
		}

		if borrower, exist := lib.borrower[transaction.BorrowId]; exist {
			borrowerName = borrower.Name
		}

		returnBook := " "
		if !transaction.DateReturn.IsZero() {
			returnBook = transaction.DateReturn.Format("2006-01-02")
		}

		fmt.Printf("%-11s | %-20s | %-20s | %-20s | %-20s\n",
			transaction.TransactionId, bookName, borrowerName, transaction.DateBorrow.Format("2006-01-02"),
			returnBook)

	}

	return nil
}

func ReturnBook(lib *Library) error {

	if len(lib.transaction) == 0 {
		fmt.Println("No one borrowed the book")
		return nil
	}

	for {
		transactionId := utils.CheckStringId("Input transaction id: ")

		transaction, exist := lib.transaction[transactionId]

		if !exist {
			fmt.Println("Not found!!!")
			continue
		}

		book, check := lib.book[transaction.BookId]

		if !check {
			fmt.Println("This book is not available in the transaction")
		} else {
			book.IsBorrow = false
			lib.book[transaction.BookId] = book
			transaction.DateReturn = time.Now()
			lib.transaction[transactionId] = transaction
			fmt.Println("Return book is successful!!!")
			return nil
		}
	}

}

func SearchBook(lib *Library) error {
	if len(lib.transaction) == 0 {
		fmt.Println("There are no books in the library")
		return nil
	}

	inputSearch := utils.CheckString("Looking for: ")

	fmt.Printf("%-11s|%-15s|%-10s|%-5s\n", "ID", "TITLE", "AUTHOR", "STATUS")
	fmt.Println(strings.Repeat("-", 40))

	for _, book := range lib.book {
		//chuoi goc
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(inputSearch)) ||
			strings.Contains(strings.ToLower(book.Author), strings.ToLower(inputSearch)) {
			fmt.Println(book.GetInfoBook())
		}

	}
	return nil

}
