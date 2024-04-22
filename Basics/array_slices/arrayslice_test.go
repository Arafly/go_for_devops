package main

import "testing"

func TestSum (t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	expected := Sum(numbers)
	want := 15

	if expected != want {
		t.Errorf("Expected %d want %d, %v", expected, want, numbers)
	}
}