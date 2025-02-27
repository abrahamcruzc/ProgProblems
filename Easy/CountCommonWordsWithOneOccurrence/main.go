package main

import "fmt"

func countWords(words1 []string, words2 []string) int {
   
    freq1 := make(map[string]int)
    freq2 := make(map[string]int)
    
    for _, word := range words1 {
        freq1[word]++
    }

    for _, word := range words2 {
        freq2[word]++
    }

    var count int

    for key, freq := range freq1 {
        if freq2[key] == 1 && freq == 1{
            count++
        }
    }

	return count
}

func main() {
	fmt.Println(countWords([]string{"leetcode","is","amazing","as","is"}, []string{"amazing","leetcode","is"})) // 2
	fmt.Println(countWords([]string{"b","bb","bbb"}, []string{"a","aa","aaa"})) // 0
	fmt.Println(countWords([]string{"a","ab"}, []string{"a","a","a","ab"})) // 1
}
