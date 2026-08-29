package main

import "testing"

func TestTwoSum1(t *testing.T) {
	problems := []struct {
		arr      []int
		target   int
		expected bool
	}{
		{
			arr:      []int{10, 20, 35, 50},
			target:   70,
			expected: true,
		},
		{
			arr:      []int{10, 20, 30},
			target:   70,
			expected: false,
		},
		{
			arr:      []int{-8, 1, 4, 6, 10, 45},
			target:   16,
			expected: true,
		},
	}

	for _, problem := range problems {
		if TwoSum1(problem.arr, problem.target) == problem.expected {
			println("Success!")
		} else {
			t.Error("Test failed.")
		}
	}

}

func TestTwoSum2(t *testing.T) {
	problems := []struct {
		arr      []int
		target   int
		expected bool
	}{
		{
			arr:      []int{10, 20, 35, 50},
			target:   70,
			expected: true,
		},
		{
			arr:      []int{10, 20, 30},
			target:   70,
			expected: false,
		},
		{
			arr:      []int{-8, 1, 4, 6, 10, 45},
			target:   16,
			expected: true,
		},
	}

	for _, problem := range problems {
		if TwoSum2(problem.arr, problem.target) == problem.expected {
			println("Success!")
		} else {
			t.Error("Test failed.")
		}
	}

}
