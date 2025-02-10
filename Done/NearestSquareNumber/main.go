package main

import (
	"math"
)

func NearestSq(n int) int {
	root := math.Sqrt(float64(n))

	if root == math.Floor(root) {
		return n
	}

	floor := int(root)
	ceil := int(math.Ceil(root))

	return func(floor, ceil int) int {
		if math.Abs(float64(n-(floor*floor))) < math.Abs(float64(n-(ceil*ceil))) {
			return floor * floor
		}
		return ceil * ceil
	}(floor, ceil)
}
