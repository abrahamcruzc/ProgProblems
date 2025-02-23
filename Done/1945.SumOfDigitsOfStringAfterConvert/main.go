package main

import (
	"fmt"
	"strconv"
	"strings"
)

func getLucky(s string, k int) int {
	var numBuilder strings.Builder
	
	for _, char := range s {
		numBuilder.WriteString(fmt.Sprintf("%d", char-'a'+1))
	}
	
	result := numBuilder.String()
	var sum int
	for k > 0 {
		sum = sumDigits(result)
		result = strconv.Itoa(sum)
		k--
	}
	
	return sum
}

func sumDigits(numString string) int {
	var sum int
	for _, digit := range numString {
		sum += int(digit-'0')
	}
	
	return sum
}

func main() {
	s := "leetcode"
	k := 2
	
	result := getLucky(s, k)
	fmt.Println(result)
}