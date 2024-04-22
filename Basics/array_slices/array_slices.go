package main

func Sum(numbers []int) int {
	sum := 0
	// The blank identifier can be assigned or declared with any value of any type, with the value discarded harmlessly.
	// It's a bit like writing to the Unix /dev/null file: it's a way to discard a value without using it.
	for _, num := range numbers {
		sum += num
	}
	return sum
}