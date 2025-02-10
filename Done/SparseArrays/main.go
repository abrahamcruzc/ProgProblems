package main

func matchingStrings(stringList []string, queries []string) []int32 {
	result := make([]int32, len(queries))
	
	for i := range queries {
		for j := range stringList {
			if queries[i] == stringList[j] {
				result[i]++
			}
		}
	}
	
	return result
}