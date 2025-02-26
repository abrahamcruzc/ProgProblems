package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	str := "ababa"
	
	got := IsPalindrome(str)
	want := true
	
	if want != got {
		t.Errorf("\nwant: %v \ngot: %v", want, got)
	}
}