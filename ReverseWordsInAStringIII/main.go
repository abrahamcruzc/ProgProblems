package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	strArr := strings.Split(s, " ")
	
	for i, word := range strArr {
		runes := []rune(word)
		for j, k := 0, len(runes)-1; j < k; j, k = j+1, k-1 {
			runes[j], runes[k] = runes[k], runes[j]
		}
		
		strArr[i] = string(runes)
	}
	
	return strings.Join(strArr, " ")
}

func main() {
	s := "Let's take LeetCode contest"
	result := reverseWords(s)
	
	fmt.Println(result)
}