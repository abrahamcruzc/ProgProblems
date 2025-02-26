package main

import "fmt"

func AmIWilson(n int) bool {
	if n < 2 {
		return false
	}
	fact := factorial(n - 1)
	return (fact+1)%(n*n) == 0
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	fact := 1
	for i := 2; i <= n; i++ {
		fact *= i
	}
	return fact
}

func main() {
	n := 563
	result := AmIWilson(n)
	fmt.Println(result)
}
