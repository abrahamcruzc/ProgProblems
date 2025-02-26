package main

import (
	"fmt"

)

func numJewelsInStones(jewels string, stones string) int {
	mapJewels := make(map[rune]struct{})
	for _, jewel := range jewels {
		mapJewels[jewel] = struct{}{}
	}
	
	var count int
	for _, stone := range stones {
		if _, exist := mapJewels[stone]; exist {
			count++
		}
	}
	
	return count
}

func main() {
	jewels := "aA"
	stones := "aAAbbbb"
	
	result := numJewelsInStones(jewels, stones)
	fmt.Println(result)
}