package main

import "testing"

func TestSum(t *testing.T) {
	// define the length of the array with [5]int or infer the fixed length with [...]int
	// []int for a slice (variable length). A function with a slice parameter will not compile with an array argument
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		// print default format %v which works well for arrays
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}
