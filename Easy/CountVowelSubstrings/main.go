package main

import (
	"fmt"
)

func countVowelSubstrings(word string) int {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}	
	var count int

	for start := range word {
		freq := make(map[byte]int)
		var uniqueVowels int

		for end := start; end < len(word); end++ {
			char := word[end]

			if !vowels[char] {
				break
			}

			if freq[char] == 0 {
				uniqueVowels++
			}

			freq[char]++

			if uniqueVowels == 5 {
				count++
			}
		}
	}

	return count
}

func main() {
	fmt.Println(countVowelSubstrings("aeiouu")) // 2
	fmt.Println(countVowelSubstrings("unicornarihan")) // 0
	fmt.Println(countVowelSubstrings("cuaieuouac")) // 7
}