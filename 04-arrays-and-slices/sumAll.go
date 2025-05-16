package main

// vardic funcitons can take a variable number of arguments
func SumAll(numberToSum ...[]int) []int {
	// Alternative version with the problem that we might get an error if we insert to
	// to a non-existing position. With the other version we need to worry less about capacity

	// numberCount := len(numberToSum)
	// sums := make([]int, numberCount)
	// for index, numbers := range numberToSum {
	// 	sums[index] = Sum(numbers)
	// }

	var sums []int
	for _, numbers := range numberToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
