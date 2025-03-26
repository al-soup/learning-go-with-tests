package main

import "testing"

// define the length of the array with [5]int or infer the fixed length with [...]int
// []int for a slice (variable length). A function with a slice parameter will not compile with an arra y argument

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}
