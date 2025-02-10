package main

func areOccurrencesEqual(s string) bool {
	m := make(map[rune]int, len(s))

	for _, char := range s {
		m[rune(char)]++
	}

	firstValue := -1
	for _, value := range m {
		if firstValue == -1 {
			firstValue = value
		} else if value != firstValue {
			return false
		}
	}

	return true
}
