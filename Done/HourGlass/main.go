package main

import (
	"math"
)

func hourglassSum(arr [][]int32) int32 {
	if len(arr) < 3 || len(arr[0]) < 3 {
		return 0
	}
	
	maxSum := int32(math.MinInt32)
	sum := int32(0)

	for i := 0; i < len(arr)-2; i++ {
		for j := 0; j < len(arr[i])-2; j++ {
			sum =
				arr[i][j] +
					arr[i][j+1] +
					arr[i][j+2] +
					arr[i+1][j+1] +
					arr[i+2][j] +
					arr[i+2][j+1] +
					arr[i+2][j+2]

			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}
