package main

// variadic functions can take a variable number of arguments.
func SumAll(numbersToSum ...[]int) (sums []int) {
	listCount := len(numbersToSum)
	// make allows you to create a slice with a starting capacity of the len of the listCount
	// make([]int, 0, 5) creates a slice with length 0 and capacity 5
	sums = make([]int, listCount)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}
