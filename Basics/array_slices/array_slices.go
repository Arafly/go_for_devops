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

// The ... syntax is a variadic function. It allows the function to accept any number of arguments.
// SumAll takes multiple slices of integers and returns a new slice containing the sum of each input slice.
func  SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, nums := range numbersToSum {
		sums[i] = Sum(nums)
	}
	return sums
}

// SumAllTails calculates the totals of the "tails" of each slice. The tail of a collection is all the items apart from the first one (the "head").
// It takes a variadic parameter `numbersToSum` which represents a slice of slices of integers and returns a slice of integers representing the sums of the tails of each input slice.
// If a slice is empty, the sum for that slice will be 0.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, nums := range numbersToSum {
		if len(nums) == 0 {
			sums = append(sums, 0)
		} else {
			// The syntax `nums[1:]` means "take the slice from the 1st element to the end of the slice".
			// This syntax is slice[low:high]. If low is omitted, it's 0. If high is omitted, it's the length of the slice.
			tail := nums[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}