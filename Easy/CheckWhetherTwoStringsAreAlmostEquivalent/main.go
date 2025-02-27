package main

import "fmt"

func checkAlmostEquivalent(word1 string, word2 string) bool {
	freq := [26]int{}

	for _, char := range word1 {
		freq[char-'a']++
	}

	for _, char := range word2 {
		freq[char-'a']--
	}

	for _, dif := range freq {
		if dif > 3 || dif < -3 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(checkAlmostEquivalent("aaaa", "bccb"))           // false
	fmt.Println(checkAlmostEquivalent("abcdeef", "abaaacc"))     // true
	fmt.Println(checkAlmostEquivalent("cccddabba", "babababab")) // true

}
