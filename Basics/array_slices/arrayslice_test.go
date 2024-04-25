package main

import "testing"

func TestSum (t *testing.T) {
	// t.Run("collection of any size", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3}
	// 	expected := Sum(numbers)
	// 	want := 6

	// 	if expected != want {
	// 		t.Errorf("Expected %d want %d, %v", expected, want, numbers)
	// 	}
	// })

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		expected := Sum(numbers)
		want := 15

		if expected != want {
			t.Errorf("Expected %d want %d, %v", expected, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	want :=  SumAll([]int{3, 5}, []int{0, 9})
	expected := []int{8, 9}

	if want != expected {
		t.Errorf("Expected %v want %v", expected, want)
	}

	// t.Run("make the sums of some slices", func(t *testing.T) {
	// 	numbers1 := []int{1, 2, 3}
	// 	numbers2 := []int{4, 5}

	// 	expected := SumAll(numbers1, numbers2)
	// 	want := []int{6, 9}

	// 	if expected != want {
	// 		t.Errorf("Expected %v want %v, %v, %v", expected, want, numbers1, numbers2)
	// 	}
	// })
}