package main

// Sums up the values of the slice but ignoring the first value
func SumAllTail(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// slice[low:high]
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
