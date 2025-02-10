package main

import (
	"sort"
)

func Comp(array1, array2 []int) bool {
	if len(array1) != len(array2) {
		return false
	} 

	sort.Ints(array1)
	sort.Ints(array2)

	for i := range array1 {
		if array1[i]*array1[i] != array2[i] {
			return false
		}
	}
	
	return true
}
