package main

func Between(a, b int) []int {
	if a > b {
		return []int{}
	}

	var arr []int
	for i := a; i <= b; i++ {
		arr = append(arr, i)
	}

	return arr
}

func main() {}
