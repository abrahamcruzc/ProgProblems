package main

import (
	"fmt"
	"unicode"
)

func minimumNumber(n int32, password string) int32 {
	const (
		minLength = 6
		requiredTypes = 4
	)
	
	var isDigit, isLowerCase, isUpperCase, isSpecial int32

	for _, char := range password {
		switch {
		case unicode.IsDigit(char):
			isDigit = 1
		case unicode.IsLower(char):
			isLowerCase = 1
		case unicode.IsUpper(char):
			isUpperCase = 1
		case unicode.IsSymbol(char) || unicode.IsPunct(char):
			isSpecial = 1
		}
	}
	totalTypeToAdd := requiredTypes - (isDigit + isLowerCase + isUpperCase + isSpecial)

	var missingLength int32

	if n <= minLength {
		missingLength = minLength - n
	}

	if missingLength > totalTypeToAdd {
		return missingLength
	}

	return totalTypeToAdd
}

func main() {
	password := "#HackerRank"
	n := int32(11)

	result := minimumNumber(n, password)
	fmt.Println(result)
}
