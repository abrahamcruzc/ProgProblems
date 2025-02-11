package main

import "fmt"

func main() {
	var w int8
	
	fmt.Scan(&w)
	
	if w > 3 && w % 2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}