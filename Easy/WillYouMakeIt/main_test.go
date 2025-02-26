package main

import "testing"

func TestZeroFuel(t *testing.T) {
	distanceToPump := 70
	mpg := 25
	fuelLeft := 1
	
	want := false
	got := ZeroFuel(distanceToPump, mpg, fuelLeft)
	
	if want != got {
		t.Errorf("\nwant: %v \ngot: %v", want, got)
	}
}