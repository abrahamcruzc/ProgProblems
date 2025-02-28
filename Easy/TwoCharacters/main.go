package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return false
		}
	}
	return true
}

func alternate (s string) int32 {
	chars := make(map[rune]struct{})

	for _, char := range s {
		chars[char] = struct{}{}
	}

	var builder strings.Builder
	for char := range chars {
		builder.WriteRune(char)
	}

	uniques := builder.String()
	var maxLength int32

	for i := range uniques {
		for j := i+1; j < len(uniques); j++ {
			builder.Reset()

			for _, char := range s {
				if byte(char) != uniques[i] && byte(char) != uniques[j] {
					builder.WriteRune(char)
				}
			}

			if isValid(builder.String()) {
				if len(builder.String()) > int(maxLength) {
					maxLength = int32(len(builder.String()))
				}
			}
		}
	}

	return maxLength
}

func main() {
	s := "cwomzxmuelmangtosqkgfdqvkzdnxerhravxndvomhbokqmvsfcaddgxgwtpgpqrmeoxvkkjunkbjeyteccpugbkvhljxsshpoymkryydtmfhaogepvbwmypeiqumcibjskmsrpllgbvc"
	fmt.Println(alternate(s))
}