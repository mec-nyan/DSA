// Sum of pair equal to target.
//
// Given a sorted array (asc) and a target, find if there exists any pair of elements (arr[i],
// arr[j]) such that their sum is equal to the target.
//
// Input: arr = [10, 20, 35, 50]; target = 70
// Output: true ([20, 50] add up to 70)

package main

func TwoSum(arr []int, target int) bool {
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

func main() {}
