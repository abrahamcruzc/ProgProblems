package main

import (
	"fmt"
	"strings"
)

func defangIPaddr(address string) string {
	return strings.Join(strings.Split(address, "."), "[.]")
}


func main() {
	ip := "1.1.1.1"
	
	result := defangIPaddr(ip)
	
	fmt.Println(result)
}