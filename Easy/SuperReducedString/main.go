package main

import (
	"fmt"
)

func superReducedString(s string) string {
	stack := []rune{}
	
	for _, char := range s {
		// Si la pila no está vacía y el último elemento es igual al actual.
		if len(stack) > 0 && stack[len(stack)-1] == char {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, char)
		}
	}
	
	if len(stack) == 0 {
		return "Empty String"
	}
	
	return string(stack)
}	

func main() {

	s := "abbaadfasgadafas"
	
	fmt.Println(superReducedString(s))
}