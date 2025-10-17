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
			fmt.Println("a")
		case 2:
			fmt.Println("a")
		case 3:
			fmt.Println("a")
		case 4:
			fmt.Println("a")
		case 5:
			fmt.Println("a")
		case 6:
			fmt.Println("a")
		case 7:
			fmt.Println("a")
		case 8:
			fmt.Println("a")
		case 9:
			return
		default:
			fmt.Println("Please choice again!!!")
		}
		utils.ReadInput("Pressing enter to continue...")
	}
}
