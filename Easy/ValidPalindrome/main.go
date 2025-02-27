package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	i, j := 0, len(s) - 1

	for i < j {
		if !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
			i++
			continue
		}

		if !unicode.IsLetter(rune(s[j])) && !unicode.IsNumber(rune(s[j])) {
			j--
			continue
		}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}

		i++
		j--
	}

	return true
}

func main() {

	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) //true
	fmt.Println(isPalindrome("race a car")) // false
	fmt.Println(isPalindrome("ab2a")) // false
	fmt.Println(isPalindrome("0P")) // false

}