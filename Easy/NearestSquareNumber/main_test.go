package main

import "testing"

func TestNearestSq(t *testing.T) {
	n := 81
	
	want := 81
	got := NearestSq(n)
	
	if want != got {
		t.Errorf("\nwant: %v \ngot: %v", want, got)
	}
}