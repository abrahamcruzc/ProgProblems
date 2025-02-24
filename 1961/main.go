package main

import (
	"fmt"
	"strings"
)

func isPrefixString(s string, words []string) bool {
	var builder strings.Builder

	for _, word := range words {
		builder.WriteString(word)
        concatenated := builder.String()

        if concatenated == s {
            return true
        }

        if len(concatenated) > len(s) {
            return false
        }
	}
    return false
}

func main() {
	s := "iloveleetcode"
	words := []string{"i", "love", "leetcode", "apples"}

	result := isPrefixString(s, words)
	fmt.Println(result)
}