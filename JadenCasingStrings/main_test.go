package main

import "testing"

func TestToJadenCase(t *testing.T ) {
	str := "How can mirrors be real if our eyes are't real"
	
	want := "How Can Mirrors Be Real If Our Eyes Are't Real" 
	got := ToJadenCase(str)
	
	if want != got {
		t.Errorf("\nwant: %v \ngot: %v", want, got)
	}
}