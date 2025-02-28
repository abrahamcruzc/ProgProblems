package main

import (
	"fmt"
	"unicode"
)

func caesarCipher(s string, k int32) string {
	k = k % 26
	encrypted := []rune(s)

	for i, char := range encrypted {
		if unicode.IsLetter(char) {
			base := 'a'
			if unicode.IsUpper(char) {
				base = 'A'
			}
			encrypted[i] = rune(base + (char-base+k)%26)
		}
	}
	return string(encrypted)
}

func main() {
	s := "abcdefghijklmnopqrstuvwxyz"
	fmt.Println(caesarCipher(s, 2))
}