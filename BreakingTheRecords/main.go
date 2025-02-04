package main

func BreakingRecords(scores []int32) []int32 {
	min := scores[0]
	max := scores[0]

	minCount := int32(0)
	maxCount := int32(0)

	for _, value := range scores {
		if value < min {
			min = value
			minCount++
		} else if value > max {
			max = value
			maxCount++
		}
	}

	return []int32{maxCount, minCount}
}
