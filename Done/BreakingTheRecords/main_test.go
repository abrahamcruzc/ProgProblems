package main

import "testing"

func TestBreakingRecords(t *testing.T) {
	scores := []int32{10, 5, 20, 20, 4, 5, 2, 25, 1}
	
	result := BreakingRecords(scores)
	expected := []int32{2, 4}
	
	if result[0] != expected[0] && result[1] != expected[1] {
		t.Errorf("Expected: %v \nGot: %v", expected, result)
	}
	
}