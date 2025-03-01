package main

import "fmt"


func migratoryBirds(arr []int32) int32 {
	frequences := make(map[int32]int)

	for _, id := range arr {
		frequences[id]++
	}

	var maxFreq int
	for _, freq := range frequences {
		if freq > maxFreq {
			maxFreq = freq
		}
	}

	var minID int32 = 6
	for id, freq := range frequences {
		if freq == maxFreq {
			if id < minID {
				minID = id
			}
		}
	}

	return minID
}

func main() {
	fmt.Println(migratoryBirds([]int32{1,4,4,4,5,3})) // 4
	fmt.Println(migratoryBirds([]int32{1,1,2,2,3})) // 1
	fmt.Println(migratoryBirds([]int32{1,2,3,4,5,4,3,2,1,3,4})) // 3
}