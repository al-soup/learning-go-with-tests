package main

func Sum(numbers []int) (sum int) {
	// range returns the index (_) and the value by iterating over the array
	for _, number := range numbers {
		sum += number
	}
	return
}
