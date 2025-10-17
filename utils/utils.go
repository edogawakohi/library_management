package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"unicode"
)

func ReadInput(promt string) string {
	fmt.Print(promt)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func GetPostiveInt(promt string) int {

	for {
		input := ReadInput(promt)
		value, err := strconv.Atoi(input)

		if err == nil && value > 0 {
			return value
		}

		fmt.Println("Please input positive integer!!!")
	}

}

func GetPostiveIntV2(promt string) int {

	for {
		input := ReadInput(promt)
		value, err := strconv.Atoi(input)

		if err == nil && value > 0 && value <= 10 {
			return value
		}

		fmt.Println("Please input positive integer and from 0 to 10!!!")
	}

}

func CheckInputScore(promt string) float64 {

	for {
		input := ReadInput(promt)
		value, err := strconv.ParseFloat(input, 64)

		if err == nil && value >= 0 && value <= 10 {
			return value
		}

		fmt.Println("Please input positive float and from 0 to 10!!!")
	}

}

func GetPositiveFloat(promt string) float64 {

	for {
		input := ReadInput(promt)
		value, err := strconv.ParseFloat(input, 64)
		var balance int64 = 1_000_000_000_000 //1 thousand VND
		if err == nil && value >= 10_000_000 && value <= float64(balance) {
			return value
		}

		fmt.Println("Please input positive float >= 10.000.000")
	}

}

func CheckString(promt string) string {

	for {
		input := ReadInput(promt)

		if input == "" {
			fmt.Println("Please input cannot empty!!!")
			continue
		}

		isNumber := false
		for _, ch := range input {
			if unicode.IsDigit(ch) {
				isNumber = true
				break
			}
		}

		if isNumber {
			fmt.Println("Input cannot contain number!!!")
			continue
		}

		isSpecial := false
		for _, ch := range input {
			if !unicode.IsDigit(ch) && !unicode.IsLetter(ch) && ch != ' ' {
				isSpecial = true
				break
			}
		}

		if isSpecial {
			fmt.Println("Input cannot contain special character!!!")
			continue
		}

		return input
	}

}

func ClearScreen() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		fmt.Println("Error clearing screen: ", err)
	}
}
