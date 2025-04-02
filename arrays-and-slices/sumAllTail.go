package main

func SumAllTail(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		// lice[low:high]
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
