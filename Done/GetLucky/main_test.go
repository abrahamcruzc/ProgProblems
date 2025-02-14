package main

import "testing"

func TestGetLucky(t *testing.T) {
	s := "leetcode"
	k := 2
	
	want := 6
	got := getLucky(s, k)
	
	if want != got {
		t.Errorf("want: %v got: %v", want, got)
	}
}