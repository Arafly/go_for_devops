package main

import (
	"reflect"
	"testing"

	"gopkg.in/check.v1"
)

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

	got :=  SumAll([]int{3, 5}, []int{0, 9})
	expected := []int{8, 9}

	// DeepEqual is a function from the reflect package. It can be used to check if two variables are the same.
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected %v want %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Expected %v want %v", got, expected)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		checkSums(t, got, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		expected := []int{0, 9}

		checkSums(t, got, expected)
	})
}