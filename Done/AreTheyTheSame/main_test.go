package main

import "testing"

func TestComp(t *testing.T) {
	array1 := []int{121, 144, 19, 161, 19, 144, 19, 11}
	array2 := []int{132, 14641, 20736, 361, 25921, 361, 20736, 361}
	
	want := false
	got := Comp(array1, array2)
	
	if want != got {
		t.Errorf("\nwant: %v \ngot: %v", want, got)
	}
}