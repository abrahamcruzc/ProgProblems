package main

import (
	"fmt"
	"strings"
)

func numOfStrings(patterns []string, word string) int {
	var counter int
	for _, pattern := range patterns {
		if strings.Contains(word, pattern) {
			counter++
		}
	}

	return counter
}

func main() {
	word := "abc"
	patterns := []string{"a", "abc", "bc", "d"}

	result := numOfStrings(patterns, word)
	fmt.Println(result)
}