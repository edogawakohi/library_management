package model

import "fmt"

type Book struct {
	Id       string
	Title    string
	Author   string
	IsBorrow bool
}

func (book Book) GetInfoBook() string {
	shortId := book.Id[:4] + "..." + book.Id[len(book.Id)-4:]
	status := "Not Borrowed"
	if book.IsBorrow {
		status = "Borrowed"
		return fmt.Sprintf("%-10s|%-10s|%-10s|%-10s\n", shortId, book.Title, book.Author, status)
	}
	return fmt.Sprintf("%-10s|%-10s|%-10s|%-10s\n", shortId, book.Title, book.Author, status)
}
