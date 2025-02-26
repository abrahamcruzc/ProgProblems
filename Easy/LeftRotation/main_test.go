package main

import "testing"

func TestRotateLeft(t *testing.T) {
	d := int32(2)
	arr := []int32{1,2,3,4,5}
	
	want := []int32{3,4,5,1,2}
	got := rotateLeft(d, arr)
	
	for i := range want {
		if want[i] != got[i] {
			t.Errorf("\nwant: %v \ngot: %v", want, got)
		}
	}
}