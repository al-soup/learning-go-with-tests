package main

// All about arrays and slices: https://dev.to/dawkaka/go-arrays-and-slices-a-deep-dive-dp8

// a slice is an array without a fixed length

func Sum(numbers []int) (sum int) {
	// range returns index and value. The blank represents a write-only value which can be disregraded.
	for _, number := range numbers {
		sum += number
	}
	// the return type does not need to be returned explicitly
	// because the function has a named return type specified
	return
}
