package main

import (
	"testing"
)

func TestBirthdays(t *testing.T) {
	s := []int32{1, 2, 1, 3, 2}
	d := int32(3)
	m := int32(2)

	result := Birthdays(s, d, m)
	expect := int32(2)

	if result != expect {
		t.Errorf("Expect: %v \nResult: %v", expect, result)
	}
}
