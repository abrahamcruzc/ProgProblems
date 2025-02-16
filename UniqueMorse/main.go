package main

import (
	"fmt"
	"strings"
)

func uniqueMorseRepresentations(words []string) int {
	morseCode := map[rune]string{
        'a': ".-", 'b': "-...", 'c': "-.-.", 'd': "-..", 'e': ".",
        'f': "..-.", 'g': "--.", 'h': "....", 'i': "..", 'j': ".---",
        'k': "-.-", 'l': ".-..", 'm': "--", 'n': "-.", 'o': "---",
        'p': ".--.", 'q': "--.-", 'r': ".-.", 's': "...", 't': "-",
        'u': "..-", 'v': "...-", 'w': ".--", 'x': "-..-", 'y': "-.--",
        'z': "--..",
    }

    mapMorseWords := make(map[string]struct{}, len(words))
	for _, word := range words {
		var builder strings.Builder
		
		for _, char := range word {
			builder.WriteString(morseCode[char])
		}
		
		mapMorseWords[builder.String()] = struct{}{}
	}
	
	return len(mapMorseWords)
}

func main() {
	words := []string{"gin", "zen", "gig", "msg"}

	result := uniqueMorseRepresentations(words)
	fmt.Println(result)
}
