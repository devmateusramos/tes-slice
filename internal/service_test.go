package internal

import (
	"errors"
	"testing"
)

func TestFindMax(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
		err      error
	}{
		{[]int{4, 7, 2, 9, 5}, 9, nil},
		{[]int{-1, -7, -2, -9, -5}, -1, nil},
		{[]int{}, 0, errors.New("slice is empty")},
		{[]int{10}, 10, nil},
	}

	for _, test := range tests {
		result, err := FindMax(test.nums)
		if result != test.expected || (err != nil && err.Error() != test.err.Error()) {
			t.Errorf(
				"FindMax(%v) = %v, %v; expected %v, %v",
				test.nums,
				result,
				err,
				test.expected,
				test.err,
			)
		}
	}
}
