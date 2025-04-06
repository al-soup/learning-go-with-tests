package main

func SumAllTail(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// lice[low:high]
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
