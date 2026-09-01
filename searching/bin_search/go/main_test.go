// main_test.go
package main

import (
	"errors"
	"log"
	"testing"
)

func TestBinSearch(t *testing.T) {
	type result struct {
		idx int
		err error
	}

	testCases := []struct {
		name string
		arr  []int
		n    int
		want result
	}{
		{
			name: "Should find 4 at pos 3.",
			arr:  []int{1, 2, 3, 4, 5},
			n:    4,
			want: result{
				idx: 3,
				err: nil,
			},
		},
		{
			name: "Should return error.",
			arr:  []int{1, 2, 3, 4, 5},
			n:    7,
			want: result{
				idx: 0,
				err: errors.New("7 is not there ..."),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := binSearch(tc.arr, tc.n)

			if got != tc.want.idx || err != nil && err.Error() != tc.want.err.Error() {
				log.Fatalf("Test failed: %s", err)
			}

			log.Println("Test passed.")

		})
	}
}
