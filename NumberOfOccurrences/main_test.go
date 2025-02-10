package main

import "testing"

func TestAreOccurrencesEqual(t *testing.T) {
	s := "abacbc"
	
	want := true
	got := areOccurrencesEqual(s)
	
	if want != got {
		t.Errorf("\nwant: %v \ngot: %v", want, got)
	}
}