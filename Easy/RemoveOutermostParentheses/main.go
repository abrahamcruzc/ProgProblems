package main

import (
	"fmt"
	"strings"
)

func removeOuterParentheses(s string) string {
	var balance int
	var start int
	var builder strings.Builder
	
	for i, parentheses := range s {
		if parentheses == '(' {
			balance++
		} else if parentheses == ')' {
			balance--
		}

		if balance == 0 {
			builder.WriteString(s[start+1 : i])
			start = i + 1
		}
	}

	return builder.String()
}

func main() {
	s := "(()())(())(()(()))"

	result := removeOuterParentheses(s)

	fmt.Println(result)
}
