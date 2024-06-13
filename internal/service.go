package internal

import (
	"errors"
)

func FindMax(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, errors.New("slice is empty")
	}

	maxNumber := nums[0]
	for _, num := range nums {
		if num > maxNumber {
			maxNumber = num
		}
	}
	return maxNumber, nil
}
