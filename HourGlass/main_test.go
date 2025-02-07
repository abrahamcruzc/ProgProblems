package main

import "testing"

func TestHourglassSum(t *testing.T) {
	matrix := [][]int32{
		{-9, -9, -9, 1, 1, 1},
		{0, -9, 0, 4, 3, 2},
		{-9, -9, -9, 1, 2, 3},
		{0, 0, 8, 6, 6, 0},
		{0, 0, 0, -2, 0, 0},
		{0, 0, 1, 2, 4, 0},
	}
	
	want := int32(28)
	got := hourglassSum(matrix)
	
	if want != got {
		t.Errorf("\nwant: %v \ngot: %v", want, got)
	}
}
