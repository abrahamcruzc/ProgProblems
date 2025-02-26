package main

import "fmt"

func romanToInt(s string) int {
    romanNumbers := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    sum := 0
    prevValue := 0

    for i := len(s) - 1; i >= 0; i-- {
        currentValue := romanNumbers[s[i]]
        if currentValue < prevValue {
            sum -= currentValue
        } else {
            sum += currentValue
        }
        prevValue = currentValue
    }

    return sum
}

func main () {
	input := "IV"
	fmt.Println(romanToInt(input))
}