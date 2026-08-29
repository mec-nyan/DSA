// Sum of pair equal to target.
//
// Given a sorted array (asc) and a target, find if there exists any pair of elements (arr[i],
// arr[j]) such that their sum is equal to the target.
//
// Input: arr = [10, 20, 35, 50]; target = 70
// Output: true ([20, 50] add up to 70)

package main

func main() {}

// Solution 1 uses the two pointers technique.
func TwoSum1(arr []int, target int) bool {
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

// Solution 2 uses a map.
func TwoSum2(arr []int, target int) bool {
	compMap := make(map[int]bool, len(arr))

	for _, i := range arr {
		if compMap[i] {
			return true
		}

		compMap[target-i] = true
	}

	return false
}
