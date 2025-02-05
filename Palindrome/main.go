package main

import (
	"strings"
)

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)

	size := len(str)

	for i := 0; i < size/2; i++ {
		if str[i] != str[size-1-i] {
			return false
		}
	}

	return true
}
