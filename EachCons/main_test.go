package main

import "testing"

func TestEachCons(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	n := 3

	got := EachCons(arr, n)
	// want := [][]int{
	// 	{1, 2},
	// 	{2, 3},
	// 	{3, 4},
	// }
	want := [][]int{
		{1, 2, 3},
		{2, 3, 4},
	}

	if len(want) != len(got) || len(want[0]) != len(got[0]) {
		t.Error("Error: the sizes are different")
	}

	for i := range want {
		for j := range want[i] {
			if want[i][j] != got[i][j] {
				t.Errorf(
					"\nIn position i: %d, j: %d, \ngot: %v \nwant: %v",
					i, j, got[i][j], want[i][j],
				)
			}
		}
	}
}
