package main

import (
	"fmt"
	"main/library"
	"main/utils"
)

func main() {
	lib := library.NewLibrary()
	for {
		fmt.Println("===Library===")
		fmt.Println("1. Add books")
		fmt.Println("2. List books")
		fmt.Println("3. Add people borrowing books")
		fmt.Println("4. List people borrowing books")
		fmt.Println("5. Borrowing books")
		fmt.Println("6. Books borrowing history")
		fmt.Println("7. Return books")
		fmt.Println("8. Search books")
		fmt.Println("9. Exit")
		choice := utils.GetPostiveInt("=> Select from menu: ")

		utils.ClearScreen()
		switch choice {
		case 1:
			fmt.Println("===Add books===")
			if err := library.AddBook(lib); err != nil {
				fmt.Printf("Error: %v\n", err)
			}

		case 2:
			fmt.Println("===List books===")
			if err := library.ListBook(lib); err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		case 3:
			fmt.Println("===Add people borrowing books===")
			if err := library.AddBorrower(lib); err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		case 4:
			fmt.Println("===List people borrowing books===")
			if err := library.ListBorrower(lib); err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		case 5:
			fmt.Println("===Borrowing books===")
			if err := library.BorrowBook(lib); err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		case 6:
			fmt.Println("===Books borrowing history===")
			if err := library.ListHistory(lib); err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		case 7:
			fmt.Println("===Return books===")
			if err := library.ReturnBook(lib); err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		case 8:
			fmt.Println("===Search books===")
			if err := library.SearchBook(lib); err != nil {
				fmt.Printf("Error: %s\n", err)
			}
		case 9:
			return
		default:
			fmt.Println("Please choice again!!!")
		}
		utils.ReadInput("Pressing enter to continue...")
	}
}
