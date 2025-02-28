package main

import "fmt"

func removeElement(nums []int, val int) int {
	writeAt := 0

	for _, num := range nums {
		if num != val {
			nums[writeAt] = num
			writeAt++
		}
	}	

	return writeAt
}

func main() {
	arr1 := []int{3,2,2,3}
	fmt.Println(removeElement(arr1, 3)) // 2
	fmt.Println(arr1)
	
	arr2 := []int{0,1,2,2,3,0,4,2}
	fmt.Println(removeElement(arr2, 2)) // 5
	fmt.Println(arr2)

}