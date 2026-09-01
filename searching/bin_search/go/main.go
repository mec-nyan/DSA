// main.go
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Printf("I 💖 Go!\n")
}

func binSearch(arr []int, n int) (int, error) {
	left := 0
	right := len(arr) - 1

	for left <= right {
		middle := (right-left)/2 + left

		if arr[middle] == n {
			return middle, nil
		}

		if arr[middle] > n {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}

	return 0, errors.New(fmt.Sprintf("%d is not there ...", n))
}
