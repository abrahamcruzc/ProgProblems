package main

import "testing"

func TestRemoveChar(t *testing.T) {
	word := "abraham"
	
	got := RemoveChar(word)
	want := "braham"
	
	if want != got {
		t.Errorf("\nwant: %v \ngot: %v", want, got)
	}
}