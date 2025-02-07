package main

import "testing"

func TestMatchingStrigs(t *testing.T) {
	stringList := []string{"ab", "ab", "abc"}
	queries := []string{"ab", "abc", "bc"}
	
	want := []int32{2,1,0}
	got := matchingStrings(stringList, queries)
	
	if len(want) != len(got) {
		t.Errorf("\nwant: %v \ngot: %v", want, got)
		return
	}
	for i := range want {
		if want[i] != got[i] {
			t.Errorf("\nwant: %v \ngot: %v", want, got)
		}
	}
}