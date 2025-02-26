package main

import "testing"

func TestStringToArray(t *testing.T) {
	str := "abraham cruz colli"
	
	want := []string{"abraham", "cruz", "colli"}
	got := StringToArray(str)
	
	for i := range want {
		if want[i] != got[i] {
			t.Errorf("\nwant: %v \ngot: %v ", want, got)
		}	
	}
}