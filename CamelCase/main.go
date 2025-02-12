package main

import (
	"fmt"
	"regexp"
	"unicode"
)

func camelCase(s string) int32 {
	rex := regexp.MustCompile(`[A-Z]`)
	
	words := rex.Split(s, -1)
	
	return int32(len(words))
}

func camelcase(s string) int32 {
	count := int32(1)
	for _, char := range s {
		if unicode.IsUpper(char) {
			count++
		}
	}
	
	return count
}

func main() {
	s := "saveChangesInTheEditor"
	
	fmt.Println(camelcase(s))
	fmt.Println(camelCase(s))
}