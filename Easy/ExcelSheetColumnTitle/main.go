package main

import "fmt"

func ConvertToTitle(columnNumber int) string {
	var result string

	for columnNumber > 0 {
		columnNumber--
		letter := 'A' + (columnNumber % 26)
		result = string(letter) + result
		columnNumber /= 26
	}

	return result
}

func main() {
	fmt.Println(ConvertToTitle(1))
	fmt.Println(ConvertToTitle(28))
	fmt.Println(ConvertToTitle(701))
	fmt.Println(ConvertToTitle(2147483647))
}