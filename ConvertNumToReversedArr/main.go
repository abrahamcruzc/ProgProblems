package main

import (
	"fmt"
	"strconv"
)

func Digitize(n int) []int {
	numberStr := strconv.Itoa(n)
	j := len(numberStr)
	arr := make([]int, j)

	for i, ch := range numberStr {
		digit, _ := strconv.Atoi(string(ch))
		arr[j-i-1] = digit
	}

	return arr
}

func main() {
	n := 35231
	result := Digitize(n)
	fmt.Println(result)

}
