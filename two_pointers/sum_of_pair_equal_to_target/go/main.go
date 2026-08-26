// Sum of pair equal to target.
//
// Given a sorted array (asc) and a target, find if there exists any pair of elements (arr[i],
// arr[j]) such that their sum is equal to the target.
//
// Input: arr = [10, 20, 35, 50]; target = 70
// Output: true ([20, 50] add up to 70)

package main

func twoSum(arr []int, target int) bool {
	left := 0
	right := len(arr) - 1

	for left < right {
		sum := arr[left] + arr[right]

		if sum == target {
			return true
		}

		if sum < target {
			left++
		} else {
			right--
		}
	}

	return false
}

func main() {
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
		if twoSum(problem.arr, problem.target) == problem.expected {
			println("== PASS.")
		} else {
			println("xx FAIL.")
		}
	}
}
