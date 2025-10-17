package main

import (
	"fmt"
	"main/utils"
)

func main() {
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
		case 2:
			fmt.Println("===List books===")
		case 3:
			fmt.Println("===Add people borrowing books===")
		case 4:
			fmt.Println("===List people borrowing books===")
		case 5:
			fmt.Println("===Borrowing books===")
		case 6:
			fmt.Println("===Books borrowing history===")
		case 7:
			fmt.Println("===Return books===")
		case 8:
			fmt.Println("===Search books===")
		case 9:
			return
		default:
			fmt.Println("Please choice again!!!")
		}
		utils.ReadInput("Pressing enter to continue...")
	}
}
