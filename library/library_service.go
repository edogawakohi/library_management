package library

import (
	"fmt"
	"main/model"
	"main/utils"
	"strings"
)

func AddBook(lib *Library) error {
	id := utils.GenerateId()
	title := utils.CheckString("Input title book: ")
	author := utils.CheckString("Input author: ")

	if _, err := lib.book[id]; err {
		return fmt.Errorf("Book with id: %s da ton tai!", id)
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

func AddBorrower() error {
	return nil
}

func ListBorrower() error {
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
